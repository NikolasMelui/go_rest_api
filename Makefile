.PHONY: build test run
build:
	go build -v ./cmd/apiserver

test:
	go test -v -race -timeout 30s ./...

run:
	go build -v ./cmd/apiserver; ./apiserver

.DEFAULT_GOAL := build
