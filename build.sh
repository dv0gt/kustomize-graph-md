#!/bin/bash

BINARY_NAME=kustomize-markdown

env GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}.exe ./cmd/main.go
env GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} ./cmd/main.go

