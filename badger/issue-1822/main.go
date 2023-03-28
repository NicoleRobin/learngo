package main

import (
	"encoding/binary"
	"log"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/skl"
)

/*
问题原因：
在执行操作时没有判断返回值是否错误
查看源码发现在一个txn中只能最多执行10万左右次操作
*/
func main() {
	nodeSize := skl.MaxNodeSize
	log.Println(nodeSize)
	db, err := badger.Open(badger.DefaultOptions("testdb"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//11 times create 10000 entries and then remove them
	for i := 0; i < 11; i++ {
		log.Println("Round ", i+1)
		err := db.Update(func(txn *badger.Txn) error {
			seq, _ := db.GetSequence([]byte("abc"), 1000)
			b := make([]byte, 8)
			count := 0
			for j := 0; j < 10000; j++ {
				s, _ := seq.Next()
				binary.LittleEndian.PutUint64(b, s)
				key := make([]byte, 8)
				copy(key, b)
				err := txn.Set(key, []byte("Hasta la vista, baby!"))
				if err != nil {
					log.Printf("txn.Set() failed, err:%s", err)
					return err
				}
				count++
			}
			log.Printf("count:%d", count)
			return nil
		})
		if err != nil {
			log.Printf("db.Update() failed, err:%s", err)
			return
		}

	}
	err = db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		count := 0
		for it.Rewind(); it.Valid(); it.Next() {
			count++
			err := txn.Delete(it.Item().KeyCopy(nil))
			if err != nil {
				log.Printf("txn.Delete() failed, err:%s", err)
				return err
			}
		}
		log.Printf("count:%d", count)
		return nil
	})
	if err != nil {
		log.Printf("db.Update() failed, err:%s", err)
		return
	}

	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		ctr := 0
		for it.Rewind(); it.Valid(); it.Next() {
			ctr++
			//log.Println(string(it.Item().Key()))
		}
		log.Println(ctr, " keys remained in the DB.")
		return nil
	})
	if err != nil {
		log.Printf("db.View() failed, err:%s", err)
		return
	}

}
