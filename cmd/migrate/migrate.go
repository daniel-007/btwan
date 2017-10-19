package main

import (
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
)

func main() {
	initDB()
	opts := badger.DefaultOptions
	opts.Dir = os.Args[1]
	opts.ValueDir = os.Args[1]
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	err = _db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("infohash"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			return db.Update(func(txn *badger.Txn) error {
				return txn.Set(k, v, 0)
			})
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	_db          *bolt.DB
	ihBucketName = []byte("infohash")
)

func initDB() error {
	file := os.Args[0]
	db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 3600 * time.Second})
	if err != nil {
		return err
	}
	_db = db
	return _db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(ihBucketName)
		return err
	})
}
