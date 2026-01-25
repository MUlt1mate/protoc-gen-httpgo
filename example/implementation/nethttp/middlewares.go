package nethttp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	serverExecutionTimeout = time.Second * 5
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
	serviceName      = "nethttp example"
	errRequestFailed = errors.New("api request failed")
)

var ServerMiddlewares = []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error){
	LoggerServerMiddleware,
	ResponseServerMiddleware,
	HeadersServerMiddleware,
	TimeoutServerMiddleware,
	ValidationServerMiddleware,
}
var ClientMiddlewares = []func(ctx context.Context, req *http.Request, handler func(ctx context.Context, req *http.Request) (resp *http.Response, err error)) (resp *http.Response, err error){
	LoggerClientMiddleware,
	ErrorClientMiddleware,
	TimeoutClientMiddleware,
}

// LoggerServerMiddleware logs request and response for server
func LoggerServerMiddleware(
	ctx context.Context, req any,
	next func(ctx context.Context, req any) (resp any, err error),
) (resp any, err error) {
	method := ctx.Value("proto_method").(string)
	// log.Printf("%s: %s: server request %s", serviceName, method, req)
	resp, err = next(ctx, req)
	if err != nil {
		log.Printf("%s: %s: server response %s", serviceName, method, resp)
	}
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
	writer, _ := ctx.Value("writer").(http.ResponseWriter)
	resp, err = next(ctx, req)
	if err == nil {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
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
		writer, _ := ctx.Value("writer").(http.ResponseWriter)
		writer.WriteHeader(http.StatusGatewayTimeout)
		return respError{Error: "timeout"}, nil
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
			writer, _ := ctx.Value("writer").(http.ResponseWriter)
			writer.WriteHeader(http.StatusBadRequest)
			return nil, err
		}
	}
	return next(ctx, req)
}

// LoggerClientMiddleware logs request and response for client
func LoggerClientMiddleware(
	ctx context.Context,
	req *http.Request,
	next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
) (resp *http.Response, err error) {
	method := ctx.Value("proto_method").(string)
	// log.Printf("%s: %s: client sending request with path %s", serviceName, method, req.URL.String())
	resp, err = next(ctx, req)
	if err != nil {
		log.Printf("%s: %s: client got response with error %s", serviceName, method, err.Error())
		return resp, err
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		var body []byte
		if body, err = io.ReadAll(resp.Body); err != nil {
			return resp, err
		}
		// Replace the body with a new reader after reading from the original
		resp.Body = io.NopCloser(bytes.NewBuffer(body))
		log.Printf("%s: %s: client got response with code %d, body %s", serviceName, method, resp.StatusCode, string(body))
	}
	return resp, err
}

// ErrorClientMiddleware checks http response code for error
func ErrorClientMiddleware(
	ctx context.Context,
	req *http.Request,
	next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
) (resp *http.Response, err error) {
	resp, err = next(ctx, req)
	if err == nil && resp.StatusCode > http.StatusBadRequest {
		return resp, fmt.Errorf("%w, code: %d", errRequestFailed, resp.StatusCode)
	}
	return resp, err
}

// TimeoutClientMiddleware sets timeout for request
func TimeoutClientMiddleware(
	ctx context.Context,
	req *http.Request,
	next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
) (resp *http.Response, err error) {
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*1)
	defer cancelFunc()
	req = req.WithContext(ctx)
	resp, err = next(ctx, req)
	return resp, err
}
