package main

import (
	"github.com/blevesearch/bleve/v2"
)

func ReadIndex(index bleve.Index) (*bleve.SearchResult, error) {
	search, err := index.Search(bleve.NewSearchRequest(bleve.NewQueryStringQuery("bodyindex1")))
	if err != nil {
		return nil, err
	}
	return search, nil
}

func ReadHighlightIndex(index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewPhraseQuery([]string{"test", "bodyindex1"}, "body")
	request := bleve.NewSearchRequest(query)
	request.Highlight = bleve.NewHighlight()
	search, err := index.Search(request)
	if err != nil {
		return nil, err
	}
	return search, nil
}
