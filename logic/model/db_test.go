package model

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
	"go-IM/pkg/tinyid"
	"testing"
)

func Test_KVWriteAndRead(t *testing.T) {
	expect := "mars"
	if err := WriteToDB("test", "name", []byte(expect)); err != nil {
		panic(err)
	}
	value, err := ReadFromDB("test", "name")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, expect, string(value))
}

func Test_ValueNotFound(t *testing.T) {
	value, err := ReadFromDB("test", "name")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 0, len(value))
}

func Test_Cursor(t *testing.T) {
	_ = WriteToDB("MyBucket", fmt.Sprint(tinyid.NextId()), []byte("mars"))
	_ = WriteToDB("MyBucket", fmt.Sprint(tinyid.NextId()), []byte("wuhan"))
	_ = WriteToDB("MyBucket", fmt.Sprint(tinyid.NextId()), []byte("tx"))

	if err := getDb().View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("MyBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
