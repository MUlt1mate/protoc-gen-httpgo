package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/middleware"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	fasthttpproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fasthttp"
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
		handler fasthttpproto.ServiceNameHTTPGoService = &implementation.Handler{}
		r                                              = router.New()
		rHttp                                          = http.NewServeMux()
	)
	if err = fasthttpproto.RegisterServiceNameHTTPGoServer(ctx, r, handler, serverMiddlewares); err != nil {
		return err
	}

	go func() {
		_ = fasthttp.ListenAndServe(":8080", r.Handler)
		_ = http.ListenAndServe(":8081", rHttp)
	}()
	return nil
}

func clientExample(ctx context.Context) (err error) {
	var (
		fastClient              *fasthttpproto.ServiceNameHTTPGoClient
		fasthttpClientTransport = &fasthttp.Client{}
		fasthttpHost            = "http://localhost:8080"
		// httpClient              *httpproto.ServiceNameHTTPGoClient
		// httpClientTransport     = &http.Client{}
		// httpHost                = "http://localhost:8081"
	)
	if fastClient, err = fasthttpproto.GetServiceNameHTTPGoClient(ctx, fasthttpClientTransport, fasthttpHost, clientMiddlewares); err != nil {
		return err
	}
	// if httpClient, err = httpproto.GetServiceNameHTTPGoClient(ctx, httpClientTransport, httpHost, clientMiddlewares); err != nil {
	// 	return err
	// }
	// sending our request
	_, _ = fastClient.RPCName(context.Background(), &common.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
	_, _ = fastClient.AllTypesTest(context.Background(), &common.AllTypesMsg{
		SliceStringValue: []string{"a", "b"},
		BytesValue:       []byte("hello world"),
		StringValue:      "hello world",
	})
	// _, _ = httpClient.RPCName(context.Background(), &common.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
	// _, _ = httpClient.AllTypesTest(context.Background(), &common.AllTypesMsg{
	// 	SliceStringValue: []string{"a", "b"},
	// 	BytesValue:       []byte("hello world"),
	// 	StringValue:      "hello world",
	// })

	_, err = fastClient.MultipartForm(context.Background(), &common.MultipartFormRequest{
		Document: &common.FileEx{
			File: []byte(`file content`),
			Name: "file.exe",
		},
		OtherField: "otherField",
	})
	if err != nil {
		log.Println(err)
	}
	return nil
}
