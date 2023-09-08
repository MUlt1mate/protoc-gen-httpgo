package proto

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"sync"
	"testing"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/middleware"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/somepackage"
)

type testCaseClient struct {
	name         string
	method       string
	uri          string
	request      interface{}
	response     interface{}
	responseErr  error
	requestBody  []byte
	mockResponse responseData
}

type testCaseServer struct {
	name           string
	method         string
	uri            string
	responseBody   []byte
	responseErr    error
	requestBody    []byte
	respStatusCode int
}

type requestData struct {
	method      string
	uri         string
	requestBody []byte
}

type responseData struct {
	body []byte
	code int
}

type readCloser struct {
	io.Reader
}

func (readCloser) Close() error { return nil }

func getReadCloser(r io.Reader) io.ReadCloser {
	return readCloser{r}
}

func TestHTTPGoClient(t *testing.T) {
	var (
		reqCh  = make(chan requestData)
		respCh = make(chan responseData)
	)
	mockServer := httptest.NewServer(http.HandlerFunc(getMockServer(reqCh, respCh)))
	defer mockServer.Close()
	var (
		client *ServiceNameHTTPGoClient
		err    error
		ctx    = context.Background()
	)
	if client, err = GetServiceNameHTTPGoClient(
		ctx,
		&fasthttp.Client{},
		mockServer.URL,
		middleware.ClientMiddlewares,
	); err != nil {
		t.Fatal(err)
	}

	tests := []testCaseClient{
		{
			name:        "RPCName Valid Request 1",
			method:      http.MethodPost,
			uri:         "/v1/test/test/1",
			request:     &InputMsgName{Int64Argument: 1, StringArgument: "test"},
			response:    &OutputMsgName{StringValue: "StringValue", IntValue: 2},
			responseErr: nil,
			requestBody: []byte(`{"int64Argument":1,"stringArgument":"test"}`),
			mockResponse: responseData{
				body: []byte(`{"intValue":2,"stringValue":"StringValue"}`),
				code: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			resp := &OutputMsgName{}
			go func(wg *sync.WaitGroup) {
				resp, err = client.RPCName(ctx, test.request.(*InputMsgName))
				wg.Done()
			}(wg)

			request := <-reqCh
			respCh <- test.mockResponse
			wg.Wait()
			compareResults(t, request, test, err)
			if !reflect.DeepEqual(test.response.(*OutputMsgName), resp) {
				t.Errorf("Expected response method '%v', but got '%v'", test.response, resp)
			}
		})
	}
	tests = []testCaseClient{
		{
			name:   "AllTypesTest Valid Request 1",
			method: http.MethodPost,
			uri:    "/v1/test/true/SECOND/1/2/3/4/5/6/7/8/9.100000/10/11/12.200000/string/bytes",
			request: &AllTypesMsg{
				BoolValue:     true,
				EnumValue:     Options_SECOND,
				Int32Value:    1,
				Sint32Value:   2,
				Uint32Value:   3,
				Int64Value:    4,
				Sint64Value:   5,
				Uint64Value:   6,
				Sfixed32Value: 7,
				Fixed32Value:  8,
				FloatValue:    9.1,
				Sfixed64Value: 10,
				Fixed64Value:  11,
				DoubleValue:   12.2,
				StringValue:   "string",
				BytesValue:    []byte("bytes"),
			},
			response: &AllTypesMsg{
				BoolValue:     true,
				EnumValue:     Options_SECOND,
				Int32Value:    2,
				Sint32Value:   3,
				Uint32Value:   4,
				Int64Value:    5,
				Sint64Value:   6,
				Uint64Value:   7,
				Sfixed32Value: 8,
				Fixed32Value:  9,
				FloatValue:    10.1,
				Sfixed64Value: 11,
				Fixed64Value:  12,
				DoubleValue:   13.2,
				StringValue:   "stringResp",
				BytesValue:    []byte("bytesResp"),
			},
			responseErr: nil,
			requestBody: []byte(`{"BoolValue":true,"EnumValue":1,"Int32Value":1,"Sint32Value":2,"Uint32Value":3,"Int64Value":4,"Sint64Value":5,"Uint64Value":6,"Sfixed32Value":7,"Fixed32Value":8,"FloatValue":9.1,"Sfixed64Value":10,"Fixed64Value":11,"DoubleValue":12.2,"StringValue":"string","BytesValue":"Ynl0ZXM="}`),
			mockResponse: responseData{
				body: []byte(`{"BoolValue":true,"EnumValue":1,"Int32Value":2,"Sint32Value":3,"Uint32Value":4,"Int64Value":5,"Sint64Value":6,"Uint64Value":7,"Sfixed32Value":8,"Fixed32Value":9,"FloatValue":10.1,"Sfixed64Value":11,"Fixed64Value":12,"DoubleValue":13.2,"StringValue":"stringResp","BytesValue":"Ynl0ZXNSZXNw"}`),
				code: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			resp := &AllTypesMsg{}
			go func(wg *sync.WaitGroup) {
				resp, err = client.AllTypesTest(ctx, test.request.(*AllTypesMsg))
				wg.Done()
			}(wg)

			request := <-reqCh
			respCh <- test.mockResponse
			wg.Wait()
			compareResults(t, request, test, err)
			if !reflect.DeepEqual(test.response.(*AllTypesMsg), resp) {
				t.Errorf("Expected response method '%v', but got '%v'", test.response, resp)
			}
		})
	}
}

type Handler struct {
}

var _ ServiceNameHTTPGoService = &Handler{}

func (h *Handler) CommonTypes(_ context.Context, _ *anypb.Any) (*emptypb.Empty, error) {
	panic("implement me")
}

func (h *Handler) Imports(_ context.Context, msg1 *somepackage.SomeCustomMsg1) (*somepackage.SomeCustomMsg2, error) {
	return &somepackage.SomeCustomMsg2{Val: msg1.Val}, nil
}

func (h *Handler) SameInputAndOutput(_ context.Context, _ *InputMsgName) (*OutputMsgName, error) {
	panic("implement me")
}

func (h *Handler) RPCName(_ context.Context, request *InputMsgName) (*OutputMsgName, error) {
	p := &OutputMsgName{
		StringValue: request.StringArgument,
		IntValue:    request.Int64Argument,
	}
	return p, nil
}

func (h *Handler) AllTypesTest(_ context.Context, msg *AllTypesMsg) (*AllTypesMsg, error) {
	p := &AllTypesMsg{
		BoolValue:     msg.BoolValue,
		EnumValue:     msg.EnumValue,
		Int32Value:    msg.Int32Value,
		Sint32Value:   msg.Sint32Value,
		Uint32Value:   msg.Uint32Value,
		Int64Value:    msg.Int64Value,
		Sint64Value:   msg.Sint64Value,
		Uint64Value:   msg.Uint64Value,
		Sfixed32Value: msg.Sfixed32Value,
		Fixed32Value:  msg.Fixed32Value,
		FloatValue:    msg.FloatValue,
		Sfixed64Value: msg.Sfixed64Value,
		Fixed64Value:  msg.Fixed64Value,
		DoubleValue:   msg.DoubleValue,
		StringValue:   msg.StringValue,
		BytesValue:    msg.BytesValue,
	}
	return p, nil
}

func TestHTTPGoServer(t *testing.T) {
	var (
		err     error
		ctx                              = context.Background()
		handler ServiceNameHTTPGoService = &Handler{}
		r                                = router.New()
	)
	if err = RegisterServiceNameHTTPGoServer(ctx, r, handler, middleware.ServerMiddlewares); err != nil {
		t.Fatal(err)
	}

	ln, err := net.Listen("tcp4", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		_ = fasthttp.Serve(ln, r.Handler)
	}()

	tests := []testCaseServer{
		{
			name:           "RPCName Valid Request 1",
			method:         http.MethodPost,
			uri:            "/v1/test/test/1",
			responseBody:   []byte(`{"stringValue":"test","intValue":1}`),
			responseErr:    nil,
			requestBody:    []byte(`{"int64Argument":1,"stringArgument":"test"}`),
			respStatusCode: http.StatusOK,
		},
		{
			name:           "imports plain",
			method:         http.MethodPost,
			uri:            "/v1/test/imports",
			responseBody:   []byte(`{}`),
			responseErr:    nil,
			requestBody:    []byte(`{}`),
			respStatusCode: http.StatusOK,
		},
		{
			name:           "query parameter",
			method:         http.MethodPost,
			uri:            "/v1/test/imports?val=test",
			responseBody:   []byte(`{"val":"test"}`),
			responseErr:    nil,
			requestBody:    []byte(`{}`),
			respStatusCode: http.StatusOK,
		},
	}
	var (
		resp       *http.Response
		requestURL *url.URL
		host       = "http://" + ln.Addr().String()
		body       []byte
		client     = http.Client{}
	)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if requestURL, err = requestURL.Parse(host + test.uri); err != nil {
				t.Fatal(err)
			}
			req := &http.Request{
				Method: test.method,
				URL:    requestURL,
				Header: http.Header{},
				Body:   getReadCloser(bytes.NewReader(test.requestBody)),
			}

			req.Header.Add("Content-Type", "application/json")
			if resp, err = client.Do(req); err != nil {
				t.Fatal(err)
			}
			if body, err = io.ReadAll(resp.Body); err != nil {
				t.Fatal(err)
			}
			if !errors.Is(test.responseErr, err) {
				t.Errorf("%s: Expected error method '%v', but got '%v'", test.name, test.responseErr, err)
			}

			if resp.StatusCode != test.respStatusCode {
				t.Errorf("%s: Expected status  '%d', but got '%d'", test.name, test.respStatusCode, resp.StatusCode)
			}

			if bytes.Compare(body, test.responseBody) != 0 {
				t.Errorf("%s: Expected responseBody body '%s', but got '%s'", test.name, string(test.responseBody), string(body))
			}
		})
	}
}

func compareResults(
	t *testing.T,
	request requestData,
	test testCaseClient,
	err error,
) {
	if request.uri != test.uri {
		t.Errorf("%s: Expected request URI '%s', but got '%s'", test.name, test.uri, request.uri)
	}

	if request.method != test.method {
		t.Errorf("%s: Expected request method '%s', but got '%s'", test.name, test.method, request.method)
	}

	if !errors.Is(test.responseErr, err) {
		t.Errorf("%s: Expected error method '%v', but got '%v'", test.name, test.responseErr, err)
	}

	if bytes.Compare(request.requestBody, test.requestBody) != 0 {
		t.Errorf("%s: Expected request body '%s', but got '%s'", test.name, string(test.requestBody), string(request.requestBody))
	}
}

func getMockServer(
	reqCh chan requestData,
	respCh chan responseData,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		reqCh <- requestData{
			method:      r.Method,
			uri:         r.URL.Path,
			requestBody: body,
		}
		res := <-respCh
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(res.code)
		_, _ = w.Write(res.body)
	}
}
