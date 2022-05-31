package main

import (
	"fmt"
	"github.com/Tonioou/go-todo-list/internal/api"
	"github.com/Tonioou/go-todo-list/internal/config"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.NewLogger()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	go func() {
		err := e.Start(":8080")
		config.Logger.Fatal(err.Error())
	}()

	todoApi := api.NewTodoApi()
	todoApi.Register(e)

	metrics := echo.New()
	metrics.HideBanner = true
	metrics.HidePort = true
	p := prometheus.NewPrometheus("echo", nil)
	e.Use(p.HandlerFunc)
	p.SetMetricsPath(metrics)

	go func() {
		err := metrics.Start(":8081")
		config.Logger.Fatal(err.Error())
	}()

	sigs := make(chan os.Signal)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
