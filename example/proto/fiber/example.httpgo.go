// source: example.proto

package proto

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	common "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	v3 "github.com/gofiber/fiber/v3"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	multipart "mime/multipart"
	url "net/url"
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
	r *v3.App,
	h ServiceNameHTTPGoService,
	middlewares []func(ctx context.Context, req any, handler func(ctx context.Context, req any) (resp any, err error)) (resp any, err error),
) error {
	var middleware = chainServerMiddlewaresExample(middlewares)

	r.Post("/v1/RPCName/:stringArgument/:int64Argument", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameRPCNameInputMsgName(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/test/:BoolValue/:EnumValue/:Int32Value/:Sint32Value/:Uint32Value/:Int64Value/:Sint64Value/:Uint64Value/:Sfixed32Value/:Fixed32Value/:FloatValue/:Sfixed64Value/:Fixed64Value/:DoubleValue/:StringValue/:BytesValue", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTypesTestAllTypesMsg(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/text/:String/:RepeatedString/:Bytes/:RepeatedBytes/:Enum/:RepeatedEnum", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTextTypesPostAllTextTypesMsg(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v2/text/:String/:RepeatedString/:Bytes/:RepeatedBytes/:Enum/:RepeatedEnum", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTextTypesGetAllTextTypesMsg(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/test/commonTypes", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameCommonTypesAny(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	// same types but different query, we need different query builder function
	r.Post("/v1/sameInputAndOutput/:stringArgument", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameSameInputAndOutputInputMsgName(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/test/optional", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameOptionalOptionalField(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v1/test/get", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMethodInputMsgName(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v1/repeated/:BoolValue/:EnumValue/:Int32Value/:Sint32Value/:Uint32Value/:Int64Value/:Sint64Value/:Uint64Value/:Sfixed32Value/:Fixed32Value/:FloatValue/:Sfixed64Value/:Fixed64Value/:DoubleValue/:StringValue/:BytesValue/:StringValueQuery", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameCheckRepeatedPathRepeatedCheck(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v2/repeated/:StringValue", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v3/repeated/:StringValue", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameCheckRepeatedPostRepeatedCheck(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v1/emptyGet", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameEmptyGetEmpty(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/emptyPost", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameEmptyPostEmpty(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/onlyStruct", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameOnlyStructInGetOnlyStruct(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/multipart", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameMultipartFormMultipartFormRequest(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/multipartall", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameMultipartFormAllTypesMultipartFormAllTypes(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v1/max/:Int32Value/:Uint32Value/:Int64Value/:Uint64Value/:FloatValue/:DoubleValue", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTypesMaxTestAllNumberTypesMsg(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v1/maxquery", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameAllTypesMaxQueryTestAllNumberTypesMsg(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	// http rule checks
	// v1/{name=messages/*}
	r.Get("/v1/messages/:name", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageGetMessageRequest(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v2/messages/:message_id", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV2GetMessageRequestV2(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Patch("/v1/messages/:message_id", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameUpdateMessageUpdateMessageRequest(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Patch("/v2/messages/:message_id", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameUpdateMessageV2MessageV2(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v3/messages/:message_id", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV3GetMessageRequestV3(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v3/users/:user_id/messages/:message_id", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV3GetMessageRequestV3(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Get("/v4/messages/base/:message_id+", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameGetMessageV4GetMessageRequestV3(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
		_, _ = fiberctx.Write(respJson)
	})

	r.Post("/v1/array", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameTopLevelArrayArray(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
			_, _ = fiberctx.Write(respJson)
		} else {
			respJson, _ := json.Marshal(resp)
			_, _ = fiberctx.Write(respJson)
		}
	})

	r.Patch("/v3/messages", func(fiberctx v3.Ctx) {
		fiberctx.Set("Content-Type", "application/json")
		input, err := buildExampleServiceNameUpdateMessageV3UpdateMessageRequest(fiberctx)
		if err != nil {
			fiberctx.Status(400)
			respJson, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
			_, _ = fiberctx.Write(respJson)
			return
		}
		ctx := context.WithValue(fiberctx.Context(), "request", fiberctx)
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
			_, _ = fiberctx.Write(respJson)
		} else {
			respJson, _ := json.Marshal(resp)
			_, _ = fiberctx.Write(respJson)
		}
	})

	return nil
}

func buildExampleServiceNameRPCNameInputMsgName(ctx v3.Ctx) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	StringArgumentStr := ctx.Params("stringArgument")
	if len(StringArgumentStr) != 0 {
		arg.StringArgument = StringArgumentStr
		if arg.StringArgument, err = url.PathUnescape(arg.StringArgument); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field stringArgument: %w", err)
		}
	}

	Int64ArgumentStr := ctx.Params("int64Argument")
	if len(Int64ArgumentStr) != 0 {
		arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter int64Argument: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameAllTypesTestAllTypesMsg(ctx v3.Ctx) (arg *common.AllTypesMsg, err error) {
	arg = &common.AllTypesMsg{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	BoolValueStr := ctx.Params("BoolValue")
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

	EnumValueStr := ctx.Params("EnumValue")
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

	Int32ValueStr := ctx.Params("Int32Value")
	if len(Int32ValueStr) != 0 {
		Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}

	Sint32ValueStr := ctx.Params("Sint32Value")
	if len(Sint32ValueStr) != 0 {
		Sint32Value, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
		}
		arg.Sint32Value = int32(Sint32Value)
	}

	Uint32ValueStr := ctx.Params("Uint32Value")
	if len(Uint32ValueStr) != 0 {
		Uint32Value, convErr := strconv.ParseUint(Uint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
		}
		arg.Uint32Value = uint32(Uint32Value)
	}

	Int64ValueStr := ctx.Params("Int64Value")
	if len(Int64ValueStr) != 0 {
		arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}

	Sint64ValueStr := ctx.Params("Sint64Value")
	if len(Sint64ValueStr) != 0 {
		arg.Sint64Value, err = strconv.ParseInt(Sint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
		}
	}

	Uint64ValueStr := ctx.Params("Uint64Value")
	if len(Uint64ValueStr) != 0 {
		arg.Uint64Value, err = strconv.ParseUint(Uint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}

	Sfixed32ValueStr := ctx.Params("Sfixed32Value")
	if len(Sfixed32ValueStr) != 0 {
		Sfixed32Value, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
		}
		arg.Sfixed32Value = int32(Sfixed32Value)
	}

	Fixed32ValueStr := ctx.Params("Fixed32Value")
	if len(Fixed32ValueStr) != 0 {
		Fixed32Value, convErr := strconv.ParseUint(Fixed32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
		}
		arg.Fixed32Value = uint32(Fixed32Value)
	}

	FloatValueStr := ctx.Params("FloatValue")
	if len(FloatValueStr) != 0 {
		FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}

	Sfixed64ValueStr := ctx.Params("Sfixed64Value")
	if len(Sfixed64ValueStr) != 0 {
		arg.Sfixed64Value, err = strconv.ParseInt(Sfixed64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
		}
	}

	Fixed64ValueStr := ctx.Params("Fixed64Value")
	if len(Fixed64ValueStr) != 0 {
		arg.Fixed64Value, err = strconv.ParseUint(Fixed64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", err)
		}
	}

	DoubleValueStr := ctx.Params("DoubleValue")
	if len(DoubleValueStr) != 0 {
		arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}

	StringValueStr := ctx.Params("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = StringValueStr
		if arg.StringValue, err = url.PathUnescape(arg.StringValue); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
		}
	}

	BytesValueStr := ctx.Params("BytesValue")
	if len(BytesValueStr) != 0 {
		arg.BytesValue = []byte(BytesValueStr)
		if BytesValueStr, err = url.PathUnescape(string(arg.BytesValue)); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field BytesValue: %w", err)
		}
		arg.BytesValue = []byte(BytesValueStr)
	}

	return arg, err
}

func buildExampleServiceNameAllTextTypesPostAllTextTypesMsg(ctx v3.Ctx) (arg *common.AllTextTypesMsg, err error) {
	arg = &common.AllTextTypesMsg{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	String_Str := ctx.Params("String")
	if len(String_Str) != 0 {
		arg.String_ = String_Str
		if arg.String_, err = url.PathUnescape(arg.String_); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field String: %w", err)
		}
	}

	RepeatedStringStr := ctx.Params("RepeatedString")
	if len(RepeatedStringStr) != 0 {
		arg.RepeatedString = strings.Split(RepeatedStringStr, ",")
		for i, value := range arg.RepeatedString {
			if arg.RepeatedString[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field RepeatedString: %w", err)
			}
		}
	}

	BytesStr := ctx.Params("Bytes")
	if len(BytesStr) != 0 {
		arg.Bytes = []byte(BytesStr)
		if BytesStr, err = url.PathUnescape(string(arg.Bytes)); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field Bytes: %w", err)
		}
		arg.Bytes = []byte(BytesStr)
	}

	RepeatedBytesStr := ctx.Params("RepeatedBytes")
	if len(RepeatedBytesStr) != 0 {
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

	EnumStr := ctx.Params("Enum")
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

	RepeatedEnumStr := ctx.Params("RepeatedEnum")
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

func buildExampleServiceNameAllTextTypesGetAllTextTypesMsg(ctx v3.Ctx) (arg *common.AllTextTypesMsg, err error) {
	arg = &common.AllTextTypesMsg{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	String_Str := ctx.Params("String")
	if len(String_Str) != 0 {
		arg.String_ = String_Str
		if arg.String_, err = url.PathUnescape(arg.String_); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field String: %w", err)
		}
	}

	RepeatedStringStr := ctx.Params("RepeatedString")
	if len(RepeatedStringStr) != 0 {
		arg.RepeatedString = strings.Split(RepeatedStringStr, ",")
		for i, value := range arg.RepeatedString {
			if arg.RepeatedString[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field RepeatedString: %w", err)
			}
		}
	}

	BytesStr := ctx.Params("Bytes")
	if len(BytesStr) != 0 {
		arg.Bytes = []byte(BytesStr)
		if BytesStr, err = url.PathUnescape(string(arg.Bytes)); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field Bytes: %w", err)
		}
		arg.Bytes = []byte(BytesStr)
	}

	RepeatedBytesStr := ctx.Params("RepeatedBytes")
	if len(RepeatedBytesStr) != 0 {
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

	EnumStr := ctx.Params("Enum")
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

	RepeatedEnumStr := ctx.Params("RepeatedEnum")
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

func buildExampleServiceNameCommonTypesAny(ctx v3.Ctx) (arg *anypb.Any, err error) {
	arg = &anypb.Any{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
		switch key {
		case "type_url":
			arg.TypeUrl = value
		case "value":
			arg.Value = []byte(value)
		default:
			return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
		}
	}
	return arg, err
}

func buildExampleServiceNameSameInputAndOutputInputMsgName(ctx v3.Ctx) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	StringArgumentStr := ctx.Params("stringArgument")
	if len(StringArgumentStr) != 0 {
		arg.StringArgument = StringArgumentStr
		if arg.StringArgument, err = url.PathUnescape(arg.StringArgument); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field stringArgument: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameOptionalOptionalField(ctx v3.Ctx) (arg *common.OptionalField, err error) {
	arg = &common.OptionalField{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	return arg, err
}

func buildExampleServiceNameGetMethodInputMsgName(ctx v3.Ctx) (arg *common.InputMsgName, err error) {
	arg = &common.InputMsgName{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	return arg, err
}

func buildExampleServiceNameCheckRepeatedPathRepeatedCheck(ctx v3.Ctx) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	BoolValueStr := ctx.Params("BoolValue")
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

	EnumValueStr := ctx.Params("EnumValue")
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

	Int32ValueStr := ctx.Params("Int32Value")
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

	Sint32ValueStr := ctx.Params("Sint32Value")
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

	Uint32ValueStr := ctx.Params("Uint32Value")
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

	Int64ValueStr := ctx.Params("Int64Value")
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

	Sint64ValueStr := ctx.Params("Sint64Value")
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

	Uint64ValueStr := ctx.Params("Uint64Value")
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

	Sfixed32ValueStr := ctx.Params("Sfixed32Value")
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

	Fixed32ValueStr := ctx.Params("Fixed32Value")
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

	FloatValueStr := ctx.Params("FloatValue")
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

	Sfixed64ValueStr := ctx.Params("Sfixed64Value")
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

	Fixed64ValueStr := ctx.Params("Fixed64Value")
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

	DoubleValueStr := ctx.Params("DoubleValue")
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

	StringValueStr := ctx.Params("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
		for i, value := range arg.StringValue {
			if arg.StringValue[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
			}
		}
	}

	BytesValueStr := ctx.Params("BytesValue")
	if len(BytesValueStr) != 0 {
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

	StringValueQueryStr := ctx.Params("StringValueQuery")
	if len(StringValueQueryStr) != 0 {
		arg.StringValueQuery = strings.Split(StringValueQueryStr, ",")
		for i, value := range arg.StringValueQuery {
			if arg.StringValueQuery[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValueQuery: %w", err)
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(ctx v3.Ctx) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	StringValueStr := ctx.Params("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
		for i, value := range arg.StringValue {
			if arg.StringValue[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameCheckRepeatedPostRepeatedCheck(ctx v3.Ctx) (arg *common.RepeatedCheck, err error) {
	arg = &common.RepeatedCheck{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	StringValueStr := ctx.Params("StringValue")
	if len(StringValueStr) != 0 {
		arg.StringValue = strings.Split(StringValueStr, ",")
		for i, value := range arg.StringValue {
			if arg.StringValue[i], err = url.PathUnescape(value); err != nil {
				return nil, fmt.Errorf("PathUnescape failed for field StringValue: %w", err)
			}
		}
	}

	return arg, err
}

func buildExampleServiceNameEmptyGetEmpty(ctx v3.Ctx) (arg *common.Empty, err error) {
	arg = &common.Empty{}
	return arg, err
}

func buildExampleServiceNameEmptyPostEmpty(ctx v3.Ctx) (arg *common.Empty, err error) {
	arg = &common.Empty{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	return arg, err
}

func buildExampleServiceNameOnlyStructInGetOnlyStruct(ctx v3.Ctx) (arg *common.OnlyStruct, err error) {
	arg = &common.OnlyStruct{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	return arg, err
}

func buildExampleServiceNameMultipartFormMultipartFormRequest(ctx v3.Ctx) (arg *common.MultipartFormRequest, err error) {
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
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
		switch key {
		case "document":
			return nil, fmt.Errorf("unsupported type message for query argument document")
		case "otherField":
			arg.OtherField = value
		default:
			return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
		}
	}
	return arg, err
}

func buildExampleServiceNameMultipartFormAllTypesMultipartFormAllTypes(ctx v3.Ctx) (arg *common.MultipartFormAllTypes, err error) {
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
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	return arg, err
}

func buildExampleServiceNameAllTypesMaxTestAllNumberTypesMsg(ctx v3.Ctx) (arg *common.AllNumberTypesMsg, err error) {
	arg = &common.AllNumberTypesMsg{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	Int32ValueStr := ctx.Params("Int32Value")
	if len(Int32ValueStr) != 0 {
		Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
		}
		arg.Int32Value = int32(Int32Value)
	}

	Uint32ValueStr := ctx.Params("Uint32Value")
	if len(Uint32ValueStr) != 0 {
		Uint32Value, convErr := strconv.ParseUint(Uint32ValueStr, 10, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
		}
		arg.Uint32Value = uint32(Uint32Value)
	}

	Int64ValueStr := ctx.Params("Int64Value")
	if len(Int64ValueStr) != 0 {
		arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
		}
	}

	Uint64ValueStr := ctx.Params("Uint64Value")
	if len(Uint64ValueStr) != 0 {
		arg.Uint64Value, err = strconv.ParseUint(Uint64ValueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", err)
		}
	}

	FloatValueStr := ctx.Params("FloatValue")
	if len(FloatValueStr) != 0 {
		FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
		if convErr != nil {
			return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
		}
		arg.FloatValue = float32(FloatValue)
	}

	DoubleValueStr := ctx.Params("DoubleValue")
	if len(DoubleValueStr) != 0 {
		arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameAllTypesMaxQueryTestAllNumberTypesMsg(ctx v3.Ctx) (arg *common.AllNumberTypesMsg, err error) {
	arg = &common.AllNumberTypesMsg{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	return arg, err
}

func buildExampleServiceNameGetMessageGetMessageRequest(ctx v3.Ctx) (arg *common.GetMessageRequest, err error) {
	arg = &common.GetMessageRequest{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
		switch key {
		case "name":
			arg.Name = value
		default:
			return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
		}
	}
	NameStr := ctx.Params("name")
	if len(NameStr) != 0 {
		arg.Name = NameStr
		if arg.Name, err = url.PathUnescape(arg.Name); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field name: %w", err)
		}
		arg.Name = fmt.Sprintf("messages/%s", arg.Name)
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV2GetMessageRequestV2(ctx v3.Ctx) (arg *common.GetMessageRequestV2, err error) {
	arg = &common.GetMessageRequestV2{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	MessageIdStr := ctx.Params("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameUpdateMessageUpdateMessageRequest(ctx v3.Ctx) (arg *common.UpdateMessageRequest, err error) {
	arg = &common.UpdateMessageRequest{}
	var body = ctx.Body()
	if len(body) > 0 {
		arg.Message = &common.MessageV2{}
		if err = json.Unmarshal(body, arg.Message); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
	MessageIdStr := ctx.Params("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameUpdateMessageV2MessageV2(ctx v3.Ctx) (arg *common.MessageV2, err error) {
	arg = &common.MessageV2{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
		switch key {
		case "message_id":
			arg.MessageId = value
		case "text":
			arg.Text = value
		default:
			return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
		}
	}
	MessageIdStr := ctx.Params("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV3GetMessageRequestV3(ctx v3.Ctx) (arg *common.GetMessageRequestV3, err error) {
	arg = &common.GetMessageRequestV3{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
		switch key {
		case "message_id":
			arg.MessageId = value
		case "user_id":
			arg.UserId = value
		default:
			return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
		}
	}
	MessageIdStr := ctx.Params("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
	}

	UserIdStr := ctx.Params("user_id")
	if len(UserIdStr) != 0 {
		arg.UserId = UserIdStr
		if arg.UserId, err = url.PathUnescape(arg.UserId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field user_id: %w", err)
		}
	}

	return arg, err
}

func buildExampleServiceNameGetMessageV4GetMessageRequestV3(ctx v3.Ctx) (arg *common.GetMessageRequestV3, err error) {
	arg = &common.GetMessageRequestV3{}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
		switch key {
		case "message_id":
			arg.MessageId = value
		case "user_id":
			arg.UserId = value
		default:
			return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
		}
	}
	MessageIdStr := ctx.Params("message_id")
	if len(MessageIdStr) != 0 {
		arg.MessageId = MessageIdStr
		if arg.MessageId, err = url.PathUnescape(arg.MessageId); err != nil {
			return nil, fmt.Errorf("PathUnescape failed for field message_id: %w", err)
		}
		arg.MessageId = fmt.Sprintf("base%s", arg.MessageId)
	}

	return arg, err
}

func buildExampleServiceNameTopLevelArrayArray(ctx v3.Ctx) (arg *common.Array, err error) {
	arg = &common.Array{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
		switch key {
		case "items[]":
			return nil, fmt.Errorf("unsupported type message for query argument items")
		default:
			return nil, fmt.Errorf("unknown query parameter %s with value %s", key, value)
		}
	}
	return arg, err
}

func buildExampleServiceNameUpdateMessageV3UpdateMessageRequest(ctx v3.Ctx) (arg *common.UpdateMessageRequest, err error) {
	arg = &common.UpdateMessageRequest{}
	var body = ctx.Body()
	if len(body) > 0 {
		if err = json.Unmarshal(body, arg); err != nil {
			return nil, err
		}
	}
	for keyB, valueB := range ctx.RequestCtx().URI().QueryArgs().All() {
		var key = string(keyB)
		var value string
		value, err = url.QueryUnescape(string(valueB))
		if err != nil {
			return nil, fmt.Errorf("failed to decode query parameter %s: %w", key, err)
		}
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
