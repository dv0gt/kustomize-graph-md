#!/bin/bash

BINARY_NAME=kustomize-markdown

mkdir -p ./bin/
env GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY_NAME}-darwin-amd64 ./cmd/main.go
env GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY_NAME}-linux-amd64 ./cmd/main.go

