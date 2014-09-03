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

	nameMapping := bleve.NewDocumentMapping().
		AddFieldMapping(
		bleve.NewFieldMapping(
			"", "text", textFieldAnalyzer,
			true, true, true, true))

	bodyMapping := bleve.NewDocumentMapping().
		AddFieldMapping(
		bleve.NewFieldMapping(
			"", "text", textFieldAnalyzer,
			true, true, true, true))

	modifiedByMapping := bleve.NewDocumentMapping().
		AddFieldMapping(
		bleve.NewFieldMapping(
			"", "text", textFieldAnalyzer,
			true, true, true, true))

	modifiedByNameMapping := bleve.NewDocumentMapping().
		AddFieldMapping(
		bleve.NewFieldMapping(
			"", "text", textFieldAnalyzer,
			true, true, true, true))

	modifiedByEmailMapping := bleve.NewDocumentMapping().
		AddFieldMapping(
		bleve.NewFieldMapping(
			"", "text", textFieldAnalyzer,
			true, true, true, true))

	modifiedByAvatarMapping := bleve.NewDocumentMapping().
		AddFieldMapping(
		bleve.NewFieldMapping(
			"", "text", textFieldAnalyzer,
			true, false, false, false))

	modifiedMapping := bleve.NewDocumentMapping().
		AddFieldMapping(
		bleve.NewFieldMapping(
			"", "datetime", textFieldAnalyzer,
			true, true, true, true))

	wikiMapping := bleve.NewDocumentMapping().
		AddSubDocumentMapping("name", nameMapping).
		AddSubDocumentMapping("body", bodyMapping).
		AddSubDocumentMapping("modified_by", modifiedByMapping).
		AddSubDocumentMapping("modified_by_name", modifiedByNameMapping).
		AddSubDocumentMapping("modified_by_email", modifiedByEmailMapping).
		AddSubDocumentMapping("modified_by_avatar", modifiedByAvatarMapping).
		AddSubDocumentMapping("modified", modifiedMapping)

	indexMapping := bleve.NewIndexMapping().
		AddDocumentMapping("wiki", wikiMapping)

	indexMapping.DefaultAnalyzer = textFieldAnalyzer

	return indexMapping
}
