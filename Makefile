.PHONY: all lint test

all: lint test

lint:
	golangci-lint run --fix ./...

test:
	go test -cover ./...
