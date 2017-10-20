package btwan

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/blevesearch/bleve"
	"github.com/yanyiwu/gojieba"
)

//var indexMapping *mapping.IndexMappingImpl
var indexer bleve.Index
var _indexChan = make(chan *MetadataInfo, 10000)

func initIndex() error {
	indexMapping := bleve.NewIndexMapping()
	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"idf":          gojieba.IDF_PATH,
			"stop_words":   gojieba.STOP_WORDS_PATH,
			"type":         "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	indexMapping.DefaultAnalyzer = "gojieba"
	indexer, err = bleve.New(workdir+"/index", indexMapping)
	if err != nil {
		panic(err)
	}
	go loop()
	go sign()
	return nil
}

func bleveSearch(q string, from, size int) (*bleve.SearchResult, error) {
	req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
	req.Highlight = bleve.NewHighlight()
	req.From = from
	req.Size = size
	return indexer.Search(req)
}
func loop() {

	var count = 0
	for meta := range _indexChan {
		if count >= 1000 {
			count = 0
		}
		if meta != nil {
			err := indexer.Index(strconv.FormatUint(meta.ID, 10), meta)
			info("index", len(_indexChan), meta, err)
			count++
		}
	}
}

func sign() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	s := <-c
	log.Println("退出信号", s)
	indexer.Close()
}
