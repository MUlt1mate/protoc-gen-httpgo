// source: no_url.proto

package proto

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	router "github.com/fasthttp/router"
	fasthttp "github.com/valyala/fasthttp"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type NoURLHTTPGoService interface {
	MethodWithoutURLAnnotation(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

func RegisterNoURLHTTPGoServer(
	_ context.Context,
	r *router.Router,
	h NoURLHTTPGoService,
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) error {
	var middleware = chainServerMiddlewaresNourl(middlewares)

	r.POST("NoURL/MethodWithoutURLAnnotation", func(ctx *fasthttp.RequestCtx) {
		input, err := buildNourlNoURLMethodWithoutURLAnnotationEmpty(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		ctx.SetUserValue("proto_service", "NoURL")
		ctx.SetUserValue("proto_method", "MethodWithoutURLAnnotation")
		handler := func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			return h.MethodWithoutURLAnnotation(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	return nil
}

func buildNourlNoURLMethodWithoutURLAnnotationEmpty(ctx *fasthttp.RequestCtx) (arg *emptypb.Empty, err error) {
	arg = &emptypb.Empty{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	return arg, err
}

func chainServerMiddlewaresNourl(
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
			return middlewares[0](ctx, req, getChainServerMiddlewareHandlerNourl(middlewares, 0, handler))
		}
	}
}

func getChainServerMiddlewareHandlerNourl(
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
	curr int,
	finalHandler func(ctx context.Context, req interface{}) (resp interface{}, err error),
) func(ctx context.Context, req interface{}) (resp interface{}, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		return middlewares[curr+1](ctx, req, getChainServerMiddlewareHandlerNourl(middlewares, curr+1, finalHandler))
	}
}

var _ NoURLHTTPGoService = &NoURLHTTPGoClient{}

type NoURLHTTPGoClient struct {
	cl          *fasthttp.Client
	host        string
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error)
	middleware  func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error)
}

func GetNoURLHTTPGoClient(
	_ context.Context,
	cl *fasthttp.Client,
	host string,
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) (*NoURLHTTPGoClient, error) {
	return &NoURLHTTPGoClient{
		cl:          cl,
		host:        host,
		middlewares: middlewares,
		middleware:  chainClientMiddlewaresNourl(middlewares),
	}, nil
}

func (p *NoURLHTTPGoClient) MethodWithoutURLAnnotation(ctx context.Context, request *emptypb.Empty) (resp *emptypb.Empty, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("NoURL/MethodWithoutURLAnnotation%s", queryArgs))
	req.Header.SetMethod("POST")
	var reqResp interface{}
	ctx = context.WithValue(ctx, "proto_service", "NoURL")
	ctx = context.WithValue(ctx, "proto_method", "MethodWithoutURLAnnotation")
	var handler = func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		resp = &fasthttp.Response{}
		err = p.cl.Do(req.(*fasthttp.Request), resp.(*fasthttp.Response))
		return resp, err
	}
	if p.middleware == nil {
		if reqResp, err = handler(ctx, req); err != nil {
			return nil, err
		}
	} else {
		if reqResp, err = p.middleware(ctx, req, handler); err != nil {
			return nil, err
		}
	}
	resp = &emptypb.Empty{}
	var respBody = reqResp.(*fasthttp.Response).Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func chainClientMiddlewaresNourl(
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
			return middlewares[0](ctx, req, getChainClientMiddlewareHandlerNourl(middlewares, 0, handler))
		}
	}
}

func getChainClientMiddlewareHandlerNourl(
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
	curr int,
	finalHandler func(ctx context.Context, req interface{}) (resp interface{}, err error),
) func(ctx context.Context, req interface{}) (resp interface{}, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		return middlewares[curr+1](ctx, req, getChainClientMiddlewareHandlerNourl(middlewares, curr+1, finalHandler))
	}
}
