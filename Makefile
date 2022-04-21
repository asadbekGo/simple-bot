.PHONY: build

CURRENT_DIR=$(shell pwd)
APP=payment-service
APP_CMD_DIR=./cmd

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

lint:
	golangci-lint -c .golangci.yaml run ./...
