// source: example.proto

package proto

import (
	context "context"
	fmt "fmt"
	common "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	gin "github.com/gin-gonic/gin"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	multipart "mime/multipart"
	strconv "strconv"
	strings "strings"
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
	r *gin.Engine,
	h ServiceNameHTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresExample(middlewares)

	r.POST("/v1/RPCName/:stringArgument/:int64Argument", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameRPCNameInputMsgName(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/test/:BoolValue/:EnumValue/:Int32Value/:Sint32Value/:Uint32Value/:Int64Value/:Sint64Value/:Uint64Value/:Sfixed32Value/:Fixed32Value/:FloatValue/:Sfixed64Value/:Fixed64Value/:DoubleValue/:StringValue/:BytesValue", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTypesTestAllTypesMsg(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/text/:String/:RepeatedString/:Bytes/:RepeatedBytes/:Enum/:RepeatedEnum", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTextTypesPostAllTextTypesMsg(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v2/text/:String/:RepeatedString/:Bytes/:RepeatedBytes/:Enum/:RepeatedEnum", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTextTypesGetAllTextTypesMsg(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/test/commonTypes", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameCommonTypesAny(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	// same types but different query, we need different query builder function
	r.POST("/v1/sameInputAndOutput/:stringArgument", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameSameInputAndOutputInputMsgName(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/test/optional", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameOptionalOptionalField(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v1/test/get", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMethodInputMsgName(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v1/repeated/:BoolValue/:EnumValue/:Int32Value/:Sint32Value/:Uint32Value/:Int64Value/:Sint64Value/:Uint64Value/:Sfixed32Value/:Fixed32Value/:FloatValue/:Sfixed64Value/:Fixed64Value/:DoubleValue/:StringValue/:BytesValue/:StringValueQuery", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameCheckRepeatedPathRepeatedCheck(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v2/repeated/:StringValue", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v3/repeated/:StringValue", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameCheckRepeatedPostRepeatedCheck(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v1/emptyGet", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameEmptyGetEmpty(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/emptyPost", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameEmptyPostEmpty(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/onlyStruct", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameOnlyStructInGetOnlyStruct(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/multipart", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameMultipartFormMultipartFormRequest(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/multipartall", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameMultipartFormAllTypesMultipartFormAllTypes(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v1/max/:Int32Value/:Uint32Value/:Int64Value/:Uint64Value/:FloatValue/:DoubleValue", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTypesMaxTestAllNumberTypesMsg(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v1/maxquery", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTypesMaxQueryTestAllNumberTypesMsg(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	// http rule checks
	// v1/{name=messages/*}
	r.GET("/v1/messages/:name", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageGetMessageRequest(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v2/messages/:message_id", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV2GetMessageRequestV2(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.PATCH("/v1/messages/:message_id", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameUpdateMessageUpdateMessageRequest(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.PATCH("/v2/messages/:message_id", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameUpdateMessageV2MessageV2(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v3/messages/:message_id", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV3GetMessageRequestV3(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v3/users/:user_id/messages/:message_id", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV3GetMessageRequestV3(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.GET("/v4/messages/base/*message_id", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV4GetMessageRequestV3(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
		ginctx.JSON(ginctx.Writer.Status(), resp)
	})

	r.POST("/v1/array", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameTopLevelArrayArray(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
			ginctx.JSON(ginctx.Writer.Status(), typedResp.Items)
		} else {
			ginctx.JSON(ginctx.Writer.Status(), resp)
		}
	})

	r.PATCH("/v3/messages", func(ginctx *gin.Context) {
		ginctx.Header("Content-Type", "application/json")
		input, err := buildExampleServiceNameUpdateMessageV3UpdateMessageRequest(ginctx)
		if err != nil {
			ginctx.JSON(400, struct{ Error string }{Error: err.Error()})
			return
		}
		ctx := context.WithValue(ginctx, "request", ginctx)
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
			ginctx.JSON(ginctx.Writer.Status(), typedResp.Message)
		} else {
			ginctx.JSON(ginctx.Writer.Status(), resp)
		}
	})

	return nil
}

func buildExampleServiceNameRPCNameInputMsgName(ctx *gin.Context) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "int64Argument":
				arg.Int64Argument, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				}
			case "stringArgument":
				arg.StringArgument = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	StringArgumentStr := ctx.Param("stringArgument")
	if len(StringArgumentStr) != 0 {
		arg.StringArgument = StringArgumentStr
	}

	Int64ArgumentStr := ctx.Param("int64Argument")
	if len(Int64ArgumentStr) != 0 {
		arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameAllTypesTestAllTypesMsg(ctx *gin.Context) (arg *common.AllTypesMsg, err error) {
	arg = &common.AllTypesMsg{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "BoolValue":
				switch value {
				case "true", "t", "1":
					arg.BoolValue = true
				case "false", "f", "0":
					arg.BoolValue = false
				default:
					return nil, fmt.Errorf("unknown bool string value %s", value)
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
						return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					}
				}
			case "Int32Value":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				arg.Int32Value = int32(Int32Value)
			case "Sint32Value":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				arg.Sint32Value = int32(Sint32Value)
			case "Uint32Value":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				arg.Uint32Value = uint32(Uint32Value)
			case "Int64Value":
				arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				}
			case "Sint64Value":
				arg.Sint64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
				}
			case "Uint64Value":
				arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				}
			case "Sfixed32Value":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				arg.Sfixed32Value = int32(Sfixed32Value)
			case "Fixed32Value":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				arg.Fixed32Value = uint32(Fixed32Value)
			case "FloatValue":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				arg.FloatValue = float32(FloatValue)
			case "Sfixed64Value":
				arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				}
			case "Fixed64Value":
				arg.Fixed64Value, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
				}
			case "DoubleValue":
				arg.DoubleValue, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				}
			case "StringValue":
				arg.StringValue = value
			case "BytesValue":
				arg.BytesValue = []byte(value)
			case "MessageValue":
				return nil, fmt.Errorf("unsupported type message for query argument MessageValue")
			case "MessageValue.int64Argument":
				if arg.MessageValue == nil {
					arg.MessageValue = &common.InputMsgName{}
				}
				arg.MessageValue.Int64Argument, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
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
					return nil, fmt.Errorf("conversion failed for parameter SliceInt32Value: %w", convErr)
				}
				arg.SliceInt32Value = append(arg.SliceInt32Value, int32(SliceInt32Value))
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	BoolValueStr := ctx.Param("BoolValue")
	if len(BoolValueStr) != 0 {
		switch BoolValueStr {
		case "true", "t", "1":
			arg.BoolValue = true
		case "false", "f", "0":
			arg.BoolValue = false
		default:
			return nil, fmt.Errorf("unknown bool string value %s", BoolValueStr)
		}
	}

	EnumValueStr := ctx.Param("EnumValue")
	if len(EnumValueStr) != 0 {
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

	Int32ValueStr := ctx.Param("Int32Value")
	if len(Int32ValueStr) != 0 {
		Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}

	Sint32ValueStr := ctx.Param("Sint32Value")
	if len(Sint32ValueStr) != 0 {
		Sint32Value, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
		}
		arg.Sint32Value = int32(Sint32Value)
	}

	Uint32ValueStr := ctx.Param("Uint32Value")
	if len(Uint32ValueStr) != 0 {
		Uint32Value, convErr := strconv.ParseUint(Uint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
		}
		arg.Uint32Value = uint32(Uint32Value)
	}

	Int64ValueStr := ctx.Param("Int64Value")
	if len(Int64ValueStr) != 0 {
		arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}

	Sint64ValueStr := ctx.Param("Sint64Value")
	if len(Sint64ValueStr) != 0 {
		arg.Sint64Value, err = strconv.ParseInt(Sint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
		}
	}

	Uint64ValueStr := ctx.Param("Uint64Value")
	if len(Uint64ValueStr) != 0 {
		arg.Uint64Value, err = strconv.ParseUint(Uint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}

	Sfixed32ValueStr := ctx.Param("Sfixed32Value")
	if len(Sfixed32ValueStr) != 0 {
		Sfixed32Value, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
		}
		arg.Sfixed32Value = int32(Sfixed32Value)
	}

	Fixed32ValueStr := ctx.Param("Fixed32Value")
	if len(Fixed32ValueStr) != 0 {
		Fixed32Value, convErr := strconv.ParseUint(Fixed32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
		}
		arg.Fixed32Value = uint32(Fixed32Value)
	}

	FloatValueStr := ctx.Param("FloatValue")
	if len(FloatValueStr) != 0 {
		FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}

	Sfixed64ValueStr := ctx.Param("Sfixed64Value")
	if len(Sfixed64ValueStr) != 0 {
		arg.Sfixed64Value, err = strconv.ParseInt(Sfixed64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
		}
	}

	Fixed64ValueStr := ctx.Param("Fixed64Value")
	if len(Fixed64ValueStr) != 0 {
		arg.Fixed64Value, err = strconv.ParseUint(Fixed64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
		}
	}

	DoubleValueStr := ctx.Param("DoubleValue")
	if len(DoubleValueStr) != 0 {
		arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}

	StringValueStr := ctx.Param("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = StringValueStr
	}

	BytesValueStr := ctx.Param("BytesValue")
	if len(BytesValueStr) != 0 {
		arg.BytesValue = []byte(BytesValueStr)
	}

	return arg, err
}

func buildExampleServiceNameAllTextTypesPostAllTextTypesMsg(ctx *gin.Context) (arg *common.AllTextTypesMsg, err error) {
	arg = &common.AllTextTypesMsg{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
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
						return nil, fmt.Errorf("conversion failed for parameter Enum: %w", convErr)
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
						return nil, fmt.Errorf("conversion failed for parameter RepeatedEnum: %w", convErr)
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
						return nil, fmt.Errorf("conversion failed for parameter OptionalEnum: %w", convErr)
					}
				}
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	String_Str := ctx.Param("String")
	if len(String_Str) != 0 {
		arg.String_ = String_Str
	}

	RepeatedStringStr := ctx.Param("RepeatedString")
	if len(RepeatedStringStr) != 0 {
		arg.RepeatedString = strings.Split(RepeatedStringStr, ",")
	}

	BytesStr := ctx.Param("Bytes")
	if len(BytesStr) != 0 {
		arg.Bytes = []byte(BytesStr)
	}

	RepeatedBytesStr := ctx.Param("RepeatedBytes")
	if len(RepeatedBytesStr) != 0 {
		RepeatedBytesStrs := strings.Split(RepeatedBytesStr, ",")
		arg.RepeatedBytes = make([][]byte, 0, len(RepeatedBytesStrs))
		for _, str := range RepeatedBytesStrs {
			arg.RepeatedBytes = append(arg.RepeatedBytes, []byte(str))
		}
	}

	EnumStr := ctx.Param("Enum")
	if len(EnumStr) != 0 {
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

	RepeatedEnumStr := ctx.Param("RepeatedEnum")
	if len(RepeatedEnumStr) != 0 {
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

func buildExampleServiceNameAllTextTypesGetAllTextTypesMsg(ctx *gin.Context) (arg *common.AllTextTypesMsg, err error) {
	arg = &common.AllTextTypesMsg{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
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
						return nil, fmt.Errorf("conversion failed for parameter Enum: %w", convErr)
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
						return nil, fmt.Errorf("conversion failed for parameter RepeatedEnum: %w", convErr)
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
						return nil, fmt.Errorf("conversion failed for parameter OptionalEnum: %w", convErr)
					}
				}
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	String_Str := ctx.Param("String")
	if len(String_Str) != 0 {
		arg.String_ = String_Str
	}

	RepeatedStringStr := ctx.Param("RepeatedString")
	if len(RepeatedStringStr) != 0 {
		arg.RepeatedString = strings.Split(RepeatedStringStr, ",")
	}

	BytesStr := ctx.Param("Bytes")
	if len(BytesStr) != 0 {
		arg.Bytes = []byte(BytesStr)
	}

	RepeatedBytesStr := ctx.Param("RepeatedBytes")
	if len(RepeatedBytesStr) != 0 {
		RepeatedBytesStrs := strings.Split(RepeatedBytesStr, ",")
		arg.RepeatedBytes = make([][]byte, 0, len(RepeatedBytesStrs))
		for _, str := range RepeatedBytesStrs {
			arg.RepeatedBytes = append(arg.RepeatedBytes, []byte(str))
		}
	}

	EnumStr := ctx.Param("Enum")
	if len(EnumStr) != 0 {
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

	RepeatedEnumStr := ctx.Param("RepeatedEnum")
	if len(RepeatedEnumStr) != 0 {
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

func buildExampleServiceNameCommonTypesAny(ctx *gin.Context) (arg *anypb.Any, err error) {
	arg = &anypb.Any{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "type_url":
				arg.TypeUrl = value
			case "value":
				arg.Value = []byte(value)
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameSameInputAndOutputInputMsgName(ctx *gin.Context) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "int64Argument":
				arg.Int64Argument, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				}
			case "stringArgument":
				arg.StringArgument = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	StringArgumentStr := ctx.Param("stringArgument")
	if len(StringArgumentStr) != 0 {
		arg.StringArgument = StringArgumentStr
	}

	return arg, err
}

func buildExampleServiceNameOptionalOptionalField(ctx *gin.Context) (arg *common.OptionalField, err error) {
	arg = &common.OptionalField{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
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
					return nil, fmt.Errorf("unknown bool string value %s", value)
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
						return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					}
				}
			case "Int32Value":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				Int32ValueValue := int32(Int32Value)
				arg.Int32Value = &Int32ValueValue
			case "Sint32Value":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				Sint32ValueValue := int32(Sint32Value)
				arg.Sint32Value = &Sint32ValueValue
			case "Uint32Value":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				Uint32ValueValue := uint32(Uint32Value)
				arg.Uint32Value = &Uint32ValueValue
			case "Int64Value":
				Int64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				}
				arg.Int64Value = &Int64Value
			case "Sint64Value":
				Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				}
				arg.Sint64Value = &Sint64Value
			case "Uint64Value":
				Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				}
				arg.Uint64Value = &Uint64Value
			case "Sfixed32Value":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				Sfixed32ValueValue := int32(Sfixed32Value)
				arg.Sfixed32Value = &Sfixed32ValueValue
			case "Fixed32Value":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				Fixed32ValueValue := uint32(Fixed32Value)
				arg.Fixed32Value = &Fixed32ValueValue
			case "FloatValue":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				FloatValueValue := float32(FloatValue)
				arg.FloatValue = &FloatValueValue
			case "Sfixed64Value":
				Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				}
				arg.Sfixed64Value = &Sfixed64Value
			case "Fixed64Value":
				Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				}
				arg.Fixed64Value = &Fixed64Value
			case "DoubleValue":
				DoubleValue, convErr := strconv.ParseFloat(value, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				}
				arg.DoubleValue = &DoubleValue
			case "StringValue":
				arg.StringValue = &value
			case "BytesValue":
				arg.BytesValue = []byte(value)
			case "MessageValue":
				return nil, fmt.Errorf("unsupported type message for query argument MessageValue")
			case "MessageValue.int64Argument":
				if arg.MessageValue == nil {
					arg.MessageValue = &common.InputMsgName{}
				}
				arg.MessageValue.Int64Argument, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				}
			case "MessageValue.stringArgument":
				if arg.MessageValue == nil {
					arg.MessageValue = &common.InputMsgName{}
				}
				arg.MessageValue.StringArgument = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameGetMethodInputMsgName(ctx *gin.Context) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "int64Argument":
				arg.Int64Argument, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
				}
			case "stringArgument":
				arg.StringArgument = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameCheckRepeatedPathRepeatedCheck(ctx *gin.Context) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "BoolValue[]":
				switch value {
				case "true", "t", "1":
					arg.BoolValue = append(arg.BoolValue, true)
				case "false", "f", "0":
					arg.BoolValue = append(arg.BoolValue, false)
				default:
					return nil, fmt.Errorf("unknown bool string value %s", value)
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
						return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					}
				}
			case "Int32Value[]":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				arg.Int32Value = append(arg.Int32Value, int32(Int32Value))
			case "Sint32Value[]":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				arg.Sint32Value = append(arg.Sint32Value, int32(Sint32Value))
			case "Uint32Value[]":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
			case "Int64Value[]":
				Int64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				}
				arg.Int64Value = append(arg.Int64Value, Int64Value)
			case "Sint64Value[]":
				Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				}
				arg.Sint64Value = append(arg.Sint64Value, Sint64Value)
			case "Uint64Value[]":
				Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				}
				arg.Uint64Value = append(arg.Uint64Value, Uint64Value)
			case "Sfixed32Value[]":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32Value))
			case "Fixed32Value[]":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
			case "FloatValue[]":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				arg.FloatValue = append(arg.FloatValue, float32(FloatValue))
			case "Sfixed64Value[]":
				Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				}
				arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64Value)
			case "Fixed64Value[]":
				Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				}
				arg.Fixed64Value = append(arg.Fixed64Value, Fixed64Value)
			case "DoubleValue[]":
				DoubleValue, convErr := strconv.ParseFloat(value, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				}
				arg.DoubleValue = append(arg.DoubleValue, DoubleValue)
			case "StringValue[]":
				arg.StringValue = append(arg.StringValue, value)
			case "BytesValue[]":
				arg.BytesValue = append(arg.BytesValue, []byte(value))
			case "StringValueQuery[]":
				arg.StringValueQuery = append(arg.StringValueQuery, value)
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	BoolValueStr := ctx.Param("BoolValue")
	if len(BoolValueStr) != 0 {
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

	EnumValueStr := ctx.Param("EnumValue")
	if len(EnumValueStr) != 0 {
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

	Int32ValueStr := ctx.Param("Int32Value")
	if len(Int32ValueStr) != 0 {
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

	Sint32ValueStr := ctx.Param("Sint32Value")
	if len(Sint32ValueStr) != 0 {
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

	Uint32ValueStr := ctx.Param("Uint32Value")
	if len(Uint32ValueStr) != 0 {
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

	Int64ValueStr := ctx.Param("Int64Value")
	if len(Int64ValueStr) != 0 {
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

	Sint64ValueStr := ctx.Param("Sint64Value")
	if len(Sint64ValueStr) != 0 {
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

	Uint64ValueStr := ctx.Param("Uint64Value")
	if len(Uint64ValueStr) != 0 {
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

	Sfixed32ValueStr := ctx.Param("Sfixed32Value")
	if len(Sfixed32ValueStr) != 0 {
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

	Fixed32ValueStr := ctx.Param("Fixed32Value")
	if len(Fixed32ValueStr) != 0 {
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

	FloatValueStr := ctx.Param("FloatValue")
	if len(FloatValueStr) != 0 {
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

	Sfixed64ValueStr := ctx.Param("Sfixed64Value")
	if len(Sfixed64ValueStr) != 0 {
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

	Fixed64ValueStr := ctx.Param("Fixed64Value")
	if len(Fixed64ValueStr) != 0 {
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

	DoubleValueStr := ctx.Param("DoubleValue")
	if len(DoubleValueStr) != 0 {
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

	StringValueStr := ctx.Param("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
	}

	BytesValueStr := ctx.Param("BytesValue")
	if len(BytesValueStr) != 0 {
		BytesValueStrs := strings.Split(BytesValueStr, ",")
		arg.BytesValue = make([][]byte, 0, len(BytesValueStrs))
		for _, str := range BytesValueStrs {
			arg.BytesValue = append(arg.BytesValue, []byte(str))
		}
	}

	StringValueQueryStr := ctx.Param("StringValueQuery")
	if len(StringValueQueryStr) != 0 {
		arg.StringValueQuery = strings.Split(StringValueQueryStr, ",")
	}

	return arg, err
}

func buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(ctx *gin.Context) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "BoolValue[]":
				switch value {
				case "true", "t", "1":
					arg.BoolValue = append(arg.BoolValue, true)
				case "false", "f", "0":
					arg.BoolValue = append(arg.BoolValue, false)
				default:
					return nil, fmt.Errorf("unknown bool string value %s", value)
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
						return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					}
				}
			case "Int32Value[]":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				arg.Int32Value = append(arg.Int32Value, int32(Int32Value))
			case "Sint32Value[]":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				arg.Sint32Value = append(arg.Sint32Value, int32(Sint32Value))
			case "Uint32Value[]":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
			case "Int64Value[]":
				Int64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				}
				arg.Int64Value = append(arg.Int64Value, Int64Value)
			case "Sint64Value[]":
				Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				}
				arg.Sint64Value = append(arg.Sint64Value, Sint64Value)
			case "Uint64Value[]":
				Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				}
				arg.Uint64Value = append(arg.Uint64Value, Uint64Value)
			case "Sfixed32Value[]":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32Value))
			case "Fixed32Value[]":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
			case "FloatValue[]":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				arg.FloatValue = append(arg.FloatValue, float32(FloatValue))
			case "Sfixed64Value[]":
				Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				}
				arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64Value)
			case "Fixed64Value[]":
				Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				}
				arg.Fixed64Value = append(arg.Fixed64Value, Fixed64Value)
			case "DoubleValue[]":
				DoubleValue, convErr := strconv.ParseFloat(value, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				}
				arg.DoubleValue = append(arg.DoubleValue, DoubleValue)
			case "StringValue[]":
				arg.StringValue = append(arg.StringValue, value)
			case "BytesValue[]":
				arg.BytesValue = append(arg.BytesValue, []byte(value))
			case "StringValueQuery[]":
				arg.StringValueQuery = append(arg.StringValueQuery, value)
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	StringValueStr := ctx.Param("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
	}

	return arg, err
}

func buildExampleServiceNameCheckRepeatedPostRepeatedCheck(ctx *gin.Context) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "BoolValue[]":
				switch value {
				case "true", "t", "1":
					arg.BoolValue = append(arg.BoolValue, true)
				case "false", "f", "0":
					arg.BoolValue = append(arg.BoolValue, false)
				default:
					return nil, fmt.Errorf("unknown bool string value %s", value)
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
						return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					}
				}
			case "Int32Value[]":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				arg.Int32Value = append(arg.Int32Value, int32(Int32Value))
			case "Sint32Value[]":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				arg.Sint32Value = append(arg.Sint32Value, int32(Sint32Value))
			case "Uint32Value[]":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
			case "Int64Value[]":
				Int64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				}
				arg.Int64Value = append(arg.Int64Value, Int64Value)
			case "Sint64Value[]":
				Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				}
				arg.Sint64Value = append(arg.Sint64Value, Sint64Value)
			case "Uint64Value[]":
				Uint64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				}
				arg.Uint64Value = append(arg.Uint64Value, Uint64Value)
			case "Sfixed32Value[]":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32Value))
			case "Fixed32Value[]":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
			case "FloatValue[]":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				arg.FloatValue = append(arg.FloatValue, float32(FloatValue))
			case "Sfixed64Value[]":
				Sfixed64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				}
				arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64Value)
			case "Fixed64Value[]":
				Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				}
				arg.Fixed64Value = append(arg.Fixed64Value, Fixed64Value)
			case "DoubleValue[]":
				DoubleValue, convErr := strconv.ParseFloat(value, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				}
				arg.DoubleValue = append(arg.DoubleValue, DoubleValue)
			case "StringValue[]":
				arg.StringValue = append(arg.StringValue, value)
			case "BytesValue[]":
				arg.BytesValue = append(arg.BytesValue, []byte(value))
			case "StringValueQuery[]":
				arg.StringValueQuery = append(arg.StringValueQuery, value)
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	StringValueStr := ctx.Param("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
	}

	return arg, err
}

func buildExampleServiceNameEmptyGetEmpty(ctx *gin.Context) (arg *common.Empty, err error) {
	arg = &common.Empty{}
	return arg, err
}

func buildExampleServiceNameEmptyPostEmpty(ctx *gin.Context) (arg *common.Empty, err error) {
	arg = &common.Empty{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	return arg, err
}

func buildExampleServiceNameOnlyStructInGetOnlyStruct(ctx *gin.Context) (arg *common.OnlyStruct, err error) {
	arg = &common.OnlyStruct{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "value":
				return nil, fmt.Errorf("unsupported type message for query argument value")
			case "value.value":
				if arg.Value == nil {
					arg.Value = &common.ArrayItem{}
				}
				arg.Value.Value = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameMultipartFormMultipartFormRequest(ctx *gin.Context) (arg *common.MultipartFormRequest, err error) {
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
	if values := form.Value["otherField"]; len(values) > 0 {
		arg.OtherField = values[0]
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "document":
				return nil, fmt.Errorf("unsupported type message for query argument document")
			case "otherField":
				arg.OtherField = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameMultipartFormAllTypesMultipartFormAllTypes(ctx *gin.Context) (arg *common.MultipartFormAllTypes, err error) {
	arg = &common.MultipartFormAllTypes{}
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	if values := form.Value["BoolValue"]; len(values) > 0 {
		switch values[0] {
		case "true", "t", "1":
			arg.BoolValue = true
		case "false", "f", "0":
			arg.BoolValue = false
		default:
			return nil, fmt.Errorf("unknown bool string value %s", values[0])
		}
	}
	if values := form.Value["EnumValue"]; len(values) > 0 {
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
	if values := form.Value["Int32Value"]; len(values) > 0 {
		Int32Value, convErr := strconv.ParseInt(values[0], 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}
	if values := form.Value["Sint32Value"]; len(values) > 0 {
		Sint32Value, convErr := strconv.ParseInt(values[0], 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
		}
		arg.Sint32Value = int32(Sint32Value)
	}
	if values := form.Value["Uint32Value"]; len(values) > 0 {
		for _, value := range values {
			Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
		}
	}
	if values := form.Value["Int64Value"]; len(values) > 0 {
		arg.Int64Value, err = strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}
	if values := form.Value["Sint64Value"]; len(values) > 0 {
		Sint64Value, convErr := strconv.ParseInt(values[0], 10, 64)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
		}
		arg.Sint64Value = &Sint64Value
	}
	if values := form.Value["Uint64Value"]; len(values) > 0 {
		arg.Uint64Value, err = strconv.ParseUint(values[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}
	if values := form.Value["Sfixed32Value"]; len(values) > 0 {
		Sfixed32Value, convErr := strconv.ParseInt(values[0], 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
		}
		arg.Sfixed32Value = int32(Sfixed32Value)
	}
	if values := form.Value["Fixed32Value"]; len(values) > 0 {
		for _, value := range values {
			Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
			if convErr != nil {
				return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
		}
	}
	if values := form.Value["FloatValue"]; len(values) > 0 {
		FloatValue, convErr := strconv.ParseFloat(values[0], 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}
	if values := form.Value["Sfixed64Value"]; len(values) > 0 {
		arg.Sfixed64Value, err = strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
		}
	}
	if values := form.Value["Fixed64Value"]; len(values) > 0 {
		Fixed64Value, convErr := strconv.ParseUint(values[0], 10, 64)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
		}
		arg.Fixed64Value = &Fixed64Value
	}
	if values := form.Value["DoubleValue"]; len(values) > 0 {
		arg.DoubleValue, err = strconv.ParseFloat(values[0], 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}
	if values := form.Value["StringValue"]; len(values) > 0 {
		arg.StringValue = values[0]
	}
	if values := form.Value["BytesValue"]; len(values) > 0 {
		arg.BytesValue = []byte(values[0])
	}
	if values := form.Value["SliceStringValue"]; len(values) > 0 {
		arg.SliceStringValue = append(arg.SliceStringValue, values...)
	}
	if values := form.Value["SliceInt32Value"]; len(values) > 0 {
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
	if values := form.Value["RepeatedStringValue"]; len(values) > 0 {
		arg.RepeatedStringValue = append(arg.RepeatedStringValue, values...)
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "BoolValue":
				switch value {
				case "true", "t", "1":
					arg.BoolValue = true
				case "false", "f", "0":
					arg.BoolValue = false
				default:
					return nil, fmt.Errorf("unknown bool string value %s", value)
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
						return nil, fmt.Errorf("conversion failed for parameter EnumValue: %w", convErr)
					}
				}
			case "Int32Value":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				arg.Int32Value = int32(Int32Value)
			case "Sint32Value":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				arg.Sint32Value = int32(Sint32Value)
			case "Uint32Value[]":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32Value))
			case "Int64Value":
				arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				}
			case "Sint64Value":
				Sint64Value, convErr := strconv.ParseInt(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				}
				arg.Sint64Value = &Sint64Value
			case "Uint64Value":
				arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				}
			case "Sfixed32Value":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				arg.Sfixed32Value = int32(Sfixed32Value)
			case "Fixed32Value[]":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32Value))
			case "FloatValue":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				arg.FloatValue = float32(FloatValue)
			case "Sfixed64Value":
				arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				}
			case "Fixed64Value":
				Fixed64Value, convErr := strconv.ParseUint(value, 10, 64)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				}
				arg.Fixed64Value = &Fixed64Value
			case "DoubleValue":
				arg.DoubleValue, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
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
					return nil, fmt.Errorf("conversion failed for parameter SliceInt32Value: %w", convErr)
				}
				arg.SliceInt32Value = append(arg.SliceInt32Value, int32(SliceInt32Value))
			case "document":
				return nil, fmt.Errorf("unsupported type message for query argument document")
			case "RepeatedStringValue[]":
				arg.RepeatedStringValue = append(arg.RepeatedStringValue, value)
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameAllTypesMaxTestAllNumberTypesMsg(ctx *gin.Context) (arg *common.AllNumberTypesMsg, err error) {
	arg = &common.AllNumberTypesMsg{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "Int32Value":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				arg.Int32Value = int32(Int32Value)
			case "Sint32Value":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				arg.Sint32Value = int32(Sint32Value)
			case "Uint32Value":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				arg.Uint32Value = uint32(Uint32Value)
			case "Int64Value":
				arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				}
			case "Sint64Value":
				arg.Sint64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
				}
			case "Uint64Value":
				arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				}
			case "Sfixed32Value":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				arg.Sfixed32Value = int32(Sfixed32Value)
			case "Fixed32Value":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				arg.Fixed32Value = uint32(Fixed32Value)
			case "FloatValue":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				arg.FloatValue = float32(FloatValue)
			case "Sfixed64Value":
				arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				}
			case "Fixed64Value":
				arg.Fixed64Value, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
				}
			case "DoubleValue":
				arg.DoubleValue, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				}
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	Int32ValueStr := ctx.Param("Int32Value")
	if len(Int32ValueStr) != 0 {
		Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}

	Uint32ValueStr := ctx.Param("Uint32Value")
	if len(Uint32ValueStr) != 0 {
		Uint32Value, convErr := strconv.ParseUint(Uint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
		}
		arg.Uint32Value = uint32(Uint32Value)
	}

	Int64ValueStr := ctx.Param("Int64Value")
	if len(Int64ValueStr) != 0 {
		arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}

	Uint64ValueStr := ctx.Param("Uint64Value")
	if len(Uint64ValueStr) != 0 {
		arg.Uint64Value, err = strconv.ParseUint(Uint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}

	FloatValueStr := ctx.Param("FloatValue")
	if len(FloatValueStr) != 0 {
		FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}

	DoubleValueStr := ctx.Param("DoubleValue")
	if len(DoubleValueStr) != 0 {
		arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameAllTypesMaxQueryTestAllNumberTypesMsg(ctx *gin.Context) (arg *common.AllNumberTypesMsg, err error) {
	arg = &common.AllNumberTypesMsg{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "Int32Value":
				Int32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				}
				arg.Int32Value = int32(Int32Value)
			case "Sint32Value":
				Sint32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				}
				arg.Sint32Value = int32(Sint32Value)
			case "Uint32Value":
				Uint32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				}
				arg.Uint32Value = uint32(Uint32Value)
			case "Int64Value":
				arg.Int64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				}
			case "Sint64Value":
				arg.Sint64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
				}
			case "Uint64Value":
				arg.Uint64Value, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
				}
			case "Sfixed32Value":
				Sfixed32Value, convErr := strconv.ParseInt(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				}
				arg.Sfixed32Value = int32(Sfixed32Value)
			case "Fixed32Value":
				Fixed32Value, convErr := strconv.ParseUint(value, 10, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				}
				arg.Fixed32Value = uint32(Fixed32Value)
			case "FloatValue":
				FloatValue, convErr := strconv.ParseFloat(value, 32)
				if convErr != nil {
					return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				}
				arg.FloatValue = float32(FloatValue)
			case "Sfixed64Value":
				arg.Sfixed64Value, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				}
			case "Fixed64Value":
				arg.Fixed64Value, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
				}
			case "DoubleValue":
				arg.DoubleValue, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				}
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameGetMessageGetMessageRequest(ctx *gin.Context) (arg *common.GetMessageRequest, err error) {
	arg = &common.GetMessageRequest{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "name":
				arg.Name = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	NameStr := ctx.Param("name")
	if len(NameStr) != 0 {
		arg.Name = NameStr
		arg.Name = fmt.Sprintf("messages/%s", arg.Name)
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV2GetMessageRequestV2(ctx *gin.Context) (arg *common.GetMessageRequestV2, err error) {
	arg = &common.GetMessageRequestV2{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "message_id":
				arg.MessageId = value
			case "revision":
				arg.Revision, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("conversion failed for parameter revision: %w", err)
				}
			case "sub":
				return nil, fmt.Errorf("unsupported type message for query argument sub")
			case "sub.subfield":
				if arg.Sub == nil {
					arg.Sub = &common.GetMessageRequestV2_SubMessage{}
				}
				arg.Sub.Subfield = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	MessageIdStr := ctx.Param("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
	}

	return arg, err
}

func buildExampleServiceNameUpdateMessageUpdateMessageRequest(ctx *gin.Context) (arg *common.UpdateMessageRequest, err error) {
	arg = &common.UpdateMessageRequest{}
	arg.Message = &common.MessageV2{}
	if err = ctx.ShouldBindBodyWithJSON(arg.Message); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "message_id":
				arg.MessageId = value
			case "message":
				return nil, fmt.Errorf("unsupported type message for query argument message")
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
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	MessageIdStr := ctx.Param("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
	}

	return arg, err
}

func buildExampleServiceNameUpdateMessageV2MessageV2(ctx *gin.Context) (arg *common.MessageV2, err error) {
	arg = &common.MessageV2{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "message_id":
				arg.MessageId = value
			case "text":
				arg.Text = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	MessageIdStr := ctx.Param("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV3GetMessageRequestV3(ctx *gin.Context) (arg *common.GetMessageRequestV3, err error) {
	arg = &common.GetMessageRequestV3{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "message_id":
				arg.MessageId = value
			case "user_id":
				arg.UserId = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	MessageIdStr := ctx.Param("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
	}

	UserIdStr := ctx.Param("user_id")
	if len(UserIdStr) != 0 {
		arg.UserId = UserIdStr
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV4GetMessageRequestV3(ctx *gin.Context) (arg *common.GetMessageRequestV3, err error) {
	arg = &common.GetMessageRequestV3{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "message_id":
				arg.MessageId = value
			case "user_id":
				arg.UserId = value
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	MessageIdStr := ctx.Param("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		arg.MessageId = fmt.Sprintf("base%s", arg.MessageId)
	}

	return arg, err
}

func buildExampleServiceNameTopLevelArrayArray(ctx *gin.Context) (arg *common.Array, err error) {
	arg = &common.Array{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "items[]":
				return nil, fmt.Errorf("unsupported type message for query argument items")
			default:
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
	return arg, err
}

func buildExampleServiceNameUpdateMessageV3UpdateMessageRequest(ctx *gin.Context) (arg *common.UpdateMessageRequest, err error) {
	arg = &common.UpdateMessageRequest{}
	if err = ctx.ShouldBindBodyWithJSON(arg); err != nil {
		return nil, err
	}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			switch key {
			case "message_id":
				arg.MessageId = value
			case "message":
				return nil, fmt.Errorf("unsupported type message for query argument message")
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
				return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
			}
		}
	}
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
