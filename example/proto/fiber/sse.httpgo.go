// source: sse.proto

package proto

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	v3 "github.com/gofiber/fiber/v3"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
)

type StreamingServiceHTTPGoService interface {
	Streaming(context.Context, *common.InputMsgName, func(*common.OutputMsgName) error) error
}

func RegisterStreamingServiceHTTPGoServer(
	_ context.Context,
	r *v3.App,
	h StreamingServiceHTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresSse(middlewares)

	r.Post("/v1/sse", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "text/event-stream")
		fiberctx.Set("Cache-Control", "no-cache")
		fiberctx.Set("Connection", "keep-alive")
		input, err := buildSseStreamingServiceStreamingInputMsgName(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
		ctx = context.WithValue(ctx, "proto_service", "StreamingService")
		ctx = context.WithValue(ctx, "proto_method", "Streaming")
		fiberctx.Abandon()
		fiberctx.RequestCtx().SetBodyStreamWriter(func(w *bufio.Writer) {
			var data []byte
			send := func(msg *common.OutputMsgName) error {
				if data, err = json.Marshal(msg); err != nil {
					return err
				}
				if _, err = fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
					return err
				}
				_ = w.Flush()
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
				_, _ = fiberctx.Write(respJson)
			}
			fiberctx.ForceRelease()
		})
	})

	return nil
}

func buildSseStreamingServiceStreamingInputMsgName(ctx v3.Ctx) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
