package proto

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/valyala/fasthttp"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/middleware"
)

type testCase struct {
	name         string
	method       string
	uri          string
	request      interface{}
	response     interface{}
	responseErr  error
	requestBody  []byte
	mockResponse responseData
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

func TestMockHTTPServer(t *testing.T) {
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

	tests := []testCase{
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
	tests = []testCase{
		{
			name:   "AllTypesTest Valid Request 1",
			method: http.MethodPost,
			uri:    "/v1/test/true/SECOND/1/2/3/4/5/6/7/8/9/10/11/12/string/bytes",
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
				FloatValue:    9.0,
				Sfixed64Value: 10,
				Fixed64Value:  11,
				DoubleValue:   12.0,
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
				FloatValue:    10,
				Sfixed64Value: 11,
				Fixed64Value:  12,
				DoubleValue:   13,
				StringValue:   "stringResp",
				BytesValue:    []byte("bytesResp"),
			},
			responseErr: nil,
			requestBody: []byte(`{"BoolValue":true,"EnumValue":1,"Int32Value":1,"Sint32Value":2,"Uint32Value":3,"Int64Value":4,"Sint64Value":5,"Uint64Value":6,"Sfixed32Value":7,"Fixed32Value":8,"FloatValue":9,"Sfixed64Value":10,"Fixed64Value":11,"DoubleValue":12,"StringValue":"string","BytesValue":"Ynl0ZXM="}`),
			mockResponse: responseData{
				body: []byte(`{"BoolValue":true,"EnumValue":1,"Int32Value":2,"Sint32Value":3,"Uint32Value":4,"Int64Value":5,"Sint64Value":6,"Uint64Value":7,"Sfixed32Value":8,"Fixed32Value":9,"FloatValue":10,"Sfixed64Value":11,"Fixed64Value":12,"DoubleValue":13,"StringValue":"stringResp","BytesValue":"Ynl0ZXNSZXNw"}`),
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

func compareResults(
	t *testing.T,
	request requestData,
	test testCase,
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
