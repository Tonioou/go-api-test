package main

import (
	"github.com/Tonioou/go-person-crud/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ginApi := gin.Default()
	api.NewPersonApi().Register(ginApi)

	go ginApi.Run(":8080")
	ginMetrics := gin.Default()
	ginMetrics.GET("/metrics", prometheusHandler())
	ginMetrics.Run(":8081")

}

func prometheusHandler() gin.HandlerFunc {
	promHandler := promhttp.Handler()
	return func(g *gin.Context) {
		promHandler.ServeHTTP(g.Writer, g.Request)
	}
}
