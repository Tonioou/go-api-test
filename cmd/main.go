package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Tonioou/go-person-crud/internal/api"
	"github.com/Tonioou/go-person-crud/internal/config"
	prometheus_middleware "github.com/Tonioou/go-person-crud/internal/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	config.NewLogger()
	gin.SetMode(gin.ReleaseMode)
	ginApi := gin.Default()

	go ginApi.Run(":8080")

	todoApi := api.NewTodoApi()
	todoApi.Register(ginApi)

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

func GetLogger() *logrus.Entry {
	entry := logrus.WithFields(
		logrus.Fields{
			"level":     config.GetConfig().LogLevel,
			"timestamp": time.Now(),
		},
	)
	return entry
}
