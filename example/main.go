package main

import (
	"context"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/middleware"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
)

var (
	serverMiddlewares = middleware.ServerMiddlewares
	clientMiddlewares = middleware.ClientMiddlewares
)

func main() {
	_ = serverExample(context.TODO())
	_ = clientExample(context.TODO())
}

func serverExample(ctx context.Context) (err error) {
	var (
		handler proto.ServiceNameHTTPGoService = &implementation.Handler{}
		r                                      = router.New()
	)
	if err = proto.RegisterServiceNameHTTPGoServer(ctx, r, handler, serverMiddlewares); err != nil {
		return err
	}

	go func() { _ = fasthttp.ListenAndServe(":8080", r.Handler) }()
	return nil
}

func clientExample(ctx context.Context) (err error) {
	var (
		client     *proto.ServiceNameHTTPGoClient
		httpClient = &fasthttp.Client{}
		host       = "http://localhost:8080"
	)
	if client, err = proto.GetServiceNameHTTPGoClient(ctx, httpClient, host, clientMiddlewares); err != nil {
		return err
	}
	// sending our request
	_, _ = client.RPCName(context.Background(), &proto.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
	return nil
}
