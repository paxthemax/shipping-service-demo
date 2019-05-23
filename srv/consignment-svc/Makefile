# Makefile
#
# Author: Pavle Batuta (pavle.batuta@protonmail.com)
# Description: Makefile for standalone and docker build and run of the service.

########################
# PARAMETERS:
########################

# General parameters:
NAME=consignment

# Golang:
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

# Go mod:
GOMOD=$(GOCMD) mod

# Targets:
SERVICE_NAME=$(NAME)-svc
BIN_PATH=bin/$(NAME)

# System:
RM_CMD=rm -f

# Protobuf:
PROTO_DEFINITION_PATH=proto/$(NAME).proto
PROTO_TARGET_PATH=proto/$(NAME).pb.go

########################
# RECEPIES:
########################

## 1: Protobuf

# Code generation:
.PHONY: proto-build
proto-build: $(PROTO_TARGET_PATH)

# Real target is the pb.go file:
$(PROTO_TARGET_PATH): $(PROTO_DEFINITION_PATH)
	$(info Compiling $(PROTO_TARGET_PATH))
	protoc -I. --go_out=plugins=micro:. $(PROTO_DEFINITION_PATH)

# Cleaning protobuf generated code:
.PHONY: proto-clean
proto-clean:
	$(RM_CMD) $(PROTO_TARGET_PATH)

## 2: Builds

# Linux target, single binary statically linked:
.PHONY: build-linux
build-linux: proto-build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	 -a \
	 -installsuffix cgo \
	 -o bin/$(NAME) .
	docker build -t $(SERVICE_NAME) .

# Build default target:
.PHONY: build
build: build-linux

## 3: Running

# Run in a docker container:
.PHONY: docker-run
docker-run:
	docker run \
		-p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		$(SERVICE_NAME)

# Default run target:
.PHONY: run
run: docker-run

## 4: Misc:

# Clean up binary:
.PHONY: clean-bin
clean-bin:
	$(RM_CMD) $(BIN_PATH)

# Default clean target:
.PHONY: clean
clean: clean-bin

# Clean ALL:
.PHONY: clean-ALL
clean-all: clean-bin proto-clean

# Fetch deps:
.PHONY: deps
deps:
	$(GOMOD) download && $(GOMOD) tidy

