package main

import (
	"log"
	"os"

	"github.com/dgraph-io/badger"
)

func main() {
	opts := badger.DefaultOptions
	opts.Dir = os.Args[1]
	opts.ValueDir = os.Args[1]
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	db.View(func(txn *badger.Txn) error {
		return nil
	})
}

func search(key string, from, size int) {

}
