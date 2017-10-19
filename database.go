package btwan

import (
	"encoding/json"
	"strconv"
	"sync/atomic"

	"github.com/dgraph-io/badger"
)

var (
	_db          *badger.DB
	ihBucketName = []byte("infohash")
)

func initDB() error {
	file := workdir + "/infohash.db"
	opts := badger.DefaultOptions
	opts.Dir = file
	opts.ValueDir = file

	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	_db = db
	return nil
}

func getMetadata(id uint64) (t *MetadataInfo, err error) {
	err = _db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(strconv.FormatUint(id, 10)))
		if err != nil {
			return err
		}
		b, err := item.Value()
		if err != nil {
			return err
		}
		t = new(MetadataInfo)
		return t.fromBytes(b)
	})
	return
}

func findMetadata(ids []uint64) (tms []*MetadataInfo, err error) {
	tms = []*MetadataInfo{}
	err = _db.View(func(txn *badger.Txn) error {
		for _, id := range ids {
			item, err := txn.Get([]byte(strconv.FormatUint(id, 10)))
			if err != nil {
				return err
			}
			b, err := item.Value()
			if err != nil {
				return err
			}
			t := new(MetadataInfo)
			err = t.fromBytes(b)
			if err != nil {
				return err
			}
			tms = append(tms, t)
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
	return _db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(p.InfoHash))
		if err == badger.ErrKeyNotFound {
			p.ID = uint64(GenrateID())
			b, err := p.toBytes()
			if err != nil {
				return err
			}
			err = txn.Set([]byte(p.InfoHash), []byte(strconv.FormatUint(p.ID, 10)), 0)
			if err != nil {
				return err
			}
			return txn.Set([]byte(strconv.FormatUint(p.ID, 10)), b, 0)
		}
		if err != nil {
			return err
		}
		b, err := item.Value()
		if err != nil {
			return err
		}
		item, err = txn.Get(b)
		if err != nil {
			return err
		}
		bb, err := item.Value()
		if err != nil {
			return err
		}
		t := new(MetadataInfo)
		err = t.fromBytes(bb)
		if err != nil {
			return err
		}
		atomic.AddUint64(&p.Degree, 1)
		bb, err = t.toBytes()
		if err != nil {
			return err
		}
		return txn.Set(b, bb, 0)
	})
}
func (p *MetadataInfo) addDegree(id uint64, c uint16) error {
	return _db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(strconv.FormatUint(id, 10)))
		if err != nil {
			return err
		}
		b, err := item.Value()
		if err != nil {
			return err
		}
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
			return txn.Set([]byte(strconv.FormatUint(id, 10)), b, 0)
		}
		return nil
	})
}
