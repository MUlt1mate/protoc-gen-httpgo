package implementation

import (
	"log"

	"github.com/valyala/fasthttp"
)

var serviceName = "example"

func LoggerMiddleware(ctx *fasthttp.RequestCtx, next func(ctx *fasthttp.RequestCtx)) {
	log.Println(serviceName, "request", string(ctx.PostBody()))
	next(ctx)
	log.Println(serviceName, "response", string(ctx.Response.Body()))
}
