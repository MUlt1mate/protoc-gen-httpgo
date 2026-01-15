// source: example2.proto

package proto

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	fasthttp "github.com/valyala/fasthttp"
	strings "strings"
)

type ServiceName2HTTPGoService interface {
	Imports(context.Context, *SomeCustomMsg) (*SomeCustomMsg, error)
}
type SecondServiceName2HTTPGoService interface {
	Imports(context.Context, *SomeCustomMsg) (*SomeCustomMsg, error)
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
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/test/imports%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
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
	err = json.Unmarshal(respBody, resp)
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
	err = json.Unmarshal(respBody, resp)
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
