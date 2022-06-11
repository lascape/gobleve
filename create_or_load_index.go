package main

import (
	"github.com/blevesearch/bleve/v2"
)

func CreateOrLoadIndex() (bleve.Index, error) {
	open, err := bleve.Open("example.bleve")
	if err != nil {
		if err != bleve.ErrorIndexPathDoesNotExist {
			return nil, err
		}
		open, err = bleve.New("example.bleve", bleve.NewIndexMapping())
		if err != nil {
			return nil, err
		}
	}
	return open, nil
}
