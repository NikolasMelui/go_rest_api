.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	go build -v ./cmd/apiserver; ./apiserver

.DEFAULT_GOAL := build
