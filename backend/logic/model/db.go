package model

import (
	"github.com/boltdb/bolt"
	"log"
	"sync"
	"time"
)

var dbonce sync.Once
var dbIns *bolt.DB

const boltPath = "goim.db"

func getDb() *bolt.DB {
	dbonce.Do(func() {
		db, err := bolt.Open(boltPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			log.Panic(err)
		}
		dbIns = db
	})
	return dbIns
}

// WriteToDB write key-value to boltdb
func WriteToDB(bucketName string, key string, value []byte) error {
	return getDb().Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return bucket.Put([]byte(key), value)
	})
}

// ReadFromDB read value from boltdb by key
func ReadFromDB(bucketName string, key string) ([]byte, error) {
	tx, err := getDb().Begin(true)
	if err != nil {
		return nil, err
	}
	bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
	if err != nil {
		return nil, err
	}
	return bucket.Get([]byte(key)), tx.Commit()
}
