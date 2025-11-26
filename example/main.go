package main

import (
	"context"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/middleware"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fasthttp"
)

var (
	serverMiddlewares = middleware.ServerMiddlewares
	clientMiddlewares = middleware.ClientMiddlewares
)

func main() {
	_ = serverExample(context.TODO())
	time.Sleep(time.Millisecond * 500)
	_ = clientExample(context.TODO())
	f := make(chan bool)
	<-f
}

func serverExample(ctx context.Context) (err error) {
	var (
		handler proto.ServiceNameHTTPGoService = &implementation.Handler{}
		r                                      = router.New()
		// r                                      = http.NewServeMux()
	)
	if err = proto.RegisterServiceNameHTTPGoServer(ctx, r, handler, serverMiddlewares); err != nil {
		return err
	}

	go func() {
		_ = fasthttp.ListenAndServe(":8080", r.Handler)
		// _ = http.ListenAndServe(":8080", r)
	}()
	return nil
}

func clientExample(ctx context.Context) (err error) {
	var (
		client *proto.ServiceNameHTTPGoClient
		// httpClient = &http.Client{}
		httpClient = &fasthttp.Client{}
		host       = "http://localhost:8080"
	)
	if client, err = proto.GetServiceNameHTTPGoClient(ctx, httpClient, host, clientMiddlewares); err != nil {
		return err
	}
	// sending our request
	_, _ = client.RPCName(context.Background(), &proto.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
	_, _ = client.AllTypesTest(context.Background(), &proto.AllTypesMsg{
		SliceStringValue: []string{"a", "b"},
		BytesValue:       []byte("hello world"),
		StringValue:      "hello world",
	})
	return nil
}
