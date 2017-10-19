package controllers

import (
	"btwan"
	"fmt"
	"log"
	"time"

	"github.com/revel/revel"
	"google.golang.org/grpc"
)

var client btwan.OwstoniServiceClient

func init() {
	var conn *grpc.ClientConn
	for {
		con, err := grpc.Dial("api.btwan.net:7700", grpc.WithInsecure())
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		conn = con
		break
	}
	client = btwan.NewOwstoniServiceClient(conn)

	revel.TemplateFuncs["divide"] = func(a, b int64) int64 {
		return a / b
	}

	revel.TemplateFuncs["size"] = func(a uint64) string {
		if a < 1024 {
			return fmt.Sprintf("%v B", a)
		}
		if float64(a)/1024.0 < 1024.0 {
			return fmt.Sprintf("%.2f KB", float64(a)/1024.0)
		}
		if float64(a)/1024.0/1024.0 < 1024.0 {
			return fmt.Sprintf("%.2f MB", float64(a)/1024.0/1024.0)
		}
		if float64(a)/1024.0/1024.0/1024.0 < 1024.0 {
			return fmt.Sprintf("%.2f GB", float64(a)/1024.0/1024.0/1024.0)
		}
		if float64(a)/1024.0/1024.0/1024.0/1024.0 < 1024.0 {
			return fmt.Sprintf("%.2f TB", float64(a)/1024.0/1024.0/1024.0/1024.0)
		}
		if float64(a)/1024.0/1024.0/1024.0/1024.0/1024.0 < 1024.0 {
			return fmt.Sprintf("%.2f PB", float64(a)/1024.0/1024.0/1024.0/1024.0/1024.0)
		}
		if float64(a)/1024.0/1024.0/1024.0/1024.0/1024.0/1024.0 < 1024.0 {
			return fmt.Sprintf("%.2f EB", float64(a)/1024.0/1024.0/1024.0/1024.0/1024.0/1024.0)
		}
		return fmt.Sprintf("%v B", a)
	}
}
