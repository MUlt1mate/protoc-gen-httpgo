package fiber

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	v3 "github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type ServerMiddleware = func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error)

// MonitoringServerMiddleware tracks request metrics with Prometheus.
func MonitoringServerMiddleware(reg prometheus.Registerer) ServerMiddleware {
	labels := []string{"service", "service_method"}
	serverHandledTotal := registerOrGet(reg, prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http_server", Name: "handled_total", Help: "The total number of handled requests by method and response code",
	}, append(labels, "resp_code")))
	serverHandledSeconds := registerOrGet(reg, prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_server", Name: "handled_seconds", Help: "Latency of handled requests by method and response code",
	}, labels))
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		start := time.Now()
		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)

		resp, err = next(ctx, req)

		requestLabels := prometheus.Labels{
			"service":        service,
			"service_method": method,
		}
		serverHandledSeconds.With(requestLabels).Observe(time.Since(start).Seconds())

		var respCode int
		switch {
		case ctx.Err() != nil:
			respCode = http.StatusGatewayTimeout
		case err != nil:
			var httpErr *HttpError
			if errors.As(err, &httpErr) {
				respCode = httpErr.Code
			} else {
				respCode = http.StatusInternalServerError
			}
		default:
			respCode = http.StatusOK
		}
		requestLabels["resp_code"] = strconv.Itoa(respCode)
		serverHandledTotal.With(requestLabels).Inc()

		return resp, err
	}
}

// TimeoutServerMiddleware sets a timeout for request processing.
func TimeoutServerMiddleware() ServerMiddleware {
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		ctx, cancel := context.WithTimeout(ctx, serverExecutionTimeout)
		defer cancel()

		type result struct {
			resp any
			err  error
		}
		done := make(chan result, 1)
		go func() {
			r, e := next(ctx, req)
			done <- result{resp: r, err: e}
		}()

		select {
		case <-ctx.Done():
			if fiberctx, ok := ctx.Value("request").(v3.Ctx); ok && fiberctx != nil {
				fiberctx.Status(http.StatusGatewayTimeout)
			}
			return respError{Error: "request timeout"}, nil
		case res := <-done:
			return res.resp, res.err
		}
	}
}

// RecoveryServerMiddleware recovers from panics.
// Sets status code and response body directly because panic recovery bypasses inner middlewares.
func RecoveryServerMiddleware() ServerMiddleware {
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("panic recovered", "panic", r, "stack", string(debug.Stack()))
				if fiberctx, ok := ctx.Value("request").(v3.Ctx); ok && fiberctx != nil {
					fiberctx.Status(http.StatusInternalServerError)
				}
				resp = respError{Error: "internal server error"}
				err = nil
			}
		}()
		return next(ctx, req)
	}
}

// ResponseServerMiddleware converts errors to response body.
func ResponseServerMiddleware() ServerMiddleware {
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		resp, err = next(ctx, req)
		if err != nil {
			resp = respError{Error: err.Error()}
			err = nil
		}
		return resp, err
	}
}

// HeadersServerMiddleware sets HTTP status code based on error type.
func HeadersServerMiddleware() ServerMiddleware {
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		fiberctx, _ := ctx.Value("request").(v3.Ctx)
		resp, err = next(ctx, req)
		if fiberctx == nil {
			return resp, err
		}
		if err == nil {
			fiberctx.Status(http.StatusOK)
			return resp, nil
		}
		var httpErr *HttpError
		if errors.As(err, &httpErr) {
			fiberctx.Status(httpErr.Code)
		} else {
			fiberctx.Status(http.StatusInternalServerError)
		}
		return resp, err
	}
}

// TracingServerMiddleware creates OpenTelemetry spans for incoming requests.
func TracingServerMiddleware(tracer trace.Tracer) ServerMiddleware {
	propagator := newB3Propagator()
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		if fiberctx, ok := ctx.Value("request").(v3.Ctx); ok && fiberctx != nil {
			ctx = propagator.Extract(ctx, &HeaderCarrier{Ctx: fiberctx})
		}

		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)

		ctx, span := tracer.Start(ctx, service+"/"+method, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		span.SetAttributes(attribute.String("component", "http"))
		if fiberctx, ok := ctx.Value("request").(v3.Ctx); ok && fiberctx != nil {
			span.SetAttributes(attribute.String("uri", fiberctx.OriginalURL()))
		}

		resp, err = next(ctx, req)
		if err != nil {
			var httpErr *HttpError
			if errors.As(err, &httpErr) && httpErr.Code >= http.StatusInternalServerError {
				span.SetStatus(codes.Error, "server error")
			}
			span.RecordError(err)
		} else {
			span.SetStatus(codes.Ok, "")
		}
		return resp, err
	}
}

// AuthServerMiddleware extracts Bearer token from Authorization header.
func AuthServerMiddleware() ServerMiddleware {
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		fiberctx, _ := ctx.Value("request").(v3.Ctx)
		if fiberctx == nil {
			return next(ctx, req)
		}
		auth := fiberctx.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			return nil, &HttpError{Code: http.StatusUnauthorized, Message: "unauthorized"}
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		ctx = context.WithValue(ctx, ContextAuthToken, token)
		return next(ctx, req)
	}
}

// LoggerServerMiddleware logs requests and responses with slog.
func LoggerServerMiddleware(l *slog.Logger) ServerMiddleware {
	httpLogger := l.WithGroup("http")
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)
		resp, err = next(ctx, req)

		attrs := []any{
			slog.String("service", service),
			slog.String("method", method),
		}
		if err != nil {
			httpLogger.Error("server request failed", append(attrs, slog.String("error", err.Error()))...)
		}
		return resp, err
	}
}

// ValidationServerMiddleware validates requests.
func ValidationServerMiddleware() ServerMiddleware {
	return func(
		ctx context.Context, req any,
		next func(ctx context.Context, req any) (resp any, err error),
	) (resp any, err error) {
		if v, ok := req.(validator); ok {
			if err = v.Validate(); err != nil {
				return nil, &HttpError{Code: http.StatusBadRequest, Message: err.Error()}
			}
		}
		return next(ctx, req)
	}
}
