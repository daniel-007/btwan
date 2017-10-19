package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"btwan"

	"github.com/shiyanhui/dht"
	"google.golang.org/grpc"
)

var client btwan.OwstoniServiceClient

func main() {
	conn, err := grpc.Dial("api.btwan.net:7700", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client = btwan.NewOwstoniServiceClient(conn)

	w := dht.NewWire(65536, 1024, 256)
	go func() {
		for resp := range w.Response() {
			metadata, err := dht.Decode(resp.MetadataInfo)
			if err != nil {
				continue
			}
			log.Println("元数据", metadata)
			info := metadata.(map[string]interface{})
			if _, ok := info["name"]; !ok {
				continue
			}
			mi := btwan.MetadataInfo{
				InfoHash:    hex.EncodeToString(resp.InfoHash),
				Name:        fmt.Sprintf("%v", info["name"]),
				Degree:      1,
				CollectTime: time.Now().Unix(),
			}

			if v, ok := info["files"]; ok {
				if files, ok := v.([]interface{}); ok {
					mi.Files = []*btwan.FileInfo{}
					for _, item := range files {
						if f, ok := item.(map[string]interface{}); ok {
							path := []string{}
							if p, ok := f["path"].([]interface{}); ok {
								for _, pp := range p {
									if vv, ok := pp.(string); ok {
										path = append(path, vv)
									}
								}
								fi := btwan.FileInfo{
									Path:   path,
									Length: uint64(f["length"].(int)),
								}
								mi.Files = append(mi.Files, &fi)
							}
						}
					}
				}

			} else if _, ok := info["length"]; ok {
				if length, ok := info["length"].(int); ok {
					mi.Length = uint64(length)
				}
			}
			if mi.Length > 0 && strings.TrimSpace(mi.Name) != "" {
				client.Index(context.Background(), &mi)
			}
		}
	}()
	go w.Run()

	config := dht.NewCrawlConfig()
	config.Address = ":6881"
	config.PrimeNodes = []string{"router.bittorrent.com:6881",
		"dht.transmissionbt.com:6881"}
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		log.Println("接收到一条Infohash", infoHash, ip, port)
		w.Request([]byte(infoHash), ip, port)
	}
	d := dht.New(config)
	recv, err := client.Recv(context.Background(), &btwan.Void{})
	go func() {
		for {
			event, err := recv.Recv()
			if err == io.EOF {
				time.Sleep(5 * time.Second)
				continue
			}
			ih := event.Attributes["infohash"]
			peers, err := d.GetPeers(ih)
			for _, peer := range peers {
				w.Request([]byte(ih), peer.IP.String(), peer.Port)
			}
		}
	}()
	d.Run()
}
