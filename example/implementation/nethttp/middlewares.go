package nethttp

import (
	"bytes"
	"context"
	"encoding/json"
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
	serviceName = "nethttp example"

	errRequestFailed = errors.New("api request failed")
	errTimeoutBody   = `{"error":"timeout"}`
)

var ServerMiddlewares = []func(ctx context.Context, arg any, handler func(ctx context.Context, arg any) (resp any, err error)) (resp any, err error){
	LoggerServerMiddleware,
	ResponseServerMiddleware,
	HeadersServerMiddleware,
	TimeoutServerMiddleware,
	ValidationServerMiddleware,
}
var ClientMiddlewares = []func(ctx context.Context, req *http.Request, handler func(ctx context.Context, req *http.Request) (resp *http.Response, err error)) (resp *http.Response, err error){
	LoggerClientMiddleware,
	HeadersClientMiddleware,
	ErrorClientMiddleware,
	TimeoutClientMiddleware,
}

// LoggerServerMiddleware logs request and response for server
func LoggerServerMiddleware(
	ctx context.Context, arg any,
	next func(ctx context.Context, arg any) (resp any, err error),
) (resp any, err error) {
	log.Printf("%s: server request %s", serviceName, arg)
	resp, err = next(ctx, arg)
	log.Printf("%s: server response %s", serviceName, resp)
	return resp, err
}

// ResponseServerMiddleware formats response for server
func ResponseServerMiddleware(
	ctx context.Context, arg any,
	next func(ctx context.Context, arg any) (resp any, err error),
) (resp any, err error) {
	var responseBody []byte
	resp, err = next(ctx, arg)
	if err != nil {
		resp = respError{Error: err.Error()}
	}
	responseBody, _ = json.Marshal(resp)
	writer, _ := ctx.Value("writer").(http.ResponseWriter)
	_, _ = writer.Write(responseBody)
	return resp, err
}

// HeadersServerMiddleware checks and sets headers for server
func HeadersServerMiddleware(
	ctx context.Context, arg any,
	next func(ctx context.Context, arg any) (resp any, err error),
) (resp any, err error) {
	writer, _ := ctx.Value("writer").(http.ResponseWriter)
	// request, _ := ctx.Value("request").(*http.Request)
	// jsonContentType := "application/json"
	// contentType := request.Header.Get("Content-Type")
	// if contentType != jsonContentType {
	// 	writer.WriteHeader(http.StatusBadRequest)
	// 	return nil, errors.New("incorrect content type")
	// }
	// writer.Header().Add("Content-Type", jsonContentType)
	resp, err = next(ctx, arg)
	if err == nil {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	return resp, err
}

// TimeoutServerMiddleware sets timeout for request
func TimeoutServerMiddleware(
	ctx context.Context, arg any,
	next func(ctx context.Context, arg any) (resp any, err error),
) (resp any, err error) {
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
		writer, _ := ctx.Value("writer").(http.ResponseWriter)
		writer.WriteHeader(http.StatusGatewayTimeout)
		writer.Header().Add("Content-Type", "application/json")
		_, _ = writer.Write([]byte(errTimeoutBody))
		return resp, err
	case <-done:
		return resp, err
	}
}

// ValidationServerMiddleware validates request
func ValidationServerMiddleware(
	ctx context.Context, arg any,
	next func(ctx context.Context, arg any) (resp any, err error),
) (resp any, err error) {
	if validatorArg, ok := arg.(validator); ok {
		if err = validatorArg.Validate(); err != nil {
			writer, _ := ctx.Value("writer").(http.ResponseWriter)
			writer.WriteHeader(http.StatusBadRequest)
			writer.Header().Add("Content-Type", "application/json")
			_, _ = writer.Write([]byte(err.Error()))
			return nil, err
		}
	}
	return next(ctx, arg)
}

// LoggerClientMiddleware logs request and response for client
func LoggerClientMiddleware(
	ctx context.Context,
	req *http.Request,
	next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
) (resp *http.Response, err error) {
	log.Printf("%s: client sending request with path %s", serviceName, req.URL.String())
	resp, err = next(ctx, req)
	if err != nil {
		log.Printf("%s: client got response with error %s", serviceName, err.Error())
		return resp, err
	}
	if resp != nil {
		var body []byte
		if body, err = io.ReadAll(resp.Body); err != nil {
			return resp, err
		}
		// Replace the body with a new reader after reading from the original
		resp.Body = io.NopCloser(bytes.NewBuffer(body))
		log.Printf("%s: client got response with code %d, body %s", serviceName, resp.StatusCode, string(body))
	}
	return resp, err
}

// HeadersClientMiddleware checks and sets headers for client
func HeadersClientMiddleware(
	ctx context.Context,
	req *http.Request,
	next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
) (resp *http.Response, err error) {
	// jsonContentType := "application/json"
	// req.(*http.Request).Header.Set("Content-Type", jsonContentType)
	resp, err = next(ctx, req)
	// if err == nil && resp.(*http.Response).Header.Get("Content-Type") != jsonContentType {
	// err = fmt.Errorf("incorrect response content type %s", resp.(*http.Response).Header.Get("Content-Type"))
	// }
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
