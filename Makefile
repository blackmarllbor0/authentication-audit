include .env

PROJECT_NAME = apteka_aprel
PROJECT_PATH = cmd/$(PROJECT_NAME).go

.PHONY:run
run:
	go run $(PROJECT_PATH)

.PHONY:build
build:
	go build -o bin/$(PROGRAM_NAME) $(PROJECT_PATH)

.PHONY:test
test:
	go test ./...

.PHONY:lint
lint:
	golangci-lint run

.PHONY: up_db
up_db:
	docker compose up -d

psql:
	docker exec -it postgres_db psql $(POSTGRES_DB) $(POSTGRES_USER)