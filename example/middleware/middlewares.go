package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

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
)

var ServerMiddlewares = []func(ctx *fasthttp.RequestCtx, handler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error)) (resp interface{}, err error){
	LoggerServerMiddleware,
	ResponseServerMiddleware,
	HeadersServerMiddleware,
}
var ClientMiddlewares = []func(req *fasthttp.Request, handler func(req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error){
	LoggerClientMiddleware,
	HeadersClientMiddleware,
	ErrorClientMiddleware,
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

// LoggerClientMiddleware logs request and response for client
func LoggerClientMiddleware(
	req *fasthttp.Request,
	next func(req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	log.Printf("%s: client sending request with path %s", serviceName, string(req.RequestURI()))
	resp, err = next(req)
	log.Printf("%s: client got response with code %d, body %s", serviceName, resp.StatusCode(), string(resp.Body()))
	return resp, err
}

// HeadersClientMiddleware checks and sets headers for client
func HeadersClientMiddleware(
	req *fasthttp.Request,
	next func(req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	jsonContentType := "application/json"
	req.Header.SetContentType(jsonContentType)
	resp, err = next(req)
	if err == nil && string(resp.Header.ContentType()) != jsonContentType {
		err = fmt.Errorf("incorrect response content type %s", string(resp.Header.ContentType()))
	}
	return resp, err
}

// ErrorClientMiddleware checks http response code for error
func ErrorClientMiddleware(
	req *fasthttp.Request,
	next func(req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	resp, err = next(req)
	if err == nil && resp.StatusCode() > http.StatusBadRequest {
		return resp, fmt.Errorf("%w, code: %d, body: %b", errorRequestFailed, resp.StatusCode(), resp.Body())
	}
	return resp, err
}
