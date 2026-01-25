package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fasthttp/router"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/valyala/fasthttp"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	fasthttpmdlwr "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/fasthttp"
	ginmdlwr "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/gin"
	nethttpmdlwr "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/nethttp"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	fastproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fasthttp"
	ginproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/gin"
	httpproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/nethttp"
)

const (
	fasthttpAddr = "localhost:8080"
	nethttpAddr  = "localhost:8081"
	ginAddr      = "localhost:8082"
)

var (
	fastClient          httpproto.ServiceNameHTTPGoService
	nethttpClient       httpproto.ServiceNameHTTPGoService
	nethttpClientForGin httpproto.ServiceNameHTTPGoService
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
	if err = clientRunRequests(ctx, nethttpClientForGin); err != nil {
		log.Println(err)
	}

	// f := make(chan bool)
	// <-f
}

func serverInit(ctx context.Context) (err error) {
	gin.SetMode(gin.ReleaseMode)
	var (
		handler        fastproto.ServiceNameHTTPGoService = &implementation.Handler{}
		fasthttpRouter                                    = router.New()
		rHttp                                             = http.NewServeMux()
		ginRouter                                         = gin.New()
	)
	if err = fastproto.RegisterServiceNameHTTPGoServer(ctx, fasthttpRouter, handler, fasthttpmdlwr.ServerMiddlewares); err != nil {
		return err
	}
	if err = httpproto.RegisterServiceNameHTTPGoServer(ctx, rHttp, handler, nethttpmdlwr.ServerMiddlewares); err != nil {
		return err
	}
	/*
		Gin has its own middleware format, but with this one you can have transport independent handler
		with context.Context that can be populated with the same keys in both HTTP and GRPC middlewares.
		And of course if you have only HTTP you can use gin middleware format and pass nil to httpgo middlewares
	*/
	if err = ginproto.RegisterServiceNameHTTPGoServer(ctx, ginRouter, handler, ginmdlwr.ServerMiddlewares); err != nil {
		return err
	}

	go func() {
		if err = fasthttp.ListenAndServe(fasthttpAddr, fasthttpRouter.Handler); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err = http.ListenAndServe(nethttpAddr, rHttp); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err = ginRouter.Run(ginAddr); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func clientInit(ctx context.Context) (err error) {
	var fasthttpClientTransport = &fasthttp.Client{}
	if fastClient, err = fastproto.GetServiceNameHTTPGoClient(
		ctx,
		fasthttpClientTransport,
		"http://"+fasthttpAddr,
		fasthttpmdlwr.ClientMiddlewares,
	); err != nil {
		return err
	}

	var httpClientTransport = &http.Client{}
	if nethttpClient, err = httpproto.GetServiceNameHTTPGoClient(
		ctx,
		httpClientTransport,
		"http://"+nethttpAddr,
		nethttpmdlwr.ClientMiddlewares,
	); err != nil {
		return err
	}

	// we use other client to check gin server
	if nethttpClientForGin, err = httpproto.GetServiceNameHTTPGoClient(
		ctx,
		httpClientTransport,
		"http://"+ginAddr,
		nethttpmdlwr.ClientMiddlewares,
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
	items := []*common.ArrayItem{
		{Value: "a"}, {Value: "b"}, {Value: "c"},
	}
	var topLevelArrayResp *common.Array
	if topLevelArrayResp, err = client.TopLevelArray(ctx, &common.Array{Items: items}); err != nil {
		return fmt.Errorf("TopLevelArray failed: %w", err)
	}
	if diff := cmp.Diff(items, topLevelArrayResp.Items, cmpopts.IgnoreUnexported(*topLevelArrayResp.Items[0])); diff != "" {
		log.Println(diff)
	}
	var respMsgV3 *common.UpdateMessageRequest
	reqMsgV3 := &common.UpdateMessageRequest{
		MessageId: "123456",
		Message:   &common.MessageV2{Text: "Hi!"},
	}
	if respMsgV3, err = client.UpdateMessageV3(ctx, reqMsgV3); err != nil {
		return fmt.Errorf("UpdateMessageV3 failed: %w", err)
	}
	if diff := cmp.Diff(reqMsgV3.Message, respMsgV3.Message, cmpopts.IgnoreUnexported(*respMsgV3, *respMsgV3.Message)); diff != "" {
		log.Println(diff)
	}

	return nil
}
