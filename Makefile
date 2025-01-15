.PHONY: all lint format

.DEFAULT_GOAL := all

SRC := $(shell find . -name "*.go" -not -path "./vendor/*")

all: lint format

lint:
	golangci-lint run --config .golangci.yml

format:
	gofmt -w $(SRC)