FROM golang:1.21-alpine as builder

WORKDIR /build

EXPOSE 8080

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./ 

RUN CGO_ENABLED=0 GOOS=linux go build  -a -installsuffix cgo -o  todo-list ./cmd/.

FROM alpine:latest
WORKDIR /app 
COPY --from=builder /build/todo-list .
COPY /internal/config/config.yaml internal/config/
CMD ["/app/todo-list"]