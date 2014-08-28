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
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/mschoch/blackfriday-text"
	"github.com/russross/blackfriday"
)

type WikiPage struct {
	Name               string    `json:"name"`
	Body               string    `json:"body"`
	ModifiedBy         string    `json:"modified_by"`
	ModifiedByName     string    `json:"modified_by_name"`
	ModifiedByEmail    string    `json:"modified_by_email"`
	ModifiedByGravatar string    `json:"modified_by_gravatar"`
	Modified           time.Time `json:"modified"`
}

func (w *WikiPage) Type() string {
	return "wiki"
}

func NewWikiFromFile(path string) (*WikiPage, error) {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cleanedUpBytes := cleanupMarkdown(fileBytes)

	name := path
	lastSlash := strings.LastIndex(path, string(os.PathSeparator))
	if lastSlash > 0 {
		name = name[lastSlash+1:]
	}
	if strings.HasSuffix(name, ".md") {
		name = name[0 : len(name)-len(".md")]
	}
	rv := WikiPage{
		Name: name,
		Body: string(cleanedUpBytes),
	}
	return &rv, nil
}

func cleanupMarkdown(input []byte) []byte {
	extensions := 0
	renderer := blackfridaytext.TextRenderer()
	output := blackfriday.Markdown(input, renderer, extensions)
	return output
}
