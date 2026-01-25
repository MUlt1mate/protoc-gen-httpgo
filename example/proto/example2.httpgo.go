// source: example2.proto

package proto

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	router "github.com/fasthttp/router"
	fasthttp "github.com/valyala/fasthttp"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	strconv "strconv"
	strings "strings"
)

type ServiceName2HTTPGoService interface {
	Imports(context.Context, *SomeCustomMsg) (*SomeCustomMsg, error)
}

func RegisterServiceName2HTTPGoServer(
	_ context.Context,
	r *router.Router,
	h ServiceName2HTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresExample2(middlewares)

	r.POST("/v1/test/imports", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExample2ServiceName2ImportsSomeCustomMsg(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName2")
		ctx = context.WithValue(ctx, "proto_method", "Imports")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.Imports(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := protojson.Marshal(resp.(proto.Message))
		_, _ = fastctx.Write(respJson)
	})

	return nil
}

func buildExample2ServiceName2ImportsSomeCustomMsg(ctx *fasthttp.RequestCtx) (arg *SomeCustomMsg, err error) {
	arg = &SomeCustomMsg{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = protojson.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "val":
			arg.Val = value
		case "option":
			if SomeOptionsValue, optValueOk := SomeOptions_value[strings.ToUpper(value)]; optValueOk {
				arg.Option = SomeOptions(SomeOptionsValue)
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := SomeOptions_name[int32(intOptionValue)]; optIntValueOk {
						arg.Option = SomeOptions(intOptionValue)
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter option: %w", convErr)
					return
				}
			}
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

type SecondServiceName2HTTPGoService interface {
	Imports(context.Context, *SomeCustomMsg) (*SomeCustomMsg, error)
}

func RegisterSecondServiceName2HTTPGoServer(
	_ context.Context,
	r *router.Router,
	h SecondServiceName2HTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresExample2(middlewares)

	r.GET("/v1/test/imports", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExample2SecondServiceName2ImportsSomeCustomMsg(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "SecondServiceName2")
		ctx = context.WithValue(ctx, "proto_method", "Imports")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.Imports(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := protojson.Marshal(resp.(proto.Message))
		_, _ = fastctx.Write(respJson)
	})

	return nil
}

func buildExample2SecondServiceName2ImportsSomeCustomMsg(ctx *fasthttp.RequestCtx) (arg *SomeCustomMsg, err error) {
	arg = &SomeCustomMsg{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "val":
			arg.Val = value
		case "option":
			if SomeOptionsValue, optValueOk := SomeOptions_value[strings.ToUpper(value)]; optValueOk {
				arg.Option = SomeOptions(SomeOptionsValue)
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := SomeOptions_name[int32(intOptionValue)]; optIntValueOk {
						arg.Option = SomeOptions(intOptionValue)
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter option: %w", convErr)
					return
				}
			}
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func chainServerMiddlewaresExample2(
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
			return middlewares[0](ctx, req, getChainServerMiddlewareHandlerExample2(middlewares, 0, handler))
		}
	}
}

func getChainServerMiddlewareHandlerExample2(
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
	curr int,
	finalHandler func(ctx context.Context, req any) (resp any, err error),
) func(ctx context.Context, req any) (resp any, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req any) (resp any, err error) {
		return middlewares[curr+1](ctx, req, getChainServerMiddlewareHandlerExample2(middlewares, curr+1, finalHandler))
	}
}

var _ ServiceName2HTTPGoService = &ServiceName2HTTPGoClient{}

type ServiceName2HTTPGoClient struct {
	cl          *fasthttp.Client
	host        string
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error)
	middleware  func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error)
}

func GetServiceName2HTTPGoClient(
	_ context.Context,
	cl *fasthttp.Client,
	host string,
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error),
) (*ServiceName2HTTPGoClient, error) {
	return &ServiceName2HTTPGoClient{
		cl:          cl,
		host:        host,
		middlewares: middlewares,
		middleware:  chainClientMiddlewaresExample2(middlewares),
	}, nil
}

func (p *ServiceName2HTTPGoClient) Imports(ctx context.Context, request *SomeCustomMsg) (resp *SomeCustomMsg, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = protojson.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/test/imports%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName2")
	ctx = context.WithValue(ctx, "proto_method", "Imports")
	var handler = func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error) {
		resp = &fasthttp.Response{}
		err = p.cl.Do(req, resp)
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
	resp = &SomeCustomMsg{}
	var respBody = reqResp.Body()
	err = protojson.Unmarshal(respBody, resp)
	return resp, err
}

var _ SecondServiceName2HTTPGoService = &SecondServiceName2HTTPGoClient{}

type SecondServiceName2HTTPGoClient struct {
	cl          *fasthttp.Client
	host        string
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error)
	middleware  func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error)
}

func GetSecondServiceName2HTTPGoClient(
	_ context.Context,
	cl *fasthttp.Client,
	host string,
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error),
) (*SecondServiceName2HTTPGoClient, error) {
	return &SecondServiceName2HTTPGoClient{
		cl:          cl,
		host:        host,
		middlewares: middlewares,
		middleware:  chainClientMiddlewaresExample2(middlewares),
	}, nil
}

func (p *SecondServiceName2HTTPGoClient) Imports(ctx context.Context, request *SomeCustomMsg) (resp *SomeCustomMsg, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{
		"val=%s",
		"option=%s",
	}
	var values = []any{
		request.Val,
		request.Option,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	req.SetRequestURI(fmt.Sprintf("%s/v1/test/imports%s", p.host, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "SecondServiceName2")
	ctx = context.WithValue(ctx, "proto_method", "Imports")
	var handler = func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error) {
		resp = &fasthttp.Response{}
		err = p.cl.Do(req, resp)
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
	resp = &SomeCustomMsg{}
	var respBody = reqResp.Body()
	err = protojson.Unmarshal(respBody, resp)
	return resp, err
}

func chainClientMiddlewaresExample2(
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error),
) func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error) {
			return middlewares[0](ctx, req, getChainClientMiddlewareHandlerExample2(middlewares, 0, handler))
		}
	}
}

func getChainClientMiddlewareHandlerExample2(
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error),
	curr int,
	finalHandler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error) {
		return middlewares[curr+1](ctx, req, getChainClientMiddlewareHandlerExample2(middlewares, curr+1, finalHandler))
	}
}
