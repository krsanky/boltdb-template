package db

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/krsanky/config"
)

var DB *bolt.DB

func Init() {
	var err error
	var db_file = config.Get("db_file")
	fmt.Printf("db_file:%s\n", db_file)
	DB, err = bolt.Open(db_file, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}
