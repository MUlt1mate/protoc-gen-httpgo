package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fasthttp/router"
	"github.com/gin-gonic/gin"
	v3 "github.com/gofiber/fiber/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	fasthttpmdlwr "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/fasthttp"
	fibermdlwr "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/fiber"
	ginmdlwr "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/gin"
	nethttpmdlwr "github.com/MUlt1mate/protoc-gen-httpgo/example/implementation/nethttp"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	fastproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fasthttp"
	fiberproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fiber"
	ginproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/gin"
	httpproto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/nethttp"
)

const (
	fasthttpAddr = "localhost:8080"
	nethttpAddr  = "localhost:8081"
	ginAddr      = "localhost:8082"
	fiberAddr    = "localhost:8083"
)

var (
	fastClient            httpproto.ServiceNameHTTPGoService
	nethttpClient         httpproto.ServiceNameHTTPGoService
	nethttpClientForGin   httpproto.ServiceNameHTTPGoService
	nethttpClientForFiber httpproto.ServiceNameHTTPGoService
)

func main() {
	log.SetFlags(log.Lshortfile + log.Ltime + log.Lmicroseconds)
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

	// wait to init
	time.Sleep(time.Millisecond * 100)

	if err = clientRunRequests(ctx, fastClient, "fasthttp"); err != nil {
		log.Println(fmt.Errorf("fasthttp client request failed: %w", err))
		os.Exit(1)
	}
	if err = clientRunRequests(ctx, nethttpClient, "nethttp"); err != nil {
		log.Println(fmt.Errorf("nethttp client request failed: %w", err))
		os.Exit(1)
	}
	if err = clientRunRequests(ctx, nethttpClientForGin, "gin"); err != nil {
		log.Println(fmt.Errorf("gin client request failed: %w", err))
		os.Exit(1)
	}
	if err = clientRunRequests(ctx, nethttpClientForFiber, "fiber"); err != nil {
		log.Println(fmt.Errorf("fiber client request failed: %w", err))
		os.Exit(1)
	}

	streamHosts := map[string]string{
		"fasthttp": fasthttpAddr,
		"nethttp":  nethttpAddr,
		"gin":      ginAddr,
		"fiber":    fiberAddr,
	}
	for library, host := range streamHosts {
		if err = clientRunStreaming(ctx, "http://"+host); err != nil {
			log.Fatalf("%s SSE request failed: %v", library, err)
		}
	}

	// f := make(chan bool)
	// <-f
}

func serverInit(ctx context.Context) (err error) {
	gin.SetMode(gin.ReleaseMode)
	var (
		logger = slog.Default()
		tracer = otel.Tracer("example")
		reg    = prometheus.DefaultRegisterer

		_              fastproto.ServiceNameHTTPGoService      = &implementation.Handler{}
		_              fastproto.StreamingServiceHTTPGoService = &implementation.Handler{}
		handler                                                = &implementation.Handler{}
		fasthttpRouter                                         = router.New()
		rHttp                                                  = http.NewServeMux()
		ginRouter                                              = gin.New()
		fiberApp                                               = v3.New()
	)
	if err = fastproto.RegisterServiceNameHTTPGoServer(ctx, fasthttpRouter, handler, fasthttpmdlwr.GetServerMiddlewares(logger, tracer, reg)); err != nil {
		return err
	}
	if err = httpproto.RegisterServiceNameHTTPGoServer(ctx, rHttp, handler, nethttpmdlwr.GetServerMiddlewares(logger, tracer, reg)); err != nil {
		return err
	}
	/*
		Gin has its own middleware format, but with this one you can have transport independent handler
		with context.Context that can be populated with the same keys in both HTTP and GRPC middlewares.
		And of course if you have only HTTP you can use gin middleware format and pass nil to httpgo middlewares
	*/
	if err = ginproto.RegisterServiceNameHTTPGoServer(ctx, ginRouter, handler, ginmdlwr.GetServerMiddlewares(logger, tracer, reg)); err != nil {
		return err
	}
	if err = fiberproto.RegisterServiceNameHTTPGoServer(ctx, fiberApp, handler, fibermdlwr.GetServerMiddlewares(logger, tracer, reg)); err != nil {
		return err
	}

	if err = fastproto.RegisterStreamingServiceHTTPGoServer(ctx, fasthttpRouter, handler, fasthttpmdlwr.GetSSEServerMiddlewares(logger, tracer, reg)); err != nil {
		return err
	}
	if err = httpproto.RegisterStreamingServiceHTTPGoServer(ctx, rHttp, handler, nethttpmdlwr.GetSSEServerMiddlewares(logger, tracer, reg)); err != nil {
		return err
	}
	if err = ginproto.RegisterStreamingServiceHTTPGoServer(ctx, ginRouter, handler, ginmdlwr.GetSSEServerMiddlewares(logger, tracer, reg)); err != nil {
		return err
	}
	if err = fiberproto.RegisterStreamingServiceHTTPGoServer(ctx, fiberApp, handler, fibermdlwr.GetSSEServerMiddlewares(logger, tracer, reg)); err != nil {
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
	go func() {
		if err = fiberApp.Listen(fiberAddr, v3.ListenConfig{DisableStartupMessage: true}); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func clientInit(ctx context.Context) (err error) {
	var fasthttpClientTransport = &fasthttp.Client{}
	var (
		logger = slog.Default()
		tracer = otel.Tracer("example")
		reg    = prometheus.DefaultRegisterer
	)
	if fastClient, err = fastproto.GetServiceNameHTTPGoClient(
		ctx,
		fasthttpClientTransport,
		"http://"+fasthttpAddr,
		fasthttpmdlwr.GetClientMiddlewares(logger, tracer, reg),
	); err != nil {
		return err
	}

	var httpClientTransport = &http.Client{}
	if nethttpClient, err = httpproto.GetServiceNameHTTPGoClient(
		ctx,
		httpClientTransport,
		"http://"+nethttpAddr,
		nethttpmdlwr.GetClientMiddlewares(logger, tracer, reg),
	); err != nil {
		return err
	}

	// we use other client to check gin server
	if nethttpClientForGin, err = httpproto.GetServiceNameHTTPGoClient(
		ctx,
		httpClientTransport,
		"http://"+ginAddr,
		nethttpmdlwr.GetClientMiddlewares(logger, tracer, reg),
	); err != nil {
		return err
	}

	// we use other client to check fiber server
	if nethttpClientForFiber, err = httpproto.GetServiceNameHTTPGoClient(
		ctx,
		httpClientTransport,
		"http://"+fiberAddr,
		nethttpmdlwr.GetClientMiddlewares(logger, tracer, reg),
	); err != nil {
		return err
	}
	return nil
}

func clientRunRequests(ctx context.Context, client httpproto.ServiceNameHTTPGoService, library string) (err error) {
	if _, err = client.RPCName(ctx, &common.InputMsgName{Int64Argument: 999, StringArgument: "rand"}); err != nil {
		return fmt.Errorf("RPCName failed: %w", err)
	}

	var allTypesResp *common.AllTypesMsg
	if allTypesResp, err = client.AllTypesTest(ctx, &implementation.AllTypesMsg); err != nil {
		return fmt.Errorf("AllTypesTest failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.AllTypesMsg, allTypesResp, cmpopts.IgnoreUnexported(implementation.AllTypesMsg)); diff != "" {
		log.Println(library, diff)
	}
	var allNumberTypesResp *common.AllNumberTypesMsg
	if allNumberTypesResp, err = client.AllTypesMaxTest(ctx, &implementation.NumberTypesMaxMsg); err != nil {
		return fmt.Errorf("AllTypesTest failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.NumberTypesMaxMsg, allNumberTypesResp, cmpopts.IgnoreUnexported(implementation.NumberTypesMaxMsg)); diff != "" {
		log.Println(library, diff)
	}
	if allNumberTypesResp, err = client.AllTypesMaxQueryTest(ctx, &implementation.NumberTypesMaxMsg); err != nil {
		return fmt.Errorf("AllTypesTest failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.NumberTypesMaxMsg, allNumberTypesResp, cmpopts.IgnoreUnexported(implementation.NumberTypesMaxMsg)); diff != "" {
		log.Println(library, diff)
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
		log.Println(library, diff)
	}
	if allTextTypesResp, err = client.AllTextTypesPost(ctx, &implementation.AllTextTypesMsg); err != nil {
		return fmt.Errorf("AllTextTypesPost failed: %w", err)
	}
	if diff := cmp.Diff(&implementation.AllTextTypesMsg, allTextTypesResp, cmpopts.IgnoreUnexported(implementation.AllTextTypesMsg)); diff != "" {
		log.Println(library, diff)
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
		log.Println(library, diff)
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
		log.Println(library, diff)
	}

	return nil
}

func clientRunStreaming(ctx context.Context, host string) error {
	req := &common.InputMsgName{Int64Argument: 3, StringArgument: "tick"}
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, host+"/v1/sse", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "text/event-stream")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if got := resp.Header.Get("Content-Type"); !strings.HasPrefix(got, "text/event-stream") {
		return fmt.Errorf("expected Content-Type text/event-stream, got %q", got)
	}

	want := []*common.OutputMsgName{
		{IntValue: 0, StringValue: "tick"},
		{IntValue: 1, StringValue: "tick"},
		{IntValue: 2, StringValue: "tick"},
	}
	got, err := parseSSE(resp.Body, len(want))
	if err != nil {
		return fmt.Errorf("read SSE: %w", err)
	}
	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(common.OutputMsgName{})); diff != "" {
		return fmt.Errorf("stream mismatch: %s", diff)
	}
	return nil
}

func parseSSE(r io.Reader, expected int) ([]*common.OutputMsgName, error) {
	br := bufio.NewReader(r)
	msgs := make([]*common.OutputMsgName, 0, expected)
	for len(msgs) < expected {
		line, err := br.ReadString('\n')
		if err != nil {
			return msgs, err
		}
		line = strings.TrimRight(line, "\r\n")
		payload, ok := strings.CutPrefix(line, "data: ")
		if !ok {
			continue
		}
		msg := &common.OutputMsgName{}
		if err = json.Unmarshal([]byte(payload), msg); err != nil {
			return msgs, fmt.Errorf("unmarshal frame %q: %w", payload, err)
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}
