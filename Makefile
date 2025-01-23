.PHONY: all lint format

.DEFAULT_GOAL := all

SRC := $(shell pwd)/src
GCI := $(shell pwd)/.golangci.yml

all: lint format

lint:
	find src -name "go.mod" -execdir golangci-lint run --config $(GCI) ./... \;

format:
	gofmt -w $(SRC)