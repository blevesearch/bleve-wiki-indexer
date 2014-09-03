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
	"github.com/blevesearch/bleve"
)

const textFieldAnalyzer = "en"

func buildIndexMapping() *bleve.IndexMapping {

	enTextFieldMapping := bleve.NewTextFieldMapping()
	enTextFieldMapping.Analyzer = textFieldAnalyzer

	storeFieldOnlyMapping := bleve.NewTextFieldMapping()
	storeFieldOnlyMapping.Index = false
	storeFieldOnlyMapping.IncludeTermVectors = false
	storeFieldOnlyMapping.IncludeInAll = false

	dateTimeMapping := bleve.NewDateTimeFieldMapping()

	wikiMapping := bleve.NewDocumentMapping()
	wikiMapping.AddFieldMappingsAt("name", enTextFieldMapping)
	wikiMapping.AddFieldMappingsAt("body", enTextFieldMapping)
	wikiMapping.AddFieldMappingsAt("modified_by", enTextFieldMapping)
	wikiMapping.AddFieldMappingsAt("modified_by_name", enTextFieldMapping)
	wikiMapping.AddFieldMappingsAt("modified_by_email", enTextFieldMapping)
	wikiMapping.AddFieldMappingsAt("modified_by_avatar", storeFieldOnlyMapping)
	wikiMapping.AddFieldMappingsAt("modified", dateTimeMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("wiki", wikiMapping)

	indexMapping.DefaultAnalyzer = textFieldAnalyzer

	return indexMapping
}
