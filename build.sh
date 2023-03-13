#!/bin/bash

BINARY_NAME=kustomize_md
BINARY_VERSION=v1

env GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}_windows_${BINARY_VERSION}.exe ./cmd/main.go
env GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}_linux_${BINARY_VERSION} ./cmd/main.go

