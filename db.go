package main

import (
	"log"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Buku test
type Buku struct {
	Title string `json:"title"`
	Price uint32 `json:"price"`
}

// SetClient Middleware untuk menambahkan client
func SetClient() gin.HandlerFunc {
	client := GetClient()

	return func(c *gin.Context) {
		c.Set("client", *client)
		c.Next()
	}
}

// GetClient fungsi untuk membuat client baru
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

// GetByKey digunakan untuk melakukan query ke aerospike berdasarkan key
func GetByKey(c *gin.Context, id string) *as.Record {
	key, err := as.NewKey("test", "buku", id)
	readPolicy := as.NewPolicy()
	client := c.MustGet("client").(as.Client)
	rec, err := client.Get(readPolicy, key)
	panicOnError(err)
	return rec
}

// CreateData digunakan untuk membuat data ke dalam aerospike
func CreateData(c *gin.Context) string {
	id := uuid.Must(uuid.NewV4()).String()
	key, err := as.NewKey("test", "buku", id)
	client := c.MustGet("client").(as.Client)
	writePolicy := as.NewWritePolicy(0, 0)
	var buku Buku
	c.BindJSON(&buku)
	bin := as.BinMap{
		"title": buku.Title,
		"price": buku.Price,
	}
	err = client.Put(writePolicy, key, bin)
	panicOnError(err)
	return id
}
