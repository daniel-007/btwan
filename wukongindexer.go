package btwan

import (
	"encoding/gob"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

var (
	_index     = engine.Engine{}
	_indexChan = make(chan *MetadataInfo, 128)
)

func close() {
	flushIndex()
	closeIndexer()
}
func closeIndexer() {
	_index.Close()
}

func _initIndexer() error {
	info("_init_indexer")
	gob.Register(MetaScoringFields{})
	_index.Init(types.EngineInitOptions{
		SegmenterDictionaries: dicts(),
		StopTokenFile:         workdir + "/dict/stop_tokens.txt",
		IndexerInitOptions: &types.IndexerInitOptions{
			//IndexType: types.LocationsIndex,
			IndexType: types.FrequenciesIndex,
		},
		NumShards:               8,
		UsePersistentStorage:    true,
		PersistentStorageFolder: workdir + "/index",
		PersistentStorageShards: 8,
	})
	go doIndex()
	return nil
}

func dicts() string {
	dict := []string{}
	filepath.Walk(workdir+"/dict", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.ToLower(filepath.Ext(info.Name())) == ".txt" && info.Name() != "stop_tokens.txt" {
			dict = append(dict, path)
		}
		return nil
	})
	return strings.Join(dict, ",")
}

func flushIndex() error {
	_index.FlushIndex()
	return nil
}

func doIndex() {
	var count = 0
	for {
		if count >= 128 {
			flushIndex()
			count = 0
		}
		meta := <-_indexChan
		info("index", meta)
		if meta != nil {
			_index.IndexDocument(meta.ID, types.DocumentIndexData{
				Content: meta.Name,
				Fields: MetaScoringFields{
					InfoHash:    meta.InfoHash,
					Name:        meta.Name,
					Files:       int64(len(meta.Files)),
					Size:        int64(meta.Length),
					Reviews:     meta.Reviews,
					Thumbs:      meta.Thumbs,
					Degree:      meta.Degree,
					Seeders:     meta.Seeders,
					Downloaders: meta.Downloaders,
					CollectTime: meta.CollectTime,
				},
			}, false)
		}
		count++
	}
}

type MetaScoringFields struct {
	InfoHash    string
	Name        string
	Tags        string
	Taxonomy    string
	Files       int64
	Size        int64
	Reviews     uint64 //评论数
	Follows     uint64 //关注数
	Thumbs      uint64 //赞
	Degree      uint64 //热度
	Seeders     uint64 //种子数
	Downloaders uint64 //下载数
	CollectTime int64
}

type MetaScoringCriteria struct {
}

const (
	SecondsInADay     = 86400
	MaxTokenProximity = 2
)

func (criteria MetaScoringCriteria) Score(
	doc types.IndexedDocument, fields interface{}) []float32 {
	if reflect.TypeOf(fields) != reflect.TypeOf(MetaScoringFields{}) {
		return []float32{}
	}
	wsf := fields.(MetaScoringFields)
	output := make([]float32, 9)
	if doc.TokenProximity > MaxTokenProximity {
		output[0] = 1.0 / float32(doc.TokenProximity)
	} else {
		output[0] = 1.0
	}

	output[1] = float32(doc.BM25 * (1 + float32(wsf.Degree)/10000))
	output[2] = float32(doc.BM25 * (1 + float32(wsf.Seeders)/10000))
	output[3] = float32(doc.BM25 * (1 + float32(wsf.Downloaders)/10000))
	output[4] = float32(doc.BM25 * (1 + float32(wsf.Thumbs)/10000))
	output[5] = float32(doc.BM25 * (1 + float32(wsf.Follows)/10000))
	output[6] = float32(doc.BM25 * (1 + float32(wsf.Reviews)/10000))
	output[7] = float32(doc.BM25 * (1 + float32(wsf.Size)/10000))
	output[8] = float32(wsf.CollectTime / (SecondsInADay * 3))
	return output
}

func removeIndex(id uint64) {
	_index.RemoveDocument(id, false)
	// flush_index()
}
func searchIndex(query string, offset, limit int) types.SearchResponse {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}

	return _index.Search(types.SearchRequest{
		Text: query,
		RankOptions: &types.RankOptions{
			ScoringCriteria: &MetaScoringCriteria{},
			OutputOffset:    offset,
			MaxOutputs:      limit,
		},
	})

}
