#!/bin/bash

BINARY_NAME=kgraph-1.0

env GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-windows.exe ./cmd/main.go
env GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux ./cmd/main.go

