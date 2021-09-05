package main

import (
	prometheus_middleware "github.com/Tonioou/go-person-crud/internal/prometheus"
	"github.com/gin-gonic/gin"
)

func main() {
	ginApi := gin.Default()

	go ginApi.Run(":8080")
	ginMetrics := gin.Default()
	ginMetrics.GET("/metrics", prometheus_middleware.PrometheusHandler())
	ginMetrics.Run(":8081")

}
