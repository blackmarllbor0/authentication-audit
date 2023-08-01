include .env

PROJECT_NAME = main
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

.PHONY:restore_db
restore_db:
	docker stop $(DOCKER_CONTAINER_NAME) && docker rm $(DOCKER_CONTAINER_NAME) && make up_db

.PHONY:psql
psql:
	docker exec -it $(DOCKER_CONTAINER_NAME) psql $(POSTGRES_DB) $(POSTGRES_USER)