package fasthttp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type ClientMiddleware = func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error)

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
		req *fasthttp.Request,
		next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
	) (resp *fasthttp.Response, err error) {
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
			respCode = fasthttp.StatusGatewayTimeout
		} else if resp != nil {
			respCode = resp.StatusCode()
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
		req *fasthttp.Request,
		next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
	) (resp *fasthttp.Response, err error) {
		service, _ := ctx.Value("proto_service").(string)
		method, _ := ctx.Value("proto_method").(string)

		ctx, span := tracer.Start(ctx, service+"/"+method, trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()

		propagator.Inject(ctx, &HeaderCarrier{Request: req})
		span.SetAttributes(
			attribute.String("component", "http"),
			attribute.String("uri", string(req.URI().RequestURI())),
		)

		resp, err = next(ctx, req)
		if err != nil {
			span.SetStatus(codes.Error, "client request failed")
			span.RecordError(err)
			return resp, err
		}
		if resp.StatusCode() >= http.StatusBadRequest {
			span.SetStatus(codes.Error, http.StatusText(resp.StatusCode()))
			span.RecordError(fmt.Errorf("status %d", resp.StatusCode()))
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
		req *fasthttp.Request,
		next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
	) (resp *fasthttp.Response, err error) {
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
		if resp != nil && resp.StatusCode() != fasthttp.StatusOK {
			httpLogger.Warn("client non-ok response", append(attrs,
				slog.Int("status_code", resp.StatusCode()),
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
		req *fasthttp.Request,
		next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
	) (resp *fasthttp.Response, err error) {
		resp, err = next(ctx, req)
		if err == nil && resp != nil && resp.StatusCode() >= fasthttp.StatusBadRequest {
			return resp, fmt.Errorf("%w, code: %d, body: %s", errRequestFailed, resp.StatusCode(), string(resp.Body()))
		}
		return resp, err
	}
}

// TimeoutClientMiddleware sets timeout for outgoing requests.
func TimeoutClientMiddleware() ClientMiddleware {
	return func(
		ctx context.Context,
		req *fasthttp.Request,
		next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
	) (resp *fasthttp.Response, err error) {
		req.SetTimeout(time.Second * 5)
		return next(ctx, req)
	}
}
