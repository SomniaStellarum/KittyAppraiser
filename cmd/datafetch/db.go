package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"encoding/binary"
	"github.com/gin-gonic/gin/json"
)

var db *bolt.DB
var sales = []byte("sales")

func init() {
	var err error
	db, err = bolt.Open("db/kittysales.db", 0644, nil)
	if err != nil {
		log.Panic("Error Opening Database: ", err, "\n")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(sales)
		return err
	})
}

func infoDb() {
	sts := db.Stats()
	fmt.Print(sts)
}

func insertSalesData(sls []salesData) error {
	err := db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(sales)

		for _, sl := range sls {
			num, err := bucket.NextSequence()
			if err != nil {
				return err
			}
			buf := make([]byte, binary.MaxVarintLen64)
			_ = binary.PutUvarint(buf, num)

			b, err := json.Marshal(sl)
			if err != nil {
				return err
			}

			err = bucket.Put(buf, b)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func printBucketStats() {
	_ = db.View(func(tx *bolt.Tx)error{
		bucket := tx.Bucket(sales)
		bstats := bucket.Stats()
		fmt.Print("Number of Entries in Bucket: ", bstats.KeyN, "\n")
		return nil
	})
}
