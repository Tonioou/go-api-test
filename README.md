# TODO List

That's a todo crud, where i'm exercising some things and trying others
like using open telemetry for tracing, writing my one openapi3 file, and
exercising golang architecture.

## How to run

You must have docker and docker compose installed on your machine

To start the database and application run, if running in a linux with make installed run:

`make start` and to shut down server `make stop`

If you don't have make on your system then run from the root folder:

`docker compose up -f build/docker-compose.yaml` to start the database and application

and `docker compose down -f build/docker-compose.yaml` to stop.

## Tips

Before the following code it's necessary to configure the TextMapPropagator just like this:
```go
otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
```

Here's a reference link to link for [b3](https://opentelemetry.io/docs/instrumentation/go/manual/#propagators-and-context).

Now, if needed to pass the open telemetry headers (w3c,b3,..) you can use the following code, i'm using
resty to do the http call, but the relevant part is related to otel.


```go
    url := "your_url"
    client := resty.New()
	req := client.
		R().
		SetContext(ctx).
		EnableTrace()
	
	// the magic happens here
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := req.Get(url)

```

## Monitoring

| Name       | Url                     |
|------------|-------------------------|
| Jaeger     | http://localhost:16686/ |
| Prometheus | http://localhost:9090/  |
| Grafana    | http://localhost:3000/  |


