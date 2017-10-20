package main

import (
	"btwan"
	"log"

	_ "github.com/yanyiwu/gojieba/bleve"
)

func main() {
	log.Println(btwan.ListenAndServe())
}
