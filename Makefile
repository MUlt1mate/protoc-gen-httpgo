.DEFAULT_GOAL := help

PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
GOWORKPATH := "/home/mult1mate/go"
INCLUDEPATH := "/usr/local/include"
.PHONY: debug

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
	@go generate ./example/proto/generate.go
	@sed -i '1d' ./example/proto/*.httpgo.go # mark file as not generated for correct linter check
	@sed -i '1d' ./example/proto/fasthttp/*.httpgo.go

test:    			## run tests
	@printf "\033[35mRunning tests...\033[0m\n"
	@go test -v ./...
	@(cd ./example/ && go test -v ./...)

lint:        		## lint code https://golangci-lint.run/welcome/install/
	@printf "\033[34mLinting code...\033[0m %s\n"
	@golangci-lint run --fix
	@(cd ./example && golangci-lint run --fix)

format:				## format code
	@printf "\033[36mFormatting code...\033[0m\n"
	@gofmt -s -w .

run: 			## run example code
	@printf "\033[37mRunning example methods...\033[0m\n"
	@(cd ./example/ && go run .)

debug: ## go install github.com/lyft/protoc-gen-star/protoc-gen-debug@latest
	@protoc -I=. -I=./example/vendor -I=/usr/local/include -I=./example/proto \
       --plugin=protoc-gen-debug=/home/mult1mate/go/bin/protoc-gen-debug \
       --debug_out="./debug:." \
       ./example/proto/*.proto

installdeps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5
	@go install github.com/mailru/easyjson/...@v0.9.0
	@mkdir ${GOWORKPATH}/bin/ -p
	@wget -q https://github.com/protocolbuffers/protobuf/releases/download/v31.1/protoc-31.1-linux-x86_64.zip -P .  \
             && unzip ./protoc-31.1-linux-x86_64.zip -d ./protoc  \
             && mv ./protoc/bin/* ${GOWORKPATH}/bin/.  \
             && sudo mv ./protoc/include/google ${INCLUDEPATH}/.  \
             && chmod 755 ${GOWORKPATH}/bin/* \
             && rm -rf ./protoc
	@export TMP_PATH=${GOWORKPATH}/src/github.com/googleapis \
         && sudo mkdir ${INCLUDEPATH}/google/api -p \
         && git clone https://github.com/googleapis/googleapis.git $TMP_PATH \
         && cd $TMP_PATH \
         && git checkout d9eae9f029427bd9ed4379d8e3cd46ca69f1a33f \
         && sudo cp google/api/annotations.proto google/api/field_behavior.proto google/api/http.proto google/api/httpbody.proto \
         ${INCLUDEPATH}/google/api \
         && rm -rf $TMP_PATH