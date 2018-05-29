package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	rec := GetByKey("1")
	fmt.Printf("%s", rec)

	c.String(200, "hallo")
}

func main() {
	r := gin.New()
	r.GET("/", rootHandler)
	r.Run(":8000")
}
