package main

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	index "github.com/blevesearch/bleve_index_api"
)

type Message struct {
	Id   int
	From string
	Body string
}

func init() {
	//open, err := bleve.New("example.bleve3", bleve.NewIndexMapping())
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//batch := open.NewBatch()
	//for i := 0; i < 1000; i++ {
	//	_ = batch.Index(strconv.Itoa(i), Message{
	//		From: "from" + strconv.Itoa(i),
	//		Body: "body" + strconv.Itoa(i),
	//	})
	//	if batch.Size() > 100 {
	//		_ = open.Batch(batch)
	//	}
	//}
}

func main() {
	index2, _ := bleve.Open("example.bleve3")
	query := bleve.NewMatchQuery("body3")
	query.SetField("From")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index2.Search(searchRequest)
	fmt.Println(searchResult)

	return
	document, _ := index2.Document("196310000")

	document.VisitFields(func(field index.Field) {
		fmt.Println(field.Value())
	})
	fmt.Println(document)
}
