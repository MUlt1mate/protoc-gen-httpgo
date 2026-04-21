// source: sse.proto

package proto

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	common "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	io "io"
	http "net/http"
	strconv "strconv"
)

type StreamingServiceHTTPGoService interface {
	Streaming(context.Context, *common.InputMsgName, func(*common.OutputMsgName) error) error
}

func RegisterStreamingServiceHTTPGoServer(
	_ context.Context,
	r *http.ServeMux,
	h StreamingServiceHTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresSse(middlewares)

	r.HandleFunc("POST /v1/sse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}
		input, err := buildSseStreamingServiceStreamingInputMsgName(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = w.Write(respJson)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "writer", w)
		ctx = context.WithValue(ctx, "request", r)
		ctx = context.WithValue(ctx, "proto_service", "StreamingService")
		ctx = context.WithValue(ctx, "proto_method", "Streaming")
		var data []byte
		send := func(msg *common.OutputMsgName) error {
			if data, err = json.Marshal(msg); err != nil {
				return err
			}
			if _, err = fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
				return err
			}
			flusher.Flush()
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
			_, _ = w.Write(respJson)
		}
	})

	return nil
}

func buildSseStreamingServiceStreamingInputMsgName(r *http.Request) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	var body []byte
	if body, err = io.ReadAll(r.Body); err != nil {
		return nil, err
	}
	_ = r.Body.Close()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for key, values := range r.URL.Query() {
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
