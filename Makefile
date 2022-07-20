.PHONY: all
all: deps build

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

.PHONY: clean
clean:
	go clean
	go clean -testcache
	rm -f $(BINARY_NAME)
	rm -f coverage.out

.PHONY: stats
stats:
	cloc --by-file --vcs=git .
