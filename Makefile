# GoVault Makefile

# Build settings
BINARY_NAME=govault
MAIN_FILE=main.go
GO_FLAGS=-ldflags="-s -w"

# Phony targets
.PHONY: all build run run-dev test clean deps fmt lint install release help

all: build

## Build: Build the application binary
build:
	go build $(GO_FLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

## Run-Dev: Run the application directly without build (e.g. make run CMD="add github")
run:
	go run $(MAIN_FILE) $(CMD)

## Run: Build then run the binary with optional command (e.g. make run-build CMD="get github")
run-build: build
	./$(BINARY_NAME) $(CMD)

## Test: Run all tests
test:
	go test -v ./...

## Clean: Remove built binaries and temporary files
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-linux
	rm -f $(BINARY_NAME)-mac
	rm -f $(BINARY_NAME)-mac-arm64
	rm -f $(BINARY_NAME).exe

## Deps: Tidy Go modules
deps:
	go mod tidy

## Fmt: Format Go code
fmt:
	go fmt ./...

## Lint: Run go vet
lint:
	go vet ./...

## Install: Install the binary globally
install:
	go install

## Release: Build binaries for multiple platforms
release:
	# Linux
	GOOS=linux GOARCH=amd64 go build $(GO_FLAGS) -o $(BINARY_NAME)-linux $(MAIN_FILE)
	# macOS (Intel)
	GOOS=darwin GOARCH=amd64 go build $(GO_FLAGS) -o $(BINARY_NAME)-mac $(MAIN_FILE)
	# macOS (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build $(GO_FLAGS) -o $(BINARY_NAME)-mac-arm64 $(MAIN_FILE)
	# Windows
	GOOS=windows GOARCH=amd64 go build $(GO_FLAGS) -o $(BINARY_NAME).exe $(MAIN_FILE)

## Help: Show available commands
help:
	@echo "GoVault Makefile Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
