
include .env
SHELL := /bin/bash
# LOCAL_BIN:=$(CURDIR)/bin
# PATH:=$(LOCAL_BIN):$(PATH)
export $(shell sed 's/=.*//' .env)

.PHONY: build
build: 
	go build -v -o ./bin/main ./cmd/server/main.go

.PHONY: run
run: 
	./bin/main

.PHONY: install
install:
	go mod download
	
.PHONY: migrate-build
migrate-build:
	go build -o bin/migrate ./cmd/migrate

migrate-up: migrate-build
	./bin/migrate --up

migrate-down: migrate-build
	./bin/migrate --down

swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init --outputTypes yaml --parseDependency --parseInternal -g ./cmd/server/main.go 
.PHONY: swagger

.DEFAULT_GOAL = build