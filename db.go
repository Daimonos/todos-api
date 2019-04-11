package main

import (
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

// InitDB initializes botldb
func InitDB(filepath string) {
	var err error
	db, err = bolt.Open(filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("todos"))
		if err != nil {
			return err
		}
		return nil
	})

}
