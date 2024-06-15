// source: proto/example.proto

package proto

import (
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	somepackage "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/somepackage"
	router "github.com/fasthttp/router"
	easyjson "github.com/mailru/easyjson"
	fasthttp "github.com/valyala/fasthttp"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strconv "strconv"
	strings "strings"
)

type ServiceNameHTTPGoService interface {
	RPCName(context.Context, *InputMsgName) (*OutputMsgName, error)
	AllTypesTest(context.Context, *AllTypesMsg) (*AllTypesMsg, error)
	CommonTypes(context.Context, *anypb.Any) (*emptypb.Empty, error)
	Imports(context.Context, *somepackage.SomeCustomMsg1) (*somepackage.SomeCustomMsg2, error)
	SameInputAndOutput(context.Context, *InputMsgName) (*OutputMsgName, error)
	Optional(context.Context, *InputMsgName) (*OptionalField, error)
	GetMethod(context.Context, *InputMsgName) (*OutputMsgName, error)
	CheckRepeated(context.Context, *RepeatedCheck) (*RepeatedCheck, error)
}

func RegisterServiceNameHTTPGoServer(
	_ context.Context,
	r *router.Router,
	h ServiceNameHTTPGoService,
	middlewares []func(ctx *fasthttp.RequestCtx, handler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error)) (resp interface{}, err error),
) error {
	var middleware = chainServerMiddlewaresExample(middlewares)

	r.POST("/v1/test/{stringArgument}/{int64Argument}", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameRPCNameInputMsgName(ctx)
			if err != nil {
				return nil, err
			}
			return h.RPCName(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	r.POST("/v1/test/{BoolValue}/{EnumValue}/{Int32Value}/{Sint32Value}/{Uint32Value}/{Int64Value}/{Sint64Value}/{Uint64Value}/{Sfixed32Value}/{Fixed32Value}/{FloatValue}/{Sfixed64Value}/{Fixed64Value}/{DoubleValue}/{StringValue}/{BytesValue}", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameAllTypesTestAllTypesMsg(ctx)
			if err != nil {
				return nil, err
			}
			return h.AllTypesTest(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	r.POST("/v1/test/commonTypes", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameCommonTypesAny(ctx)
			if err != nil {
				return nil, err
			}
			return h.CommonTypes(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	r.POST("/v1/test/imports", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameImportsSomeCustomMsg1(ctx)
			if err != nil {
				return nil, err
			}
			return h.Imports(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	// same types but different query, we need different query builder function
	r.POST("/v1/test/{stringArgument}", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameSameInputAndOutputInputMsgName(ctx)
			if err != nil {
				return nil, err
			}
			return h.SameInputAndOutput(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	r.POST("/v1/test/optional", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameOptionalInputMsgName(ctx)
			if err != nil {
				return nil, err
			}
			return h.Optional(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	r.GET("/v1/test/get", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameGetMethodInputMsgName(ctx)
			if err != nil {
				return nil, err
			}
			return h.GetMethod(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	r.GET("/v1/repeated/{stringValueArg}", func(ctx *fasthttp.RequestCtx) {
		handler := func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
			input, err := buildExampleServiceNameCheckRepeatedRepeatedCheck(ctx)
			if err != nil {
				return nil, err
			}
			return h.CheckRepeated(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx)
			return
		}
		_, _ = middleware(ctx, handler)
	})

	return nil
}

func buildExampleServiceNameRPCNameInputMsgName(ctx *fasthttp.RequestCtx) (arg *InputMsgName, err error) {
	arg = &InputMsgName{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "int64Argument":
			Int64ArgumentStr := string(value)
			arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Argument: %w", err)
				return
			}
		case "stringArgument":
			arg.StringArgument = string(value)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	StringArgumentStr, ok := ctx.UserValue("stringArgument").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringArgument")
	}
	arg.StringArgument = StringArgumentStr

	Int64ArgumentStr, ok := ctx.UserValue("int64Argument").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Int64Argument")
	}
	arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("conversion failed for parameter Int64Argument: %w", err)
	}

	return arg, err
}

func buildExampleServiceNameAllTypesTestAllTypesMsg(ctx *fasthttp.RequestCtx) (arg *AllTypesMsg, err error) {
	arg = &AllTypesMsg{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "BoolValue":
			BoolValueStr := string(value)
			switch BoolValueStr {
			case "true", "t", "1":
				arg.BoolValue = true
			case "false", "f", "0":
				arg.BoolValue = false
			default:
				err = fmt.Errorf("unknown bool string value %s", BoolValueStr)
				return
			}
		case "EnumValue":
			EnumValueStr := string(value)
			if OptionsValue, ok := Options_value[strings.ToUpper(EnumValueStr)]; ok {
				arg.EnumValue = Options(OptionsValue)
			} else {
				if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
					if _, ok = Options_name[int32(intOptionValue)]; ok {
						arg.EnumValue = Options(intOptionValue)
					}
				}
			}
		case "Int32Value":
			Int32ValueStr := string(value)
			Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = int32(Int32Value)
		case "Sint32Value":
			Sint32ValueStr := string(value)
			Sint32Value, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = int32(Sint32Value)
		case "Uint32Value":
			Uint32ValueStr := string(value)
			Uint32Value, convErr := strconv.ParseInt(Uint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = uint32(Uint32Value)
		case "Int64Value":
			Int64ValueStr := string(value)
			arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
				return
			}
		case "Sint64Value":
			Sint64ValueStr := string(value)
			arg.Sint64Value, err = strconv.ParseInt(Sint64ValueStr, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
				return
			}
		case "Uint64Value":
			Uint64ValueStr := string(value)
			Uint64Value, convErr := strconv.ParseInt(Uint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = uint64(Uint64Value)
		case "Sfixed32Value":
			Sfixed32ValueStr := string(value)
			Sfixed32Value, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = int32(Sfixed32Value)
		case "Fixed32Value":
			Fixed32ValueStr := string(value)
			Fixed32Value, convErr := strconv.ParseInt(Fixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = uint32(Fixed32Value)
		case "FloatValue":
			FloatValueStr := string(value)
			FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = float32(FloatValue)
		case "Sfixed64Value":
			Sfixed64ValueStr := string(value)
			arg.Sfixed64Value, err = strconv.ParseInt(Sfixed64ValueStr, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
				return
			}
		case "Fixed64Value":
			Fixed64ValueStr := string(value)
			Fixed64Value, convErr := strconv.ParseInt(Fixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = uint64(Fixed64Value)
		case "DoubleValue":
			DoubleValueStr := string(value)
			arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
				return
			}
		case "StringValue":
			arg.StringValue = string(value)
		case "BytesValue":
			arg.BytesValue = value
		case "MessageValue":
			err = fmt.Errorf("unsupported type message for query argument MessageValue")
			return
		case "SliceStringValue":
			SliceStringValue := string(value)
			arg.SliceStringValue = strings.Split(SliceStringValue, ",")
		case "SliceInt32Value":
			err = fmt.Errorf("unsupported type repeated int32 for query argument SliceInt32Value")
			return
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	BoolValueStr, ok := ctx.UserValue("BoolValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter BoolValue")
	}
	switch BoolValueStr {
	case "true", "t", "1":
		arg.BoolValue = true
	case "false", "f", "0":
		arg.BoolValue = false
	default:
		return nil, fmt.Errorf("unknown bool string value %s", BoolValueStr)
	}

	EnumValueStr, ok := ctx.UserValue("EnumValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter EnumValue")
	}
	OptionsValue, ok := Options_value[strings.ToUpper(EnumValueStr)]
	if ok {
		arg.EnumValue = Options(OptionsValue)
	} else {
		if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
			if _, ok = Options_name[int32(intOptionValue)]; ok {
				arg.EnumValue = Options(intOptionValue)
			}
		}
	}

	Int32ValueStr, ok := ctx.UserValue("Int32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Int32Value")
	}
	Int32Value, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
	}
	arg.Int32Value = int32(Int32Value)

	Sint32ValueStr, ok := ctx.UserValue("Sint32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sint32Value")
	}
	Sint32Value, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
	}
	arg.Sint32Value = int32(Sint32Value)

	Uint32ValueStr, ok := ctx.UserValue("Uint32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Uint32Value")
	}
	Uint32Value, convErr := strconv.ParseInt(Uint32ValueStr, 10, 32)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
	}
	arg.Uint32Value = uint32(Uint32Value)

	Int64ValueStr, ok := ctx.UserValue("Int64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Int64Value")
	}
	arg.Int64Value, err = strconv.ParseInt(Int64ValueStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("conversion failed for parameter Int64Value: %w", err)
	}

	Sint64ValueStr, ok := ctx.UserValue("Sint64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sint64Value")
	}
	arg.Sint64Value, err = strconv.ParseInt(Sint64ValueStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("conversion failed for parameter Sint64Value: %w", err)
	}

	Uint64ValueStr, ok := ctx.UserValue("Uint64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Uint64Value")
	}
	Uint64Value, convErr := strconv.ParseInt(Uint64ValueStr, 10, 64)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
	}
	arg.Uint64Value = uint64(Uint64Value)

	Sfixed32ValueStr, ok := ctx.UserValue("Sfixed32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sfixed32Value")
	}
	Sfixed32Value, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
	}
	arg.Sfixed32Value = int32(Sfixed32Value)

	Fixed32ValueStr, ok := ctx.UserValue("Fixed32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Fixed32Value")
	}
	Fixed32Value, convErr := strconv.ParseInt(Fixed32ValueStr, 10, 32)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
	}
	arg.Fixed32Value = uint32(Fixed32Value)

	FloatValueStr, ok := ctx.UserValue("FloatValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter FloatValue")
	}
	FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
	}
	arg.FloatValue = float32(FloatValue)

	Sfixed64ValueStr, ok := ctx.UserValue("Sfixed64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sfixed64Value")
	}
	arg.Sfixed64Value, err = strconv.ParseInt(Sfixed64ValueStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", err)
	}

	Fixed64ValueStr, ok := ctx.UserValue("Fixed64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Fixed64Value")
	}
	Fixed64Value, convErr := strconv.ParseInt(Fixed64ValueStr, 10, 64)
	if convErr != nil {
		return nil, fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
	}
	arg.Fixed64Value = uint64(Fixed64Value)

	DoubleValueStr, ok := ctx.UserValue("DoubleValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter DoubleValue")
	}
	arg.DoubleValue, err = strconv.ParseFloat(DoubleValueStr, 64)
	if err != nil {
		return nil, fmt.Errorf("conversion failed for parameter DoubleValue: %w", err)
	}

	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringValue")
	}
	arg.StringValue = StringValueStr

	BytesValueStr, ok := ctx.UserValue("BytesValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter BytesValue")
	}
	arg.BytesValue = []byte(BytesValueStr)

	return arg, err
}

func buildExampleServiceNameCommonTypesAny(ctx *fasthttp.RequestCtx) (arg *anypb.Any, err error) {
	arg = &anypb.Any{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "typeUrl":
			arg.TypeUrl = string(value)
		case "value":
			arg.Value = value
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameImportsSomeCustomMsg1(ctx *fasthttp.RequestCtx) (arg *somepackage.SomeCustomMsg1, err error) {
	arg = &somepackage.SomeCustomMsg1{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "val":
			arg.Val = string(value)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameSameInputAndOutputInputMsgName(ctx *fasthttp.RequestCtx) (arg *InputMsgName, err error) {
	arg = &InputMsgName{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "int64Argument":
			Int64ArgumentStr := string(value)
			arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Argument: %w", err)
				return
			}
		case "stringArgument":
			arg.StringArgument = string(value)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	StringArgumentStr, ok := ctx.UserValue("stringArgument").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringArgument")
	}
	arg.StringArgument = StringArgumentStr

	return arg, err
}

func buildExampleServiceNameOptionalInputMsgName(ctx *fasthttp.RequestCtx) (arg *InputMsgName, err error) {
	arg = &InputMsgName{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "int64Argument":
			Int64ArgumentStr := string(value)
			arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Argument: %w", err)
				return
			}
		case "stringArgument":
			arg.StringArgument = string(value)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameGetMethodInputMsgName(ctx *fasthttp.RequestCtx) (arg *InputMsgName, err error) {
	arg = &InputMsgName{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "int64Argument":
			Int64ArgumentStr := string(value)
			arg.Int64Argument, err = strconv.ParseInt(Int64ArgumentStr, 10, 64)
			if err != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Argument: %w", err)
				return
			}
		case "stringArgument":
			arg.StringArgument = string(value)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameCheckRepeatedRepeatedCheck(ctx *fasthttp.RequestCtx) (arg *RepeatedCheck, err error) {
	arg = &RepeatedCheck{}
	if body := ctx.PostBody(); len(body) > 0 {
		if argEJ, ok := interface{}(arg).(easyjson.Unmarshaler); ok {
			if err = easyjson.Unmarshal(body, argEJ); err != nil {
				return nil, err
			}
		} else {
			if err = json.Unmarshal(body, arg); err != nil {
				return nil, err
			}
		}
	}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "stringValueArg":
			StringValueArg := string(value)
			arg.StringValueArg = strings.Split(StringValueArg, ",")
		case "stringValueQuery":
			StringValueQuery := string(value)
			arg.StringValueQuery = strings.Split(StringValueQuery, ",")
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	StringValueArgStr, ok := ctx.UserValue("stringValueArg").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringValueArg")
	}
	arg.StringValueArg = strings.Split(StringValueArgStr, ",")
	return arg, err
}

func chainServerMiddlewaresExample(
	middlewares []func(ctx *fasthttp.RequestCtx, handler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error)) (resp interface{}, err error),
) func(ctx *fasthttp.RequestCtx, handler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error)) (resp interface{}, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx *fasthttp.RequestCtx, handler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error)) (resp interface{}, err error) {
			return middlewares[0](ctx, getChainServerMiddlewareHandlerExample(middlewares, 0, handler))
		}
	}
}

func getChainServerMiddlewareHandlerExample(
	middlewares []func(ctx *fasthttp.RequestCtx, handler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error)) (resp interface{}, err error),
	curr int,
	finalHandler func(ctx *fasthttp.RequestCtx) (resp interface{}, err error),
) func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx *fasthttp.RequestCtx) (resp interface{}, err error) {
		return middlewares[curr+1](ctx, getChainServerMiddlewareHandlerExample(middlewares, curr+1, finalHandler))
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

func (p *ServiceNameHTTPGoClient) RPCName(ctx context.Context, request *InputMsgName) (resp *OutputMsgName, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/%s/%d", request.StringArgument, request.Int64Argument))
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
	resp = &OutputMsgName{}
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (p *ServiceNameHTTPGoClient) AllTypesTest(ctx context.Context, request *AllTypesMsg) (resp *AllTypesMsg, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/%t/%s/%d/%d/%d/%d/%d/%d/%d/%d/%f/%d/%d/%f/%s/%s", request.BoolValue, request.EnumValue, request.Int32Value, request.Sint32Value, request.Uint32Value, request.Int64Value, request.Sint64Value, request.Uint64Value, request.Sfixed32Value, request.Fixed32Value, request.FloatValue, request.Sfixed64Value, request.Fixed64Value, request.DoubleValue, request.StringValue, request.BytesValue))
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
	resp = &AllTypesMsg{}
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (p *ServiceNameHTTPGoClient) CommonTypes(ctx context.Context, request *anypb.Any) (resp *emptypb.Empty, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/commonTypes"))
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
	resp = &emptypb.Empty{}
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (p *ServiceNameHTTPGoClient) Imports(ctx context.Context, request *somepackage.SomeCustomMsg1) (resp *somepackage.SomeCustomMsg2, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
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
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
	return resp, err
}

// SameInputAndOutput same types but different query, we need different query builder function
func (p *ServiceNameHTTPGoClient) SameInputAndOutput(ctx context.Context, request *InputMsgName) (resp *OutputMsgName, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/%s", request.StringArgument))
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
	resp = &OutputMsgName{}
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (p *ServiceNameHTTPGoClient) Optional(ctx context.Context, request *InputMsgName) (resp *OptionalField, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/optional"))
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
	resp = &OptionalField{}
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (p *ServiceNameHTTPGoClient) GetMethod(ctx context.Context, request *InputMsgName) (resp *OutputMsgName, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/get"))
	req.Header.SetMethod("GET")
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
	resp = &OutputMsgName{}
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (p *ServiceNameHTTPGoClient) CheckRepeated(ctx context.Context, request *RepeatedCheck) (resp *RepeatedCheck, err error) {
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req := &fasthttp.Request{}
	req.SetBody(body)
	StringValueArgRequest := strings.Join(request.StringValueArg, ",")
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/repeated/%s", StringValueArgRequest))
	req.Header.SetMethod("GET")
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
	resp = &RepeatedCheck{}
	if respEJ, ok := interface{}(resp).(easyjson.Unmarshaler); ok {
		if err = easyjson.Unmarshal(reqResp.Body(), respEJ); err != nil {
			return nil, err
		}
	} else {
		if err = json.Unmarshal(reqResp.Body(), resp); err != nil {
			return nil, err
		}
	}
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
