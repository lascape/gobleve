package main

import (
	"strconv"

	"github.com/blevesearch/bleve/v2"
)

type Meta struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
	From string `json:"from"`
}

//WriteIndex 写入单条索引，如果索引已经存在则删除重建
func WriteIndex(index bleve.Index) error {
	for i := 0; i < 10; i++ {
		idStr := strconv.Itoa(i)
		_ = index.Delete(idStr)
		if err := index.Index(idStr, Meta{
			Id:   i,
			Body: "test bodyindex" + idStr,
			From: "test bodyindex" + idStr,
		}); err != nil {
			return err
		}
	}
	return nil
}

//WriteBatchIndex 批量写入索引数据
func WriteBatchIndex(index bleve.Index) error {
	batchDel := index.NewBatch()
	batchAdd := index.NewBatch()
	for i := 0; i < 10; i++ {
		idStr := strconv.Itoa(i)
		batchDel.Delete(idStr)
		if err := batchAdd.Index(idStr, Meta{
			Id:   i,
			Body: "test bodyindex" + idStr,
			From: "test bodyindex" + idStr,
		}); err != nil {
			return err
		}
	}
	_ = index.Batch(batchDel)
	return index.Batch(batchAdd)
}
