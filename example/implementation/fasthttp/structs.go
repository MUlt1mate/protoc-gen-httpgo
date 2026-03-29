package fasthttp

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	serverExecutionTimeout = time.Second * 5
	ContextFastHTTPCtx     = "request"
	ContextAuthToken       = "auth_token"
)

type (
	respError struct {
		Error string
	}
	HttpError struct {
		Code    int
		Message string
	}
	validator interface {
		Validate() error
	}
)

func (e *HttpError) Error() string {
	return e.Message
}

// HeaderCarrier adapts fasthttp.Request headers to OTel propagation.TextMapCarrier.
type HeaderCarrier struct {
	Request *fasthttp.Request
}

func (c *HeaderCarrier) Get(key string) string {
	return string(c.Request.Header.Peek(key))
}

func (c *HeaderCarrier) Set(key, value string) {
	c.Request.Header.Set(key, value)
}

func (c *HeaderCarrier) Keys() []string {
	keys := make([]string, 0, c.Request.Header.Len())
	for key := range c.Request.Header.All() {
		keys = append(keys, string(key))
	}
	return keys
}

func registerOrGet[T prometheus.Collector](reg prometheus.Registerer, c T) T {
	if err := reg.Register(c); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			return are.ExistingCollector.(T)
		}
		panic(err)
	}
	return c
}

func newB3Propagator() propagation.TextMapPropagator {
	return b3.New(b3.WithInjectEncoding(b3.B3SingleHeader))
}

// GetServerMiddlewares returns the default server middleware chain.
func GetServerMiddlewares(
	l *slog.Logger, tracer trace.Tracer, reg prometheus.Registerer,
) []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
	return []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error){
		MonitoringServerMiddleware(reg),
		TimeoutServerMiddleware(),
		RecoveryServerMiddleware(),
		ResponseServerMiddleware(),
		HeadersServerMiddleware(),
		TracingServerMiddleware(tracer),
		LoggerServerMiddleware(l),
		ValidationServerMiddleware(),
	}
}

// GetServerSecureMiddlewares returns the server middleware chain with authentication.
func GetServerSecureMiddlewares(
	l *slog.Logger, tracer trace.Tracer, reg prometheus.Registerer,
) []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
	return []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error){
		MonitoringServerMiddleware(reg),
		TimeoutServerMiddleware(),
		RecoveryServerMiddleware(),
		ResponseServerMiddleware(),
		HeadersServerMiddleware(),
		TracingServerMiddleware(tracer),
		AuthServerMiddleware(),
		LoggerServerMiddleware(l),
		ValidationServerMiddleware(),
	}
}

// GetClientMiddlewares returns the default client middleware chain.
func GetClientMiddlewares(
	l *slog.Logger, tracer trace.Tracer, reg prometheus.Registerer,
) []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error) {
	return []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error){
		MonitoringClientMiddleware(reg),
		TracingClientMiddleware(tracer),
		LoggerClientMiddleware(l),
		ErrorClientMiddleware(),
		TimeoutClientMiddleware(),
	}
}
