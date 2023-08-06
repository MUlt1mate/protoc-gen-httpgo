package implementation

import (
	"log"

	"github.com/valyala/fasthttp"
)

var serviceName = "example"

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
