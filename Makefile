.PHONY: all
all: deps build test

# BINARY
BINARY_DIR=bin
BINARY_NAME=$(BINARY_DIR)/paket

# FLAGS
HEAD = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date "+%Y-%m-%dT%H:%M:%S")
BUILD_OS = $(shell uname)
LINKER_FLAGS = -ldflags "-X main.commit=$(HEAD) -X main.date=$(BUILD_DATE) -X main.hostOS=$(BUILD_OS)"
UNAME := $(shell uname)

# COMMANDS
.PHONY: deps
deps:
	go mod tidy

.PHONY: build
build:
	go build -o $(BINARY_NAME) $(LINKER_FLAGS) ./cmd

.PHONY: test
test:
	go test -cover -covermode=atomic ./...

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...

.PHONY: coverage-cli
coverage-cli:
	go tool cover -func=coverage.out

.PHONY: coverage-html
coverage-html:
	go tool cover -html=coverage.out

.PHONY: check
check:
	pre-commit run --all-files

.PHONY: clean
clean:
	go clean
	go clean -testcache
	rm -f $(BINARY_NAME)
	rm -f coverage.out

.PHONY: stats
stats:
	cloc --vcs=git .
