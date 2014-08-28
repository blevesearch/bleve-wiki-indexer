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
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"regexp"

	bleveHttp "github.com/couchbaselabs/bleve/http"
)

var dir = flag.String("dir", "", "directory to monitor and index")
var indexPath = flag.String("index", "wiki.bleve", "path to store index")
var pathFilter = flag.String("pathFilter", `\.md$`, "regular expression that file names must match")
var staticEtag = flag.String("staticEtag", "", "A static etag value.")
var staticPath = flag.String("static", "static/", "Path to the static content")
var bindAddr = flag.String("addr", ":8099", "http listen address")
var pathRegexp *regexp.Regexp

func main() {

	flag.Parse()

	if *dir == "" {
		log.Fatalf("must specify a directory to watch")
	}

	// cleanup the dir
	*dir = filepath.Clean(*dir)

	var err error
	if *pathFilter != "" {
		pathRegexp, err = regexp.Compile(*pathFilter)
		if err != nil {
			log.Fatal(err)
		}
	}

	// open the index
	index := openIndex(*indexPath)

	// open the git repo
	repo := openGitRepo(*dir)

	// create a router to serve static files
	router := staticFileRouter()

	// add the API
	bleveHttp.RegisterIndexName("wiki", index)
	searchHandler := bleveHttp.NewSearchHandler("wiki")
	router.Handle("/api/search", searchHandler).Methods("POST")

	// start a watcher on the directory
	watcher := startWatching(*dir, index, repo)
	defer watcher.Close()

	// walk the directory to ensure current
	walkForIndexing(*dir, index, repo)

	http.Handle("/", router)
	log.Printf("Listening on %v", *bindAddr)
	log.Fatal(http.ListenAndServe(*bindAddr, nil))
}

func pathMatch(path string) bool {
	if pathRegexp != nil {
		return pathRegexp.MatchString(path)
	}
	return true
}
