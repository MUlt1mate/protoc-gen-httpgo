// source: example.proto

package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
)

type ServiceNameHTTPGoService interface {
	RPCName(context.Context, *common.InputMsgName) (*common.OutputMsgName, error)
	AllTypesTest(context.Context, *common.AllTypesMsg) (*common.AllTypesMsg, error)
	AllTextTypesPost(context.Context, *common.AllTextTypesMsg) (*common.AllTextTypesMsg, error)
	AllTextTypesGet(context.Context, *common.AllTextTypesMsg) (*common.AllTextTypesMsg, error)
	CommonTypes(context.Context, *anypb.Any) (*emptypb.Empty, error)
	SameInputAndOutput(context.Context, *common.InputMsgName) (*common.OutputMsgName, error)
	Optional(context.Context, *common.OptionalField) (*common.OptionalField, error)
	GetMethod(context.Context, *common.InputMsgName) (*common.OutputMsgName, error)
	CheckRepeatedPath(context.Context, *common.RepeatedCheck) (*common.RepeatedCheck, error)
	CheckRepeatedQuery(context.Context, *common.RepeatedCheck) (*common.RepeatedCheck, error)
	CheckRepeatedPost(context.Context, *common.RepeatedCheck) (*common.RepeatedCheck, error)
	EmptyGet(context.Context, *common.Empty) (*common.Empty, error)
	EmptyPost(context.Context, *common.Empty) (*common.Empty, error)
	OnlyStructInGet(context.Context, *common.OnlyStruct) (*common.Empty, error)
	MultipartForm(context.Context, *common.MultipartFormRequest) (*common.Empty, error)
	MultipartFormAllTypes(context.Context, *common.MultipartFormAllTypes) (*common.Empty, error)
	AllTypesMaxTest(context.Context, *common.AllNumberTypesMsg) (*common.AllNumberTypesMsg, error)
	AllTypesMaxQueryTest(context.Context, *common.AllNumberTypesMsg) (*common.AllNumberTypesMsg, error)
	GetMessage(context.Context, *common.GetMessageRequest) (*common.Message, error)
	GetMessageV2(context.Context, *common.GetMessageRequestV2) (*common.Message, error)
	UpdateMessage(context.Context, *common.UpdateMessageRequest) (*common.Message, error)
	UpdateMessageV2(context.Context, *common.MessageV2) (*common.MessageV2, error)
	GetMessageV3(context.Context, *common.GetMessageRequestV3) (*common.MessageV2, error)
	GetMessageV4(context.Context, *common.GetMessageRequestV3) (*common.MessageV2, error)
	TopLevelArray(context.Context, *common.Array) (*common.Array, error)
	UpdateMessageV3(context.Context, *common.UpdateMessageRequest) (*common.UpdateMessageRequest, error)
}

func RegisterServiceNameHTTPGoServer(
	_ context.Context,
	r *router.Router,
	h ServiceNameHTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresExample(middlewares)

	r.POST("/v1/RPCName/{stringArgument}/{int64Argument}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameRPCNameInputMsgName(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "RPCName")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.RPCName(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/test/{BoolValue}/{EnumValue}/{Int32Value}/{Sint32Value}/{Uint32Value}/{Int64Value}/{Sint64Value}/{Uint64Value}/{Sfixed32Value}/{Fixed32Value}/{FloatValue}/{Sfixed64Value}/{Fixed64Value}/{DoubleValue}/{StringValue}/{BytesValue}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameAllTypesTestAllTypesMsg(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "AllTypesTest")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.AllTypesTest(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/text/{String}/{RepeatedString}/{Bytes}/{RepeatedBytes}/{Enum}/{RepeatedEnum}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameAllTextTypesPostAllTextTypesMsg(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "AllTextTypesPost")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.AllTextTypesPost(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v2/text/{String}/{RepeatedString}/{Bytes}/{RepeatedBytes}/{Enum}/{RepeatedEnum}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameAllTextTypesGetAllTextTypesMsg(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "AllTextTypesGet")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.AllTextTypesGet(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/test/commonTypes", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameCommonTypesAny(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "CommonTypes")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.CommonTypes(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	// same types but different query, we need different query builder function
	r.POST("/v1/sameInputAndOutput/{stringArgument}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameSameInputAndOutputInputMsgName(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "SameInputAndOutput")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.SameInputAndOutput(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/test/optional", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameOptionalOptionalField(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "Optional")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.Optional(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v1/test/get", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameGetMethodInputMsgName(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "GetMethod")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.GetMethod(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v1/repeated/{BoolValue}/{EnumValue}/{Int32Value}/{Sint32Value}/{Uint32Value}/{Int64Value}/{Sint64Value}/{Uint64Value}/{Sfixed32Value}/{Fixed32Value}/{FloatValue}/{Sfixed64Value}/{Fixed64Value}/{DoubleValue}/{StringValue}/{BytesValue}/{StringValueQuery}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameCheckRepeatedPathRepeatedCheck(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "CheckRepeatedPath")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.CheckRepeatedPath(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v2/repeated/{StringValue}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "CheckRepeatedQuery")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.CheckRepeatedQuery(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v3/repeated/{StringValue}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameCheckRepeatedPostRepeatedCheck(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "CheckRepeatedPost")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.CheckRepeatedPost(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v1/emptyGet", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameEmptyGetEmpty(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "EmptyGet")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.EmptyGet(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/emptyPost", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameEmptyPostEmpty(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "EmptyPost")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.EmptyPost(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/onlyStruct", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameOnlyStructInGetOnlyStruct(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "OnlyStructInGet")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.OnlyStructInGet(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/multipart", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameMultipartFormMultipartFormRequest(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "MultipartForm")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.MultipartForm(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/multipartall", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameMultipartFormAllTypesMultipartFormAllTypes(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "MultipartFormAllTypes")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.MultipartFormAllTypes(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v1/max/{Int32Value}/{Uint32Value}/{Int64Value}/{Uint64Value}/{FloatValue}/{DoubleValue}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameAllTypesMaxTestAllNumberTypesMsg(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "AllTypesMaxTest")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.AllTypesMaxTest(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v1/maxquery", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameAllTypesMaxQueryTestAllNumberTypesMsg(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "AllTypesMaxQueryTest")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.AllTypesMaxQueryTest(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	// http rule checks
	// v1/{name=messages/*}
	r.GET("/v1/messages/{name}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameGetMessageGetMessageRequest(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "GetMessage")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.GetMessage(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v2/messages/{message_id}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameGetMessageV2GetMessageRequestV2(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "GetMessageV2")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.GetMessageV2(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.PATCH("/v1/messages/{message_id}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameUpdateMessageUpdateMessageRequest(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "UpdateMessage")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.UpdateMessage(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.PATCH("/v2/messages/{message_id}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameUpdateMessageV2MessageV2(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "UpdateMessageV2")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.UpdateMessageV2(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v3/messages/{message_id}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameGetMessageV3GetMessageRequestV3(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "GetMessageV3")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.GetMessageV3(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v3/users/{user_id}/messages/{message_id}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameGetMessageV3GetMessageRequestV3(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "GetMessageV3")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.GetMessageV3(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.GET("/v4/messages/base/{message_id:*}", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameGetMessageV4GetMessageRequestV3(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "GetMessageV4")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.GetMessageV4(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		respJson, _ := json.Marshal(resp)
		_, _ = fastctx.Write(respJson)
	})

	r.POST("/v1/array", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameTopLevelArrayArray(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "TopLevelArray")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.TopLevelArray(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		if typedResp, ok := resp.(*common.Array); ok {
			respJson, _ := json.Marshal(typedResp.Items)
			_, _ = fastctx.Write(respJson)
		} else {
			respJson, _ := json.Marshal(resp)
			_, _ = fastctx.Write(respJson)
		}
	})

	r.PATCH("/v3/messages", func(fastctx *fasthttp.RequestCtx) {
		fastctx.Response.Header.SetContentType("application/json")
		input, err := buildExampleServiceNameUpdateMessageV3UpdateMessageRequest(fastctx)
		if err != nil {
			fastctx.SetStatusCode(fasthttp.StatusBadRequest)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fastctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fastctx, "request", fastctx)
		ctx = context.WithValue(ctx, "proto_service", "ServiceName")
		ctx = context.WithValue(ctx, "proto_method", "UpdateMessageV3")
		handler := func(ctx context.Context, req any) (resp any, err error) {
			return h.UpdateMessageV3(ctx, input)
		}
		var resp any
		if middleware == nil {
			resp, _ = handler(ctx, input)
		} else {
			resp, _ = middleware(ctx, input, handler)
		}
		if typedResp, ok := resp.(*common.UpdateMessageRequest); ok {
			respJson, _ := json.Marshal(typedResp.Message)
			_, _ = fastctx.Write(respJson)
		} else {
			respJson, _ := json.Marshal(resp)
			_, _ = fastctx.Write(respJson)
		}
	})

	return nil
}

func buildExampleServiceNameRPCNameInputMsgName(ctx *fasthttp.RequestCtx) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "int64Argument":
			arg.Int64Argument, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				return
			}
		case "stringArgument":
			arg.StringArgument = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	StringArgumentStr, ok := ctx.UserValue("stringArgument").(string)
	if ok && len(StringArgumentStr) != 0 {
		arg.StringArgument = StringArgumentStr
		if arg.StringArgument, err = url.PathUnescape(arg.StringArgument); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field stringArgument: %w", err)
		}
	}

	Int64ArgumentStr, ok := ctx.UserValue("int64Argument").(string)
	if ok && len(Int64ArgumentStr) != 0 {
		arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameAllTypesTestAllTypesMsg(ctx *fasthttp.RequestCtx) (arg *common.AllTypesMsg, err error) {
	arg = &common.AllTypesMsg{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "BoolValue":
			switch value {
			case "true", "t", "1":
				arg.BoolValue = true
			case "false", "f", "0":
				arg.BoolValue = false
			default:
				err = fmt.Errorf("unknown bool string value %s", value)
				return
			}
		case "EnumValue":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.EnumValue = common.Options(OptionsValue)
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = common.Options(intOptionValue)
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					return
				}
			}
		case "Int32Value":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = int32(Int32Value)
		case "Sint32Value":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = int32(Sint32Value)
		case "Uint32Value":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = uint32(Uint32Value)
		case "Int64Value":
			arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				return
			}
		case "Sint64Value":
			arg.Sint64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
				return
			}
		case "Uint64Value":
			arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				return
			}
		case "Sfixed32Value":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = int32(Sfixed32Value)
		case "Fixed32Value":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = uint32(Fixed32Value)
		case "FloatValue":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = float32(FloatValue)
		case "Sfixed64Value":
			arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				return
			}
		case "Fixed64Value":
			arg.Fixed64Value, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
				return
			}
		case "DoubleValue":
			arg.DoubleValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				return
			}
		case "StringValue":
			arg.StringValue = value
		case "BytesValue":
			arg.BytesValue = []byte(value)
		case "MessageValue":
			err = fmt.Errorf("unsupported type message for query argument MessageValue")
			return
		case "MessageValue.int64Argument":
			if arg.MessageValue == nil {
				arg.MessageValue = &common.InputMsgName{}
			}
			arg.MessageValue.Int64Argument, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				return
			}
		case "MessageValue.stringArgument":
			if arg.MessageValue == nil {
				arg.MessageValue = &common.InputMsgName{}
			}
			arg.MessageValue.StringArgument = value
		case "SliceStringValue[]":
			arg.SliceStringValue = append(arg.SliceStringValue, value)
		case "SliceInt32Value[]":
			SliceInt32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter SliceInt32Value: %w", convErr)
				return
			}
			arg.SliceInt32Value = append(arg.SliceInt32Value, int32(SliceInt32Value))
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	BoolValueStr, ok := ctx.UserValue("BoolValue").(string)
	if ok && len(BoolValueStr) != 0 {
		switch BoolValueStr {
		case "true", "t", "1":
			arg.BoolValue = true
		case "false", "f", "0":
			arg.BoolValue = false
		default:
			return nil, fmt.Errorf("unknown bool string value %s", BoolValueStr)
		}
	}

	EnumValueStr, ok := ctx.UserValue("EnumValue").(string)
	if ok && len(EnumValueStr) != 0 {
		if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(EnumValueStr)]; optValueOk {
			arg.EnumValue = common.Options(OptionsValue)
		} else {
			if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
				if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
					arg.EnumValue = common.Options(intOptionValue)
				}
			} else {
				return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
			}
		}
	}

	Int32ValueStr, ok := ctx.UserValue("Int32Value").(string)
	if ok && len(Int32ValueStr) != 0 {
		Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}

	Sint32ValueStr, ok := ctx.UserValue("Sint32Value").(string)
	if ok && len(Sint32ValueStr) != 0 {
		Sint32Value, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
		}
		arg.Sint32Value = int32(Sint32Value)
	}

	Uint32ValueStr, ok := ctx.UserValue("Uint32Value").(string)
	if ok && len(Uint32ValueStr) != 0 {
		Uint32Value, convErr := strconv.ParseUint(Uint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
		}
		arg.Uint32Value = uint32(Uint32Value)
	}

	Int64ValueStr, ok := ctx.UserValue("Int64Value").(string)
	if ok && len(Int64ValueStr) != 0 {
		arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}

	Sint64ValueStr, ok := ctx.UserValue("Sint64Value").(string)
	if ok && len(Sint64ValueStr) != 0 {
		arg.Sint64Value, err = strconv.ParseInt(Sint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
		}
	}

	Uint64ValueStr, ok := ctx.UserValue("Uint64Value").(string)
	if ok && len(Uint64ValueStr) != 0 {
		arg.Uint64Value, err = strconv.ParseUint(Uint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}

	Sfixed32ValueStr, ok := ctx.UserValue("Sfixed32Value").(string)
	if ok && len(Sfixed32ValueStr) != 0 {
		Sfixed32Value, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
		}
		arg.Sfixed32Value = int32(Sfixed32Value)
	}

	Fixed32ValueStr, ok := ctx.UserValue("Fixed32Value").(string)
	if ok && len(Fixed32ValueStr) != 0 {
		Fixed32Value, convErr := strconv.ParseUint(Fixed32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
		}
		arg.Fixed32Value = uint32(Fixed32Value)
	}

	FloatValueStr, ok := ctx.UserValue("FloatValue").(string)
	if ok && len(FloatValueStr) != 0 {
		FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}

	Sfixed64ValueStr, ok := ctx.UserValue("Sfixed64Value").(string)
	if ok && len(Sfixed64ValueStr) != 0 {
		arg.Sfixed64Value, err = strconv.ParseInt(Sfixed64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
		}
	}

	Fixed64ValueStr, ok := ctx.UserValue("Fixed64Value").(string)
	if ok && len(Fixed64ValueStr) != 0 {
		arg.Fixed64Value, err = strconv.ParseUint(Fixed64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
		}
	}

	DoubleValueStr, ok := ctx.UserValue("DoubleValue").(string)
	if ok && len(DoubleValueStr) != 0 {
		arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}

	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if ok && len(StringValueStr) != 0 {
		arg.StringValue = StringValueStr
		if arg.StringValue, err = url.PathUnescape(arg.StringValue); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
		}
	}

	BytesValueStr, ok := ctx.UserValue("BytesValue").(string)
	if ok && len(BytesValueStr) != 0 {
		arg.BytesValue = []byte(BytesValueStr)
		if BytesValueStr, err = url.PathUnescape(string(arg.BytesValue)); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field BytesValue: %w", err)
		}
		arg.BytesValue = []byte(BytesValueStr)
	}

	return arg, err
}

func buildExampleServiceNameAllTextTypesPostAllTextTypesMsg(ctx *fasthttp.RequestCtx) (arg *common.AllTextTypesMsg, err error) {
	arg = &common.AllTextTypesMsg{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "String":
			arg.String_ = value
		case "RepeatedString[]":
			arg.RepeatedString = append(arg.RepeatedString, value)
		case "OptionalString":
			arg.OptionalString = &value
		case "Bytes":
			arg.Bytes = []byte(value)
		case "RepeatedBytes[]":
			arg.RepeatedBytes = append(arg.RepeatedBytes, []byte(value))
		case "OptionalBytes":
			arg.OptionalBytes = []byte(value)
		case "Enum":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.Enum = common.Options(OptionsValue)
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.Enum = common.Options(intOptionValue)
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter Enum: %w", convErr)
					return
				}
			}
		case "RepeatedEnum[]":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(intOptionValue))
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter RepeatedEnum: %w", convErr)
					return
				}
			}
		case "OptionalEnum":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				OptionalEnum := common.Options(OptionsValue)
				arg.OptionalEnum = &OptionalEnum
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						OptionalEnum := common.Options(intOptionValue)
						arg.OptionalEnum = &OptionalEnum
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter OptionalEnum: %w", convErr)
					return
				}
			}
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	String_Str, ok := ctx.UserValue("String").(string)
	if ok && len(String_Str) != 0 {
		arg.String_ = String_Str
		if arg.String_, err = url.PathUnescape(arg.String_); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field String: %w", err)
		}
	}

	RepeatedStringStr, ok := ctx.UserValue("RepeatedString").(string)
	if ok && len(RepeatedStringStr) != 0 {
		arg.RepeatedString = strings.Split(RepeatedStringStr, ",")
		for i, value := range arg.RepeatedString {
			if arg.RepeatedString[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field RepeatedString: %w", err)
			}
		}
	}

	BytesStr, ok := ctx.UserValue("Bytes").(string)
	if ok && len(BytesStr) != 0 {
		arg.Bytes = []byte(BytesStr)
		if BytesStr, err = url.PathUnescape(string(arg.Bytes)); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field Bytes: %w", err)
		}
		arg.Bytes = []byte(BytesStr)
	}

	RepeatedBytesStr, ok := ctx.UserValue("RepeatedBytes").(string)
	if ok && len(RepeatedBytesStr) != 0 {
		RepeatedBytesStrs := strings.Split(RepeatedBytesStr, ",")
		arg.RepeatedBytes = make([][]byte, 0, len(RepeatedBytesStrs))
		for _, str := range RepeatedBytesStrs {
			arg.RepeatedBytes = append(arg.RepeatedBytes, []byte(str))
		}
		for i, value := range arg.RepeatedBytes {
			if RepeatedBytesStr, err = url.PathUnescape(string(value)); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field RepeatedBytes: %w", err)
			}
			arg.RepeatedBytes[i] = []byte(RepeatedBytesStr)
		}
	}

	EnumStr, ok := ctx.UserValue("Enum").(string)
	if ok && len(EnumStr) != 0 {
		if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(EnumStr)]; optValueOk {
			arg.Enum = common.Options(OptionsValue)
		} else {
			if intOptionValue, convErr := strconv.ParseInt(EnumStr, 10, 32); convErr == nil {
				if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
					arg.Enum = common.Options(intOptionValue)
				}
			} else {
				return nil, fmt.Errorf("conversion failed for parameter Enum: %w", convErr)
			}
		}
	}

	RepeatedEnumStr, ok := ctx.UserValue("RepeatedEnum").(string)
	if ok && len(RepeatedEnumStr) != 0 {
		RepeatedEnumStrs := strings.Split(RepeatedEnumStr, ",")
		arg.RepeatedEnum = make([]common.Options, 0, len(RepeatedEnumStrs))
		for _, str := range RepeatedEnumStrs {
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(str)]; optValueOk {
				arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(str, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(intOptionValue))
					}
				} else {
					return nil, fmt.Errorf("conversion failed for parameter RepeatedEnum: %w", convErr)
				}
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameAllTextTypesGetAllTextTypesMsg(ctx *fasthttp.RequestCtx) (arg *common.AllTextTypesMsg, err error) {
	arg = &common.AllTextTypesMsg{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "String":
			arg.String_ = value
		case "RepeatedString[]":
			arg.RepeatedString = append(arg.RepeatedString, value)
		case "OptionalString":
			arg.OptionalString = &value
		case "Bytes":
			arg.Bytes = []byte(value)
		case "RepeatedBytes[]":
			arg.RepeatedBytes = append(arg.RepeatedBytes, []byte(value))
		case "OptionalBytes":
			arg.OptionalBytes = []byte(value)
		case "Enum":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.Enum = common.Options(OptionsValue)
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.Enum = common.Options(intOptionValue)
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter Enum: %w", convErr)
					return
				}
			}
		case "RepeatedEnum[]":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(intOptionValue))
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter RepeatedEnum: %w", convErr)
					return
				}
			}
		case "OptionalEnum":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				OptionalEnum := common.Options(OptionsValue)
				arg.OptionalEnum = &OptionalEnum
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						OptionalEnum := common.Options(intOptionValue)
						arg.OptionalEnum = &OptionalEnum
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter OptionalEnum: %w", convErr)
					return
				}
			}
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	String_Str, ok := ctx.UserValue("String").(string)
	if ok && len(String_Str) != 0 {
		arg.String_ = String_Str
		if arg.String_, err = url.PathUnescape(arg.String_); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field String: %w", err)
		}
	}

	RepeatedStringStr, ok := ctx.UserValue("RepeatedString").(string)
	if ok && len(RepeatedStringStr) != 0 {
		arg.RepeatedString = strings.Split(RepeatedStringStr, ",")
		for i, value := range arg.RepeatedString {
			if arg.RepeatedString[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field RepeatedString: %w", err)
			}
		}
	}

	BytesStr, ok := ctx.UserValue("Bytes").(string)
	if ok && len(BytesStr) != 0 {
		arg.Bytes = []byte(BytesStr)
		if BytesStr, err = url.PathUnescape(string(arg.Bytes)); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field Bytes: %w", err)
		}
		arg.Bytes = []byte(BytesStr)
	}

	RepeatedBytesStr, ok := ctx.UserValue("RepeatedBytes").(string)
	if ok && len(RepeatedBytesStr) != 0 {
		RepeatedBytesStrs := strings.Split(RepeatedBytesStr, ",")
		arg.RepeatedBytes = make([][]byte, 0, len(RepeatedBytesStrs))
		for _, str := range RepeatedBytesStrs {
			arg.RepeatedBytes = append(arg.RepeatedBytes, []byte(str))
		}
		for i, value := range arg.RepeatedBytes {
			if RepeatedBytesStr, err = url.PathUnescape(string(value)); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field RepeatedBytes: %w", err)
			}
			arg.RepeatedBytes[i] = []byte(RepeatedBytesStr)
		}
	}

	EnumStr, ok := ctx.UserValue("Enum").(string)
	if ok && len(EnumStr) != 0 {
		if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(EnumStr)]; optValueOk {
			arg.Enum = common.Options(OptionsValue)
		} else {
			if intOptionValue, convErr := strconv.ParseInt(EnumStr, 10, 32); convErr == nil {
				if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
					arg.Enum = common.Options(intOptionValue)
				}
			} else {
				return nil, fmt.Errorf("conversion failed for parameter Enum: %w", convErr)
			}
		}
	}

	RepeatedEnumStr, ok := ctx.UserValue("RepeatedEnum").(string)
	if ok && len(RepeatedEnumStr) != 0 {
		RepeatedEnumStrs := strings.Split(RepeatedEnumStr, ",")
		arg.RepeatedEnum = make([]common.Options, 0, len(RepeatedEnumStrs))
		for _, str := range RepeatedEnumStrs {
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(str)]; optValueOk {
				arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(str, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.RepeatedEnum = append(arg.RepeatedEnum, common.Options(intOptionValue))
					}
				} else {
					return nil, fmt.Errorf("conversion failed for parameter RepeatedEnum: %w", convErr)
				}
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameCommonTypesAny(ctx *fasthttp.RequestCtx) (arg *anypb.Any, err error) {
	arg = &anypb.Any{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "type_url":
			arg.TypeUrl = value
		case "value":
			arg.Value = []byte(value)
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameSameInputAndOutputInputMsgName(ctx *fasthttp.RequestCtx) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "int64Argument":
			arg.Int64Argument, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				return
			}
		case "stringArgument":
			arg.StringArgument = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	StringArgumentStr, ok := ctx.UserValue("stringArgument").(string)
	if ok && len(StringArgumentStr) != 0 {
		arg.StringArgument = StringArgumentStr
		if arg.StringArgument, err = url.PathUnescape(arg.StringArgument); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field stringArgument: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameOptionalOptionalField(ctx *fasthttp.RequestCtx) (arg *common.OptionalField, err error) {
	arg = &common.OptionalField{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "BoolValue":
			switch value {
			case "true", "t", "1":
				BoolValue := true
				arg.BoolValue = &BoolValue
			case "false", "f", "0":
				BoolValue := false
				arg.BoolValue = &BoolValue
			default:
				err = fmt.Errorf("unknown bool string value %s", value)
				return
			}
		case "EnumValue":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				EnumValue := common.Options(OptionsValue)
				arg.EnumValue = &EnumValue
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						EnumValue := common.Options(intOptionValue)
						arg.EnumValue = &EnumValue
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					return
				}
			}
		case "Int32Value":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			Int32ValueValue := int32(Int32Value)
			arg.Int32Value = &Int32ValueValue
		case "Sint32Value":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			Sint32ValueValue := int32(Sint32Value)
			arg.Sint32Value = &Sint32ValueValue
		case "Uint32Value":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			Uint32ValueValue := uint32(Uint32Value)
			arg.Uint32Value = &Uint32ValueValue
		case "Int64Value":
			Int64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = &Int64Value
		case "Sint64Value":
			Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = &Sint64Value
		case "Uint64Value":
			Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = &Uint64Value
		case "Sfixed32Value":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			Sfixed32ValueValue := int32(Sfixed32Value)
			arg.Sfixed32Value = &Sfixed32ValueValue
		case "Fixed32Value":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			Fixed32ValueValue := uint32(Fixed32Value)
			arg.Fixed32Value = &Fixed32ValueValue
		case "FloatValue":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			FloatValueValue := float32(FloatValue)
			arg.FloatValue = &FloatValueValue
		case "Sfixed64Value":
			Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = &Sfixed64Value
		case "Fixed64Value":
			Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = &Fixed64Value
		case "DoubleValue":
			DoubleValue, convErr := strconv.ParseFloat(value, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = &DoubleValue
		case "StringValue":
			arg.StringValue = &value
		case "BytesValue":
			arg.BytesValue = []byte(value)
		case "MessageValue":
			err = fmt.Errorf("unsupported type message for query argument MessageValue")
			return
		case "MessageValue.int64Argument":
			if arg.MessageValue == nil {
				arg.MessageValue = &common.InputMsgName{}
			}
			arg.MessageValue.Int64Argument, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				return
			}
		case "MessageValue.stringArgument":
			if arg.MessageValue == nil {
				arg.MessageValue = &common.InputMsgName{}
			}
			arg.MessageValue.StringArgument = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameGetMethodInputMsgName(ctx *fasthttp.RequestCtx) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "int64Argument":
			arg.Int64Argument, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				return
			}
		case "stringArgument":
			arg.StringArgument = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameCheckRepeatedPathRepeatedCheck(ctx *fasthttp.RequestCtx) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "BoolValue[]":
			switch value {
			case "true", "t", "1":
				arg.BoolValue = append(arg.BoolValue, true)
			case "false", "f", "0":
				arg.BoolValue = append(arg.BoolValue, false)
			default:
				err = fmt.Errorf("unknown bool string value %s", value)
				return
			}
		case "EnumValue[]":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.EnumValue = append(arg.EnumValue, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = append(arg.EnumValue, common.Options(intOptionValue))
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					return
				}
			}
		case "Int32Value[]":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = append(arg.Int32Value, int32(Int32Value))
		case "Sint32Value[]":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = append(arg.Sint32Value, int32(Sint32Value))
		case "Uint32Value[]":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
		case "Int64Value[]":
			Int64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = append(arg.Int64Value, Int64Value)
		case "Sint64Value[]":
			Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = append(arg.Sint64Value, Sint64Value)
		case "Uint64Value[]":
			Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = append(arg.Uint64Value, Uint64Value)
		case "Sfixed32Value[]":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32Value))
		case "Fixed32Value[]":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
		case "FloatValue[]":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = append(arg.FloatValue, float32(FloatValue))
		case "Sfixed64Value[]":
			Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64Value)
		case "Fixed64Value[]":
			Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = append(arg.Fixed64Value, Fixed64Value)
		case "DoubleValue[]":
			DoubleValue, convErr := strconv.ParseFloat(value, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = append(arg.DoubleValue, DoubleValue)
		case "StringValue[]":
			arg.StringValue = append(arg.StringValue, value)
		case "BytesValue[]":
			arg.BytesValue = append(arg.BytesValue, []byte(value))
		case "StringValueQuery[]":
			arg.StringValueQuery = append(arg.StringValueQuery, value)
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	BoolValueStr, ok := ctx.UserValue("BoolValue").(string)
	if ok && len(BoolValueStr) != 0 {
		BoolValueStrs := strings.Split(BoolValueStr, ",")
		arg.BoolValue = make([]bool, 0, len(BoolValueStrs))
		for _, str := range BoolValueStrs {
			switch str {
			case "true", "t", "1":
				arg.BoolValue = append(arg.BoolValue, true)
			case "false", "f", "0":
				arg.BoolValue = append(arg.BoolValue, false)
			default:
				return nil, fmt.Errorf("unknown bool string value %s", str)
			}
		}
	}

	EnumValueStr, ok := ctx.UserValue("EnumValue").(string)
	if ok && len(EnumValueStr) != 0 {
		EnumValueStrs := strings.Split(EnumValueStr, ",")
		arg.EnumValue = make([]common.Options, 0, len(EnumValueStrs))
		for _, str := range EnumValueStrs {
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(str)]; optValueOk {
				arg.EnumValue = append(arg.EnumValue, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(str, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = append(arg.EnumValue, common.Options(intOptionValue))
					}
				} else {
					return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
				}
			}
		}
	}

	Int32ValueStr, ok := ctx.UserValue("Int32Value").(string)
	if ok && len(Int32ValueStr) != 0 {
		Int32ValueStrs := strings.Split(Int32ValueStr, ",")
		arg.Int32Value = make([]int32, 0, len(Int32ValueStrs))
		for _, str := range Int32ValueStrs {
			Int32Value, convErr := strconv.ParseInt(str, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
			}
			arg.Int32Value = append(arg.Int32Value, int32(Int32Value))
		}
	}

	Sint32ValueStr, ok := ctx.UserValue("Sint32Value").(string)
	if ok && len(Sint32ValueStr) != 0 {
		Sint32ValueStrs := strings.Split(Sint32ValueStr, ",")
		arg.Sint32Value = make([]int32, 0, len(Sint32ValueStrs))
		for _, str := range Sint32ValueStrs {
			Sint32Value, convErr := strconv.ParseInt(str, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
			}
			arg.Sint32Value = append(arg.Sint32Value, int32(Sint32Value))
		}
	}

	Uint32ValueStr, ok := ctx.UserValue("Uint32Value").(string)
	if ok && len(Uint32ValueStr) != 0 {
		Uint32ValueStrs := strings.Split(Uint32ValueStr, ",")
		arg.Uint32Value = make([]uint32, 0, len(Uint32ValueStrs))
		for _, str := range Uint32ValueStrs {
			Uint32Value, convErr := strconv.ParseUint(str, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
		}
	}

	Int64ValueStr, ok := ctx.UserValue("Int64Value").(string)
	if ok && len(Int64ValueStr) != 0 {
		Int64ValueStrs := strings.Split(Int64ValueStr, ",")
		arg.Int64Value = make([]int64, 0, len(Int64ValueStrs))
		for _, str := range Int64ValueStrs {
			Int64Value, convErr := strconv.ParseInt(str, 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
			}
			arg.Int64Value = append(arg.Int64Value, Int64Value)
		}
	}

	Sint64ValueStr, ok := ctx.UserValue("Sint64Value").(string)
	if ok && len(Sint64ValueStr) != 0 {
		Sint64ValueStrs := strings.Split(Sint64ValueStr, ",")
		arg.Sint64Value = make([]int64, 0, len(Sint64ValueStrs))
		for _, str := range Sint64ValueStrs {
			Sint64Value, convErr := strconv.ParseInt(str, 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
			}
			arg.Sint64Value = append(arg.Sint64Value, Sint64Value)
		}
	}

	Uint64ValueStr, ok := ctx.UserValue("Uint64Value").(string)
	if ok && len(Uint64ValueStr) != 0 {
		Uint64ValueStrs := strings.Split(Uint64ValueStr, ",")
		arg.Uint64Value = make([]uint64, 0, len(Uint64ValueStrs))
		for _, str := range Uint64ValueStrs {
			Uint64Value, convErr := strconv.ParseUint(str, 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
			}
			arg.Uint64Value = append(arg.Uint64Value, Uint64Value)
		}
	}

	Sfixed32ValueStr, ok := ctx.UserValue("Sfixed32Value").(string)
	if ok && len(Sfixed32ValueStr) != 0 {
		Sfixed32ValueStrs := strings.Split(Sfixed32ValueStr, ",")
		arg.Sfixed32Value = make([]int32, 0, len(Sfixed32ValueStrs))
		for _, str := range Sfixed32ValueStrs {
			Sfixed32Value, convErr := strconv.ParseInt(str, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
			}
			arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32Value))
		}
	}

	Fixed32ValueStr, ok := ctx.UserValue("Fixed32Value").(string)
	if ok && len(Fixed32ValueStr) != 0 {
		Fixed32ValueStrs := strings.Split(Fixed32ValueStr, ",")
		arg.Fixed32Value = make([]uint32, 0, len(Fixed32ValueStrs))
		for _, str := range Fixed32ValueStrs {
			Fixed32Value, convErr := strconv.ParseUint(str, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
		}
	}

	FloatValueStr, ok := ctx.UserValue("FloatValue").(string)
	if ok && len(FloatValueStr) != 0 {
		FloatValueStrs := strings.Split(FloatValueStr, ",")
		arg.FloatValue = make([]float32, 0, len(FloatValueStrs))
		for _, str := range FloatValueStrs {
			FloatValue, convErr := strconv.ParseFloat(str, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
			}
			arg.FloatValue = append(arg.FloatValue, float32(FloatValue))
		}
	}

	Sfixed64ValueStr, ok := ctx.UserValue("Sfixed64Value").(string)
	if ok && len(Sfixed64ValueStr) != 0 {
		Sfixed64ValueStrs := strings.Split(Sfixed64ValueStr, ",")
		arg.Sfixed64Value = make([]int64, 0, len(Sfixed64ValueStrs))
		for _, str := range Sfixed64ValueStrs {
			Sfixed64Value, convErr := strconv.ParseInt(str, 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
			}
			arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64Value)
		}
	}

	Fixed64ValueStr, ok := ctx.UserValue("Fixed64Value").(string)
	if ok && len(Fixed64ValueStr) != 0 {
		Fixed64ValueStrs := strings.Split(Fixed64ValueStr, ",")
		arg.Fixed64Value = make([]uint64, 0, len(Fixed64ValueStrs))
		for _, str := range Fixed64ValueStrs {
			Fixed64Value, convErr := strconv.ParseUint(str, 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
			}
			arg.Fixed64Value = append(arg.Fixed64Value, Fixed64Value)
		}
	}

	DoubleValueStr, ok := ctx.UserValue("DoubleValue").(string)
	if ok && len(DoubleValueStr) != 0 {
		DoubleValueStrs := strings.Split(DoubleValueStr, ",")
		arg.DoubleValue = make([]float64, 0, len(DoubleValueStrs))
		for _, str := range DoubleValueStrs {
			DoubleValue, convErr := strconv.ParseFloat(str, 64)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
			}
			arg.DoubleValue = append(arg.DoubleValue, DoubleValue)
		}
	}

	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if ok && len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
		for i, value := range arg.StringValue {
			if arg.StringValue[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
			}
		}
	}

	BytesValueStr, ok := ctx.UserValue("BytesValue").(string)
	if ok && len(BytesValueStr) != 0 {
		BytesValueStrs := strings.Split(BytesValueStr, ",")
		arg.BytesValue = make([][]byte, 0, len(BytesValueStrs))
		for _, str := range BytesValueStrs {
			arg.BytesValue = append(arg.BytesValue, []byte(str))
		}
		for i, value := range arg.BytesValue {
			if BytesValueStr, err = url.PathUnescape(string(value)); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field BytesValue: %w", err)
			}
			arg.BytesValue[i] = []byte(BytesValueStr)
		}
	}

	StringValueQueryStr, ok := ctx.UserValue("StringValueQuery").(string)
	if ok && len(StringValueQueryStr) != 0 {
		arg.StringValueQuery = strings.Split(StringValueQueryStr, ",")
		for i, value := range arg.StringValueQuery {
			if arg.StringValueQuery[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValueQuery: %w", err)
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(ctx *fasthttp.RequestCtx) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "BoolValue[]":
			switch value {
			case "true", "t", "1":
				arg.BoolValue = append(arg.BoolValue, true)
			case "false", "f", "0":
				arg.BoolValue = append(arg.BoolValue, false)
			default:
				err = fmt.Errorf("unknown bool string value %s", value)
				return
			}
		case "EnumValue[]":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.EnumValue = append(arg.EnumValue, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = append(arg.EnumValue, common.Options(intOptionValue))
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					return
				}
			}
		case "Int32Value[]":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = append(arg.Int32Value, int32(Int32Value))
		case "Sint32Value[]":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = append(arg.Sint32Value, int32(Sint32Value))
		case "Uint32Value[]":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
		case "Int64Value[]":
			Int64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = append(arg.Int64Value, Int64Value)
		case "Sint64Value[]":
			Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = append(arg.Sint64Value, Sint64Value)
		case "Uint64Value[]":
			Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = append(arg.Uint64Value, Uint64Value)
		case "Sfixed32Value[]":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32Value))
		case "Fixed32Value[]":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
		case "FloatValue[]":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = append(arg.FloatValue, float32(FloatValue))
		case "Sfixed64Value[]":
			Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64Value)
		case "Fixed64Value[]":
			Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = append(arg.Fixed64Value, Fixed64Value)
		case "DoubleValue[]":
			DoubleValue, convErr := strconv.ParseFloat(value, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = append(arg.DoubleValue, DoubleValue)
		case "StringValue[]":
			arg.StringValue = append(arg.StringValue, value)
		case "BytesValue[]":
			arg.BytesValue = append(arg.BytesValue, []byte(value))
		case "StringValueQuery[]":
			arg.StringValueQuery = append(arg.StringValueQuery, value)
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if ok && len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
		for i, value := range arg.StringValue {
			if arg.StringValue[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameCheckRepeatedPostRepeatedCheck(ctx *fasthttp.RequestCtx) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "BoolValue[]":
			switch value {
			case "true", "t", "1":
				arg.BoolValue = append(arg.BoolValue, true)
			case "false", "f", "0":
				arg.BoolValue = append(arg.BoolValue, false)
			default:
				err = fmt.Errorf("unknown bool string value %s", value)
				return
			}
		case "EnumValue[]":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.EnumValue = append(arg.EnumValue, common.Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = append(arg.EnumValue, common.Options(intOptionValue))
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					return
				}
			}
		case "Int32Value[]":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = append(arg.Int32Value, int32(Int32Value))
		case "Sint32Value[]":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = append(arg.Sint32Value, int32(Sint32Value))
		case "Uint32Value[]":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
		case "Int64Value[]":
			Int64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = append(arg.Int64Value, Int64Value)
		case "Sint64Value[]":
			Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = append(arg.Sint64Value, Sint64Value)
		case "Uint64Value[]":
			Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = append(arg.Uint64Value, Uint64Value)
		case "Sfixed32Value[]":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32Value))
		case "Fixed32Value[]":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
		case "FloatValue[]":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = append(arg.FloatValue, float32(FloatValue))
		case "Sfixed64Value[]":
			Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64Value)
		case "Fixed64Value[]":
			Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = append(arg.Fixed64Value, Fixed64Value)
		case "DoubleValue[]":
			DoubleValue, convErr := strconv.ParseFloat(value, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = append(arg.DoubleValue, DoubleValue)
		case "StringValue[]":
			arg.StringValue = append(arg.StringValue, value)
		case "BytesValue[]":
			arg.BytesValue = append(arg.BytesValue, []byte(value))
		case "StringValueQuery[]":
			arg.StringValueQuery = append(arg.StringValueQuery, value)
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if ok && len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
		for i, value := range arg.StringValue {
			if arg.StringValue[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameEmptyGetEmpty(ctx *fasthttp.RequestCtx) (arg *common.Empty, err error) {
	arg = &common.Empty{}
	return arg, err
}

func buildExampleServiceNameEmptyPostEmpty(ctx *fasthttp.RequestCtx) (arg *common.Empty, err error) {
	arg = &common.Empty{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	return arg, err
}

func buildExampleServiceNameOnlyStructInGetOnlyStruct(ctx *fasthttp.RequestCtx) (arg *common.OnlyStruct, err error) {
	arg = &common.OnlyStruct{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "value":
			err = fmt.Errorf("unsupported type message for query argument value")
			return
		case "value.value":
			if arg.Value == nil {
				arg.Value = &common.ArrayItem{}
			}
			arg.Value.Value = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameMultipartFormMultipartFormRequest(ctx *fasthttp.RequestCtx) (arg *common.MultipartFormRequest, err error) {
	arg = &common.MultipartFormRequest{}
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	if file, ok := form.File["document"]; ok && len(file) > 0 {
		var f multipart.File
		f, err = file[0].Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: document: %w", err)
		}
		arg.Document = &common.FileEx{
			File:    make([]byte, file[0].Size),
			Name:    file[0].Filename,
			Headers: make(map[string]string, len(file[0].Header)),
		}
		for key, value := range file[0].Header {
			arg.Document.Headers[key] = value[0]
		}
		_, err = f.Read(arg.Document.File)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: document: %w", err)
		}
	}
	if values, ok := form.Value["otherField"]; ok && len(values) > 0 {
		arg.OtherField = values[0]
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "document":
			err = fmt.Errorf("unsupported type message for query argument document")
			return
		case "otherField":
			arg.OtherField = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameMultipartFormAllTypesMultipartFormAllTypes(ctx *fasthttp.RequestCtx) (arg *common.MultipartFormAllTypes, err error) {
	arg = &common.MultipartFormAllTypes{}
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	if values, ok := form.Value["BoolValue"]; ok && len(values) > 0 {
		switch values[0] {
		case "true", "t", "1":
			arg.BoolValue = true
		case "false", "f", "0":
			arg.BoolValue = false
		default:
			return nil, fmt.Errorf("unknown bool string value %s", values[0])
		}
	}
	if values, ok := form.Value["EnumValue"]; ok && len(values) > 0 {
		if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(values[0])]; optValueOk {
			arg.EnumValue = common.Options(OptionsValue)
		} else {
			if intOptionValue, convErr := strconv.ParseInt(values[0], 10, 32); convErr == nil {
				if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
					arg.EnumValue = common.Options(intOptionValue)
				}
			} else {
				return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
			}
		}
	}
	if values, ok := form.Value["Int32Value"]; ok && len(values) > 0 {
		Int32Value, convErr := strconv.ParseInt(values[0], 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}
	if values, ok := form.Value["Sint32Value"]; ok && len(values) > 0 {
		Sint32Value, convErr := strconv.ParseInt(values[0], 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
		}
		arg.Sint32Value = int32(Sint32Value)
	}
	if values, ok := form.Value["Uint32Value"]; ok && len(values) > 0 {
		for _, value := range values {
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
		}
	}
	if values, ok := form.Value["Int64Value"]; ok && len(values) > 0 {
		arg.Int64Value, err = strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}
	if values, ok := form.Value["Sint64Value"]; ok && len(values) > 0 {
		Sint64Value, convErr := strconv.ParseInt(values[0], 10, 64)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
		}
		arg.Sint64Value = &Sint64Value
	}
	if values, ok := form.Value["Uint64Value"]; ok && len(values) > 0 {
		arg.Uint64Value, err = strconv.ParseUint(values[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}
	if values, ok := form.Value["Sfixed32Value"]; ok && len(values) > 0 {
		Sfixed32Value, convErr := strconv.ParseInt(values[0], 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
		}
		arg.Sfixed32Value = int32(Sfixed32Value)
	}
	if values, ok := form.Value["Fixed32Value"]; ok && len(values) > 0 {
		for _, value := range values {
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
		}
	}
	if values, ok := form.Value["FloatValue"]; ok && len(values) > 0 {
		FloatValue, convErr := strconv.ParseFloat(values[0], 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}
	if values, ok := form.Value["Sfixed64Value"]; ok && len(values) > 0 {
		arg.Sfixed64Value, err = strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
		}
	}
	if values, ok := form.Value["Fixed64Value"]; ok && len(values) > 0 {
		Fixed64Value, convErr := strconv.ParseUint(values[0], 10, 64)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
		}
		arg.Fixed64Value = &Fixed64Value
	}
	if values, ok := form.Value["DoubleValue"]; ok && len(values) > 0 {
		arg.DoubleValue, err = strconv.ParseFloat(values[0], 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}
	if values, ok := form.Value["StringValue"]; ok && len(values) > 0 {
		arg.StringValue = values[0]
	}
	if values, ok := form.Value["BytesValue"]; ok && len(values) > 0 {
		arg.BytesValue = []byte(values[0])
	}
	if values, ok := form.Value["SliceStringValue"]; ok && len(values) > 0 {
		arg.SliceStringValue = append(arg.SliceStringValue, values...)
	}
	if values, ok := form.Value["SliceInt32Value"]; ok && len(values) > 0 {
		for _, value := range values {
			SliceInt32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter SliceInt32Value: %w", convErr)
			}
			arg.SliceInt32Value = append(arg.SliceInt32Value, int32(SliceInt32Value))
		}
	}
	if file, ok := form.File["document"]; ok && len(file) > 0 {
		var f multipart.File
		f, err = file[0].Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: document: %w", err)
		}
		arg.Document = &common.FileEx{
			File:    make([]byte, file[0].Size),
			Name:    file[0].Filename,
			Headers: make(map[string]string, len(file[0].Header)),
		}
		for key, value := range file[0].Header {
			arg.Document.Headers[key] = value[0]
		}
		_, err = f.Read(arg.Document.File)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: document: %w", err)
		}
	}
	if values, ok := form.Value["RepeatedStringValue"]; ok && len(values) > 0 {
		arg.RepeatedStringValue = append(arg.RepeatedStringValue, values...)
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "BoolValue":
			switch value {
			case "true", "t", "1":
				arg.BoolValue = true
			case "false", "f", "0":
				arg.BoolValue = false
			default:
				err = fmt.Errorf("unknown bool string value %s", value)
				return
			}
		case "EnumValue":
			if OptionsValue, optValueOk := common.Options_value[strings.ToUpper(value)]; optValueOk {
				arg.EnumValue = common.Options(OptionsValue)
			} else {
				if intOptionValue, convErr := strconv.ParseInt(value, 10, 32); convErr == nil {
					if _, optIntValueOk := common.Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = common.Options(intOptionValue)
					}
				} else {
					err = fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					return
				}
			}
		case "Int32Value":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = int32(Int32Value)
		case "Sint32Value":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = int32(Sint32Value)
		case "Uint32Value[]":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
		case "Int64Value":
			arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				return
			}
		case "Sint64Value":
			Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = &Sint64Value
		case "Uint64Value":
			arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				return
			}
		case "Sfixed32Value":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = int32(Sfixed32Value)
		case "Fixed32Value[]":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
		case "FloatValue":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = float32(FloatValue)
		case "Sfixed64Value":
			arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				return
			}
		case "Fixed64Value":
			Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = &Fixed64Value
		case "DoubleValue":
			arg.DoubleValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				return
			}
		case "StringValue":
			arg.StringValue = value
		case "BytesValue":
			arg.BytesValue = []byte(value)
		case "SliceStringValue[]":
			arg.SliceStringValue = append(arg.SliceStringValue, value)
		case "SliceInt32Value[]":
			SliceInt32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter SliceInt32Value: %w", convErr)
				return
			}
			arg.SliceInt32Value = append(arg.SliceInt32Value, int32(SliceInt32Value))
		case "document":
			err = fmt.Errorf("unsupported type message for query argument document")
			return
		case "RepeatedStringValue[]":
			arg.RepeatedStringValue = append(arg.RepeatedStringValue, value)
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameAllTypesMaxTestAllNumberTypesMsg(ctx *fasthttp.RequestCtx) (arg *common.AllNumberTypesMsg, err error) {
	arg = &common.AllNumberTypesMsg{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "Int32Value":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = int32(Int32Value)
		case "Sint32Value":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = int32(Sint32Value)
		case "Uint32Value":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = uint32(Uint32Value)
		case "Int64Value":
			arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				return
			}
		case "Sint64Value":
			arg.Sint64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
				return
			}
		case "Uint64Value":
			arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				return
			}
		case "Sfixed32Value":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = int32(Sfixed32Value)
		case "Fixed32Value":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = uint32(Fixed32Value)
		case "FloatValue":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = float32(FloatValue)
		case "Sfixed64Value":
			arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				return
			}
		case "Fixed64Value":
			arg.Fixed64Value, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
				return
			}
		case "DoubleValue":
			arg.DoubleValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				return
			}
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	Int32ValueStr, ok := ctx.UserValue("Int32Value").(string)
	if ok && len(Int32ValueStr) != 0 {
		Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}

	Uint32ValueStr, ok := ctx.UserValue("Uint32Value").(string)
	if ok && len(Uint32ValueStr) != 0 {
		Uint32Value, convErr := strconv.ParseUint(Uint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
		}
		arg.Uint32Value = uint32(Uint32Value)
	}

	Int64ValueStr, ok := ctx.UserValue("Int64Value").(string)
	if ok && len(Int64ValueStr) != 0 {
		arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}

	Uint64ValueStr, ok := ctx.UserValue("Uint64Value").(string)
	if ok && len(Uint64ValueStr) != 0 {
		arg.Uint64Value, err = strconv.ParseUint(Uint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}

	FloatValueStr, ok := ctx.UserValue("FloatValue").(string)
	if ok && len(FloatValueStr) != 0 {
		FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}

	DoubleValueStr, ok := ctx.UserValue("DoubleValue").(string)
	if ok && len(DoubleValueStr) != 0 {
		arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameAllTypesMaxQueryTestAllNumberTypesMsg(ctx *fasthttp.RequestCtx) (arg *common.AllNumberTypesMsg, err error) {
	arg = &common.AllNumberTypesMsg{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "Int32Value":
			Int32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = int32(Int32Value)
		case "Sint32Value":
			Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = int32(Sint32Value)
		case "Uint32Value":
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = uint32(Uint32Value)
		case "Int64Value":
			arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				return
			}
		case "Sint64Value":
			arg.Sint64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
				return
			}
		case "Uint64Value":
			arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				return
			}
		case "Sfixed32Value":
			Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = int32(Sfixed32Value)
		case "Fixed32Value":
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = uint32(Fixed32Value)
		case "FloatValue":
			FloatValue, convErr := strconv.ParseFloat(value, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = float32(FloatValue)
		case "Sfixed64Value":
			arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				return
			}
		case "Fixed64Value":
			arg.Fixed64Value, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
				return
			}
		case "DoubleValue":
			arg.DoubleValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				return
			}
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameGetMessageGetMessageRequest(ctx *fasthttp.RequestCtx) (arg *common.GetMessageRequest, err error) {
	arg = &common.GetMessageRequest{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "name":
			arg.Name = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	NameStr, ok := ctx.UserValue("name").(string)
	if ok && len(NameStr) != 0 {
		arg.Name = NameStr
		if arg.Name, err = url.PathUnescape(arg.Name); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field name: %w", err)
		}
		arg.Name = fmt.Sprintf("messages/%s", arg.Name)
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV2GetMessageRequestV2(ctx *fasthttp.RequestCtx) (arg *common.GetMessageRequestV2, err error) {
	arg = &common.GetMessageRequestV2{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "message_id":
			arg.MessageId = value
		case "revision":
			arg.Revision, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter revision: %w", err)
				return
			}
		case "sub":
			err = fmt.Errorf("unsupported type message for query argument sub")
			return
		case "sub.subfield":
			if arg.Sub == nil {
				arg.Sub = &common.GetMessageRequestV2_SubMessage{}
			}
			arg.Sub.Subfield = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	MessageIdStr, ok := ctx.UserValue("message_id").(string)
	if ok && len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameUpdateMessageUpdateMessageRequest(ctx *fasthttp.RequestCtx) (arg *common.UpdateMessageRequest, err error) {
	arg = &common.UpdateMessageRequest{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		arg.Message = &common.MessageV2{}
		if err = json.Unmarshal(body, arg.Message); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "message_id":
			arg.MessageId = value
		case "message":
			err = fmt.Errorf("unsupported type message for query argument message")
			return
		case "message.message_id":
			if arg.Message == nil {
				arg.Message = &common.MessageV2{}
			}
			arg.Message.MessageId = value
		case "message.text":
			if arg.Message == nil {
				arg.Message = &common.MessageV2{}
			}
			arg.Message.Text = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	MessageIdStr, ok := ctx.UserValue("message_id").(string)
	if ok && len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameUpdateMessageV2MessageV2(ctx *fasthttp.RequestCtx) (arg *common.MessageV2, err error) {
	arg = &common.MessageV2{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "message_id":
			arg.MessageId = value
		case "text":
			arg.Text = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	MessageIdStr, ok := ctx.UserValue("message_id").(string)
	if ok && len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV3GetMessageRequestV3(ctx *fasthttp.RequestCtx) (arg *common.GetMessageRequestV3, err error) {
	arg = &common.GetMessageRequestV3{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "message_id":
			arg.MessageId = value
		case "user_id":
			arg.UserId = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	MessageIdStr, ok := ctx.UserValue("message_id").(string)
	if ok && len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	UserIdStr, ok := ctx.UserValue("user_id").(string)
	if ok && len(UserIdStr) != 0 {
		arg.UserId = UserIdStr
		if arg.UserId, err = url.PathUnescape(arg.UserId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field user_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV4GetMessageRequestV3(ctx *fasthttp.RequestCtx) (arg *common.GetMessageRequestV3, err error) {
	arg = &common.GetMessageRequestV3{}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "message_id":
			arg.MessageId = value
		case "user_id":
			arg.UserId = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	MessageIdStr, ok := ctx.UserValue("message_id").(string)
	if ok && len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
		arg.MessageId = fmt.Sprintf("base/%s", arg.MessageId)
	}

	return arg, err
}

func buildExampleServiceNameTopLevelArrayArray(ctx *fasthttp.RequestCtx) (arg *common.Array, err error) {
	arg = &common.Array{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "items[]":
			err = fmt.Errorf("unsupported type message for query argument items")
			return
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameUpdateMessageV3UpdateMessageRequest(ctx *fasthttp.RequestCtx) (arg *common.UpdateMessageRequest, err error) {
	arg = &common.UpdateMessageRequest{}
	var body = ctx.PostBody()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	ctx.QueryArgs().VisitAll(func(keyB, valueB []byte) {
		var key = string(keyB)
		var value = string(valueB)
		switch key {
		case "message_id":
			arg.MessageId = value
		case "message":
			err = fmt.Errorf("unsupported type message for query argument message")
			return
		case "message.message_id":
			if arg.Message == nil {
				arg.Message = &common.MessageV2{}
			}
			arg.Message.MessageId = value
		case "message.text":
			if arg.Message == nil {
				arg.Message = &common.MessageV2{}
			}
			arg.Message.Text = value
		default:
			err = fmt.Errorf("unknown query parameter %s with value %s", key, value)
			return
		}
	})
	return arg, err
}

func chainServerMiddlewaresExample(
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error) {
			return middlewares[0](ctx, req, getChainServerMiddlewareHandlerExample(middlewares, 0, handler))
		}
	}
}

func getChainServerMiddlewareHandlerExample(
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
	curr int,
	finalHandler func(ctx context.Context, req any) (resp any, err error),
) func(ctx context.Context, req any) (resp any, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req any) (resp any, err error) {
		return middlewares[curr+1](ctx, req, getChainServerMiddlewareHandlerExample(middlewares, curr+1, finalHandler))
	}
}

var _ ServiceNameHTTPGoService = &ServiceNameHTTPGoClient{}

type ServiceNameHTTPGoClient struct {
	cl          *fasthttp.Client
	host        string
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error)
	middleware  func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error)
}

func GetServiceNameHTTPGoClient(
	_ context.Context,
	cl *fasthttp.Client,
	host string,
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error),
) (*ServiceNameHTTPGoClient, error) {
	return &ServiceNameHTTPGoClient{
		cl:          cl,
		host:        host,
		middlewares: middlewares,
		middleware:  chainClientMiddlewaresExample(middlewares),
	}, nil
}

func (p *ServiceNameHTTPGoClient) RPCName(ctx context.Context, request *common.InputMsgName) (resp *common.OutputMsgName, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/RPCName/%s/%d%s", p.host, request.StringArgument, request.Int64Argument, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "RPCName")
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
	resp = &common.OutputMsgName{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) AllTypesTest(ctx context.Context, request *common.AllTypesMsg) (resp *common.AllTypesMsg, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/test/%t/%s/%d/%d/%d/%d/%d/%d/%d/%d/%f/%d/%d/%f/%s/%s%s", p.host, request.BoolValue, request.EnumValue, request.Int32Value, request.Sint32Value, request.Uint32Value, request.Int64Value, request.Sint64Value, request.Uint64Value, request.Sfixed32Value, request.Fixed32Value, request.FloatValue, request.Sfixed64Value, request.Fixed64Value, request.DoubleValue, request.StringValue, request.BytesValue, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "AllTypesTest")
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
	resp = &common.AllTypesMsg{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) AllTextTypesPost(ctx context.Context, request *common.AllTextTypesMsg) (resp *common.AllTextTypesMsg, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	RepeatedStringRequest := strings.Join(request.RepeatedString, ",")
	RepeatedBytesStrs := make([]string, len(request.RepeatedBytes))
	for i, v := range request.RepeatedBytes {
		RepeatedBytesStrs[i] = string(v)
	}
	RepeatedBytesRequest := strings.Join(RepeatedBytesStrs, ",")
	RepeatedEnumStrs := make([]string, len(request.RepeatedEnum))
	for i, v := range request.RepeatedEnum {
		RepeatedEnumStrs[i] = v.String()
	}
	RepeatedEnumRequest := strings.Join(RepeatedEnumStrs, ",")
	req.SetRequestURI(fmt.Sprintf("%s/v1/text/%s/%s/%s/%s/%s/%s%s", p.host, request.String_, RepeatedStringRequest, request.Bytes, RepeatedBytesRequest, request.Enum, RepeatedEnumRequest, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "AllTextTypesPost")
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
	resp = &common.AllTextTypesMsg{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) AllTextTypesGet(ctx context.Context, request *common.AllTextTypesMsg) (resp *common.AllTextTypesMsg, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{}
	var values = []any{}
	if request.OptionalString != nil {
		parameters = append(parameters, "OptionalString=%s")
		values = append(values, *request.OptionalString)
	}
	if request.OptionalBytes != nil {
		parameters = append(parameters, "OptionalBytes=%s")
		values = append(values, request.OptionalBytes)
	}
	if request.OptionalEnum != nil {
		parameters = append(parameters, "OptionalEnum=%s")
		values = append(values, *request.OptionalEnum)
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	RepeatedStringRequest := strings.Join(request.RepeatedString, ",")
	RepeatedBytesStrs := make([]string, len(request.RepeatedBytes))
	for i, v := range request.RepeatedBytes {
		RepeatedBytesStrs[i] = string(v)
	}
	RepeatedBytesRequest := strings.Join(RepeatedBytesStrs, ",")
	RepeatedEnumStrs := make([]string, len(request.RepeatedEnum))
	for i, v := range request.RepeatedEnum {
		RepeatedEnumStrs[i] = v.String()
	}
	RepeatedEnumRequest := strings.Join(RepeatedEnumStrs, ",")
	req.SetRequestURI(fmt.Sprintf("%s/v2/text/%s/%s/%s/%s/%s/%s%s", p.host, request.String_, RepeatedStringRequest, request.Bytes, RepeatedBytesRequest, request.Enum, RepeatedEnumRequest, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "AllTextTypesGet")
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
	resp = &common.AllTextTypesMsg{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) CommonTypes(ctx context.Context, request *anypb.Any) (resp *emptypb.Empty, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/test/commonTypes%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "CommonTypes")
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
	resp = &emptypb.Empty{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

// SameInputAndOutput same types but different query, we need different query builder function
func (p *ServiceNameHTTPGoClient) SameInputAndOutput(ctx context.Context, request *common.InputMsgName) (resp *common.OutputMsgName, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/sameInputAndOutput/%s%s", p.host, request.StringArgument, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "SameInputAndOutput")
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
	resp = &common.OutputMsgName{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) Optional(ctx context.Context, request *common.OptionalField) (resp *common.OptionalField, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/test/optional%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "Optional")
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
	resp = &common.OptionalField{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) GetMethod(ctx context.Context, request *common.InputMsgName) (resp *common.OutputMsgName, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{
		"int64Argument=%d",
		"stringArgument=%s",
	}
	var values = []any{
		request.Int64Argument,
		request.StringArgument,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	req.SetRequestURI(fmt.Sprintf("%s/v1/test/get%s", p.host, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "GetMethod")
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
	resp = &common.OutputMsgName{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) CheckRepeatedPath(ctx context.Context, request *common.RepeatedCheck) (resp *common.RepeatedCheck, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	BoolValueStrs := make([]string, len(request.BoolValue))
	for i, v := range request.BoolValue {
		BoolValueStrs[i] = strconv.FormatBool(v)
	}
	BoolValueRequest := strings.Join(BoolValueStrs, ",")
	EnumValueStrs := make([]string, len(request.EnumValue))
	for i, v := range request.EnumValue {
		EnumValueStrs[i] = v.String()
	}
	EnumValueRequest := strings.Join(EnumValueStrs, ",")
	Int32ValueStrs := make([]string, len(request.Int32Value))
	for i, v := range request.Int32Value {
		Int32ValueStrs[i] = strconv.FormatInt(int64(v), 10)
	}
	Int32ValueRequest := strings.Join(Int32ValueStrs, ",")
	Sint32ValueStrs := make([]string, len(request.Sint32Value))
	for i, v := range request.Sint32Value {
		Sint32ValueStrs[i] = strconv.FormatInt(int64(v), 10)
	}
	Sint32ValueRequest := strings.Join(Sint32ValueStrs, ",")
	Uint32ValueStrs := make([]string, len(request.Uint32Value))
	for i, v := range request.Uint32Value {
		Uint32ValueStrs[i] = strconv.FormatInt(int64(v), 10)
	}
	Uint32ValueRequest := strings.Join(Uint32ValueStrs, ",")
	Int64ValueStrs := make([]string, len(request.Int64Value))
	for i, v := range request.Int64Value {
		Int64ValueStrs[i] = strconv.FormatInt(v, 10)
	}
	Int64ValueRequest := strings.Join(Int64ValueStrs, ",")
	Sint64ValueStrs := make([]string, len(request.Sint64Value))
	for i, v := range request.Sint64Value {
		Sint64ValueStrs[i] = strconv.FormatInt(v, 10)
	}
	Sint64ValueRequest := strings.Join(Sint64ValueStrs, ",")
	Uint64ValueStrs := make([]string, len(request.Uint64Value))
	for i, v := range request.Uint64Value {
		Uint64ValueStrs[i] = strconv.FormatUint(v, 10)
	}
	Uint64ValueRequest := strings.Join(Uint64ValueStrs, ",")
	Sfixed32ValueStrs := make([]string, len(request.Sfixed32Value))
	for i, v := range request.Sfixed32Value {
		Sfixed32ValueStrs[i] = strconv.FormatInt(int64(v), 10)
	}
	Sfixed32ValueRequest := strings.Join(Sfixed32ValueStrs, ",")
	Fixed32ValueStrs := make([]string, len(request.Fixed32Value))
	for i, v := range request.Fixed32Value {
		Fixed32ValueStrs[i] = strconv.FormatInt(int64(v), 10)
	}
	Fixed32ValueRequest := strings.Join(Fixed32ValueStrs, ",")
	FloatValueStrs := make([]string, len(request.FloatValue))
	for i, v := range request.FloatValue {
		FloatValueStrs[i] = strconv.FormatFloat(float64(v), 'f', -1, 64)
	}
	FloatValueRequest := strings.Join(FloatValueStrs, ",")
	Sfixed64ValueStrs := make([]string, len(request.Sfixed64Value))
	for i, v := range request.Sfixed64Value {
		Sfixed64ValueStrs[i] = strconv.FormatInt(v, 10)
	}
	Sfixed64ValueRequest := strings.Join(Sfixed64ValueStrs, ",")
	Fixed64ValueStrs := make([]string, len(request.Fixed64Value))
	for i, v := range request.Fixed64Value {
		Fixed64ValueStrs[i] = strconv.FormatUint(v, 10)
	}
	Fixed64ValueRequest := strings.Join(Fixed64ValueStrs, ",")
	DoubleValueStrs := make([]string, len(request.DoubleValue))
	for i, v := range request.DoubleValue {
		DoubleValueStrs[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	DoubleValueRequest := strings.Join(DoubleValueStrs, ",")
	StringValueRequest := strings.Join(request.StringValue, ",")
	BytesValueStrs := make([]string, len(request.BytesValue))
	for i, v := range request.BytesValue {
		BytesValueStrs[i] = string(v)
	}
	BytesValueRequest := strings.Join(BytesValueStrs, ",")
	StringValueQueryRequest := strings.Join(request.StringValueQuery, ",")
	req.SetRequestURI(fmt.Sprintf("%s/v1/repeated/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s%s", p.host, BoolValueRequest, EnumValueRequest, Int32ValueRequest, Sint32ValueRequest, Uint32ValueRequest, Int64ValueRequest, Sint64ValueRequest, Uint64ValueRequest, Sfixed32ValueRequest, Fixed32ValueRequest, FloatValueRequest, Sfixed64ValueRequest, Fixed64ValueRequest, DoubleValueRequest, StringValueRequest, BytesValueRequest, StringValueQueryRequest, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "CheckRepeatedPath")
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
	resp = &common.RepeatedCheck{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) CheckRepeatedQuery(ctx context.Context, request *common.RepeatedCheck) (resp *common.RepeatedCheck, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{}
	var values = []any{}
	for _, v := range request.BoolValue {
		parameters = append(parameters, "BoolValue[]=%t")
		values = append(values, v)
	}
	for _, v := range request.EnumValue {
		parameters = append(parameters, "EnumValue[]=%s")
		values = append(values, v)
	}
	for _, v := range request.Int32Value {
		parameters = append(parameters, "Int32Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Sint32Value {
		parameters = append(parameters, "Sint32Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Uint32Value {
		parameters = append(parameters, "Uint32Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Int64Value {
		parameters = append(parameters, "Int64Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Sint64Value {
		parameters = append(parameters, "Sint64Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Uint64Value {
		parameters = append(parameters, "Uint64Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Sfixed32Value {
		parameters = append(parameters, "Sfixed32Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Fixed32Value {
		parameters = append(parameters, "Fixed32Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.FloatValue {
		parameters = append(parameters, "FloatValue[]=%f")
		values = append(values, v)
	}
	for _, v := range request.Sfixed64Value {
		parameters = append(parameters, "Sfixed64Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.Fixed64Value {
		parameters = append(parameters, "Fixed64Value[]=%d")
		values = append(values, v)
	}
	for _, v := range request.DoubleValue {
		parameters = append(parameters, "DoubleValue[]=%f")
		values = append(values, v)
	}
	for _, v := range request.BytesValue {
		parameters = append(parameters, "BytesValue[]=%s")
		values = append(values, v)
	}
	for _, v := range request.StringValueQuery {
		parameters = append(parameters, "StringValueQuery[]=%s")
		values = append(values, v)
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	StringValueRequest := strings.Join(request.StringValue, ",")
	req.SetRequestURI(fmt.Sprintf("%s/v2/repeated/%s%s", p.host, StringValueRequest, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "CheckRepeatedQuery")
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
	resp = &common.RepeatedCheck{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) CheckRepeatedPost(ctx context.Context, request *common.RepeatedCheck) (resp *common.RepeatedCheck, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	StringValueRequest := strings.Join(request.StringValue, ",")
	req.SetRequestURI(fmt.Sprintf("%s/v3/repeated/%s%s", p.host, StringValueRequest, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "CheckRepeatedPost")
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
	resp = &common.RepeatedCheck{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) EmptyGet(ctx context.Context, request *common.Empty) (resp *common.Empty, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	req.SetRequestURI(fmt.Sprintf("%s/v1/emptyGet%s", p.host, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "EmptyGet")
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
	resp = &common.Empty{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) EmptyPost(ctx context.Context, request *common.Empty) (resp *common.Empty, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/emptyPost%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "EmptyPost")
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
	resp = &common.Empty{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) OnlyStructInGet(ctx context.Context, request *common.OnlyStruct) (resp *common.Empty, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/onlyStruct%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "OnlyStructInGet")
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
	resp = &common.Empty{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) MultipartForm(ctx context.Context, request *common.MultipartFormRequest) (resp *common.Empty, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("document", request.Document.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file document:  %w", err)
	}
	if _, err = part.Write(request.Document.File); err != nil {
		return nil, fmt.Errorf("failed to write data to part document: %w", err)
	}
	if err = writer.WriteField("otherField", request.OtherField); err != nil {
		return nil, fmt.Errorf("failed to write field otherField:  %w", err)
	}
	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}
	req.SetBody(requestBody.Bytes())
	req.SetRequestURI(fmt.Sprintf("%s/v1/multipart%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "MultipartForm")
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
	resp = &common.Empty{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) MultipartFormAllTypes(ctx context.Context, request *common.MultipartFormAllTypes) (resp *common.Empty, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	if err = writer.WriteField("BoolValue", strconv.FormatBool(request.BoolValue)); err != nil {
		return nil, fmt.Errorf("failed to write field BoolValue:  %w", err)
	}
	if err = writer.WriteField("EnumValue", request.EnumValue.String()); err != nil {
		return nil, fmt.Errorf("failed to write field EnumValue:  %w", err)
	}
	if err = writer.WriteField("Int32Value", strconv.FormatInt(int64(request.Int32Value), 10)); err != nil {
		return nil, fmt.Errorf("failed to write field Int32Value:  %w", err)
	}
	if err = writer.WriteField("Sint32Value", strconv.FormatInt(int64(request.Sint32Value), 10)); err != nil {
		return nil, fmt.Errorf("failed to write field Sint32Value:  %w", err)
	}
	for _, value := range request.Uint32Value {
		if err = writer.WriteField("Uint32Value", strconv.FormatInt(int64(value), 10)); err != nil {
			return nil, fmt.Errorf("failed to write field Uint32Value:  %w", err)
		}
	}
	if err = writer.WriteField("Int64Value", strconv.FormatInt(request.Int64Value, 10)); err != nil {
		return nil, fmt.Errorf("failed to write field Int64Value:  %w", err)
	}
	if request.Sint64Value != nil {
		if err = writer.WriteField("Sint64Value", strconv.FormatInt(*request.Sint64Value, 10)); err != nil {
			return nil, fmt.Errorf("failed to write field Sint64Value:  %w", err)
		}
	}
	if err = writer.WriteField("Uint64Value", strconv.FormatUint(request.Uint64Value, 10)); err != nil {
		return nil, fmt.Errorf("failed to write field Uint64Value:  %w", err)
	}
	if err = writer.WriteField("Sfixed32Value", strconv.FormatInt(int64(request.Sfixed32Value), 10)); err != nil {
		return nil, fmt.Errorf("failed to write field Sfixed32Value:  %w", err)
	}
	for _, value := range request.Fixed32Value {
		if err = writer.WriteField("Fixed32Value", strconv.FormatInt(int64(value), 10)); err != nil {
			return nil, fmt.Errorf("failed to write field Fixed32Value:  %w", err)
		}
	}
	if err = writer.WriteField("FloatValue", strconv.FormatFloat(float64(request.FloatValue), 'f', -1, 64)); err != nil {
		return nil, fmt.Errorf("failed to write field FloatValue:  %w", err)
	}
	if err = writer.WriteField("Sfixed64Value", strconv.FormatInt(request.Sfixed64Value, 10)); err != nil {
		return nil, fmt.Errorf("failed to write field Sfixed64Value:  %w", err)
	}
	if request.Fixed64Value != nil {
		if err = writer.WriteField("Fixed64Value", strconv.FormatUint(*request.Fixed64Value, 10)); err != nil {
			return nil, fmt.Errorf("failed to write field Fixed64Value:  %w", err)
		}
	}
	if err = writer.WriteField("DoubleValue", strconv.FormatFloat(request.DoubleValue, 'f', -1, 64)); err != nil {
		return nil, fmt.Errorf("failed to write field DoubleValue:  %w", err)
	}
	if err = writer.WriteField("StringValue", request.StringValue); err != nil {
		return nil, fmt.Errorf("failed to write field StringValue:  %w", err)
	}
	if err = writer.WriteField("BytesValue", string(request.BytesValue)); err != nil {
		return nil, fmt.Errorf("failed to write field BytesValue:  %w", err)
	}
	for _, value := range request.SliceStringValue {
		if err = writer.WriteField("SliceStringValue", value); err != nil {
			return nil, fmt.Errorf("failed to write field SliceStringValue:  %w", err)
		}
	}
	for _, value := range request.SliceInt32Value {
		if err = writer.WriteField("SliceInt32Value", strconv.FormatInt(int64(value), 10)); err != nil {
			return nil, fmt.Errorf("failed to write field SliceInt32Value:  %w", err)
		}
	}
	part, err := writer.CreateFormFile("document", request.Document.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file document:  %w", err)
	}
	if _, err = part.Write(request.Document.File); err != nil {
		return nil, fmt.Errorf("failed to write data to part document: %w", err)
	}
	for _, value := range request.RepeatedStringValue {
		if err = writer.WriteField("RepeatedStringValue", value); err != nil {
			return nil, fmt.Errorf("failed to write field RepeatedStringValue:  %w", err)
		}
	}
	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}
	req.SetBody(requestBody.Bytes())
	req.SetRequestURI(fmt.Sprintf("%s/v1/multipartall%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "MultipartFormAllTypes")
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
	resp = &common.Empty{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) AllTypesMaxTest(ctx context.Context, request *common.AllNumberTypesMsg) (resp *common.AllNumberTypesMsg, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{
		"Sint32Value=%d",
		"Sint64Value=%d",
		"Sfixed32Value=%d",
		"Fixed32Value=%d",
		"Sfixed64Value=%d",
		"Fixed64Value=%d",
	}
	var values = []any{
		request.Sint32Value,
		request.Sint64Value,
		request.Sfixed32Value,
		request.Fixed32Value,
		request.Sfixed64Value,
		request.Fixed64Value,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	req.SetRequestURI(fmt.Sprintf("%s/v1/max/%d/%d/%d/%d/%f/%f%s", p.host, request.Int32Value, request.Uint32Value, request.Int64Value, request.Uint64Value, request.FloatValue, request.DoubleValue, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "AllTypesMaxTest")
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
	resp = &common.AllNumberTypesMsg{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) AllTypesMaxQueryTest(ctx context.Context, request *common.AllNumberTypesMsg) (resp *common.AllNumberTypesMsg, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{
		"Int32Value=%d",
		"Sint32Value=%d",
		"Uint32Value=%d",
		"Int64Value=%d",
		"Sint64Value=%d",
		"Uint64Value=%d",
		"Sfixed32Value=%d",
		"Fixed32Value=%d",
		"FloatValue=%f",
		"Sfixed64Value=%d",
		"Fixed64Value=%d",
		"DoubleValue=%f",
	}
	var values = []any{
		request.Int32Value,
		request.Sint32Value,
		request.Uint32Value,
		request.Int64Value,
		request.Sint64Value,
		request.Uint64Value,
		request.Sfixed32Value,
		request.Fixed32Value,
		request.FloatValue,
		request.Sfixed64Value,
		request.Fixed64Value,
		request.DoubleValue,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	req.SetRequestURI(fmt.Sprintf("%s/v1/maxquery%s", p.host, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "AllTypesMaxQueryTest")
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
	resp = &common.AllNumberTypesMsg{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

// GetMessage http rule checks
// v1/{name=messages/*}
func (p *ServiceNameHTTPGoClient) GetMessage(ctx context.Context, request *common.GetMessageRequest) (resp *common.Message, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	req.SetRequestURI(fmt.Sprintf("%s/v1/messages/%s%s", p.host, request.Name, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "GetMessage")
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
	resp = &common.Message{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) GetMessageV2(ctx context.Context, request *common.GetMessageRequestV2) (resp *common.Message, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{
		"revision=%d",
		"sub.subfield=%s",
	}
	var values = []any{
		request.Revision,
		request.Sub.Subfield,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	req.SetRequestURI(fmt.Sprintf("%s/v2/messages/%s%s", p.host, request.MessageId, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "GetMessageV2")
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
	resp = &common.Message{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) UpdateMessage(ctx context.Context, request *common.UpdateMessageRequest) (resp *common.Message, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request.Message)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/messages/%s%s", p.host, request.MessageId, queryArgs))
	req.Header.SetMethod("PATCH")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "UpdateMessage")
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
	resp = &common.Message{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) UpdateMessageV2(ctx context.Context, request *common.MessageV2) (resp *common.MessageV2, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v2/messages/%s%s", p.host, request.MessageId, queryArgs))
	req.Header.SetMethod("PATCH")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "UpdateMessageV2")
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
	resp = &common.MessageV2{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) GetMessageV3(ctx context.Context, request *common.GetMessageRequestV3) (resp *common.MessageV2, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{
		"user_id=%s",
	}
	var values = []any{
		request.UserId,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	req.SetRequestURI(fmt.Sprintf("%s/v3/messages/%s%s", p.host, request.MessageId, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "GetMessageV3")
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
	resp = &common.MessageV2{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) GetMessageV4(ctx context.Context, request *common.GetMessageRequestV3) (resp *common.MessageV2, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var parameters = []string{
		"user_id=%s",
	}
	var values = []any{
		request.UserId,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	queryArgs = strings.ReplaceAll(queryArgs, "[]", "%5B%5D")
	req.SetRequestURI(fmt.Sprintf("%s/v4/messages/base/%s%s", p.host, request.MessageId, queryArgs))
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "GetMessageV4")
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
	resp = &common.MessageV2{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, resp)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) TopLevelArray(ctx context.Context, request *common.Array) (resp *common.Array, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v1/array%s", p.host, queryArgs))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "TopLevelArray")
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
	resp = &common.Array{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, &resp.Items)
	return resp, err
}

func (p *ServiceNameHTTPGoClient) UpdateMessageV3(ctx context.Context, request *common.UpdateMessageRequest) (resp *common.UpdateMessageRequest, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	var queryArgs string
	var body []byte
	body, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(fmt.Sprintf("%s/v3/messages%s", p.host, queryArgs))
	req.Header.SetMethod("PATCH")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var reqResp *fasthttp.Response
	ctx = context.WithValue(ctx, "proto_service", "ServiceName")
	ctx = context.WithValue(ctx, "proto_method", "UpdateMessageV3")
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
	resp = &common.UpdateMessageRequest{}
	var respBody = reqResp.Body()
	err = json.Unmarshal(respBody, &resp.Message)
	return resp, err
}

func chainClientMiddlewaresExample(
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error),
) func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error) {
			return middlewares[0](ctx, req, getChainClientMiddlewareHandlerExample(middlewares, 0, handler))
		}
	}
}

func getChainClientMiddlewareHandlerExample(
	middlewares []func(ctx context.Context, req *fasthttp.Request, handler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error),
	curr int,
	finalHandler func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error),
) func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req *fasthttp.Request) (resp *fasthttp.Response, err error) {
		return middlewares[curr+1](ctx, req, getChainClientMiddlewareHandlerExample(middlewares, curr+1, finalHandler))
	}
}
