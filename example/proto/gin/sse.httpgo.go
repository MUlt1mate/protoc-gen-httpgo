// source: sse.proto

package proto

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	common "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	gin "github.com/gin-gonic/gin"
	strconv "strconv"
)

type StreamingServiceHTTPGoService interface {
	Streaming(context.Context, *common.InputMsgName, func(*common.OutputMsgName) error) error
}

func RegisterStreamingServiceHTTPGoServer(
	_ context.Context,
	r *gin.Engine,
	h StreamingServiceHTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresSse(middlewares)

	r.POST("/v1/sse", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "text/event-stream")
		ginctx.Header("Cache-Control", "no-cache")
		ginctx.Header("Connection", "keep-alive")
		input, err := buildSseStreamingServiceStreamingInputMsgName(ginctx)
		if err != nil {
			ginctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = ginctx.Writer.Write(respJson)
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
		ctx = context.WithValue(ctx, "proto_service", "StreamingService")
		ctx = context.WithValue(ctx, "proto_method", "Streaming")
		var data []byte
		send := func(msg *common.OutputMsgName) error {
			if data, err = json.Marshal(msg); err != nil {
				return err
			}
			if _, err = fmt.Fprintf(ginctx.Writer, "data: %s\n\n", data); err != nil {
				return err
			}
			ginctx.Writer.Flush()
			return nil
		}
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return nil, h.Streaming(ctx, input, send)
		}
		if middleware == nil {
			_, err = handler(ctx, input)
		} else {
			_, err = middleware(ctx, input, handler)
		}
		if err != nil {
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fmt.Fprintf(ginctx.Writer, "event: error\ndata: %s\n\n", respJson)
			ginctx.Writer.Flush()
		}
	})

	return nil
}

func buildSseStreamingServiceStreamingInputMsgName(ctx *gin.Context) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "int64Argument":
				arg.Int64Argument, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				}
			case "stringArgument":
				arg.StringArgument = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func chainServerMiddlewaresSse(
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
			return middlewares[0](ctx, req, getChainServerMiddlewareHandlerSse(middlewares, 0, handler))
		}
	}
}

func getChainServerMiddlewareHandlerSse(
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
	curr int,
	finalHandler func(ctx context.Context, req any) (resp any, err error),
) func(ctx context.Context, req any) (resp any, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req any) (resp any, err error) {
		return middlewares[curr+1](ctx, req, getChainServerMiddlewareHandlerSse(middlewares, curr+1, finalHandler))
	}
}
