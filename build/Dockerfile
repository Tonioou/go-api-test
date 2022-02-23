FROM golang:1.17-alpine as builder

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

CMD ["/app/todo-list"]