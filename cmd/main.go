package main

import (
	"context"
	"fmt"
	"github.com/Tonioou/go-todo-list/internal/api"
	"github.com/Tonioou/go-todo-list/internal/config"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.NewLogger()

	exp, err := newExporter(context.Background())
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(5))),
	)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			config.Logger.Fatal(err.Error())
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(otelecho.Middleware("go-todo-list"))
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

	sigs := make(chan os.Signal, 1)

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

// newExporter returns a console exporter.
func newExporter(ctx context.Context) (trace.SpanExporter, error) {
	client := otlptracegrpc.NewClient(otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(config.GetConfig().Otel.GrpcEndpoint))
	return otlptrace.New(ctx, client)
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.GetConfig().ServiceName),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}
