# protoc-gen-httpgo

![workflow](https://github.com/MUlt1mate/protoc-gen-httpgo/actions/workflows/go.yml/badge.svg)
![go-report](https://goreportcard.com/badge/github.com/MUlt1mate/protoc-gen-httpgo)
[![Go Reference](https://pkg.go.dev/badge/github.com/MUlt1mate/protoc-gen-httpgo.svg)](https://pkg.go.dev/github.com/MUlt1mate/protoc-gen-httpgo)

**httpgo** is a protoc plugin that generates native HTTP server and client code from your proto files.  
It is a lightweight, high-performance alternative to [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

Choose httpgo if you want to:

- **Eliminate Proxy Overhead**: Serve HTTP directly from your application without a transcoding layer.
- **Dual-Stack Support**: Use HTTP and gRPC simultaneously with minimal performance impact.
- **Protobuf-First Design**: Define your API in Protobuf while leveraging the full flexibility of the HTTP ecosystem.

## Table of Contents

- [Performance](#performance)
- [Features](#features)
- [Usage](#usage)
    - [Installation](#installation)
    - [Definition](#definition)
    - [Generation](#generation)
        - [Parameters](#parameters)
    - [Implementation](#implementation)
        - [Server](#server)
        - [Client](#client)
        - [Middlewares](#middlewares)
        - [Error handling](#error-handling)
        - [Conventions](#conventions)
        - [Files](#files)
    - [Server-side streaming (SSE)](#server-side-streaming-sse)
- [TODO](#todo)

## Performance

Check [benchmark](./benchmark/README.md)

- 30% faster than grpc-gateway + grpc
- 95% reduction in memory overhead

## Features

- Generation of both server and client code
    - Supports [gin](https://github.com/gin-gonic/gin) (server only)
    - Supports [fiber](https://github.com/gofiber/fiber) (server only)
    - Supports [fasthttp](https://github.com/valyala/fasthttp)
    - Supports [net/http](https://pkg.go.dev/net/http)
- Provides multiple options for Marshaling/Unmarshaling:
    - Uses the native `encoding/json` by default
    - Optional usage of [protojson](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson) for better
      protocol buffer support
- Uses standard
  [google.api.http](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.api#httprule)
  definitions for path mapping
    - supports all HttpRule fields
- Supports automatic URI generation
- Supports a wide range of data types in path parameters
- Supports middlewares
- Supports multipart form with files
- Supports server-side streaming RPCs over Server-Sent Events (SSE)
- Zero additional dependencies in generated code

## Usage

### Installation

```bash
 go install github.com/MUlt1mate/protoc-gen-httpgo@latest
 ```

### Definition

Use proto with RPC to define methods

```protobuf
import "google/api/annotations.proto";

service TestService {
  rpc TestMethod (TestMessage) returns (TestMessage) {
    option (google.api.http) = {
      get: "/v1/test/{field1}"
    };
  }
}

message TestMessage {
  string field1 = 1;
  string field2 = 2;
}
```

### Generation

```bash  
protoc -I=. --httpgo_out=paths=source_relative:. example/proto/example.proto
```  

#### Parameters

| Name            | Values                        | Description                                                                                                      |
|-----------------|-------------------------------|------------------------------------------------------------------------------------------------------------------|
| paths           | source_relative, import       | Inherited from protogen, see [docs](https://protobuf.dev/reference/go/go-generated/#invocation) for more details |
| marshaller      | json, protojson               | Specifies the data marshaling/unmarshaling package. Uses `encoding/json` by default.                             |
| only            | server, client                | Use to generate either the server or client code exclusively                                                     |
| autoURI         | false, true                   | Create method URI if annotation is missing.                                                                      |
| bodylessMethods | GET;DELETE                    | List of semicolon separated http methods that should not have a body.                                            |
| library         | gin, fiber, fasthttp, nethttp | Server library                                                                                                   |
| stream          | none, sse                     | Wire format for server-side streaming RPCs. `sse` emits Server-Sent Events handlers; requires `only=server`.     |

Example of parameters usage:

```bash
protoc -I=. --httpgo_out=paths=source_relative,only=server,autoURI=true,library=fasthttp:. example/proto/example.proto
```

The plugin will create an example.httpgo.go file with the following:

- `Register{ServiceName}HTTPGoServer` - function to register server handlers
- `{ServiceName}HTTPGoService` - interface with all client methods
- `Get{ServiceName}HTTPGoClient` - client constructor that implements the above interface

### Implementation

#### Server

```go
package main

import (
	"context"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func serverExample(ctx context.Context) (err error) {
	var (
		handler proto.TestServiceHTTPGoService = &implementation.Handler{}
		r                                      = router.New()
	)
	if err = proto.RegisterTestServiceHTTPGoServer(ctx, r, handler, serverMiddlewares); err != nil {
		return err
	}
	go func() { _ = fasthttp.ListenAndServe(":8080", r.Handler) }()
	return nil
}

```

#### Client

```go
package main

import (
	"context"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
	"github.com/valyala/fasthttp"
)

func clientExample(ctx context.Context) (err error) {
	var (
		client     *proto.TestServiceHTTPGoClient
		httpClient = &fasthttp.Client{}
		host       = "http://localhost:8080"
	)
	if client, err = proto.GetTestServiceHTTPGoClient(ctx, httpClient, host, clientMiddlewares); err != nil {
		return err
	}
	// sending our request
	_, _ = client.TestMethod(context.Background(), &proto.TestMessage{Field1: "value", Field2: "rand"})
	return nil
}

```

#### Middlewares

You can define custom middlewares with specific arguments and return values.
Pass a slice of middlewares to the constructor, and they will be invoked in the specified order.
There are ready-to-use examples with 9 server middlewares (monitoring, timeout, recovery, response, headers,
tracing, auth, logging, validation) across all four implementations:
[fasthttp](example/implementation/fasthttp/README.md),
[fiber](example/implementation/fiber/README.md),
[gin](example/implementation/gin/README.md),
[nethttp](example/implementation/nethttp/README.md).
Additionally, 5 client middlewares (monitoring, tracing, logging, error handling, timeout) are included in the
[fasthttp](example/implementation/fasthttp/README.md) and [nethttp](example/implementation/nethttp/README.md) examples;
the fiber and gin examples are server-only.

```go
package fasthttp

import (
	"context"
	"log"

	"github.com/valyala/fasthttp"
)

var ServerMiddlewares = []func(ctx context.Context, arg any, handler func(ctx context.Context, arg any) (resp any, err error)) (resp any, err error){
	LoggerServerMiddleware,
}
var ClientMiddlewares = []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error){
	LoggerClientMiddleware,
}

func LoggerServerMiddleware(
	ctx context.Context, arg any,
	next func(ctx context.Context, arg any) (resp any, err error),
) (resp any, err error) {
	log.Println("server request", arg)
	resp, err = next(ctx, arg)
	log.Println("server response", resp)
	return resp, err
}

func LoggerClientMiddleware(
	ctx context.Context,
	req *fasthttp.Request,
	next func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
	log.Println("client request", string(req.URL.String()))
	resp, err = next(ctx, req)
	log.Println("client response", string(resp.Body()))
	return resp, err
}

```

See [example](https://github.com/MUlt1mate/protoc-gen-httpgo/tree/main/example) for more details.

#### Error handling

The generated server code does not inspect the `error` returned by your handler: it marshals whatever value the handler
returned and writes it to the response body. If you want a non-nil error to produce an error response, handle it in a
middleware — check `err` after calling `next`, and write the desired payload (for example by replacing `resp` with an
error struct, or setting the response status code directly via the transport-specific context stored on `ctx`).

The example implementations in `ResponseServerMiddleware` does exactly this: after calling `next`, if `err` is
non-nil it replaces `resp` with an error struct (`respError{Error: err.Error()}`) and clears `err` so the generated
code marshals the error payload into the response body.  
Pair it with `HeadersServerMiddleware` if you also need to map errors to HTTP status codes.

For SSE handlers the response body is the event stream itself, so there is nothing for a middleware to overwrite after
the stream completes. Middlewares still see the handler's `err` after `next` returns and can use it for logging,
metrics, or propagating cancellation — but they cannot turn it into a structured error body.

#### Conventions

Golang protobuf generator can produce fields with different case:

```protobuf
message InputMsgName {
  int64 value = 1;
}
```

```go
package main

import "google.golang.org/protobuf/runtime/protoimpl"

type InputMsgName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

```

We defined **v**alue and got **V**alue. This works just fine, but keep in mind that server will only check for arguments
with proto names.

* /v1/test?value=1 - correct
* /v1/test?Value=1 - incorrect

#### Files

To send and receive files you need to define file field as a message with the following fields.
The message name can be anything — the plugin identifies file fields by their structure, not by name.

```protobuf
message Request {
  FileMsg document = 1;
  FileMsg anotherDocument = 2;
}

message FileMsg {
  bytes file = 1;
  string name = 2;
  map<string, string> headers = 3;
}
```

### Server-side streaming (SSE)

Mark an RPC as server-streaming in your proto and generate with `stream=sse`:

```protobuf
service TickerService {
  rpc Ticks (TickRequest) returns (stream Tick) {
    option (google.api.http) = {
      get: "/v1/ticks"
    };
  }
}
```

```bash
protoc -I=. --httpgo_out=paths=source_relative,only=server,stream=sse,library=nethttp:. ticker.proto
```

Streaming methods appear on the generated service interface with a callback instead of a single return value:

```go
Ticks(context.Context, *TickRequest, func (*Tick) error) error
```

The generated handler sets the SSE response headers:

- `Content-Type: text/event-stream`
- `Cache-Control: no-cache`
- `Connection: keep-alive`

and writes each message as a `data: <json>\n\n` frame and flush.   
Call `send` once per message from your implementation; return when the stream is complete or the client disconnects (
observable via the passed `context`).

Client-streaming and bidirectional-streaming RPCs are still skipped.

## TODO

- implement more web servers
- buf
- WS streaming
- optionally ignore unknown query parameters
