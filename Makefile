.PHONY: build lint

BINARY_NAME=git-paca
BUILD_DIR=bin

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/git_paca

lint:
	golangci-lint run
