package main

import (
	"bytes"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"log"
	"reflect"
	"unsafe"
)

type Slice struct {
	Array unsafe.Pointer
	Len   int
	Cap   int
}

/*
原因：
txn中的item会被重用，因此每次获取到的item值应该先拷贝保存下来后续使用
*/
func main() {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		panic(err)
	}
	n := 200

	// Write 1,000 keys
	err = db.Update(func(txn *badger.Txn) error {
		for i := 0; i < n; i++ {
			err = txn.Set([]byte(fmt.Sprintf("%v", i)), bytes.Repeat([]byte{0}, 1024))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// Check the number of keys
	if getKeyCount(db) != n {
		panic("expected 200 elements")
	}

	// Delete all the elements
	var keys [][]byte
	err = db.Update(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek([]byte{}); it.ValidForPrefix([]byte{}); it.Next() {
			key := it.Item().Key()
			// keys = append(keys, it.Item().KeyCopy(nil))
			keys = append(keys, key)
			// fmt.Println(string(it.Item().Key()))
			sh := (*reflect.SliceHeader)(unsafe.Pointer(&key))
			log.Printf("%+v, key:%s", sh, key)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	containsDups(keys)

	err = db.Update(func(txn *badger.Txn) error {
		for _, k := range keys {
			err = txn.Delete(k)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// Check again
	if i := getKeyCount(db); i != 0 {
		panic(fmt.Sprintf("expected 0 elements, got %v", i))
	}
}

func getKeyCount(db *badger.DB) (i int) {
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek([]byte{}); it.ValidForPrefix([]byte{}); it.Next() {
			i++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

func containsDups(in [][]byte) {
	for i, k := range in {
		for j, v := range in {
			if bytes.Equal(k, v) && i != j {
				shK := (*reflect.SliceHeader)(unsafe.Pointer(&k))
				shV := (*reflect.SliceHeader)(unsafe.Pointer(&v))
				log.Printf("found dup %v=%v, k:%s, v:%s, shK:%+v, shV:%+v", i, j, k, v, shK, shV)
			}
		}
	}
}
