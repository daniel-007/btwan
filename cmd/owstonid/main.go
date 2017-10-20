package main

import (
	"btwan"
	"log"

	_ "github.com/julianshen/blevesego"
)

func main() {
	log.Println(btwan.ListenAndServe())
}
