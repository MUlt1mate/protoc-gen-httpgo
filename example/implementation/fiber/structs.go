package fiber

import (
	"context"
	"errors"
	"log/slog"
	"time"

	v3 "github.com/gofiber/fiber/v3"
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

// HeaderCarrier adapts fiber.Ctx request headers to OTel propagation.TextMapCarrier.
type HeaderCarrier struct {
	Ctx v3.Ctx
}

func (c *HeaderCarrier) Get(key string) string {
	return c.Ctx.Get(key)
}

func (c *HeaderCarrier) Set(key, value string) {
	c.Ctx.Set(key, value)
}

func (c *HeaderCarrier) Keys() []string {
	headers := c.Ctx.GetReqHeaders()
	keys := make([]string, 0, len(headers))
	for k := range headers {
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
