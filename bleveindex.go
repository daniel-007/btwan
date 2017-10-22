package btwan

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/blevesearch/bleve"
)

//var indexMapping *mapping.IndexMappingImpl
var indexer bleve.Index
var _indexChan = make(chan *MetadataInfo, 10000)

func initIndex() error {
	indexMapping := bleve.NewIndexMapping()
	err := indexMapping.AddCustomTokenizer("sego",
		map[string]interface{}{
			"dictpath": workdir + "/dict/dictionary.txt",
			"type":     "sego",
		},
	)
	if err != nil {
		panic(err)
	}
	err = indexMapping.AddCustomAnalyzer("sego",
		map[string]interface{}{
			"type":      "sego",
			"tokenizer": "sego",
		},
	)
	if err != nil {
		panic(err)
	}
	indexMapping.DefaultAnalyzer = "sego"
	indexer, err = bleve.Open(workdir + "/index")

	if err != nil {
		indexer, err = bleve.New(workdir+"/index", indexMapping)
	}
	if err != nil {
		panic(err)
	}
	go loop()
	go sign()
	return nil
}

func bleveSearch(q string, from, size int) (*bleve.SearchResult, error) {
	_suggestChan <- q
	req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
	req.Highlight = bleve.NewHighlight()
	req.From = from
	req.Size = size
	return indexer.Search(req)
}
func loop() {
	for meta := range _indexChan {
		err := indexer.Index(strconv.FormatUint(meta.ID, 10), meta)
		if err != nil {
			info("index", len(_indexChan), meta, err)
		}
	}
}

func sign() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	s := <-c
	log.Println("退出信号", s)
	indexer.Close()
	dumpSuggest()
	os.Exit(0)
}
