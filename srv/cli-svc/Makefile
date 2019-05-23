# Makefile
#
# Author: Pavle Batuta (pavle.batuta@protonmail.com)
# Description: Makefile for standalone and docker build and run of the service.

########################
# PARAMETERS:
########################

# General parameters:
NAME=cli

# Golang:
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Targets:
SERVICE_NAME=$(NAME)-svc
BIN_PATH=bin/$(NAME)

# System:
RM_CMD=rm -f

########################
# RECEPIES:
########################

## 1: Builds

# Linux target, single binary, static linkage:
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	 -a \
	 -installsuffix cgo \
	 -o bin/$(NAME) .
	docker build -t $(SERVICE_NAME) .

# Build default target:
.PHONY: build
build: build-linux

## 2: Running

# Run in a docker container:
.PHONY: docker-run
docker-run:
	docker run \
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

# Fetch deps:
.PHONY: deps
deps:
	$(GOMOD) download && $(GOMOD) tidy

