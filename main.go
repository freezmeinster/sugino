package main

import (
	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	rec := GetByKey(c, "1")
	c.JSON(200, rec)
}

func detailHandler(c *gin.Context) {
	key := c.Param("key")
	rec := GetByKey(c, key)
	c.JSON(200, rec.Bins)
}

func createHandler(c *gin.Context) {
	id := CreateData(c)
	c.JSON(200, gin.H{"document_id": id})
}

func main() {
	r := gin.New()
	r.Use(SetClient())
	r.GET("/", rootHandler)
	r.POST("/", createHandler)
	r.GET("/:key", detailHandler)
	r.Run(":8000")
}
