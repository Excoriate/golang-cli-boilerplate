SHELL := /bin/bash
DIRS=$(shell ls)
GO=go
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GO_BUILD_SCRIPT := $(shell pwd)/scripts/golang/go_build.sh
PRE_COMMIT_INIT_SCRIPT := $(shell pwd)/scripts/hooks/pre-commit-init.sh
OUTPUT_DIR := $(shell pwd)/dist
CLI_NAME := golang-cli-boilerplate

.DEFAULT_GOAL := help

# ROOT_DIR: root directory of the code base
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/. && pwd -P))
endif

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --abbrev=0 --dirty --always --tags | sed 's/-/./g')
endif

GIT_COMMIT:=$(shell git rev-parse HEAD)

# ==============================================================================
# Targets Pre-commit
# ==============================================================================
pre-commit:
	@echo "===========> Pre-commit"
	@pre-commit run --all-files

pre-commit-init:
	@echo "===========> Pre-commit init"
	@$(PRE_COMMIT_INIT_SCRIPT) --hook-type=commit
	@$(PRE_COMMIT_INIT_SCRIPT) --hook-type=prepush

# ==============================================================================
# Targets Golang
# ==============================================================================

## tidy: tidy go.mod
.PHONY: tidy
tidy:
	@$(GO) mod tidy

## fmt: Run go fmt against code.
.PHONY: fmt
fmt:
	@$(GO) fmt -x ./...

## vet: Run go vet against code.
.PHONY: vet
vet:
	@$(GO) vet ./...

## lint: Run go lint against code.
.PHONY: lint
lint:
	@golangci-lint run -v --config .golangci.yaml

## style: Code style -> fmt,vet,lint
.PHONY: style
style: fmt vet lint

## test: Run unit test
.PHONY: test
test:
	@echo "===========> Run unit test"
	@$(GO) test -race -v ./...

## Build Go Binary
.PHONY: build
build:
	@echo "===========> Building binary"
	@$(GO_BUILD_SCRIPT) golang-cli-boilerplate

run:
	@echo "===========> Running binary"
	@./$(CLI_NAME) $(ARGS)


# ==============================================================================
# Targets Docker
# ==============================================================================
.PHONY: docker-build
dockerfile_default := Dockerfile          ## Default dockerfile name to Dockerfile

docker-build-refresh:
	@echo "Building Docker image entirely from scratch"
	./scripts/containers/build-and-run.sh --image=$(CLI_NAME) --rebuild=true --action=build --dockerfile=$(dockerfile_default)

docker-build:
	@echo "Building Docker image only if it doesn't exist previously."
	./scripts/containers/build-and-run.sh --image=$(CLI_NAME) --rebuild=false --action=build --dockerfile=$(dockerfile_default)

docker-build-run:
	@echo "Building and running the CLI on Docker"
	./scripts/containers/build-and-run.sh --image=$(CLI_NAME) --rebuild=true --action=all --dockerfile=$(dockerfile_default)

docker-run:
	@echo "Running the CLI on Docker from an already built image"
	./scripts/containers/build-and-run.sh --image=$(CLI_NAME)  --rebuild=false --action=run
