package main

import (
	"log"
	"os"
	"strconv"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/keyword"
	"github.com/yanyiwu/gojieba"
	_ "github.com/yanyiwu/gojieba/bleve"
)

var (
	idx bleve.Index
)

func main1() {
	err := initIndex()
	if err != nil {
		log.Fatalln(err)
	}
	from, _ := strconv.Atoi(os.Args[3])
	size, _ := strconv.Atoi(os.Args[4])
	search(os.Args[2], from, size)
}

func search(key string, from, size int) {
	var r bleve.SearchRequest
	r.Fields = []string{"ID", "Name", "Files", "Size", "CreateDate", "Body"}
	r.From = from
	r.Size = size

	r.Query = bleve.NewMatchQuery(key)
	//r.SortBy([]string{"Size", "CreateDate"})
	result, _ := idx.Search(&r)
	log.Println(result.String())
}

func initIndex() error {
	file := os.Args[1]
	keyField := bleve.NewTextFieldMapping()
	keyField.Analyzer = keyword.Name
	notAnalyzerField := bleve.NewDocumentDisabledMapping()

	numberField := bleve.NewNumericFieldMapping()

	txtField := bleve.NewTextFieldMapping()
	txtField.Analyzer = "gojieba"

	docMapping := bleve.NewDocumentMapping()
	docMapping.AddFieldMappingsAt("Id", keyField)
	docMapping.AddFieldMappingsAt("Name", txtField)
	docMapping.AddFieldMappingsAt("CreateDate", numberField)
	docMapping.AddFieldMappingsAt("Size", numberField)
	docMapping.AddFieldMappingsAt("Files", numberField)
	docMapping.AddSubDocumentMapping("Body", notAnalyzerField)

	mapping := bleve.NewIndexMapping()
	mapping.AddDocumentMapping("infohash", docMapping)
	err := mapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"type":         "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	err = mapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	mapping.DefaultAnalyzer = "gojieba"
	index, err := bleve.Open(file)
	if err == bleve.ErrorIndexPathDoesNotExist {
		index, err = bleve.New(file, mapping)
	}

	idx = index
	return err
}
