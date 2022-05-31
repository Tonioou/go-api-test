.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY: compose-up
compose-up:
	docker compose -f build/docker-compose.yaml up -d

.PHONY: compose-down
compose-down:
	docker compose -f build/docker-compose.yaml down