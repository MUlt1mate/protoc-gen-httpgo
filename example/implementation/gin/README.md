# Gin Middlewares

Go source: [server_middlewares.go](server_middlewares.go) · [structs.go](structs.go)

Gin supports **server-side middlewares only** (no client generation). Middlewares are executed as a chain. The first
middleware in the slice is the outermost (wraps all others). On the return path, execution order is reversed.

All middlewares use the **builder pattern**: each builder function captures dependencies (logger, tracer, registerer)
and returns a middleware function matching the generated code's signature.

Factory functions assemble the full chain:

```go
logger := slog.Default()
tracer := otel.Tracer("my-service")
reg := prometheus.DefaultRegisterer

serverMiddlewares := ginmdlwr.GetServerMiddlewares(logger, tracer, reg)
```

## Context Values

Set by generated code before the middleware chain is called:

| Key               | Type          | Description               |
|-------------------|---------------|---------------------------|
| `"request"`       | `*gin.Context`| The gin request context   |
| `"proto_service"` | `string`      | Protobuf service name     |
| `"proto_method"`  | `string`      | Protobuf method name      |

Set by middlewares:

| Key            | Type     | Set by               | Description                            |
|----------------|----------|----------------------|----------------------------------------|
| `"auth_token"` | `string` | AuthServerMiddleware | Bearer token from Authorization header |

## Types

| Type            | Description                                                                        |
|-----------------|------------------------------------------------------------------------------------|
| `respError`     | JSON response body with an `Error` string field                                    |
| `HttpError`     | Error carrying an HTTP status code, used to communicate status between middlewares |
| `validator`     | Interface with `Validate() error`, checked by ValidationServerMiddleware           |
| `HeaderCarrier` | Adapts `*http.Request` headers to OTel `propagation.TextMapCarrier`                |

## Server Middlewares

### Default chain (`GetServerMiddlewares`)

```
Request  ->  Monitoring -> Timeout -> Recovery -> Response -> Headers -> Tracing -> Logger -> Validation -> Handler
Response <-  Monitoring <- Timeout <- Recovery <- Response <- Headers <- Tracing <- Logger <- Validation <- Handler
```

### Secure chain (`GetServerSecureMiddlewares`)

Same as default but with Auth between Tracing and Logger:

```
... -> Tracing -> Auth -> Logger -> ...
```

### 1. MonitoringServerMiddleware (outermost)

Tracks request metrics with Prometheus:

- `http_server_handled_total` counter (labels: service, service_method, resp_code)
- `http_server_handled_seconds` histogram (labels: service, service_method)

Registers metrics at construction time using `registerOrGet` (returns existing collector on duplicate registration).
Status code is derived from `ctx.Err()` (504) or the error type (HttpError.Code / 500 / 200).

**Dependencies:** `prometheus.Registerer`.

### 2. TimeoutServerMiddleware

Sets a 5-second deadline on the request. Runs the rest of the chain in a goroutine and waits for either completion or
timeout.

- On timeout: writes HTTP 504 via `ginctx.Writer.WriteHeader`, returns `respError{"request timeout"}`
- Uses a buffered channel with a `result` struct to avoid data races

**Dependencies:** None. Must be outside Recovery so the goroutine's panics are caught by Recovery inside.

### 3. RecoveryServerMiddleware

Catches panics from any inner middleware or the handler using `defer/recover`. Logs the panic value and full stack
trace via `slog.Error`.

- On panic: writes HTTP 500 directly via `ginctx.Writer.WriteHeader`, returns `respError{"internal server error"}` with `err = nil`
- Sets status code and response body directly because panic recovery bypasses all inner middlewares (Headers, Response)

**Dependencies:** Must be inside TimeoutServerMiddleware so that `recover()` executes within the goroutine where the
panic occurs.

### 4. ResponseServerMiddleware

Converts errors returned by inner middlewares into a `respError` JSON body. Nils out the error after conversion.

- On error: replaces `resp` with `respError{err.Error()}`, sets `err = nil`
- On success: passes through unchanged

### 5. HeadersServerMiddleware

Sets the HTTP status code on the response based on the error type.

- `err == nil`: 200 OK
- `err` is `*HttpError`: uses `HttpError.Code`
- `err` is any other error: 500 Internal Server Error

**Dependencies:** Depends on `HttpError` returned by ValidationServerMiddleware, AuthServerMiddleware, and
RecoveryServerMiddleware.

### 6. TracingServerMiddleware

Creates OpenTelemetry spans for incoming requests using B3 propagation.

- Extracts trace context from incoming request headers via B3 propagator (reads `ginctx.Request`)
- Starts a server span with attributes: component, uri
- Records errors on span (only 5xx as Error status)

**Dependencies:** `trace.Tracer`.

### 7. AuthServerMiddleware (secure chain only)

Extracts the Bearer token from the `Authorization` header and stores it in context under `ContextAuthToken`.

- On missing or invalid Authorization header: returns `HttpError{401, "unauthorized"}`

**Dependencies:** Reads `*gin.Context` from context via `ginctx.GetHeader`. Returns `HttpError` which
HeadersServerMiddleware reads.

### 8. LoggerServerMiddleware

Structured logging with `slog` under the "http" group. Logs: service, method.

- On error: logs at Error level with error message

**Dependencies:** `*slog.Logger`.

### 9. ValidationServerMiddleware (innermost)

Checks if the request implements the `validator` interface and calls `Validate()`.

- On validation failure: returns `*HttpError{Code: 400, Message: ...}`

**Dependencies:** Returns `*HttpError` which HeadersServerMiddleware reads to set the correct 400 status code.
