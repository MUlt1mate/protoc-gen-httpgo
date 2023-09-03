package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

type (
	respError struct {
		Error string
	}
)

var (
	serviceName = "example"

	errorRequestFailed = errors.New("api request failed")
	errTimeoutBody     = `{"error":"timeout"}`
)

var ServerMiddlewares = []func(ctx *fasthttp.RequestCtx, handler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error)) (resp interface{}, err error){
	LoggerServerMiddleware,
	ResponseServerMiddleware,
	HeadersServerMiddleware,
	TimeoutServerMiddleware,
}
var ClientMiddlewares = []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error){
	LoggerClientMiddleware,
	HeadersClientMiddleware,
	ErrorClientMiddleware,
	TimeoutClientMiddleware,
}

// LoggerServerMiddleware logs request and response for server
func LoggerServerMiddleware(
	ctx *fasthttp.RequestCtx,
	next func(ctx *fasthttp.RequestCtx) (resp interface{}, err error),
) (resp interface{}, err error) {
	log.Println(serviceName, "server request", string(ctx.PostBody()))
	resp, err = next(ctx)
	log.Println(serviceName, "server response", resp)
	return resp, err
}

// ResponseServerMiddleware format response for server
func ResponseServerMiddleware(
	ctx *fasthttp.RequestCtx,
	next func(ctx *fasthttp.RequestCtx) (resp interface{}, err error),
) (resp interface{}, err error) {
	var responseBody []byte
	resp, err = next(ctx)
	if err != nil {
		resp = respError{Error: err.Error()}
	}
	if _, ok := resp.(easyjson.Marshaler); ok {
		responseBody, err = easyjson.Marshal(resp.(easyjson.Marshaler))
	} else {
		responseBody, err = json.Marshal(resp)
	}
	_, _ = ctx.Write(responseBody)
	return resp, err
}

// HeadersServerMiddleware checks and sets headers for server
func HeadersServerMiddleware(
	ctx *fasthttp.RequestCtx,
	next func(ctx *fasthttp.RequestCtx) (resp interface{}, err error),
) (resp interface{}, err error) {
	jsonContentType := "application/json"
	contentType := string(ctx.Request.Header.ContentType())
	if contentType != jsonContentType {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return nil, errors.New("incorrect content type")
	}
	ctx.SetContentType(jsonContentType)
	resp, err = next(ctx)
	if err == nil {
		ctx.SetStatusCode(fasthttp.StatusOK)
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	return resp, err
}

// TimeoutServerMiddleware sets timeout for request
func TimeoutServerMiddleware(
	ctx *fasthttp.RequestCtx,
	next func(ctx *fasthttp.RequestCtx) (resp interface{}, err error),
) (resp interface{}, err error) {
	h := fasthttp.TimeoutWithCodeHandler(func(ctx *fasthttp.RequestCtx) {
		resp, err = next(ctx)
	}, time.Second*5, errTimeoutBody, http.StatusGatewayTimeout)
	h(ctx)

	return resp, err
}

// LoggerClientMiddleware logs request and response for client
func LoggerClientMiddleware(
	ctx context.Context,
	req *fasthttp.Request,
	next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	log.Printf("%s: client sending request with path %s", serviceName, string(req.RequestURI()))
	resp, err = next(ctx, req)
	log.Printf("%s: client got response with code %d, body %s", serviceName, resp.StatusCode(), string(resp.Body()))
	return resp, err
}

// HeadersClientMiddleware checks and sets headers for client
func HeadersClientMiddleware(
	ctx context.Context,
	req *fasthttp.Request,
	next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	jsonContentType := "application/json"
	req.Header.SetContentType(jsonContentType)
	resp, err = next(ctx, req)
	if err == nil && string(resp.Header.ContentType()) != jsonContentType {
		err = fmt.Errorf("incorrect response content type %s", string(resp.Header.ContentType()))
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
	if err == nil && resp.StatusCode() > http.StatusBadRequest {
		return resp, fmt.Errorf("%w, code: %d, body: %b", errorRequestFailed, resp.StatusCode(), resp.Body())
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
