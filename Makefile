SHELL := /bin/bash
DIRS=$(shell ls)
GO=go
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GO_BUILD_SCRIPT := $(shell pwd)/scripts/golang/go_build.sh
PRE_COMMIT_INIT_SCRIPT := $(shell pwd)/scripts/hooks/pre-commit-init.sh
OUTPUT_DIR := $(shell pwd)/dist
BINARY := golang-cli-boilerplate
AIR_SETUP_SCRIPT := $(shell pwd)/scripts/golang/air.sh
AIR_BINARY := $(HOME)/go/bin/air
SRC_DIR := $(shell pwd)

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
	@$(PRE_COMMIT_INIT_SCRIPT) --hook-type=auto-update

# ==============================================================================
# Targets Golang
# ==============================================================================

## tidy: tidy go.mod
.PHONY: go-tidy
go-tidy:
	@$(GO) mod tidy

## fmt: Run go fmt against code.
.PHONY: go-fmt
go-fmt:
	@$(GO) fmt -x ./...

## vet: Run go vet against code.
.PHONY: go-vet
go-vet:
	@$(GO) vet ./...

## lint: Run go lint against code.
.PHONY: go-lint
go-lint:
	@golangci-lint run -v --config .golangci.yaml

## style: Code style -> fmt,vet,lint
.PHONY: go-style
go-style: go-fmt go-vet go-lint

## test: Run unit test
.PHONY: go-test
go-test:
	@echo "===========> Run unit test"
	@$(GO) test -race -v ./...

## Build Go Binary
.PHONY: go-build
go-build:
	@echo "===========> Building binary"
	@$(GO_BUILD_SCRIPT) --binary go-service-boilerplate --path ./main.go

## Run Go source code
.PHONY: go-run
go-run:
	@echo "===========> Running source code"
	@$(GO) run $(SRC_DIR)/main.go $(ARGS)

.PHONY: go-run-bin
go-run-bin:
	@echo "===========> Running binary"
	@./$(BINARY) $(ARGS)

.PHONY: go-ci
go-ci: go-style go-test go-tidy go-build


# ==============================================================================
# Targets Docker
# ==============================================================================
.PHONY: docker-build
dockerfile_default := Dockerfile          ## Default dockerfile name to Dockerfile

.PHONY: docker-build-refresh
docker-build-refresh:
	@echo "Building Docker image entirely from scratch"
	./scripts/containers/build-and-run.sh --image=$(BINARY) --rebuild=true --action=build --dockerfile=$(dockerfile_default)

.PHONY: docker-build
docker-build:
	@echo "Building Docker image only if it doesn't exist previously."
	./scripts/containers/build-and-run.sh --image=$(BINARY) --rebuild=false --action=build --dockerfile=$(dockerfile_default)

.PHONY: docker-run
docker-build-run:
	@echo "Building and running the App on Docker"
	./scripts/containers/build-and-run.sh --image=$(BINARY) --rebuild=true --action=all --dockerfile=$(dockerfile_default)

.PHONY: docker-run
docker-run:
	@echo "Running the App on Docker from an already built image"
	./scripts/containers/build-and-run.sh --image=$(BINARY)  --rebuild=false --action=run

# ==============================================================================
# Other developer tools
# ==============================================================================
.PHONY: air-setup
air-setup:
	@echo "===========> Setting up air"
	@$(AIR_SETUP_SCRIPT) --force

.PHONY: air-run
air-run:
	@echo "===========> Running air"
	@$(AIR_BINARY) -c .air.toml
