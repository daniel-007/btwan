package main

import (
	"btwan"
	"log"
	"time"

	"google.golang.org/grpc"
)

var client btwan.OwstoniServiceClient

func _init() {
	var conn *grpc.ClientConn
	for {
		con, err := grpc.Dial(indexer, grpc.WithInsecure())
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		conn = con
		break
	}
	client = btwan.NewOwstoniServiceClient(conn)
}
