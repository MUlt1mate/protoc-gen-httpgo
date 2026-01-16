package fasthttp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	serverExecutionTimeout = time.Second * 5
	ContextFastHTTPCtx     = "request"
)

type (
	respError struct {
		Error string
	}
	validator interface {
		Validate() error
	}
)

var (
	serviceName = "fasthttp example"

	errRequestFailed = errors.New("api request failed")
	errTimeoutBody   = `{"error":"timeout"}`
)

var ServerMiddlewares = []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error){
	LoggerServerMiddleware,
	ResponseServerMiddleware,
	HeadersServerMiddleware,
	TimeoutServerMiddleware,
	ValidationServerMiddleware,
}
var ClientMiddlewares = []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error){
	LoggerClientMiddleware,
	ErrorClientMiddleware,
	TimeoutClientMiddleware,
}

// LoggerServerMiddleware logs request and response for server
func LoggerServerMiddleware(
	ctx context.Context, req any,
	next func(ctx context.Context, req any) (resp any, err error),
) (resp any, err error) {
	log.Printf("%s: server request %s", serviceName, req)
	resp, err = next(ctx, req)
	log.Printf("%s: server response %s", serviceName, resp)
	return resp, err
}

// ResponseServerMiddleware formats response for server
func ResponseServerMiddleware(
	ctx context.Context, req any,
	next func(ctx context.Context, req any) (resp any, err error),
) (resp any, err error) {
	resp, err = next(ctx, req)
	if err != nil {
		resp = respError{Error: err.Error()}
	}
	return resp, err
}

// HeadersServerMiddleware checks and sets headers for server
func HeadersServerMiddleware(
	ctx context.Context, req any,
	next func(ctx context.Context, req any) (resp any, err error),
) (resp any, err error) {
	fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
	resp, err = next(ctx, req)
	if err == nil {
		fastCtx.SetStatusCode(fasthttp.StatusOK)
	} else {
		fastCtx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	return resp, err
}

// TimeoutServerMiddleware sets timeout for request
func TimeoutServerMiddleware(
	ctx context.Context, req any,
	next func(ctx context.Context, req any) (resp any, err error),
) (resp any, err error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, serverExecutionTimeout)
	defer cancel()
	var done = make(chan struct{})
	go func() {
		resp, err = next(ctx, req)
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
		fastCtx.SetStatusCode(fasthttp.StatusGatewayTimeout)
		_, _ = fastCtx.WriteString(errTimeoutBody)
		return resp, err
	case <-done:
		return resp, err
	}
}

// ValidationServerMiddleware validates request
func ValidationServerMiddleware(
	ctx context.Context, req any,
	next func(ctx context.Context, req any) (resp any, err error),
) (resp any, err error) {
	if validatorArg, ok := req.(validator); ok {
		if err = validatorArg.Validate(); err != nil {
			fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
			fastCtx.SetStatusCode(fasthttp.StatusBadRequest)
			return nil, err
		}
	}
	return next(ctx, req)
}

// LoggerClientMiddleware logs request and response for client
func LoggerClientMiddleware(
	ctx context.Context,
	req *fasthttp.Request,
	next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	log.Printf("%s: client sending request with path %s", serviceName, req.RequestURI())
	resp, err = next(ctx, req)
	if err != nil {
		log.Printf("%s: client got response with error %s", serviceName, err.Error())
		return resp, err
	}
	if resp != nil {
		log.Printf("%s: client got response with code %d, body %s", serviceName, resp.StatusCode(), string(resp.Body()))
	}
	return resp, err
}

// ErrorClientMiddleware checks http response code for error
func ErrorClientMiddleware(
	ctx context.Context,
	req *fasthttp.Request,
	next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	resp, err = next(ctx, req)
	if err == nil && resp.StatusCode() > fasthttp.StatusBadRequest {
		return resp, fmt.Errorf("%w, code: %d", errRequestFailed, resp.StatusCode())
	}
	return resp, err
}

// TimeoutClientMiddleware sets timeout for request
func TimeoutClientMiddleware(
	ctx context.Context,
	req *fasthttp.Request,
	next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	req.SetTimeout(time.Second * 1)
	resp, err = next(ctx, req)
	return resp, err
}
