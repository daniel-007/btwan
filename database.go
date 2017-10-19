package btwan

import (
	"encoding/json"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/boltdb/bolt"
)

var (
	_db          *bolt.DB
	ihBucketName = []byte("infohash")
)

func initDB() error {
	file := workdir + "/infohash.db"
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

func getMetadata(id uint64) (t *MetadataInfo, err error) {
	err = _db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ihBucketName).Get([]byte(strconv.FormatUint(id, 10)))
		t = new(MetadataInfo)
		return t.fromBytes(b)
	})
	return
}

func findMetadata(ids []uint64) (tms []*MetadataInfo, err error) {
	tms = []*MetadataInfo{}
	err = _db.View(func(tx *bolt.Tx) error {
		for _, id := range ids {
			b := tx.Bucket(ihBucketName).Get([]byte(strconv.FormatUint(id, 10)))
			t := new(MetadataInfo)
			err = t.fromBytes(b)
			if err == nil {
				tms = append(tms, t)
			}
		}
		return nil
	})
	return
}

func (p *MetadataInfo) jsonString() string {
	b, _ := json.MarshalIndent(p, "", "  ")
	return string(b)
}

func (p *MetadataInfo) toBytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *MetadataInfo) fromBytes(b []byte) error {
	return json.Unmarshal(b, p)
}
func (p *MetadataInfo) save() error {
	return _db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ihBucketName).Get([]byte(p.InfoHash))
		if len(b) > 0 {
			bb := tx.Bucket(ihBucketName).Get(b)
			t := new(MetadataInfo)
			err := t.fromBytes(bb)
			if err != nil {
				return err
			}
			atomic.AddUint64(&p.Degree, 1)
			bb, err = t.toBytes()
			if err != nil {
				return err
			}
			return tx.Bucket(ihBucketName).Put(b, bb)
		}
		p.ID = uint64(GenrateID())
		b, err := p.toBytes()
		if err != nil {
			return err
		}
		err = tx.Bucket(ihBucketName).Put([]byte(p.InfoHash), []byte(strconv.FormatUint(p.ID, 10)))
		if err != nil {
			return err
		}
		return tx.Bucket(ihBucketName).Put([]byte(strconv.FormatUint(p.ID, 10)), b)

	})
}
func (p *MetadataInfo) addDegree(id uint64, c uint16) error {
	return _db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ihBucketName).Get([]byte(strconv.FormatUint(id, 10)))
		if b != nil && len(b) > 0 {
			t := new(MetadataInfo)
			err := t.fromBytes(b)
			if err != nil {
				return err
			}
			atomic.AddUint64(&p.Degree, uint64(c))
			b, err = t.toBytes()
			if err != nil {
				return err
			}
			return tx.Bucket(ihBucketName).Put([]byte(strconv.FormatUint(id, 10)), b)
		}
		return nil
	})
}
