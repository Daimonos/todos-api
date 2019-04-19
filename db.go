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
}
