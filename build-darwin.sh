#!/bin/bash

BINARY_NAME=kustomize-markdown

mkdir -p ./bin/
env GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY_NAME} ./main.go
