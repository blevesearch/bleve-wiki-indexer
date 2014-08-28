//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package main

import (
	"log"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/libgit2/git2go"
	"gopkg.in/fsnotify.v1"
)

func startWatching(path string, index bleve.Index, repo *git.Repository) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// start a go routine to process events
	go func() {
		idleTimer := time.NewTimer(10 * time.Second)
		queuedEvents := make([]fsnotify.Event, 0)
		for {
			select {
			case ev := <-watcher.Events:
				queuedEvents = append(queuedEvents, ev)
				idleTimer.Reset(10 * time.Second)
			case err := <-watcher.Errors:
				log.Fatal(err)
			case <-idleTimer.C:
				for _, ev := range queuedEvents {
					if pathMatch(ev.Name) {
						switch ev.Op {
						case fsnotify.Remove, fsnotify.Rename:
							// delete the path
							processDelete(index, repo, ev.Name)
						case fsnotify.Create, fsnotify.Write:
							// update the path
							processUpdate(index, repo, ev.Name)
						default:
							// ignore
						}
					}
				}
				queuedEvents = make([]fsnotify.Event, 0)
				idleTimer.Reset(10 * time.Second)
			}
		}
	}()

	// now actually watch the path requested
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("watching '%s' for changes...", path)

	return watcher
}
