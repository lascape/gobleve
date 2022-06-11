# govleve

原项目地址 [blevesearch/bleve](https://github.com/blevesearch/bleve)

**bleve** 是一个实现了高性能磁盘检索的go语言组件库
* 索引任何`go`数据结构
* 强大的配置支持
* 支持的字段类型：
  * 文本、数字、日期
* 支持的查询类型：
    * 术语、短语、匹配、匹配短语、前缀
    * 合取、析取、布尔
    * 数值范围、日期范围
    * 地理空间
    * 人工输入的简单查询语法
* tf-idf 评分
* 搜索结果匹配突出显示
* 支持聚合方面：
    * 术语方面
    * 数值范围分面
    * 日期范围方面

### 创建或者加载索引

```go
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
```

### 写入索引数据

> 如果是重复写入某些数据，需要删除之后重新写入

#### 单条写入索引

```go
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
```

#### 批量写入索引（推荐）

```go
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
```

### 读取索引数据
#### 检索文本索引数据
```go
func ReadIndex(index bleve.Index) (*bleve.SearchResult, error) {
	search, err := index.Search(bleve.NewSearchRequest(bleve.NewQueryStringQuery("bodyindex1")))
	if err != nil {
		return nil, err
	}
	return search, nil
}
```
#### 高亮检索短语索引数据
将结果高亮打印出来
```go
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
```