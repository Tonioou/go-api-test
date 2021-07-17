package main

import (
	"github.com/Tonioou/go-api-test/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.NewPersonApi().Register(r)
	r.Run(":8080")
}
