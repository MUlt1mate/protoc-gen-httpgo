package gin

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	serviceName = "gin example"
)

var ServerMiddlewares = []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error){
	LoggerServerMiddleware,
	ResponseServerMiddleware,
	HeadersServerMiddleware,
	TimeoutServerMiddleware,
	ValidationServerMiddleware,
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
	ginctx, _ := ctx.Value("request").(*gin.Context)
	resp, err = next(ctx, req)
	if err == nil {
		ginctx.Writer.WriteHeader(http.StatusOK)
	} else {
		ginctx.Writer.WriteHeader(http.StatusInternalServerError)
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
		ginctx, _ := ctx.Value("request").(*gin.Context)
		ginctx.Writer.WriteHeader(http.StatusGatewayTimeout)
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
			ginctx, _ := ctx.Value("request").(*gin.Context)
			ginctx.Writer.WriteHeader(http.StatusBadRequest)
			return nil, err
		}
	}
	return next(ctx, req)
}
