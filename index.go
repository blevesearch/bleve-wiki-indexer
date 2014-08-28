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
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/libgit2/git2go"
)

func openIndex(path string) bleve.Index {
	index, err := bleve.Open(path)
	if err == bleve.ERROR_INDEX_PATH_DOES_NOT_EXIST {
		log.Printf("Creating new index...")
		// create a mapping
		indexMapping := buildIndexMapping()
		index, err = bleve.New(path, indexMapping)
		if err != nil {
			log.Fatal(err)
		}
	} else if err == nil {
		log.Printf("Opening existing index...")
	} else {
		log.Fatal(err)
	}
	return index
}

func processUpdate(index bleve.Index, repo *git.Repository, path string) {
	log.Printf("updated: %s", path)
	rp := relativePath(path)
	wiki, err := NewWikiFromFile(path)
	if err != nil {
		log.Print(err)
	} else {
		doGitStuff(repo, rp, wiki)
		index.Index(rp, wiki)
	}
}

func processDelete(index bleve.Index, repo *git.Repository, path string) {
	log.Printf("delete: %s", path)
	rp := relativePath(path)
	err := index.Delete(rp)
	if err != nil {
		log.Print(err)
	}
}

func relativePath(path string) string {
	if strings.HasPrefix(path, *dir) {
		path = path[len(*dir)+1:]
	}
	return path
}
