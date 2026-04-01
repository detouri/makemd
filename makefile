BINARY := makemd
PKG := github.com/detouri/makemd
DIST := dist
BUILD_DIR := $(DIST)/build
RELEASE_DIR := $(DIST)/release
VERSION ?= $(shell cat VERSION 2>/dev/null || echo dev)
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w -X 'github.com/detouri/makemd/internal/cli.version=$(VERSION)'

.PHONY: help deps fmt test coverage build cross checksums set-version clean

help:
	@echo "make deps        - install dependencies"
	@echo "make fmt         - format code"
	@echo "make test        - run tests"
	@echo "make coverage    - run tests with coverage output"
	@echo "make build       - build local binary"
	@echo "make cross       - build macOS, Linux, and Windows binaries"
	@echo "make checksums   - create checksums for release artifacts"
	@echo "make set-version VERSION=x.y.z - sync versioned files"
	@echo "make clean       - remove build artifacts"

deps:
	go mod tidy

fmt:
	gofmt -w ./cmd ./internal

test:
	go test ./...

coverage:
	@mkdir -p coverage
	go test ./... -covermode=atomic -coverprofile=coverage/coverage.out
	go tool cover -func=coverage/coverage.out
	go tool cover -html=coverage/coverage.out

build:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY) ./cmd
	cp $(BUILD_DIR)/$(BINARY) $(DIST)/$(BINARY)

cross:
	@mkdir -p $(RELEASE_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(RELEASE_DIR)/$(BINARY)_linux_amd64 ./cmd
	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(RELEASE_DIR)/$(BINARY)_linux_arm64 ./cmd
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(RELEASE_DIR)/$(BINARY)_darwin_amd64 ./cmd
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(RELEASE_DIR)/$(BINARY)_darwin_arm64 ./cmd
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(RELEASE_DIR)/$(BINARY)_windows_amd64.exe ./cmd
	GOOS=windows GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(RELEASE_DIR)/$(BINARY)_windows_arm64.exe ./cmd

checksums:
	@mkdir -p $(RELEASE_DIR)
	cd $(RELEASE_DIR) && shasum -a 256 * > checksums.txt

set-version:
	@test -n "$(VERSION)" || (echo "VERSION must be set" && exit 1)
	./scripts/set-version.sh "$(VERSION)"

clean:
	rm -rf $(DIST) coverage
