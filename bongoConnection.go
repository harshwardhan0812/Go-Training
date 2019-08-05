package main

import (
	"github.com/go-bongo/bongo"
	"log"
)

var Connection *bongo.Connection
var err error
func CreateConnection() {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database: "awesomeDB",
	}
	Connection, err = bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}
}
