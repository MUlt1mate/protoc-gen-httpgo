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
	tests := []testCase{
		{
			name:        "Valid Request 1",
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
	var (
		client *ServiceNameHTTPGoClient
		err    error
		ctx    = context.Background()
		reqCh  = make(chan requestData)
		respCh = make(chan responseData)
	)
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	defer mockServer.Close()
	if client, err = GetServiceNameHTTPGoClient(
		ctx,
		&fasthttp.Client{},
		mockServer.URL,
		middleware.ClientMiddlewares,
	); err != nil {
		t.Fatal(err)
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
			if request.uri != test.uri {
				t.Errorf("Expected request URI '%s', but got '%s'", test.uri, request.uri)
			}

			if request.method != test.method {
				t.Errorf("Expected request method '%s', but got '%s'", test.method, request.method)
			}
			if !reflect.DeepEqual(test.response.(*OutputMsgName), resp) {
				t.Errorf("Expected response method '%v', but got '%v'", test.response, resp)
			}
			if !errors.Is(test.responseErr, err) {
				t.Errorf("Expected error method '%v', but got '%v'", test.responseErr, err)
			}

			if bytes.Compare(request.requestBody, test.requestBody) != 0 {
				t.Errorf("Expected request body '%s', but got '%s'", string(test.requestBody), string(request.requestBody))
			}
		})
	}
}
