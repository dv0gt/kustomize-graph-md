#!/bin/bash

BINARY_NAME=kustomize-markdown

mkdir -p ./bin/
env GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY_NAME} ./cmd/main.go

