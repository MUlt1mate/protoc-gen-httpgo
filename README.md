# protoc-gen-httpgo

This is a protoc plugin that generates HTTP server and client code from proto files.

## Features

- Generation of both server and client code
- Provides multiple options for Marshaling/Unmarshaling:
    - Uses the native `encoding/json` by default
    - Optional usage of [easyjson](https://github.com/mailru/easyjson) for performance
- Utilizes google.api.http for defining HTTP paths
- Supports a wide range of data types in path parameters
- Supports middlewares

## Usage

### Generation

```bash  
protoc -I=. --httpgo_out=. --httpgo_opt=paths=source_relative example/proto/example.proto
```  

#### Parameters

| Name       | Values                 | Description                                                                                                                                  |
|------------|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------|
| paths      | source_relative,import | Inherited from protogen, see [docs](https://protobuf.dev/reference/go/go-generated/#invocation) for more details                             |
| marshaller | easyjson               | Specifies the data marshaling/unmarshaling package. Uses `encoding/json` by default. Can be set to easyjson with fallback to `encoding/json` |
| only       | server,client          | Use to generate either the server or client code exclusively                                                                                 |

Example of parameters usage:

```bash
protoc -I=. --httpgo_out=.  --httpgo_opt=paths=source_relative,marshaller=easyjson,only=server example/proto/example.proto
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
		handler proto.ServiceNameHTTPGoService = &implementation.Handler{}
		r                                      = router.New()
	)
	if err = proto.RegisterServiceNameHTTPGoServer(ctx, r, handler, serverMiddlewares); err != nil {
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
    client     *proto.ServiceNameHTTPGoClient
    httpClient = &fasthttp.Client{}
    host       = "http://localhost:8080"
  )
  if client, err = proto.GetServiceNameHTTPGoClient(ctx, httpClient, host, clientMiddlewares); err != nil {
    return err
  }
  // sending our request
  _, _ = client.RPCName(context.Background(), &proto.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
  return nil
}

```

#### Middlewares

You can define custom middlewares with specific arguments and return values.  
Pass a slice of middlewares to the constructor, and they will be invoked in the specified order.  
There are [middleware examples](https://github.com/MUlt1mate/protoc-gen-httpgo/blob/main/example/middleware/middlewares.go) for logs, timeout, headers, etc.   

```go
package implementation

import (
  "log"

  "github.com/valyala/fasthttp"
)

var ServerMiddlewares = []func(ctx *fasthttp.RequestCtx, next func(ctx *fasthttp.RequestCtx)){
  LoggerServerMiddleware,
}
var ClientMiddlewares = []func(req *fasthttp.Request, handler func(req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error){
  LoggerClientMiddleware,
}

func LoggerServerMiddleware(
        ctx *fasthttp.RequestCtx,
        next func(ctx *fasthttp.RequestCtx),
) {
  log.Println(serviceName, "server request", string(ctx.PostBody()))
  next(ctx)
  log.Println(serviceName, "server response", string(ctx.Response.Body()))
}

func LoggerClientMiddleware(
        req *fasthttp.Request,
        next func(req *fasthttp.Request) (resp *fasthttp.Response, err error),
) (resp *fasthttp.Response, err error) {
  log.Println(serviceName, "client request", string(req.RequestURI()))
  resp, err = next(req)
  log.Println(serviceName, "client response", string(resp.Body()))
  return resp, err
}

```

See [example](https://github.com/MUlt1mate/protoc-gen-httpgo/tree/main/example) for more details.

## TODO

- Implement comprehensive test cases
