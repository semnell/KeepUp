.DEFAULT_GOAL := help

# Variables
GO := go
GOFMT := gofmt
GOFMT_FLAGS := -s
BINARY_NAME := keepup

# Targets
.PHONY: all
all: fmt test build docker-build up ## Run fmt, tests, build, and docker-build targets

.PHONY: fmt
fmt: ## Format code
	$(GOFMT) $(GOFMT_FLAGS) -w .

.PHONY: test
test: ## Run tests
	$(GO) test ./...

.PHONY: build
build: ## Build the application
	$(GO) build -o $(BINARY_NAME) main.go

.PHONY: clean
clean: ## Clean up generated files
	$(GO) clean
	rm -f $(BINARY_NAME)

.PHONY: up
up: faktory-up ## Start Docker containers
	docker-compose up -d

.PHONY: down
down: ## Stop Docker containers
	docker-compose down

.PHONY: faktory-up
faktory-up: ## Start Faktory Docker container
	docker-compose up faktory -d
	echo "Waiting for FAKTORY to start..."
	sleep 5

.PHONY: faktory-down
faktory-down: ## Stop Faktory Docker container
	docker-compose stop faktory

.PHONY: revive
revive: ## Run revive linter
	revive -exclude vendor/... ./...

.PHONY: run
run: faktory-up ## Run the application
	$(GO) run main.go

.PHONY: run-server-standalone
run-server-standalone: faktory-up ## Run the server standalone
	$(GO) run main.go server

.PHONY: run-worker-standalone
run-worker-standalone: faktory-up ## Run the worker standalone
	$(GO) run main.go worker

.PHONY: docker-build
docker-build: ## Build Docker image
	docker build . -f docker/Dockerfile --target dev

.PHONY: coverage
coverage: generate-coverage ## Run tests with coverage

.PHONY: generate-coverage
generate-coverage: ## Generate code coverage
	go test -cover ./... -args -test.gocoverdir="${PWD}/coverage/unit"
	go tool covdata percent -i=./coverage/unit
	go tool covdata textfmt -i=./coverage/unit -o coverage/profile
	go tool cover -func coverage/profile
	rm coverage/unit/cov*

.PHONY: help
help: ## Display this help message
	@echo "Usage: make [target]"
	@echo "Targets:"
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		print "  " $$1 \
	} \
	/^## / { \
		print "        " substr($$0, 4) \
	}' $(MAKEFILE_LIST)
