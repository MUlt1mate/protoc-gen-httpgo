.DEFAULT_GOAL := help

PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
GOWORKPATH := "/home/mult1mate/go"
INCLUDEPATH := "/usr/local/include"
.PHONY: debug

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

run:
	@go build -o benchmark .
	@touch output_grpcgateway.txt output_httpgo.txt
	@truncate -s 0 output_grpcgateway.txt output_httpgo.txt
	@./benchmark grpc-gateway >> output_grpcgateway.txt & GOPID=$$! ; \
	 sleep 1 ; \
	 go test ./... -bench=. >> output_grpcgateway.txt ; \
	 go test ./... -bench=. -count=10 > bench_grpcgateway.txt ; \
	 echo " " >> output_grpcgateway.txt ; \
	 curl --silent -o allocs.out http://localhost:6060/debug/pprof/allocs ; \
	 go tool pprof -top allocs.out | head --lines=15 >> output_grpcgateway.txt ; \
	 kill $$GOPID
	@./benchmark httpgo >> output_httpgo.txt & GOPID=$$! ; \
	 sleep 1 ; \
	 go test ./... -bench=. >> output_httpgo.txt ; \
	 go test ./... -bench=. -count=10 > bench_httpgo.txt ; \
	 echo " " >> output_httpgo.txt ; \
	 curl --silent -o allocs.out http://localhost:6060/debug/pprof/allocs ; \
	 go tool pprof -top allocs.out | head --lines=15 >> output_httpgo.txt ; \
	 kill $$GOPID
	@benchstat bench_grpcgateway.txt bench_httpgo.txt > benchstat.txt

gen:				## generate example
	@printf "\033[33mGenerating code...\033[0m\n"
	@go generate ./generate.go

installdeps:
	@go install golang.org/x/perf/cmd/benchstat@latest
	@go install github.com/MUlt1mate/protoc-gen-httpgo@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.1
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
