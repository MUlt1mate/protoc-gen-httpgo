package main

import (
	"context"

	"github.com/MUlt1mate/protoc-gen-httpgo/benchmark/proto"
)

type Handler struct {
	proto.UnimplementedAPIMeasureServer
}

var _ proto.APIMeasureServer = &Handler{}
var _ proto.APIMeasureHTTPGoService = &Handler{}

func (l *Handler) Measure(_ context.Context, req *proto.MeasureRequest) (*proto.MeasureRequest, error) {
	return req, nil
}
