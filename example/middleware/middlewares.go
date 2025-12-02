package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

const (
	serverExecutionTimeout        = time.Second * 5
	ContextFastHTTPCtx     ctxKey = "fastCtx"
)

type (
	ctxKey    string
	respError struct {
		Error string
	}
	validator interface {
		Validate() error
	}
)

var (
	serviceName = "example"

	errRequestFailed = errors.New("api request failed")
	errTimeoutBody   = `{"error":"timeout"}`
)

var ServerMiddlewares = []func(ctx context.Context, arg interface{}, handler func(ctx context.Context, arg interface{}) (resp interface{}, err error)) (resp interface{}, err error){
	ContextServerMiddleware,
	LoggerServerMiddleware,
	ResponseServerMiddleware,
	HeadersServerMiddleware,
	TimeoutServerMiddleware,
	ValidationServerMiddleware,
}
var ClientMiddlewares = []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error){
	LoggerClientMiddleware,
	HeadersClientMiddleware,
	ErrorClientMiddleware,
	TimeoutClientMiddleware,
}

// ContextServerMiddleware wraps fasthttp.RequestCtx from original ctx for further usage
// This middleware should be executed first
func ContextServerMiddleware(
	ctx context.Context, arg interface{},
	next func(ctx context.Context, arg interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	if fastCtx, ok := ctx.(*fasthttp.RequestCtx); ok {
		// wrap fasthttp.RequestCtx from original ctx for further usage
		// Calls of ctx.Value will eventually call *fasthttp.RequestCtx.Value
		ctx = context.WithValue(fastCtx, ContextFastHTTPCtx, fastCtx)
	} else {
		// We expect this middleware to be executed first
		// Fail instantly, so we can resolve this quick during development
		panic(fmt.Sprintf("incorrect ctx type, expected *fasthttp.RequestCtx, got %T", ctx))
	}
	return next(ctx, arg)
}

// LoggerServerMiddleware logs request and response for server
func LoggerServerMiddleware(
	ctx context.Context, arg interface{},
	next func(ctx context.Context, arg interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	log.Println(serviceName, "server request", arg)
	resp, err = next(ctx, arg)
	log.Println(serviceName, "server response", resp)
	return resp, err
}

// ResponseServerMiddleware formats response for server
func ResponseServerMiddleware(
	ctx context.Context, arg interface{},
	next func(ctx context.Context, arg interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	var responseBody []byte
	resp, err = next(ctx, arg)
	if err != nil {
		resp = respError{Error: err.Error()}
	}
	if ejMarsh, ok := resp.(easyjson.Marshaler); ok {
		responseBody, _ = easyjson.Marshal(ejMarsh)
	} else {
		responseBody, _ = json.Marshal(resp)
	}
	fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
	_, _ = fastCtx.Write(responseBody)
	return resp, err
}

// HeadersServerMiddleware checks and sets headers for server
func HeadersServerMiddleware(
	ctx context.Context, arg interface{},
	next func(ctx context.Context, arg interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
	jsonContentType := "application/json"
	// contentType := string(fastCtx.Request.Header.ContentType())
	// if contentType != jsonContentType {
	// 	fastCtx.SetStatusCode(fasthttp.StatusBadRequest)
	// 	return nil, errors.New("incorrect content type")
	// }
	fastCtx.SetContentType(jsonContentType)
	resp, err = next(ctx, arg)
	if err == nil {
		fastCtx.SetStatusCode(fasthttp.StatusOK)
	} else {
		fastCtx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	return resp, err
}

// TimeoutServerMiddleware sets timeout for request
func TimeoutServerMiddleware(
	ctx context.Context, arg interface{},
	next func(ctx context.Context, arg interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, serverExecutionTimeout)
	defer cancel()
	var done = make(chan struct{})
	go func() {
		resp, err = next(ctx, arg)
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
		fastCtx.SetStatusCode(fasthttp.StatusGatewayTimeout)
		fastCtx.SetContentType("application/json")
		_, _ = fastCtx.WriteString(errTimeoutBody)
		return resp, err
	case <-done:
		return resp, err
	}
}

// ValidationServerMiddleware validates request
func ValidationServerMiddleware(
	ctx context.Context, arg interface{},
	next func(ctx context.Context, arg interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	if validatorArg, ok := arg.(validator); ok {
		if err = validatorArg.Validate(); err != nil {
			fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
			fastCtx.SetStatusCode(fasthttp.StatusBadRequest)
			fastCtx.SetContentType("application/json")
			_, _ = fastCtx.WriteString(err.Error())
			return nil, err
		}
	}
	return next(ctx, arg)
}

// LoggerClientMiddleware logs request and response for client
func LoggerClientMiddleware(
	ctx context.Context,
	req interface{},
	next func(ctx context.Context, req interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	log.Printf("%s: client sending request with path %s", serviceName, req.(*fasthttp.Request).RequestURI())
	resp, err = next(ctx, req)
	if err != nil {
		log.Printf("%s: client got response with error %s", serviceName, err.Error())
		return resp, err
	}
	if resp != nil {
		respTyped := resp.(*fasthttp.Response)
		log.Printf("%s: client got response with code %d, body %s", serviceName, respTyped.StatusCode(), string(respTyped.Body()))
	}
	return resp, err
}

// HeadersClientMiddleware checks and sets headers for client
func HeadersClientMiddleware(
	ctx context.Context,
	req interface{},
	next func(ctx context.Context, req interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	// jsonContentType := "application/json"
	// req.(*fasthttp.Request).Header.SetContentType(jsonContentType)
	resp, err = next(ctx, req)
	// if err == nil && string(resp.(*fasthttp.Response).Header.ContentType()) != jsonContentType {
	// 	err = fmt.Errorf("incorrect response content type %s", string(resp.(*fasthttp.Response).Header.ContentType()))
	// }
	return resp, err
}

// ErrorClientMiddleware checks http response code for error
func ErrorClientMiddleware(
	ctx context.Context,
	req interface{},
	next func(ctx context.Context, req interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	resp, err = next(ctx, req)
	if err == nil && resp.(*fasthttp.Response).StatusCode() > fasthttp.StatusBadRequest {
		return resp, fmt.Errorf("%w, code: %d", errRequestFailed, resp.(*fasthttp.Response).StatusCode())
	}
	return resp, err
}

// TimeoutClientMiddleware sets timeout for request
func TimeoutClientMiddleware(
	ctx context.Context,
	req interface{},
	next func(ctx context.Context, req interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
	req.(*fasthttp.Request).SetTimeout(time.Second * 1)
	resp, err = next(ctx, req)
	return resp, err
}
