.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY: start
start:
	docker compose -f build/docker-compose.yaml up

.PHONY: stop
stop:
	docker compose -f build/docker-compose.yaml down
	docker rmi $(docker images -q --filter "reference=*todo*" )