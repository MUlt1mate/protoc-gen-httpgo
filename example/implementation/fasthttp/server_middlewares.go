package fasthttp

import (
	"context"
	"errors"
	"log/slog"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
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
		if ctx.Err() != nil {
			respCode = fasthttp.StatusGatewayTimeout
		} else if fastCtx, ok := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx); ok && fastCtx != nil {
			respCode = fastCtx.Response.StatusCode()
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
			if fastCtx, ok := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx); ok && fastCtx != nil {
				fastCtx.SetStatusCode(fasthttp.StatusGatewayTimeout)
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
				if fastCtx, ok := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx); ok && fastCtx != nil {
					fastCtx.SetStatusCode(fasthttp.StatusInternalServerError)
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
		fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
		resp, err = next(ctx, req)
		if fastCtx == nil {
			return resp, err
		}
		if err == nil {
			fastCtx.SetStatusCode(fasthttp.StatusOK)
			return resp, nil
		}
		var httpErr *HttpError
		if errors.As(err, &httpErr) {
			fastCtx.SetStatusCode(httpErr.Code)
		} else {
			fastCtx.SetStatusCode(fasthttp.StatusInternalServerError)
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
		fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
		if fastCtx != nil {
			ctx = propagator.Extract(ctx, &HeaderCarrier{Request: &fastCtx.Request})
		}

		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)

		ctx, span := tracer.Start(ctx, service+"/"+method, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		span.SetAttributes(
			attribute.String("component", "http"),
		)
		if fastCtx != nil {
			span.SetAttributes(attribute.String("uri", string(fastCtx.Request.URI().RequestURI())))
		}

		resp, err = next(ctx, req)
		if err != nil {
			var httpErr *HttpError
			if errors.As(err, &httpErr) && httpErr.Code >= fasthttp.StatusInternalServerError {
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
		fastCtx, _ := ctx.Value(ContextFastHTTPCtx).(*fasthttp.RequestCtx)
		if fastCtx == nil {
			return next(ctx, req)
		}
		auth := string(fastCtx.Request.Header.Peek("Authorization"))
		if !strings.HasPrefix(auth, "Bearer ") {
			return nil, &HttpError{Code: fasthttp.StatusUnauthorized, Message: "unauthorized"}
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
				return nil, &HttpError{Code: fasthttp.StatusBadRequest, Message: err.Error()}
			}
		}
		return next(ctx, req)
	}
}
