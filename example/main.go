package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fasthttp/router"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/valyala/fasthttp"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	fasthttpimpl "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/fasthttp"
	nethttpimpl "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/nethttp"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	fasthttpproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fasthttp"
	httpproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/nethttp"
)

const (
	fasthttpAddr = "localhost:8080"
	nethttpAddr  = "localhost:8081"
)

var (
	fastClient    httpproto.ServiceNameHTTPGoService
	nethttpClient httpproto.ServiceNameHTTPGoService
)

func main() {
	log.SetFlags(log.Lshortfile)
	var (
		err error
		ctx = context.TODO()
	)
	if err = serverInit(ctx); err != nil {
		log.Fatal(err)
	}
	if err = clientInit(ctx); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond * 100)

	if err = clientRunRequests(ctx, fastClient); err != nil {
		log.Println(err)
	}
	if err = clientRunRequests(ctx, nethttpClient); err != nil {
		log.Println(err)
	}

	// f := make(chan bool)
	// <-f
}

func serverInit(ctx context.Context) (err error) {
	var (
		handler fasthttpproto.ServiceNameHTTPGoService = &implementation.Handler{}
		r                                              = router.New()
		rHttp                                          = http.NewServeMux()
	)
	if err = fasthttpproto.RegisterServiceNameHTTPGoServer(ctx, r, handler, fasthttpimpl.ServerMiddlewares); err != nil {
		return err
	}
	if err = httpproto.RegisterServiceNameHTTPGoServer(ctx, rHttp, handler, nethttpimpl.ServerMiddlewares); err != nil {
		return err
	}

	go func() {
		if err = fasthttp.ListenAndServe(fasthttpAddr, r.Handler); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err = http.ListenAndServe(nethttpAddr, rHttp); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func clientInit(ctx context.Context) (err error) {
	var fasthttpClientTransport = &fasthttp.Client{}
	if fastClient, err = fasthttpproto.GetServiceNameHTTPGoClient(
		ctx,
		fasthttpClientTransport,
		"http://"+fasthttpAddr,
		fasthttpimpl.ClientMiddlewares,
	); err != nil {
		return err
	}

	var httpClientTransport = &http.Client{}
	if nethttpClient, err = httpproto.GetServiceNameHTTPGoClient(
		ctx,
		httpClientTransport,
		"http://"+nethttpAddr,
		nethttpimpl.ClientMiddlewares,
	); err != nil {
		return err
	}
	return nil
}

func clientRunRequests(ctx context.Context, client httpproto.ServiceNameHTTPGoService) (err error) {
	if _, err = client.RPCName(ctx, &common.InputMsgName{Int64Argument: 999, StringArgument: "rand"}); err != nil {
		return fmt.Errorf("RPCName failed: %w", err)
	}

	var allTypesResp *common.AllTypesMsg
	if allTypesResp, err = client.AllTypesTest(ctx, &implementation.AllTypesMsg); err != nil {
		return fmt.Errorf("AllTypesTest failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.AllTypesMsg, allTypesResp, cmpopts.IgnoreUnexported(implementation.AllTypesMsg)); diff != "" {
		log.Println(diff)
	}
	var allNumberTypesResp *common.AllNumberTypesMsg
	if allNumberTypesResp, err = client.AllTypesMaxTest(ctx, &implementation.NumberTypesMaxMsg); err != nil {
		return fmt.Errorf("AllTypesTest failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.NumberTypesMaxMsg, allNumberTypesResp, cmpopts.IgnoreUnexported(implementation.NumberTypesMaxMsg)); diff != "" {
		log.Println(diff)
	}
	if allNumberTypesResp, err = client.AllTypesMaxQueryTest(ctx, &implementation.NumberTypesMaxMsg); err != nil {
		return fmt.Errorf("AllTypesTest failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.NumberTypesMaxMsg, allNumberTypesResp, cmpopts.IgnoreUnexported(implementation.NumberTypesMaxMsg)); diff != "" {
		log.Println(diff)
	}

	if _, err = client.MultipartForm(ctx, &implementation.MultipartFormRequestMsg); err != nil {
		return fmt.Errorf("MultipartForm failed: %w", err)
	}
	if _, err = client.MultipartFormAllTypes(ctx, &implementation.MultipartFormRequestAllTypesMsg); err != nil {
		return fmt.Errorf("MultipartFormAllTypes failed: %w", err)
	}

	var allTextTypesResp *common.AllTextTypesMsg
	if allTextTypesResp, err = client.AllTextTypesGet(ctx, &implementation.AllTextTypesMsg); err != nil {
		return fmt.Errorf("AllTextTypesGet failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.AllTextTypesMsg, allTextTypesResp, cmpopts.IgnoreUnexported(implementation.AllTextTypesMsg)); diff != "" {
		log.Println(diff)
	}
	if allTextTypesResp, err = client.AllTextTypesPost(ctx, &implementation.AllTextTypesMsg); err != nil {
		return fmt.Errorf("AllTextTypesPost failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.AllTextTypesMsg, allTextTypesResp, cmpopts.IgnoreUnexported(implementation.AllTextTypesMsg)); diff != "" {
		log.Println(diff)
	}

	// http rule checks
	if _, err = client.GetMessage(ctx, &common.GetMessageRequest{Name: "123456"}); err != nil {
		return fmt.Errorf("GetMessage failed: %w", err)
	}
	if _, err = client.GetMessageV2(ctx, &common.GetMessageRequestV2{
		MessageId: "123456",
		Sub:       &common.GetMessageRequestV2_SubMessage{Subfield: "foo"},
	}); err != nil {
		return fmt.Errorf("GetMessageV2 failed: %w", err)
	}
	if _, err = client.UpdateMessage(ctx, &common.UpdateMessageRequest{
		MessageId: "123456",
		Message:   &common.MessageV2{Text: "Hi!"},
	}); err != nil {
		return fmt.Errorf("UpdateMessage failed: %w", err)
	}
	if _, err = client.UpdateMessageV2(ctx, &common.MessageV2{
		Text:      "Hi!",
		MessageId: "123456",
	}); err != nil {
		return fmt.Errorf("UpdateMessageV2 failed: %w", err)
	}
	if _, err = client.GetMessageV3(ctx, &common.GetMessageRequestV3{
		MessageId: "234567",
		UserId:    "",
	}); err != nil {
		return fmt.Errorf("GetMessageV3 failed: %w", err)
	}
	if _, err = client.GetMessageV4(ctx, &common.GetMessageRequestV3{
		MessageId: "seg1/seg2.ext",
	}); err != nil {
		return fmt.Errorf("GetMessageV4 failed: %w", err)
	}

	return nil
}
