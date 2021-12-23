package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	sigs := make(chan os.Signal)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

func configureLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}
