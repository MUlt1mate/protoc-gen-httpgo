package gin

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	serverExecutionTimeout = time.Second * 5
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

// HeaderCarrier adapts http.Request headers to OTel propagation.TextMapCarrier.
type HeaderCarrier struct {
	Request *http.Request
}

func (c *HeaderCarrier) Get(key string) string {
	return c.Request.Header.Get(key)
}

func (c *HeaderCarrier) Set(key, value string) {
	c.Request.Header.Set(key, value)
}

func (c *HeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(c.Request.Header))
	for k := range c.Request.Header {
		keys = append(keys, k)
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
