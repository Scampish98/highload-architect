include .env

# Database migrations
DSN:="postgres://$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable&user=$(DB_USER)&password=$(DB_PASS)"
MIGRATE:=docker run --rm -v $(shell pwd)/database/migrations:/migrations --network host migrate/migrate:v4.18.3 -path=/migrations/ -database $(DSN)

GOIMPORTS:=docker run --rm -v $(shell pwd):/src -w /src cytopia/goimports

# Runs docker-compose if exists else docker compose
ifeq ($(shell [[ `which docker-compose` != "" ]] && echo true ),true)
	DOCKER_COMPOSE=docker-compose
else
	DOCKER_COMPOSE=docker compose
endif


.PHONY: imports
imports:
	$(GOIMPORTS) -local "highload-architect/" -w .
.PHONY: swag
swag:
	$(shell go env GOPATH)/bin/swag init -g internal/app/http-server/server.go
	$(shell go env GOPATH)/bin/swag fmt && gofmt -s -w .

.PHONY: run
run:
	$(DOCKER_COMPOSE) -f docker-compose.yaml up --detach --build --remove-orphans

.PHONY: run-monitoring
run-monitoring:
	$(DOCKER_COMPOSE) -f docker-compose.monitoring.yaml up --detach --build --remove-orphans

.PHONY: migrate
migrate: migrate-up ## alias for migrate-up

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@echo "Create migration \"$(name)\""
	@$(MIGRATE) create -ext sql -dir /migrations/ $(name)

.PHONY: migrate-up
migrate-up: ## run all new database migrations
	@echo "Running all new database migrations..."
	$(MIGRATE) -verbose up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) -verbose down 1

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-force
migrate-force:
	@$(MIGRATE) force $(v)
	@echo "Forced to version $(v)"

.PHONY: hadolint
hadolint:
	@docker run --rm -i hadolint/hadolint < ./build/Dockerfile