package main

import (
	"fmt"
	"github.com/Tonioou/go-todo-list/internal/config"
)

func main() {
	cfg := config.NewConfig()
	config.NewLogger()

	fmt.Println(cfg)
	// otel related
	//exp, err := newExporter(context.Background())
	//if err != nil {
	//	config.Logger.Fatal(err.Error())
	//}
	//tp := trace.NewTracerProvider(
	//	trace.WithBatcher(exp),
	//	trace.WithResource(newResource()),
	//	trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(5))),
	//)
	//
	//defer func() {
	//	if err := tp.Shutdown(context.Background()); err != nil {
	//		config.Logger.Fatal(err.Error())
	//	}
	//}()
	//otel.SetTracerProvider(tp)
	//otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	//// http server related
	//e := echo.New()
	//
	//v := validator.New(validator.WithRequiredStructEnabled())
	//
	//e.HideBanner = true
	//e.HidePort = true
	//e.Validator = model.NewCustomValidator(v)
	//e.Use(middleware.Gzip())
	//e.Use(otelecho.Middleware("go-todo-list"))
	//go func() {
	//	err := e.Start(":8080")
	//	config.Logger.Fatal(err.Error())
	//}()
	//
	//// dependency management
	//pgClient := client.NewPgClient(cfg)
	//todoRepository := repository.NewTodoRepository(pgClient)
	//todoService := service.NewTodoService(todoRepository)
	//todoApi := api.NewTodoApi(todoService)
	//
	//// register routes
	//todoApi.Register(e)
	//
	//// metrics server
	//metrics := echo.New()
	//metrics.HideBanner = true
	//metrics.HidePort = true
	//p := prometheus.NewPrometheus("echo", nil)
	//e.Use(p.HandlerFunc)
	//p.SetMetricsPath(metrics)
	//
	//go func() {
	//	err := metrics.Start(":8081")
	//	config.Logger.Fatal(err.Error())
	//}()
	//
	//sigs := make(chan os.Signal, 1)
	//
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//
	//done := make(chan bool, 1)
	//
	//go func() {
	//	sig := <-sigs
	//	fmt.Println(sig)
	//	done <- true
	//}()
	//
	//fmt.Println("awaiting signal")
	//<-done
	//fmt.Println("exiting")
}

// newExporter returns a console exporter.
//func newExporter(ctx context.Context) (trace.SpanExporter, error) {
//	c := otlptracegrpc.NewClient(otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(config.GetConfig().Otel.GrpcEndpoint))
//	return otlptrace.New(ctx, c)
//}
//
//// newResource returns a resource describing this application.
//func newResource() *resource.Resource {
//	r, _ := resource.Merge(
//		resource.Default(),
//		resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceNameKey.String(config.GetConfig().ServiceName),
//			semconv.ServiceVersionKey.String("v0.1.0"),
//			attribute.String("environment", "demo"),
//		),
//	)
//	return r
//}
