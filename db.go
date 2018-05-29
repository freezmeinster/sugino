package main

import (
	"log"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
)

func SetClient() gin.HandlerFunc {
	client := GetClient()

	return func(c *gin.Context) {
		c.Set("client", *client)
		c.Next()
	}
}

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

func GetByKey(c *gin.Context, id string) *as.Record {
	key, err := as.NewKey("test", "buku", id)
	readPolicy := as.NewPolicy()
	client := c.MustGet("client").(as.Client)
	rec, err := client.Get(readPolicy, key)
	panicOnError(err)
	return rec
}
