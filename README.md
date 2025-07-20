# protoc-gen-httpgo

![workflow](https://github.com/MUlt1mate/protoc-gen-httpgo/actions/workflows/go.yml/badge.svg)
![go-report](https://goreportcard.com/badge/github.com/MUlt1mate/protoc-gen-httpgo)

This is a protoc plugin that generates HTTP server and client code from proto files.

## Features

- Generation of both server and client code
    - Supports [net/http](https://pkg.go.dev/net/http)
    - Supports [fasthttp](https://github.com/valyala/fasthttp)
- Provides multiple options for Marshaling/Unmarshaling:
    - Uses the native `encoding/json` by default
    - Optional usage of [easyjson](https://github.com/mailru/easyjson) for performance
    - Optional usage of [protojson](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson) for better
      protocol buffer support
- Utilizes google.api.http for defining HTTP paths (also can generate it)
- Supports a wide range of data types in path parameters
- Supports middlewares

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
protoc -I=. --httpgo_out=paths=source_relative,context=native:. example/proto/example.proto
```  

#### Parameters

| Name            | Values                  | Description                                                                                                      |
|-----------------|-------------------------|------------------------------------------------------------------------------------------------------------------|
| paths           | source_relative, import | Inherited from protogen, see [docs](https://protobuf.dev/reference/go/go-generated/#invocation) for more details |
| marshaller      | easyjson, protojson     | Specifies the data marshaling/unmarshaling package. Uses `encoding/json` by default.                             |
| only            | server, client          | Use to generate either the server or client code exclusively                                                     |
| autoURI         | false, true             | Create method URI if annotation is missing.                                                                      |
| bodylessMethods | GET;DELETE              | List of semicolon separated http methods that should not have a body.                                            |
| context         | native, fasthttp        | Type of ctx struct in server middlewares. fasthttp is default for backward compatibility. native is recommended  |
| library         | net/http, fasthttp      | Server library                                                                                                   |

Example of parameters usage:

```bash
protoc -I=. --httpgo_out=paths=source_relative,marshaller=easyjson,only=server,autoURI=true,context=native,library=fasthttp:. example/proto/example.proto
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
There
are [middleware examples](https://github.com/MUlt1mate/protoc-gen-httpgo/blob/main/example/middleware/middlewares.go)
for logs, timeout, headers, etc.

```go
package implementation

import (
	"context"
	"log"

	"github.com/valyala/fasthttp"
)

var ServerMiddlewares = []func(ctx context.Context, arg interface{}, handler func(ctx context.Context, arg interface{}) (resp interface{}, err error)) (resp interface{}, err error){
	LoggerServerMiddleware,
}
var ClientMiddlewares = []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error){
	LoggerClientMiddleware,
}

func LoggerServerMiddleware(
	ctx context.Context, arg interface{},
	next func(ctx context.Context, arg interface{}) (resp interface{}, err error),
) (resp interface{}, err error) {
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
	log.Println("client request", string(req.RequestURI()))
	resp, err = next(ctx, req)
	log.Println("client response", string(resp.Body()))
	return resp, err
}

```

See [example](https://github.com/MUlt1mate/protoc-gen-httpgo/tree/main/example) for more details.

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

## TODO

- Improve test cases
- Implement more web servers
    - native net/http
    - gin
    - chi
- File upload
- Full support
  for [httprule](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.api#httprule)
    - response_body supports for client natively, for server only in custom middleware
- dependabot
- buf
- verbose levels
- benchmark
- websocket
- production ready middleware example
- optionally ignore unknown query parameters
