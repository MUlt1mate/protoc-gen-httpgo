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
	http2 "github.com/MUlt1mate/protoc-gen-httpgo/example/middleware/http"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	fasthttpproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fasthttp"
	httpproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/nethttp"
)

var (
	serverMiddlewares = middleware.ServerMiddlewares
	clientMiddlewares = middleware.ClientMiddlewares
)

func main() {
	_ = serverExample(context.TODO())
	time.Sleep(time.Millisecond * 500)
	_ = clientExample(context.TODO())
	// f := make(chan bool)
	// <-f
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
	if err = httpproto.RegisterServiceNameHTTPGoServer(ctx, rHttp, handler, http2.ServerMiddlewares); err != nil {
		return err
	}

	go func() {
		if err = fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err = http.ListenAndServe(":8081", rHttp); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func clientExample(ctx context.Context) (err error) {
	var (
		fastClient              *fasthttpproto.ServiceNameHTTPGoClient
		fasthttpClientTransport = &fasthttp.Client{}
		fasthttpHost            = "http://localhost:8080"
		httpClient              *httpproto.ServiceNameHTTPGoClient
		httpClientTransport     = &http.Client{}
		httpHost                = "http://localhost:8081"
	)
	if fastClient, err = fasthttpproto.GetServiceNameHTTPGoClient(ctx, fasthttpClientTransport, fasthttpHost, clientMiddlewares); err != nil {
		return err
	}
	if httpClient, err = httpproto.GetServiceNameHTTPGoClient(ctx, httpClientTransport, httpHost, http2.ClientMiddlewares); err != nil {
		return err
	}
	// sending our request
	// _, _ = fastClient.RPCName(context.Background(), &common.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
	// _, _ = fastClient.AllTypesTest(context.Background(), &common.AllTypesMsg{
	// 	SliceStringValue: []string{"a a", "b b"},
	// 	BytesValue:       []byte("hello world"),
	// 	StringValue:      "hello world",
	// })
	// _, err = fastClient.MultipartForm(context.Background(), &common.MultipartFormRequest{
	// 	Document: &common.FileEx{
	// 		File: []byte(`file content`),
	// 		Name: "file.exe",
	// 	},
	// 	OtherField: "otherField",
	// })
	dd := "d d"
	hh := []byte("h h")
	opt := common.Options_FIRST
	allTextTypesMsg := &common.AllTextTypesMsg{
		String_:        "a a",
		RepeatedString: []string{"b b", "c c"},
		OptionalString: &dd,
		Bytes:          []byte("e e"),
		RepeatedBytes:  [][]byte{[]byte("f f"), []byte("g g")},
		OptionalBytes:  hh,
		Enum:           common.Options_FIRST,
		RepeatedEnum:   []common.Options{common.Options_FIRST, common.Options_SECOND},
		OptionalEnum:   &opt,
	}
	_, err = fastClient.AllTextTypesGet(context.Background(), allTextTypesMsg)
	_, err = fastClient.AllTextTypesPost(context.Background(), allTextTypesMsg)

	// _, _ = httpClient.RPCName(context.Background(), &common.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
	// _, _ = httpClient.AllTypesTest(context.Background(), &common.AllTypesMsg{
	// 	SliceStringValue: []string{"a a", "b b"},
	// 	BytesValue:       []byte("hello world"),
	// 	StringValue:      "hello world",
	// })
	// _, err = httpClient.MultipartForm(context.Background(), &common.MultipartFormRequest{
	// 	Document: &common.FileEx{
	// 		File: []byte(`file content`),
	// 		Name: "file.exe",
	// 	},
	// 	OtherField: "otherField",
	// })

	_, err = httpClient.AllTextTypesGet(context.Background(), allTextTypesMsg)
	_, err = httpClient.AllTextTypesPost(context.Background(), allTextTypesMsg)

	if err != nil {
		log.Println(err)
	}
	return nil
}
