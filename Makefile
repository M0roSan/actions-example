.DEFAULT_GOAL := all
all: get build lint test 

get:
	go get ./...

build:
	go build -v -o foo

lint: .tools/golangci-lint
	.tools/golangci-lint run

test:
	go test -parallel=10 -v -covermode=count -coverprofile=coverage_unit.out ./... $(TESTARGS)

tools: .tools .tools/gocover-cobertura .tools/gocovmerge .tools/golangci-lint 

.tools:
	@mkdir -p .tools

.tools/gocover-cobertura:
	export GOBIN=$(shell pwd)/.tools; go install github.com/boumenot/gocover-cobertura@v1.1.0

.tools/gocovmerge:
	export GOBIN=$(shell pwd)/.tools; go install github.com/wadey/gocovmerge@master

.tools/golangci-lint:
	export GOBIN=$(shell pwd)/.tools; go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0