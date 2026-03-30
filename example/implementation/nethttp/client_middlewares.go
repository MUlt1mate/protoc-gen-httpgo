package nethttp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type ClientMiddleware = func(ctx context.Context, req *http.Request, handler func(ctx context.Context, req *http.Request) (resp *http.Response, err error)) (resp *http.Response, err error)

var errRequestFailed = errors.New("api request failed")

// MonitoringClientMiddleware tracks outbound request metrics with Prometheus.
func MonitoringClientMiddleware(reg prometheus.Registerer) ClientMiddleware {
	labels := []string{"service", "service_method"}
	clientHandledTotal := registerOrGet(reg, prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http_client", Name: "handled_total", Help: "The total number of handled requests by method and response code",
	}, append(labels, "resp_code")))
	clientHandledSeconds := registerOrGet(reg, prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_client", Name: "handled_seconds", Help: "Latency of handled requests by method and response code",
	}, labels))
	return func(
		ctx context.Context,
		req *http.Request,
		next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
	) (resp *http.Response, err error) {
		start := time.Now()
		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)

		resp, err = next(ctx, req)

		requestLabels := prometheus.Labels{
			"service":        service,
			"service_method": method,
		}
		clientHandledSeconds.With(requestLabels).Observe(time.Since(start).Seconds())

		var respCode int
		if ctx.Err() != nil {
			respCode = http.StatusGatewayTimeout
		} else if resp != nil {
			respCode = resp.StatusCode
		}
		requestLabels["resp_code"] = strconv.Itoa(respCode)
		clientHandledTotal.With(requestLabels).Inc()

		return resp, err
	}
}

// TracingClientMiddleware creates OpenTelemetry spans for outgoing requests.
func TracingClientMiddleware(tracer trace.Tracer) ClientMiddleware {
	propagator := newB3Propagator()
	return func(
		ctx context.Context,
		req *http.Request,
		next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
	) (resp *http.Response, err error) {
		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)

		ctx, span := tracer.Start(ctx, service+"/"+method, trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()

		propagator.Inject(ctx, &HeaderCarrier{Request: req})
		span.SetAttributes(
			attribute.String("component", "http"),
			attribute.String("uri", req.URL.RequestURI()),
		)

		resp, err = next(ctx, req)
		if err != nil {
			span.SetStatus(codes.Error, "client request failed")
			span.RecordError(err)
			return resp, err
		}
		if resp.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, http.StatusText(resp.StatusCode))
			span.RecordError(fmt.Errorf("status %d", resp.StatusCode))
			return resp, err
		}
		span.SetStatus(codes.Ok, "")
		return resp, err
	}
}

// LoggerClientMiddleware logs outbound requests and responses with slog.
func LoggerClientMiddleware(l *slog.Logger) ClientMiddleware {
	httpLogger := l.WithGroup("http")
	return func(
		ctx context.Context,
		req *http.Request,
		next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
	) (resp *http.Response, err error) {
		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)

		resp, err = next(ctx, req)

		attrs := []any{
			slog.String("service", service),
			slog.String("method", method),
		}
		if err != nil {
			httpLogger.Error("client request failed", append(attrs, slog.String("error", err.Error()))...)
			return resp, err
		}
		if resp != nil && resp.StatusCode != http.StatusOK {
			httpLogger.Warn("client non-ok response", append(attrs,
				slog.Int("status_code", resp.StatusCode),
			)...)
			return resp, err
		}
		httpLogger.Debug("client request completed", attrs...)
		return resp, err
	}
}

// ErrorClientMiddleware converts HTTP error responses to Go errors.
func ErrorClientMiddleware() ClientMiddleware {
	return func(
		ctx context.Context,
		req *http.Request,
		next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
	) (resp *http.Response, err error) {
		resp, err = next(ctx, req)
		if err == nil && resp != nil && resp.StatusCode >= http.StatusBadRequest {
			body, readErr := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			resp.Body = io.NopCloser(bytes.NewReader(body))
			if readErr != nil {
				return resp, fmt.Errorf("%w, code: %d, body read failed: %w", errRequestFailed, resp.StatusCode, readErr)
			}
			return resp, fmt.Errorf("%w, code: %d, body: %s", errRequestFailed, resp.StatusCode, string(body))
		}
		return resp, err
	}
}

// TimeoutClientMiddleware sets timeout for outgoing requests.
func TimeoutClientMiddleware() ClientMiddleware {
	return func(
		ctx context.Context,
		req *http.Request,
		next func(ctx context.Context, req *http.Request) (resp *http.Response, err error),
	) (resp *http.Response, err error) {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		req = req.WithContext(ctx)
		return next(ctx, req)
	}
}
