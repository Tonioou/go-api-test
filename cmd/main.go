package main

import (
	"fmt"
	"github.com/Tonioou/go-person-crud/internal/api"
	"github.com/Tonioou/go-person-crud/internal/config"
	prometheus_middleware "github.com/Tonioou/go-person-crud/internal/prometheus"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.NewLogger()
	gin.SetMode(gin.ReleaseMode)
	ginApi := gin.Default()

	go func() {
		err := ginApi.Run(":8080")
		config.Logger.Fatal(err.Error())
	}()

	todoApi := api.NewTodoApi()
	todoApi.Register(ginApi)

	ginMetrics := gin.Default()
	ginMetrics.GET("/metrics", prometheus_middleware.PrometheusHandler())
	go func() {
		err := ginMetrics.Run(":8081")
		config.Logger.Fatal(err.Error())
	}()

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
