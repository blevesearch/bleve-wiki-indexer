## bleve-wiki-indexer

[Search the bleve wiki](http://wikisearch.blevesearch.com/)

This application uses the [bleve](http://www.blevesearch.com) library to build an index of the bleve [wiki](https://github.com/blevesearch/bleve/wiki).

The application monitors the specified directory for changes, and upon finiding changes it reindexes those files.  Keeping the directory up to date is the responsibility of another application.

The application exposes a search interface on port 8099.

## Status

[![Build Status](https://drone.io/github.com/blevesearch/bleve-wiki-indexer/status.png)](https://drone.io/github.com/blevesearch/bleve-wiki-indexer/latest)