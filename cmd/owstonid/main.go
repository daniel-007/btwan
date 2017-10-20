package main

import (
	"log"

	"btwan"

	_ "github.com/yanyiwu/gojieba/bleve"
)

func main() {
	log.Println(btwan.ListenAndServe())
}
