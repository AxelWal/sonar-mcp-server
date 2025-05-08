CURRENT_VERSION = $(shell git describe --tags --always --dirty)
VERSION ?= $(CURRENT_VERSION)

.PHONY: default
default: build

:PHONY: build
build:
	CGO_ENABLED=0 go build

clean:
	go clean -i ./...
	rm -rf dist/