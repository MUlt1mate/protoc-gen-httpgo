.DEFAULT_GOAL := help

PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

all: build install gen test format lint run		## all commands

build:      		## build
	@printf "\033[32mBuilding...\033[0m\n"
	@go build .

install:			## install plugin
	@printf "\033[31mInstalling plugin...\033[0m\n"
	@go install .
gen:				## generate example
	@printf "\033[33mGenerating code...\033[0m\n"
	@go generate ./example/generate.go
	@sed -i '1d' ./example/proto/*.httpgo.go # mark file as not generated for correct linter check

test:    			## run tests
	@printf "\033[35mRunning tests...\033[0m\n"
	@go test -v ./...
	@(cd ./example/ && go test -v ./...)

lint:        		## lint code
	@printf "\033[34mLinting code...\033[0m %s\n"
	@golangci-lint run
	@(cd ./example && golangci-lint run)

format:				## format code
	@printf "\033[36mFormatting code...\033[0m\n"
	@gofmt -s -w .

run: 			## run example code
	@printf "\033[37mRunning example methods...\033[0m\n"
	@(cd ./example/ && go run .)
