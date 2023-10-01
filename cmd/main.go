package main

import (
	"context"
	"fmt"
	"github.com/Tonioou/go-todo-list/internal/model"
	logger "github.com/Tonioou/go-todo-list/pkg"
	"github.com/go-playground/validator/v10"
	slogecho "github.com/samber/slog-echo"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Tonioou/go-todo-list/internal/api"
	"github.com/Tonioou/go-todo-list/internal/client"
	"github.com/Tonioou/go-todo-list/internal/config"
	"github.com/Tonioou/go-todo-list/internal/repository"
	"github.com/Tonioou/go-todo-list/internal/service"
	"github.com/labstack/echo-contrib/echoprometheus"
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
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	// open telemetry
	exp, err := newExporter(ctx, cfg)
	if err != nil {
		logger.Fatal("failed to create exporter")
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource(cfg)),
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(5))),
	)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Fatal("failed to create exporter", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{},
			propagation.Baggage{}),
	)

	// validator
	v := validator.New(validator.WithRequiredStructEnabled())

	// http server related
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Validator = model.NewCustomValidator(v)
	e.Use(middleware.Gzip())
	e.Use(otelecho.Middleware(cfg.Service.Name))
	e.Use(slogecho.NewWithFilters(
		logger.Logger(),
		slogecho.IgnoreStatus(401, 404),
	))
	e.Use(echoprometheus.NewMiddleware(strings.Replace(cfg.Service.Name, "-", "_", -1)))

	go func() {
		err := e.Start(":8080")
		logger.Fatal("failed to create exporter", err)
	}()

	// client
	pgRWClient := client.NewPgClient(ctx, &cfg.Postgres.RW)

	// repository
	todoRepository := repository.NewTodoRepository(pgRWClient)

	// service
	todoService := service.NewTodoService(todoRepository)

	// routes
	todoApi := api.NewTodoApi(todoService)

	// register routes
	todoApi.Register(e)

	// metrics server
	metrics := echo.New()
	metrics.HideBanner = true
	metrics.HidePort = true
	metrics.GET("/metrics", echoprometheus.NewHandler())

	go func() {
		err := metrics.Start(":8081")
		logger.Fatal("failed to create exporter", err)
	}()

	// graceful shutdown
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
func newExporter(ctx context.Context, cfg *config.Configs) (trace.SpanExporter, error) {
	c := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.Otel.Exporter.GRPC.Endpoint),
	)
	return otlptrace.New(ctx, c)
}

// newResource returns a resource describing this application.
func newResource(cfg *config.Configs) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Service.Name),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", cfg.Service.Env),
		),
	)
	return r
}
