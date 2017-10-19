package btwan

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/shiyanhui/dht"
)

var (
	d    *dht.DHT
	down *dht.Wire
)

func initDown() {
	config := dht.NewCrawlConfig()
	config.Address = ":4000"
	down = dht.NewWire(65535, 64, 16)
	d = dht.New(config)
	go d.Run()
	go down.Run()
	go func() {
		for resp := range down.Response() {
			metadata, err := dht.Decode(resp.MetadataInfo)
			if err != nil {
				continue
			}
			info := metadata.(map[string]interface{})
			if _, ok := info["name"]; !ok {
				continue
			}
			mi := MetadataInfo{
				InfoHash:    hex.EncodeToString(resp.InfoHash),
				Name:        info["name"].(string),
				Degree:      1,
				CollectTime: time.Now().Unix(),
			}

			if v, ok := info["files"]; ok {
				files := v.([]interface{})
				mi.Files = []*FileInfo{}
				for _, item := range files {
					f := item.(map[string]interface{})
					p := f["path"].([]interface{})
					path := []string{}
					for _, pp := range p {
						path = append(path, pp.(string))
					}
					fi := FileInfo{
						Path:   path,
						Length: uint64(f["length"].(int)),
					}
					mi.Files = append(mi.Files, &fi)
				}
			} else if _, ok := info["length"]; ok {
				mi.Length = uint64(info["length"].(int))
			}
			_indexChan<-&mi
		}
	}()
}

func downMetainfo(ih string) {
	var i = 0
	for {
		if i >= 5 {
			return
		}
		peers, err := d.GetPeers(ih)
		i++
		if err != nil {
			time.Sleep(time.Second * 2)
			continue
		}
		fmt.Println("Found peers:", peers)
		for _, peer := range peers {
			down.Request([]byte(ih), peer.IP.String(), peer.Port)
		}
		return
	}
}
