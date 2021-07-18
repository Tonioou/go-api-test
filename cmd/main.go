package main

import (
	"github.com/Tonioou/go-api-test/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.NewPersonApi().Register(r)
	r.Run(":8080")
}
