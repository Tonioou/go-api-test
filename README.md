# TODO List

That's a simples CRUD of a Todo card. The main purpose is to exercise somethings and also learn others.

## What you'll find here?
* Go 
* Docker
* Rest API
* Postgres
* OpenTelemetry
* Prometheus (soon..)
* Grafana (soon..)
* Kubernetes (soon..)
* Kong API Gateway (soon..)

## How to run in Docker

You must have docker and docker compose installed on your machine

To start the database and application run, if running in a linux with make installed run:

`make start` and to shut down server `make stop`

If you don't have make on your system then run from the root folder:

`docker compose up -f build/docker-compose.yaml` to start the database and application

and `docker compose down -f build/docker-compose.yaml` to stop.

## How to run in K8S with Kind

You'll also need to have docker installed on your machine.
Further information soon. 

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

## Observability

| Name       | Url                     |
|------------|-------------------------|
| Jaeger     | http://localhost:16686/ |
| Prometheus | http://localhost:9090/  |
| Grafana    | http://localhost:3000/  |


