package main

import (
	"log"

	as "github.com/aerospike/aerospike-client-go"
)

func GetClient() *as.Client {
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return client
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetByKey(id string) *as.Record {
	key, err := as.NewKey("test", "buku", id)
	readPolicy := as.NewPolicy()
	client := GetClient()
	rec, err := client.Get(readPolicy, key)
	panicOnError(err)
	return rec
}
