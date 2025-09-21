# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=financial-aggregator-api
BINARY_UNIX=$(BINARY_NAME)_unix
DOCKER_IMAGE=financial-aggregator-api
DOCKER_TAG=latest

# Build the application
.PHONY: build
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./...

# Build for Linux (useful for Docker)
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_UNIX) -v ./...

# Run the application
.PHONY: run
run:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./...
	./bin/$(BINARY_NAME)

# Run in development mode (with hot reload using air)
.PHONY: dev
dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not installed. Installing..."; \
		$(GOGET) -u github.com/cosmtrek/air; \
		air; \
	fi

# Run tests
.PHONY: test
test:
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Run tests with race detection
.PHONY: test-race
test-race:
	$(GOTEST) -v -race ./...

# Run benchmarks
.PHONY: benchmark
benchmark:
	$(GOTEST) -v -bench=. ./...

# Clean build artifacts
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f bin/$(BINARY_UNIX)
	rm -f coverage.out coverage.html

# Download dependencies
.PHONY: deps
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Update dependencies
.PHONY: deps-update
deps-update:
	$(GOMOD) tidy
	$(GOMOD) download

# Run linter
.PHONY: lint
lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Installing..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2; \
		golangci-lint run; \
	fi

# Format code
.PHONY: fmt
fmt:
	$(GOCMD) fmt ./...

# Vet code
.PHONY: vet
vet:
	$(GOCMD) vet ./...

# Run all checks (test, lint, vet, fmt)
.PHONY: check
check: fmt vet lint test

# Docker commands
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: docker-clean
docker-clean:
	docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG) || true

# Development setup
.PHONY: setup
setup: deps
	@echo "Setting up development environment..."
	@if ! command -v air > /dev/null; then \
		echo "Installing air for hot reload..."; \
		$(GOGET) -u github.com/cosmtrek/air; \
	fi
	@if ! command -v golangci-lint > /dev/null; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2; \
	fi
	@echo "Development environment ready!"

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  build-linux   - Build for Linux (useful for Docker)"
	@echo "  run           - Build and run the application"
	@echo "  dev           - Run in development mode with hot reload"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  test-race     - Run tests with race detection"
	@echo "  benchmark     - Run benchmarks"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Download dependencies"
	@echo "  deps-update   - Update dependencies"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  vet           - Vet code"
	@echo "  check         - Run all checks (fmt, vet, lint, test)"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container"
	@echo "  docker-clean  - Clean Docker image"
	@echo "  setup         - Set up development environment"
	@echo "  help          - Show this help message"
