# If the "VERSION" environment variable is not set, then use this value instead
VERSION?=1.0.0
TIME=$(shell date +%FT%T%z)
GOVERSION=$(shell go version | awk '{print $$3}' | sed s/go//)

LDFLAGS=-ldflags "\
	-X github.com/softwarespot/poker-evaluator/internal/version.Version=${VERSION} \
	-X github.com/softwarespot/poker-evaluator/internal/version.Time=${TIME} \
	-X github.com/softwarespot/poker-evaluator/internal/version.User=${USER} \
	-X github.com/softwarespot/poker-evaluator/internal/version.GoVersion=${GOVERSION} \
	-s \
	-w \
"

build:
	@echo building to bin/poker-evaluator
	@go build $(LDFLAGS) -o ./bin/poker-evaluator

test:
	@go test -cover -v ./...

.PHONY: build test
