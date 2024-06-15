package proto_test

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

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/middleware"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
)

type testCaseClient struct {
	name                string
	expectedMethod      string
	expectedURI         string
	request             interface{}
	exptectedResponse   interface{}
	expectedResponseErr error
	expectedRequestBody []byte
	mockResponse        responseData
}

type testCaseServer struct {
	name                   string
	method                 string
	uri                    string
	expectedResponseBody   []byte
	expectedResponseErr    error
	requestBody            []byte
	expectedRespStatusCode int
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
		client *proto.ServiceNameHTTPGoClient
		err    error
		ctx    = context.Background()
	)
	if client, err = proto.GetServiceNameHTTPGoClient(
		ctx,
		&fasthttp.Client{},
		mockServer.URL,
		middleware.ClientMiddlewares,
	); err != nil {
		t.Fatal(err)
	}

	tests := []testCaseClient{
		{
			name:                "RPCName Valid Request 1",
			expectedMethod:      http.MethodPost,
			expectedURI:         "/v1/test/test/1",
			request:             &proto.InputMsgName{Int64Argument: 1, StringArgument: "test"},
			exptectedResponse:   &proto.OutputMsgName{StringValue: "StringValue", IntValue: 2},
			expectedResponseErr: nil,
			expectedRequestBody: []byte(`{"int64Argument":1,"stringArgument":"test"}`),
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
			resp := &proto.OutputMsgName{}
			go func(wg *sync.WaitGroup) {
				resp, err = client.RPCName(ctx, test.request.(*proto.InputMsgName))
				wg.Done()
			}(wg)

			request := <-reqCh
			respCh <- test.mockResponse
			wg.Wait()
			compareClientResults(t, request, test, err)
			if !reflect.DeepEqual(test.exptectedResponse.(*proto.OutputMsgName), resp) {
				t.Errorf("Expected response method '%v', but got '%v'", test.exptectedResponse, resp)
			}
		})
	}
	tests = []testCaseClient{
		{
			name:           "AllTypesTest Valid Request 1",
			expectedMethod: http.MethodPost,
			expectedURI:    "/v1/test/true/SECOND/1/2/3/4/5/6/7/8/9.100000/10/11/12.200000/string/bytes",
			request: &proto.AllTypesMsg{
				BoolValue:        true,
				EnumValue:        proto.Options_SECOND,
				Int32Value:       1,
				Sint32Value:      2,
				Uint32Value:      3,
				Int64Value:       4,
				Sint64Value:      5,
				Uint64Value:      6,
				Sfixed32Value:    7,
				Fixed32Value:     8,
				FloatValue:       9.1,
				Sfixed64Value:    10,
				Fixed64Value:     11,
				DoubleValue:      12.2,
				StringValue:      "string",
				BytesValue:       []byte("bytes"),
				SliceStringValue: []string{"a", "b", "c"},
			},
			exptectedResponse: &proto.AllTypesMsg{
				BoolValue:        true,
				EnumValue:        proto.Options_SECOND,
				Int32Value:       2,
				Sint32Value:      3,
				Uint32Value:      4,
				Int64Value:       5,
				Sint64Value:      6,
				Uint64Value:      7,
				Sfixed32Value:    8,
				Fixed32Value:     9,
				FloatValue:       10.1,
				Sfixed64Value:    11,
				Fixed64Value:     12,
				DoubleValue:      13.2,
				StringValue:      "stringResp",
				BytesValue:       []byte("bytesResp"),
				SliceStringValue: []string{"a", "b", "c"},
			},
			expectedResponseErr: nil,
			expectedRequestBody: []byte(`{"BoolValue":true,"EnumValue":1,"Int32Value":1,"Sint32Value":2,"Uint32Value":3,"Int64Value":4,"Sint64Value":5,"Uint64Value":6,"Sfixed32Value":7,"Fixed32Value":8,"FloatValue":9.1,"Sfixed64Value":10,"Fixed64Value":11,"DoubleValue":12.2,"StringValue":"string","BytesValue":"Ynl0ZXM=","SliceStringValue":["a","b","c"]}`),
			mockResponse: responseData{
				body: []byte(`{"BoolValue":true,"EnumValue":1,"Int32Value":2,"Sint32Value":3,"Uint32Value":4,"Int64Value":5,"Sint64Value":6,"Uint64Value":7,"Sfixed32Value":8,"Fixed32Value":9,"FloatValue":10.1,"Sfixed64Value":11,"Fixed64Value":12,"DoubleValue":13.2,"StringValue":"stringResp","BytesValue":"Ynl0ZXNSZXNw","SliceStringValue":["a","b","c"]}`),
				code: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			resp := &proto.AllTypesMsg{}
			go func(wg *sync.WaitGroup) {
				resp, err = client.AllTypesTest(ctx, test.request.(*proto.AllTypesMsg))
				wg.Done()
			}(wg)

			request := <-reqCh
			respCh <- test.mockResponse
			wg.Wait()
			compareClientResults(t, request, test, err)
			if !reflect.DeepEqual(test.exptectedResponse.(*proto.AllTypesMsg), resp) {
				t.Errorf("Expected response method '%v', but got '%v'", test.exptectedResponse, resp)
			}
		})
	}
	tests = []testCaseClient{
		{
			name:           "RepeatedCheck valid",
			expectedMethod: http.MethodGet,
			request: &proto.RepeatedCheck{
				StringValueArg:   []string{"1", "2", "3"},
				StringValueQuery: []string{"a", "b", "c"},
			},
			expectedURI:         "/v1/repeated/1,2,3",
			expectedRequestBody: []byte(`{"stringValueArg":["1","2","3"],"stringValueQuery":["a","b","c"]}`),
			mockResponse: responseData{
				body: []byte(`{"stringValueArg":["1","2","3"],"stringValueQuery":["a","b","c"]}`),
				code: http.StatusOK,
			},
			exptectedResponse: &proto.RepeatedCheck{
				StringValueArg:   []string{"1", "2", "3"},
				StringValueQuery: []string{"a", "b", "c"},
			},
			expectedResponseErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			resp := &proto.RepeatedCheck{}
			go func(wg *sync.WaitGroup) {
				resp, err = client.CheckRepeated(ctx, test.request.(*proto.RepeatedCheck))
				wg.Done()
			}(wg)

			request := <-reqCh
			respCh <- test.mockResponse
			wg.Wait()
			compareClientResults(t, request, test, err)
			if !reflect.DeepEqual(test.exptectedResponse.(*proto.RepeatedCheck), resp) {
				t.Errorf("Expected response method '%v', but got '%v'", test.exptectedResponse, resp)
			}
		})
	}
}

func compareClientResults(
	t *testing.T,
	request requestData,
	test testCaseClient,
	err error,
) {
	if request.uri != test.expectedURI {
		t.Errorf("%s: Expected request URI '%s', but got '%s'", test.name, test.expectedURI, request.uri)
	}

	if request.method != test.expectedMethod {
		t.Errorf("%s: Expected request method '%s', but got '%s'", test.name, test.expectedMethod, request.method)
	}

	if !errors.Is(test.expectedResponseErr, err) {
		t.Errorf("%s: Expected error method '%v', but got '%v'", test.name, test.expectedResponseErr, err)
	}

	if !bytes.Equal(request.requestBody, test.expectedRequestBody) {
		t.Errorf("%s: Expected request body '%s', but got '%s'", test.name, string(test.expectedRequestBody), string(request.requestBody))
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

func TestHTTPGoServer(t *testing.T) {
	var (
		err     error
		ctx                                    = context.Background()
		handler proto.ServiceNameHTTPGoService = &implementation.Handler{}
		r                                      = router.New()
	)
	if err = proto.RegisterServiceNameHTTPGoServer(ctx, r, handler, middleware.ServerMiddlewares); err != nil {
		t.Fatal(err)
	}

	ln, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		_ = fasthttp.Serve(ln, r.Handler)
	}()

	tests := []testCaseServer{
		{
			name:                   "RPCName Valid Request 1",
			method:                 http.MethodPost,
			uri:                    "/v1/test/test/1",
			expectedResponseBody:   []byte(`{"stringValue":"test","intValue":1}`),
			expectedResponseErr:    nil,
			requestBody:            []byte(`{"int64Argument":1,"stringArgument":"test"}`),
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "imports plain",
			method:                 http.MethodPost,
			uri:                    "/v1/test/imports",
			expectedResponseBody:   []byte(`{}`),
			expectedResponseErr:    nil,
			requestBody:            []byte(`{}`),
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "query parameter",
			method:                 http.MethodPost,
			uri:                    "/v1/test/imports?val=test",
			expectedResponseBody:   []byte(`{"val":"test"}`),
			expectedResponseErr:    nil,
			requestBody:            []byte(`{}`),
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "get method",
			method:                 http.MethodGet,
			uri:                    "/v1/test/get?int64Argument=1&stringArgument=stringValue",
			expectedResponseBody:   []byte(`{"stringValue":"stringValue","intValue":1}`),
			expectedResponseErr:    nil,
			requestBody:            nil,
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "repeated",
			method:                 http.MethodGet,
			uri:                    "/v1/repeated/1,2,3",
			expectedResponseBody:   []byte(`{"stringValueArg":["1","2","3"],"stringValueQuery":["a","b","c"]}`),
			expectedResponseErr:    nil,
			requestBody:            []byte(`{"stringValueQuery":["a","b","c"]}`),
			expectedRespStatusCode: http.StatusOK,
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
			defer func() {
				if resp != nil && resp.Body != nil {
					_ = resp.Body.Close()
				}
			}()
			if resp, err = client.Do(req); err != nil {
				t.Fatal(err)
			}
			if body, err = io.ReadAll(resp.Body); err != nil {
				t.Fatal(err)
			}
			if !errors.Is(test.expectedResponseErr, err) {
				t.Errorf("%s: Expected error method '%v', but got '%v'", test.name, test.expectedResponseErr, err)
			}

			if resp.StatusCode != test.expectedRespStatusCode {
				t.Errorf("%s: Expected status  '%d', but got '%d'", test.name, test.expectedRespStatusCode, resp.StatusCode)
			}

			if !bytes.Equal(body, test.expectedResponseBody) {
				t.Errorf("%s: Expected responseBody body '%s', but got '%s'", test.name, string(test.expectedResponseBody), string(body))
			}
		})
	}
}
