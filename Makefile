.PHONY: all lint format
.DEFAULT_GOAL := all

GOFILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

all: lint format

lint:
	golangci-lint run $(GOFILES) --config .golangci.yml

format:
	gofmt -w $(GOFILES)