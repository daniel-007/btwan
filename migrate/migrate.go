package btwan

import (
	"btwan"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
)

func main() {
	initDB()
	opts := badger.DefaultOptions
	opts.Dir = "infohash"
	opts.ValueDir = "infohash"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	_db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("infohash"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return db.Update(func(txn *badger.Txn) error {
				return txn.Set(k, v, 0)
			})
		}

		return nil
	})
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

func getMetadata(id uint64) (t *btwan.MetadataInfo, err error) {
	err = _db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ihBucketName).Get([]byte(strconv.FormatUint(id, 10)))
		t = new(btwan.MetadataInfo)
		return json.Unmarshal(b, t)
	})
	return
}

// func (p *btwan.MetadataInfo) save() error {
// 	return _db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket(ihBucketName).Get([]byte(p.InfoHash))
// 		if len(b) > 0 {
// 			bb := tx.Bucket(ihBucketName).Get(b)
// 			t := new(btwan.MetadataInfo)
// 			err := t.fromBytes(bb)
// 			if err != nil {
// 				return err
// 			}
// 			atomic.AddUint64(&p.Degree, 1)
// 			bb, err = t.toBytes()
// 			if err != nil {
// 				return err
// 			}
// 			return tx.Bucket(ihBucketName).Put(b, bb)
// 		}
// 		p.ID = uint64(GenrateID())
// 		b, err := p.toBytes()
// 		if err != nil {
// 			return err
// 		}
// 		err = tx.Bucket(ihBucketName).Put([]byte(p.InfoHash), []byte(strconv.FormatUint(p.ID, 10)))
// 		if err != nil {
// 			return err
// 		}
// 		return tx.Bucket(ihBucketName).Put([]byte(strconv.FormatUint(p.ID, 10)), b)

// 	})
// }
// func (p *btwan.MetadataInfo) addDegree(id uint64, c uint16) error {
// 	return _db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket(ihBucketName).Get([]byte(strconv.FormatUint(id, 10)))
// 		if b != nil && len(b) > 0 {
// 			t := new(btwan.MetadataInfo)
// 			err := t.fromBytes(b)
// 			if err != nil {
// 				return err
// 			}
// 			atomic.AddUint64(&p.Degree, uint64(c))
// 			b, err = t.toBytes()
// 			if err != nil {
// 				return err
// 			}
// 			return tx.Bucket(ihBucketName).Put([]byte(strconv.FormatUint(id, 10)), b)
// 		}
// 		return nil
// 	})
// }
