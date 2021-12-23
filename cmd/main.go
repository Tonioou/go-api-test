package main

import (
	"os"

	prometheus_middleware "github.com/Tonioou/go-person-crud/internal/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	configureLog()
	ginApi := gin.Default()

	go ginApi.Run(":8080")
	ginMetrics := gin.Default()
	ginMetrics.GET("/metrics", prometheus_middleware.PrometheusHandler())
	go ginMetrics.Run(":8081")

}

func configureLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}
