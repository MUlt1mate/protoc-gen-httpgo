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
	expectedResponse    interface{}
	expectedResponseErr error
	expectedRequestBody []byte
	mockResponse        responseData
	methodName          string
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
			methodName:          "RPCName",
			expectedMethod:      http.MethodPost,
			expectedURI:         "/v1/test/test/1",
			request:             &proto.InputMsgName{Int64Argument: 1, StringArgument: "test"},
			expectedResponse:    &proto.OutputMsgName{StringValue: "StringValue", IntValue: 2},
			expectedResponseErr: nil,
			expectedRequestBody: []byte(`{"int64Argument":1,"stringArgument":"test"}`),
			mockResponse: responseData{
				body: []byte(`{"intValue":2,"stringValue":"StringValue"}`),
				code: http.StatusOK,
			},
		},
		{
			name:           "CheckRepeatedQuery valid",
			methodName:     "CheckRepeatedQuery",
			expectedMethod: http.MethodGet,
			request: &proto.RepeatedCheck{
				StringValue:      []string{"1", "2", "3"},
				StringValueQuery: []string{"a", "b", "c"},
			},
			expectedURI:         "/v1/repeated/1,2,3?StringValueQuery[]=a&StringValueQuery[]=b&StringValueQuery[]=c",
			expectedRequestBody: nil,
			mockResponse: responseData{
				body: []byte(`{"StringValue":["1","2","3"],"StringValueQuery":["a","b","c"]}`),
				code: http.StatusOK,
			},
			expectedResponse: &proto.RepeatedCheck{
				StringValue:      []string{"1", "2", "3"},
				StringValueQuery: []string{"a", "b", "c"},
			},
			expectedResponseErr: nil,
		},
		{
			name:           "CheckRepeatedPath valid",
			methodName:     "CheckRepeatedPath",
			expectedMethod: http.MethodGet,
			request: &proto.RepeatedCheck{
				BoolValue:        []bool{true, true},
				EnumValue:        []proto.Options{proto.Options_FIRST, proto.Options_SECOND},
				Int32Value:       []int32{2, 3},
				Sint32Value:      []int32{4, 5},
				Uint32Value:      []uint32{6, 7},
				Int64Value:       []int64{8, 9},
				Sint64Value:      []int64{10, 11},
				Uint64Value:      []uint64{12, 13},
				Sfixed32Value:    []int32{14, 15},
				Fixed32Value:     []uint32{16, 17},
				FloatValue:       []float32{18, 19},
				Sfixed64Value:    []int64{20, 21},
				Fixed64Value:     []uint64{22, 23},
				DoubleValue:      []float64{24, 25},
				StringValue:      []string{"a", "b"},
				BytesValue:       [][]byte{[]byte("c"), []byte("d")},
				StringValueQuery: []string{"e", "f"},
			},
			expectedURI:         "/v1/repeated/true,true/0,1/2,3/4,5/6,7/8,9/10,11/12,13/14,15/16,17/18,19/20,21/22,23/24,25/a,b/c,d/e,f",
			expectedRequestBody: nil,
			mockResponse: responseData{
				body: []byte(`{"BoolValue":[true,true],"EnumValue":[0,1],"Int32Value":[2,3],"Sint32Value":[4,5],"Uint32Value":[6,7],"Int64Value":[8,9],"Sint64Value":[10,11],"Uint64Value":[12,13],"Sfixed32Value":[14,15],"Fixed32Value":[16,17],"FloatValue":[18,19],"Sfixed64Value":[20,21],"Fixed64Value":[22,23],"DoubleValue":[24,25],"StringValue":["a","b"],"BytesValue":["Yyxk"],"StringValueQuery":["e","f"]}`),
				code: http.StatusOK,
			},
			expectedResponse: &proto.RepeatedCheck{
				BoolValue:        []bool{true, true},
				EnumValue:        []proto.Options{proto.Options_FIRST, proto.Options_SECOND},
				Int32Value:       []int32{2, 3},
				Sint32Value:      []int32{4, 5},
				Uint32Value:      []uint32{6, 7},
				Int64Value:       []int64{8, 9},
				Sint64Value:      []int64{10, 11},
				Uint64Value:      []uint64{12, 13},
				Sfixed32Value:    []int32{14, 15},
				Fixed32Value:     []uint32{16, 17},
				FloatValue:       []float32{18, 19},
				Sfixed64Value:    []int64{20, 21},
				Fixed64Value:     []uint64{22, 23},
				DoubleValue:      []float64{24, 25},
				StringValue:      []string{"a", "b"},
				BytesValue:       [][]byte{[]byte("c,d")}, // differs from request because delimiter being handled like []byte itself
				StringValueQuery: []string{"e", "f"},
			},
			expectedResponseErr: nil,
		},
		{
			name:           "CheckRepeatedPost valid",
			methodName:     "CheckRepeatedPost",
			expectedMethod: http.MethodPost,
			request: &proto.RepeatedCheck{
				BoolValue:        []bool{true, true},
				EnumValue:        []proto.Options{proto.Options_FIRST, proto.Options_SECOND},
				Int32Value:       []int32{2, 3},
				Sint32Value:      []int32{4, 5},
				Uint32Value:      []uint32{6, 7},
				Int64Value:       []int64{8, 9},
				Sint64Value:      []int64{10, 11},
				Uint64Value:      []uint64{12, 13},
				Sfixed32Value:    []int32{14, 15},
				Fixed32Value:     []uint32{16, 17},
				FloatValue:       []float32{18, 19},
				Sfixed64Value:    []int64{20, 21},
				Fixed64Value:     []uint64{22, 23},
				DoubleValue:      []float64{24, 25},
				StringValue:      []string{"a", "b"},
				BytesValue:       [][]byte{[]byte("c"), []byte("d")},
				StringValueQuery: []string{"e", "f"},
			},
			expectedURI:         "/v1/repeated/a,b",
			expectedRequestBody: []byte(`{"BoolValue":[true,true],"EnumValue":[0,1],"Int32Value":[2,3],"Sint32Value":[4,5],"Uint32Value":[6,7],"Int64Value":[8,9],"Sint64Value":[10,11],"Uint64Value":[12,13],"Sfixed32Value":[14,15],"Fixed32Value":[16,17],"FloatValue":[18,19],"Sfixed64Value":[20,21],"Fixed64Value":[22,23],"DoubleValue":[24,25],"StringValue":["a","b"],"BytesValue":["Yw==","ZA=="],"StringValueQuery":["e","f"]}`),
			mockResponse: responseData{
				body: []byte(`{"BoolValue":[true,true],"EnumValue":[0,1],"Int32Value":[2,3],"Sint32Value":[4,5],"Uint32Value":[6,7],"Int64Value":[8,9],"Sint64Value":[10,11],"Uint64Value":[12,13],"Sfixed32Value":[14,15],"Fixed32Value":[16,17],"FloatValue":[18,19],"Sfixed64Value":[20,21],"Fixed64Value":[22,23],"DoubleValue":[24,25],"StringValue":["a","b"],"BytesValue":["Yw==","ZA=="],"StringValueQuery":["e","f"]}`),
				code: http.StatusOK,
			},
			expectedResponse: &proto.RepeatedCheck{
				BoolValue:        []bool{true, true},
				EnumValue:        []proto.Options{proto.Options_FIRST, proto.Options_SECOND},
				Int32Value:       []int32{2, 3},
				Sint32Value:      []int32{4, 5},
				Uint32Value:      []uint32{6, 7},
				Int64Value:       []int64{8, 9},
				Sint64Value:      []int64{10, 11},
				Uint64Value:      []uint64{12, 13},
				Sfixed32Value:    []int32{14, 15},
				Fixed32Value:     []uint32{16, 17},
				FloatValue:       []float32{18, 19},
				Sfixed64Value:    []int64{20, 21},
				Fixed64Value:     []uint64{22, 23},
				DoubleValue:      []float64{24, 25},
				StringValue:      []string{"a", "b"},
				BytesValue:       [][]byte{[]byte("c"), []byte("d")},
				StringValueQuery: []string{"e", "f"},
			},
			expectedResponseErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			wg.Add(1)

			method := reflect.ValueOf(client).MethodByName(test.methodName)
			if !method.IsValid() {
				t.Fatalf("Method %s does not exist on client", test.methodName)
			}
			resp := reflect.New(reflect.TypeOf(test.expectedResponse).Elem()).Interface()

			go func(wg *sync.WaitGroup) {
				defer wg.Done()

				results := method.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(test.request)})
				if respErr, ok := results[1].Interface().(error); ok {
					err = respErr
				} else {
					err = nil
				}
				reflect.ValueOf(resp).Elem().Set(results[0].Elem())
			}(wg)

			request := <-reqCh
			respCh <- test.mockResponse
			wg.Wait()
			compareClientResults(t, request, test, err)
			if !reflect.DeepEqual(test.expectedResponse, resp) {
				t.Errorf("Expected response '%v', \nbut got '%v'", test.expectedResponse, resp)
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
		t.Errorf("%s: Expected request URI '%s', \nbut got '%s'", test.name, test.expectedURI, request.uri)
	}

	if request.method != test.expectedMethod {
		t.Errorf("%s: Expected request method '%s', \nbut got '%s'", test.name, test.expectedMethod, request.method)
	}

	if !errors.Is(test.expectedResponseErr, err) {
		t.Errorf("%s: Expected error method '%v', \nbut got '%v'", test.name, test.expectedResponseErr, err)
	}

	if !bytes.Equal(request.requestBody, test.expectedRequestBody) {
		t.Errorf("%s: Expected request body '%s', \nbut got '%s'", test.name, string(test.expectedRequestBody), string(request.requestBody))
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
			uri:         r.RequestURI,
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
			requestBody:            []byte(`{"int64Argument":1,"stringArgument":"test"}`),
			expectedResponseBody:   []byte(`{"stringValue":"test","intValue":1}`),
			expectedResponseErr:    nil,
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "imports plain",
			method:                 http.MethodPost,
			uri:                    "/v1/test/imports",
			requestBody:            []byte(`{}`),
			expectedResponseBody:   []byte(`{}`),
			expectedResponseErr:    nil,
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "query parameter",
			method:                 http.MethodPost,
			uri:                    "/v1/test/imports?val=test",
			requestBody:            []byte(`{}`),
			expectedResponseBody:   []byte(`{"val":"test"}`),
			expectedResponseErr:    nil,
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "get method",
			method:                 http.MethodGet,
			uri:                    "/v1/test/get?int64Argument=1&stringArgument=stringValue",
			requestBody:            nil,
			expectedResponseBody:   []byte(`{"stringValue":"stringValue","intValue":1}`),
			expectedResponseErr:    nil,
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "all repeated types in query",
			method:                 http.MethodGet,
			uri:                    "/v1/repeated/a,b?BoolValue[]=true&BoolValue[]=true&EnumValue[]=FIRST&EnumValue[]=1&Int32Value[]=2&Int32Value[]=3&Sint32Value[]=4&Sint32Value[]=5&Uint32Value[]=6&Uint32Value[]=7&Int64Value[]=8&Int64Value[]=9&Sint64Value[]=10&Sint64Value[]=11&Uint64Value[]=12&Uint64Value[]=13&Sfixed32Value[]=14&Sfixed32Value[]=15&Fixed32Value[]=16&Fixed32Value[]=17&FloatValue[]=18&FloatValue[]=19&Sfixed64Value[]=20&Sfixed64Value[]=21&Fixed64Value[]=22&Fixed64Value[]=23&DoubleValue[]=24&DoubleValue[]=25&BytesValue[]=c&BytesValue[]=d&StringValueQuery[]=e&StringValueQuery[]=f",
			requestBody:            nil,
			expectedResponseBody:   []byte(`{"BoolValue":[true,true],"EnumValue":[0,1],"Int32Value":[2,3],"Sint32Value":[4,5],"Uint32Value":[6,7],"Int64Value":[8,9],"Sint64Value":[10,11],"Uint64Value":[12,13],"Sfixed32Value":[14,15],"Fixed32Value":[16,17],"FloatValue":[18,19],"Sfixed64Value":[20,21],"Fixed64Value":[22,23],"DoubleValue":[24,25],"StringValue":["a","b"],"BytesValue":["Yw==","ZA=="],"StringValueQuery":["e","f"]}`),
			expectedResponseErr:    nil,
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "all repeated types in body",
			method:                 http.MethodPost,
			uri:                    "/v1/repeated/a,b",
			requestBody:            []byte(`{"BoolValue":[true,true],"EnumValue":[0,1],"Int32Value":[2,3],"Sint32Value":[4,5],"Uint32Value":[6,7],"Int64Value":[8,9],"Sint64Value":[10,11],"Uint64Value":[12,13],"Sfixed32Value":[14,15],"Fixed32Value":[16,17],"FloatValue":[18,19],"Sfixed64Value":[20,21],"Fixed64Value":[22,23],"DoubleValue":[24,25],"BytesValue":["Yw==","ZA=="],"StringValueQuery":["e","f"]}`),
			expectedResponseBody:   []byte(`{"BoolValue":[true,true],"EnumValue":[0,1],"Int32Value":[2,3],"Sint32Value":[4,5],"Uint32Value":[6,7],"Int64Value":[8,9],"Sint64Value":[10,11],"Uint64Value":[12,13],"Sfixed32Value":[14,15],"Fixed32Value":[16,17],"FloatValue":[18,19],"Sfixed64Value":[20,21],"Fixed64Value":[22,23],"DoubleValue":[24,25],"StringValue":["a","b"],"BytesValue":["Yw==","ZA=="],"StringValueQuery":["e","f"]}`),
			expectedResponseErr:    nil,
			expectedRespStatusCode: http.StatusOK,
		},
		{
			name:                   "all repeated types in path",
			method:                 http.MethodGet,
			uri:                    "/v1/repeated/t,true/FIRST,1/2,3/4,5/6,7/8,9/10,11/12,13/14,15/16,17/18,19/20,21/22,23/24,25/a,b/c,d/e,f",
			requestBody:            nil,
			expectedResponseBody:   []byte(`{"BoolValue":[true,true],"EnumValue":[0,1],"Int32Value":[2,3],"Sint32Value":[4,5],"Uint32Value":[6,7],"Int64Value":[8,9],"Sint64Value":[10,11],"Uint64Value":[12,13],"Sfixed32Value":[14,15],"Fixed32Value":[16,17],"FloatValue":[18,19],"Sfixed64Value":[20,21],"Fixed64Value":[22,23],"DoubleValue":[24,25],"StringValue":["a","b"],"BytesValue":["Yyxk"],"StringValueQuery":["e","f"]}`),
			expectedResponseErr:    nil,
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
				t.Errorf("%s: Expected error method '%v', \nbut got '%v'", test.name, test.expectedResponseErr, err)
			}

			if resp.StatusCode != test.expectedRespStatusCode {
				t.Errorf("%s: Expected status  '%d', \nbut got '%d'", test.name, test.expectedRespStatusCode, resp.StatusCode)
			}

			if !bytes.Equal(body, test.expectedResponseBody) {
				t.Errorf("%s: Expected responseBody body '%s', \nbut got '%s'", test.name, string(test.expectedResponseBody), string(body))
			}
		})
	}
}
