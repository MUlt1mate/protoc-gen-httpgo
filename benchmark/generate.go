//go:generate protoc -I=. -I=./vendor --go_out=paths=source_relative:. --httpgo_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=logtostderr=true,paths=source_relative:. ./proto/api.proto

package main

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
)
