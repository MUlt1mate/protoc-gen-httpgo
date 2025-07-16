// source: proto/example2.proto

package proto

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	somepackage "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/somepackage"
	io "io"
	http "net/http"
	url "net/url"
)

type ServiceName2HTTPGoService interface {
	Imports(context.Context, *somepackage.SomeCustomMsg1) (*somepackage.SomeCustomMsg2, error)
}
type SecondServiceName2HTTPGoService interface {
	Imports(context.Context, *somepackage.SomeCustomMsg1) (*somepackage.SomeCustomMsg2, error)
}

var _ ServiceName2HTTPGoService = &ServiceName2HTTPGoClient{}

type ServiceName2HTTPGoClient struct {
	cl          *http.Client
	host        string
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error)
	middleware  func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error)
}

func GetServiceName2HTTPGoClient(
	_ context.Context,
	cl *http.Client,
	host string,
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) (*ServiceName2HTTPGoClient, error) {
	return &ServiceName2HTTPGoClient{
		cl:          cl,
		host:        host,
		middlewares: middlewares,
		middleware:  chainClientMiddlewaresExample2(middlewares),
	}, nil
}

func (p *ServiceName2HTTPGoClient) Imports(ctx context.Context, request *somepackage.SomeCustomMsg1) (resp *somepackage.SomeCustomMsg2, err error) {
	req := &http.Request{Header: make(http.Header)}
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	u, err := url.Parse(fmt.Sprintf("%s/v1/test/imports%s", p.host, queryArgs))
	if err != nil {
		return nil, err
	}
	req.URL = u
	req.Method = http.MethodPost
	var reqResp interface{}
	ctx = context.WithValue(ctx, "proto_service", "ServiceName2")
	ctx = context.WithValue(ctx, "proto_method", "Imports")
	var handler = func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		resp, err = p.cl.Do(req.(*http.Request))
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
	var respBody []byte
	if respBody, err = io.ReadAll(reqResp.(*http.Response).Body); err != nil {
		return nil, err
	}
	_ = reqResp.(*http.Response).Body.Close()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

var _ SecondServiceName2HTTPGoService = &SecondServiceName2HTTPGoClient{}

type SecondServiceName2HTTPGoClient struct {
	cl          *http.Client
	host        string
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error)
	middleware  func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error)
}

func GetSecondServiceName2HTTPGoClient(
	_ context.Context,
	cl *http.Client,
	host string,
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) (*SecondServiceName2HTTPGoClient, error) {
	return &SecondServiceName2HTTPGoClient{
		cl:          cl,
		host:        host,
		middlewares: middlewares,
		middleware:  chainClientMiddlewaresExample2(middlewares),
	}, nil
}

func (p *SecondServiceName2HTTPGoClient) Imports(ctx context.Context, request *somepackage.SomeCustomMsg1) (resp *somepackage.SomeCustomMsg2, err error) {
	req := &http.Request{Header: make(http.Header)}
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	u, err := url.Parse(fmt.Sprintf("%s/v1/test/imports%s", p.host, queryArgs))
	if err != nil {
		return nil, err
	}
	req.URL = u
	req.Method = http.MethodPost
	var reqResp interface{}
	ctx = context.WithValue(ctx, "proto_service", "SecondServiceName2")
	ctx = context.WithValue(ctx, "proto_method", "Imports")
	var handler = func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		resp, err = p.cl.Do(req.(*http.Request))
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
	var respBody []byte
	if respBody, err = io.ReadAll(reqResp.(*http.Response).Body); err != nil {
		return nil, err
	}
	_ = reqResp.(*http.Response).Body.Close()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func chainClientMiddlewaresExample2(
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
			return middlewares[0](ctx, req, getChainClientMiddlewareHandlerExample2(middlewares, 0, handler))
		}
	}
}

func getChainClientMiddlewareHandlerExample2(
	middlewares []func(ctx context.Context, req interface{}, handler func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
	curr int,
	finalHandler func(ctx context.Context, req interface{}) (resp interface{}, err error),
) func(ctx context.Context, req interface{}) (resp interface{}, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		return middlewares[curr+1](ctx, req, getChainClientMiddlewareHandlerExample2(middlewares, curr+1, finalHandler))
	}
}
