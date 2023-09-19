// source: proto/example2.proto

package proto

import (
	context "context"
	json "encoding/json"
	fmt "fmt"

	fasthttp "github.com/valyala/fasthttp"

	somepackage "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/somepackage"
)

type ServiceName2HTTPGoService interface {
	Imports(context.Context, *somepackage.SomeCustomMsg1) (*somepackage.SomeCustomMsg2, error)
}
type SecondServiceName2HTTPGoService interface {
	Imports(context.Context, *somepackage.SomeCustomMsg1) (*somepackage.SomeCustomMsg2, error)
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

func (p *ServiceName2HTTPGoClient) Imports(ctx context.Context, request *somepackage.SomeCustomMsg1) (resp *somepackage.SomeCustomMsg2, err error) {
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/imports"))
	req.Header.SetMethod("POST")
	var reqResp *fasthttp.Response
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
	resp = &somepackage.SomeCustomMsg2{}
	err = json.Unmarshal(reqResp.Body(), resp)
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

func (p *SecondServiceName2HTTPGoClient) Imports(ctx context.Context, request *somepackage.SomeCustomMsg1) (resp *somepackage.SomeCustomMsg2, err error) {
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/imports"))
	req.Header.SetMethod("POST")
	var reqResp *fasthttp.Response
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
	resp = &somepackage.SomeCustomMsg2{}
	err = json.Unmarshal(reqResp.Body(), resp)
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
