.PHONY: build lint test

BINARY_NAME=git-paca
BUILD_DIR=bin

build:
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)

lint:
	golangci-lint run

test:
	go test ./... -v
