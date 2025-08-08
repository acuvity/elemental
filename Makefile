MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

export GO111MODULE = on

default: lint test

install-tools:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install github.com/securego/gosec/cmd/gosec@master

lint:
	golangci-lint run \
		--timeout=5m \
		--disable=govet  \
		--enable=errcheck \
		--enable=ineffassign \
		--enable=unused \
		--enable=unconvert \
		--enable=misspell \
		--enable=prealloc \
		--enable=nakedret \
		--enable=unparam \
		--enable=nilerr \
		--enable=bodyclose \
		--enable=errorlint \
		./...
sec:
	gosec -quiet ./...

.PHONY: test
test:
	go test ./... -race -cover -covermode=atomic -coverprofile=unit_coverage.out

remod:
	go get go.acuvity.ai/regolithe@master
	go mod tidy
