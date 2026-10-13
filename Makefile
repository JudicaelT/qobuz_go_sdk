MAKEFLAGS += --silent

.PHONY: help

help: # Show available make commands
	@awk 'BEGIN {FS = ":.*#"} /^[a-zA-Z0-9_-]+:.*#/ { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

test: # Run unit tests
	@echo -e "\e[0;32mRunning tests...\e[0m"
	go test -cover ./...

test-coverage: # Run tests and output coverage.out file
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lint: # Format source files
	gofmt -d -w .
