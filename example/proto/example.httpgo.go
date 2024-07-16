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
	Optional(context.Context, *OptionalField) (*OptionalField, error)
	GetMethod(context.Context, *InputMsgName) (*OutputMsgName, error)
	CheckRepeatedPath(context.Context, *RepeatedCheck) (*RepeatedCheck, error)
	CheckRepeatedQuery(context.Context, *RepeatedCheck) (*RepeatedCheck, error)
	CheckRepeatedPost(context.Context, *RepeatedCheck) (*RepeatedCheck, error)
	EmptyGet(context.Context, *Empty) (*Empty, error)
	EmptyPost(context.Context, *Empty) (*Empty, error)
}

func RegisterServiceNameHTTPGoServer(
	_ context.Context,
	r *router.Router,
	h ServiceNameHTTPGoService,
	middlewares []func(ctx *fasthttp.RequestCtx, req interface{}, handler func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) error {
	var middleware = chainServerMiddlewaresExample(middlewares)

	r.POST("/v1/test/{stringArgument}/{int64Argument}", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameRPCNameInputMsgName(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.RPCName(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.POST("/v1/test/{BoolValue}/{EnumValue}/{Int32Value}/{Sint32Value}/{Uint32Value}/{Int64Value}/{Sint64Value}/{Uint64Value}/{Sfixed32Value}/{Fixed32Value}/{FloatValue}/{Sfixed64Value}/{Fixed64Value}/{DoubleValue}/{StringValue}/{BytesValue}", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameAllTypesTestAllTypesMsg(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.AllTypesTest(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.POST("/v1/test/commonTypes", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameCommonTypesAny(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.CommonTypes(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.POST("/v1/test/imports", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameImportsSomeCustomMsg1(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.Imports(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	// same types but different query, we need different query builder function
	r.POST("/v1/test/{stringArgument}", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameSameInputAndOutputInputMsgName(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.SameInputAndOutput(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.POST("/v1/test/optional", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameOptionalOptionalField(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.Optional(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.GET("/v1/test/get", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameGetMethodInputMsgName(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.GetMethod(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.GET("/v1/repeated/{BoolValue}/{EnumValue}/{Int32Value}/{Sint32Value}/{Uint32Value}/{Int64Value}/{Sint64Value}/{Uint64Value}/{Sfixed32Value}/{Fixed32Value}/{FloatValue}/{Sfixed64Value}/{Fixed64Value}/{DoubleValue}/{StringValue}/{BytesValue}/{StringValueQuery}", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameCheckRepeatedPathRepeatedCheck(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.CheckRepeatedPath(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.GET("/v1/repeated/{StringValue}", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.CheckRepeatedQuery(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.POST("/v1/repeated/{StringValue}", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameCheckRepeatedPostRepeatedCheck(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.CheckRepeatedPost(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.GET("/v1/emptyGet", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameEmptyGetEmpty(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.EmptyGet(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
	})

	r.POST("/v1/emptyPost", func(ctx *fasthttp.RequestCtx) {
		input, err := buildExampleServiceNameEmptyPostEmpty(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
			return
		}
		handler := func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
			return h.EmptyPost(ctx, input)
		}
		if middleware == nil {
			_, _ = handler(ctx, input)
			return
		}
		_, _ = middleware(ctx, input, handler)
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
			Int64ArgumentValue, convErr := strconv.ParseInt(Int64ArgumentStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Argument: %w", convErr)
				return
			}
			arg.Int64Argument = Int64ArgumentValue
		case "stringArgument":
			StringArgumentValue := string(value)
			arg.StringArgument = StringArgumentValue
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
				BoolValueValue := true
				arg.BoolValue = BoolValueValue
			case "false", "f", "0":
				BoolValueValue := false
				arg.BoolValue = BoolValueValue
			default:
				err = fmt.Errorf("unknown bool string value %s", BoolValueStr)
				return
			}
		case "EnumValue":
			EnumValueStr := string(value)
			if OptionsValue, optValueOk := Options_value[strings.ToUpper(EnumValueStr)]; optValueOk {
				EnumValueOptionsValue := Options(OptionsValue)
				arg.EnumValue = EnumValueOptionsValue
			} else {
				if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
					if _, optIntValueOk := Options_name[int32(intOptionValue)]; optIntValueOk {
						EnumValueOptionsValue := Options(intOptionValue)
						arg.EnumValue = EnumValueOptionsValue
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
			Int32ValueValue := int32(Int32Value)
			arg.Int32Value = Int32ValueValue
		case "Sint32Value":
			Sint32ValueStr := string(value)
			Sint32Value, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			Sint32ValueValue := int32(Sint32Value)
			arg.Sint32Value = Sint32ValueValue
		case "Uint32Value":
			Uint32ValueStr := string(value)
			Uint32Value, convErr := strconv.ParseInt(Uint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			Uint32ValueValue := uint32(Uint32Value)
			arg.Uint32Value = Uint32ValueValue
		case "Int64Value":
			Int64ValueStr := string(value)
			Int64ValueValue, convErr := strconv.ParseInt(Int64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = Int64ValueValue
		case "Sint64Value":
			Sint64ValueStr := string(value)
			Sint64ValueValue, convErr := strconv.ParseInt(Sint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = Sint64ValueValue
		case "Uint64Value":
			Uint64ValueStr := string(value)
			Uint64Value, convErr := strconv.ParseInt(Uint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			Uint64ValueValue := uint64(Uint64Value)
			arg.Uint64Value = Uint64ValueValue
		case "Sfixed32Value":
			Sfixed32ValueStr := string(value)
			Sfixed32Value, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			Sfixed32ValueValue := int32(Sfixed32Value)
			arg.Sfixed32Value = Sfixed32ValueValue
		case "Fixed32Value":
			Fixed32ValueStr := string(value)
			Fixed32Value, convErr := strconv.ParseInt(Fixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			Fixed32ValueValue := uint32(Fixed32Value)
			arg.Fixed32Value = Fixed32ValueValue
		case "FloatValue":
			FloatValueStr := string(value)
			FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			FloatValueValue := float32(FloatValue)
			arg.FloatValue = FloatValueValue
		case "Sfixed64Value":
			Sfixed64ValueStr := string(value)
			Sfixed64ValueValue, convErr := strconv.ParseInt(Sfixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = Sfixed64ValueValue
		case "Fixed64Value":
			Fixed64ValueStr := string(value)
			Fixed64Value, convErr := strconv.ParseInt(Fixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			Fixed64ValueValue := uint64(Fixed64Value)
			arg.Fixed64Value = Fixed64ValueValue
		case "DoubleValue":
			DoubleValueStr := string(value)
			DoubleValueValue, convErr := strconv.ParseFloat(DoubleValueStr, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = DoubleValueValue
		case "StringValue":
			StringValueValue := string(value)
			arg.StringValue = StringValueValue
		case "BytesValue":
			arg.BytesValue = value
		case "MessageValue":
			err = fmt.Errorf("unsupported type message for query argument MessageValue")
			return
		case "SliceStringValue[]":
			SliceStringValue := string(value)
			arg.SliceStringValue = append(arg.SliceStringValue, SliceStringValue)
		case "SliceInt32Value[]":
			SliceInt32ValueStr := string(value)
			SliceInt32ValueVal, convErr := strconv.ParseInt(SliceInt32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter SliceInt32Value: %w", convErr)
				return
			}
			arg.SliceInt32Value = append(arg.SliceInt32Value, int32(SliceInt32ValueVal))
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
	OptionsValue, optValueOk := Options_value[strings.ToUpper(EnumValueStr)]
	if optValueOk {
		arg.EnumValue = Options(OptionsValue)
	} else {
		if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
			if _, optIntValueOk := Options_name[int32(intOptionValue)]; optIntValueOk {
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
			TypeUrlValue := string(value)
			arg.TypeUrl = TypeUrlValue
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
			ValValue := string(value)
			arg.Val = ValValue
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
			Int64ArgumentValue, convErr := strconv.ParseInt(Int64ArgumentStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Argument: %w", convErr)
				return
			}
			arg.Int64Argument = Int64ArgumentValue
		case "stringArgument":
			StringArgumentValue := string(value)
			arg.StringArgument = StringArgumentValue
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

func buildExampleServiceNameOptionalOptionalField(ctx *fasthttp.RequestCtx) (arg *OptionalField, err error) {
	arg = &OptionalField{}
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
				BoolValueValue := true
				arg.BoolValue = &BoolValueValue
			case "false", "f", "0":
				BoolValueValue := false
				arg.BoolValue = &BoolValueValue
			default:
				err = fmt.Errorf("unknown bool string value %s", BoolValueStr)
				return
			}
		case "EnumValue":
			EnumValueStr := string(value)
			if OptionsValue, optValueOk := Options_value[strings.ToUpper(EnumValueStr)]; optValueOk {
				EnumValueOptionsValue := Options(OptionsValue)
				arg.EnumValue = &EnumValueOptionsValue
			} else {
				if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
					if _, optIntValueOk := Options_name[int32(intOptionValue)]; optIntValueOk {
						EnumValueOptionsValue := Options(intOptionValue)
						arg.EnumValue = &EnumValueOptionsValue
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
			Int32ValueValue := int32(Int32Value)
			arg.Int32Value = &Int32ValueValue
		case "Sint32Value":
			Sint32ValueStr := string(value)
			Sint32Value, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			Sint32ValueValue := int32(Sint32Value)
			arg.Sint32Value = &Sint32ValueValue
		case "Uint32Value":
			Uint32ValueStr := string(value)
			Uint32Value, convErr := strconv.ParseInt(Uint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			Uint32ValueValue := uint32(Uint32Value)
			arg.Uint32Value = &Uint32ValueValue
		case "Int64Value":
			Int64ValueStr := string(value)
			Int64ValueValue, convErr := strconv.ParseInt(Int64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = &Int64ValueValue
		case "Sint64Value":
			Sint64ValueStr := string(value)
			Sint64ValueValue, convErr := strconv.ParseInt(Sint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = &Sint64ValueValue
		case "Uint64Value":
			Uint64ValueStr := string(value)
			Uint64Value, convErr := strconv.ParseInt(Uint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			Uint64ValueValue := uint64(Uint64Value)
			arg.Uint64Value = &Uint64ValueValue
		case "Sfixed32Value":
			Sfixed32ValueStr := string(value)
			Sfixed32Value, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			Sfixed32ValueValue := int32(Sfixed32Value)
			arg.Sfixed32Value = &Sfixed32ValueValue
		case "Fixed32Value":
			Fixed32ValueStr := string(value)
			Fixed32Value, convErr := strconv.ParseInt(Fixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			Fixed32ValueValue := uint32(Fixed32Value)
			arg.Fixed32Value = &Fixed32ValueValue
		case "FloatValue":
			FloatValueStr := string(value)
			FloatValue, convErr := strconv.ParseFloat(FloatValueStr, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			FloatValueValue := float32(FloatValue)
			arg.FloatValue = &FloatValueValue
		case "Sfixed64Value":
			Sfixed64ValueStr := string(value)
			Sfixed64ValueValue, convErr := strconv.ParseInt(Sfixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = &Sfixed64ValueValue
		case "Fixed64Value":
			Fixed64ValueStr := string(value)
			Fixed64Value, convErr := strconv.ParseInt(Fixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			Fixed64ValueValue := uint64(Fixed64Value)
			arg.Fixed64Value = &Fixed64ValueValue
		case "DoubleValue":
			DoubleValueStr := string(value)
			DoubleValueValue, convErr := strconv.ParseFloat(DoubleValueStr, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = &DoubleValueValue
		case "StringValue":
			StringValueValue := string(value)
			arg.StringValue = &StringValueValue
		case "BytesValue":
			arg.BytesValue = value
		case "MessageValue":
			err = fmt.Errorf("unsupported type message for query argument MessageValue")
			return
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameGetMethodInputMsgName(ctx *fasthttp.RequestCtx) (arg *InputMsgName, err error) {
	arg = &InputMsgName{}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "int64Argument":
			Int64ArgumentStr := string(value)
			Int64ArgumentValue, convErr := strconv.ParseInt(Int64ArgumentStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Argument: %w", convErr)
				return
			}
			arg.Int64Argument = Int64ArgumentValue
		case "stringArgument":
			StringArgumentValue := string(value)
			arg.StringArgument = StringArgumentValue
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameCheckRepeatedPathRepeatedCheck(ctx *fasthttp.RequestCtx) (arg *RepeatedCheck, err error) {
	arg = &RepeatedCheck{}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "BoolValue[]":
			BoolValueStr := string(value)
			switch BoolValueStr {
			case "true", "t", "1":
				arg.BoolValue = append(arg.BoolValue, true)
			case "false", "f", "0":
				arg.BoolValue = append(arg.BoolValue, false)
			default:
				err = fmt.Errorf("unknown bool string value %s", BoolValueStr)
				return
			}
		case "EnumValue[]":
			EnumValueStr := string(value)
			if OptionsValue, optValueOk := Options_value[strings.ToUpper(EnumValueStr)]; optValueOk {
				arg.EnumValue = append(arg.EnumValue, Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
					if _, optIntValueOk := Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = append(arg.EnumValue, Options(intOptionValue))
					}
				}
			}
		case "Int32Value[]":
			Int32ValueStr := string(value)
			Int32ValueVal, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = append(arg.Int32Value, int32(Int32ValueVal))
		case "Sint32Value[]":
			Sint32ValueStr := string(value)
			Sint32ValueVal, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = append(arg.Sint32Value, int32(Sint32ValueVal))
		case "Uint32Value[]":
			Uint32ValueStr := string(value)
			Uint32ValueVal, convErr := strconv.ParseInt(Uint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32ValueVal))
		case "Int64Value[]":
			Int64ValueStr := string(value)
			Int64ValueVal, convErr := strconv.ParseInt(Int64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = append(arg.Int64Value, Int64ValueVal)
		case "Sint64Value[]":
			Sint64ValueStr := string(value)
			Sint64ValueVal, convErr := strconv.ParseInt(Sint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = append(arg.Sint64Value, Sint64ValueVal)
		case "Uint64Value[]":
			Uint64ValueStr := string(value)
			Uint64ValueVal, convErr := strconv.ParseUint(Uint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = append(arg.Uint64Value, Uint64ValueVal)
		case "Sfixed32Value[]":
			Sfixed32ValueStr := string(value)
			Sfixed32ValueVal, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32ValueVal))
		case "Fixed32Value[]":
			Fixed32ValueStr := string(value)
			Fixed32ValueVal, convErr := strconv.ParseInt(Fixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32ValueVal))
		case "FloatValue[]":
			FloatValueStr := string(value)
			FloatValueVal, convErr := strconv.ParseFloat(FloatValueStr, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = append(arg.FloatValue, float32(FloatValueVal))
		case "Sfixed64Value[]":
			Sfixed64ValueStr := string(value)
			Sfixed64ValueVal, convErr := strconv.ParseInt(Sfixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64ValueVal)
		case "Fixed64Value[]":
			Fixed64ValueStr := string(value)
			Fixed64ValueVal, convErr := strconv.ParseUint(Fixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = append(arg.Fixed64Value, Fixed64ValueVal)
		case "DoubleValue[]":
			DoubleValueStr := string(value)
			DoubleValueVal, convErr := strconv.ParseFloat(DoubleValueStr, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = append(arg.DoubleValue, DoubleValueVal)
		case "StringValue[]":
			StringValue := string(value)
			arg.StringValue = append(arg.StringValue, StringValue)
		case "BytesValue[]":
			BytesValue := value
			arg.BytesValue = append(arg.BytesValue, BytesValue)
		case "StringValueQuery[]":
			StringValueQuery := string(value)
			arg.StringValueQuery = append(arg.StringValueQuery, StringValueQuery)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	BoolValueStr, ok := ctx.UserValue("BoolValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter BoolValue")
	}
	BoolValueStrs := strings.Split(BoolValueStr, ",")
	for _, str := range BoolValueStrs {
		switch str {
		case "true", "t", "1":
			arg.BoolValue = append(arg.BoolValue, true)
		case "false", "f", "0":
			arg.BoolValue = append(arg.BoolValue, false)
		default:
			err = fmt.Errorf("unknown bool string value %s", str)
			return nil, err
		}
	}
	EnumValueStr, ok := ctx.UserValue("EnumValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter EnumValue")
	}
	EnumValueStrs := strings.Split(EnumValueStr, ",")
	for _, str := range EnumValueStrs {
		if OptionsValue, optValueOk := Options_value[strings.ToUpper(str)]; optValueOk {
			arg.EnumValue = append(arg.EnumValue, Options(OptionsValue))
		} else {
			if intOptionValue, convErr := strconv.ParseInt(str, 10, 32); convErr == nil {
				if _, optIntValueOk := Options_name[int32(intOptionValue)]; optIntValueOk {
					arg.EnumValue = append(arg.EnumValue, Options(intOptionValue))
				}
			}
		}
	}
	Int32ValueStr, ok := ctx.UserValue("Int32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Int32Value")
	}
	Int32ValueStrs := strings.Split(Int32ValueStr, ",")
	for _, str := range Int32ValueStrs {
		Int32ValueVal, convErr := strconv.ParseInt(str, 10, 32)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
			return nil, err
		}
		arg.Int32Value = append(arg.Int32Value, int32(Int32ValueVal))
	}
	Sint32ValueStr, ok := ctx.UserValue("Sint32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sint32Value")
	}
	Sint32ValueStrs := strings.Split(Sint32ValueStr, ",")
	for _, str := range Sint32ValueStrs {
		Sint32ValueVal, convErr := strconv.ParseInt(str, 10, 32)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
			return nil, err
		}
		arg.Sint32Value = append(arg.Sint32Value, int32(Sint32ValueVal))
	}
	Uint32ValueStr, ok := ctx.UserValue("Uint32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Uint32Value")
	}
	Uint32ValueStrs := strings.Split(Uint32ValueStr, ",")
	for _, str := range Uint32ValueStrs {
		Uint32ValueVal, convErr := strconv.ParseInt(str, 10, 32)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
			return nil, err
		}
		arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32ValueVal))
	}
	Int64ValueStr, ok := ctx.UserValue("Int64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Int64Value")
	}
	Int64ValueStrs := strings.Split(Int64ValueStr, ",")
	for _, str := range Int64ValueStrs {
		Int64ValueVal, convErr := strconv.ParseInt(str, 10, 64)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
			return nil, err
		}
		arg.Int64Value = append(arg.Int64Value, Int64ValueVal)
	}
	Sint64ValueStr, ok := ctx.UserValue("Sint64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sint64Value")
	}
	Sint64ValueStrs := strings.Split(Sint64ValueStr, ",")
	for _, str := range Sint64ValueStrs {
		Sint64ValueVal, convErr := strconv.ParseInt(str, 10, 64)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
			return nil, err
		}
		arg.Sint64Value = append(arg.Sint64Value, Sint64ValueVal)
	}
	Uint64ValueStr, ok := ctx.UserValue("Uint64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Uint64Value")
	}
	Uint64ValueStrs := strings.Split(Uint64ValueStr, ",")
	for _, str := range Uint64ValueStrs {
		Uint64ValueVal, convErr := strconv.ParseUint(str, 10, 64)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
			return nil, err
		}
		arg.Uint64Value = append(arg.Uint64Value, Uint64ValueVal)
	}
	Sfixed32ValueStr, ok := ctx.UserValue("Sfixed32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sfixed32Value")
	}
	Sfixed32ValueStrs := strings.Split(Sfixed32ValueStr, ",")
	for _, str := range Sfixed32ValueStrs {
		Sfixed32ValueVal, convErr := strconv.ParseInt(str, 10, 32)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
			return nil, err
		}
		arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32ValueVal))
	}
	Fixed32ValueStr, ok := ctx.UserValue("Fixed32Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Fixed32Value")
	}
	Fixed32ValueStrs := strings.Split(Fixed32ValueStr, ",")
	for _, str := range Fixed32ValueStrs {
		Fixed32ValueVal, convErr := strconv.ParseInt(str, 10, 32)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
			return nil, err
		}
		arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32ValueVal))
	}
	FloatValueStr, ok := ctx.UserValue("FloatValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter FloatValue")
	}
	FloatValueStrs := strings.Split(FloatValueStr, ",")
	for _, str := range FloatValueStrs {
		FloatValueVal, convErr := strconv.ParseFloat(str, 32)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
			return nil, err
		}
		arg.FloatValue = append(arg.FloatValue, float32(FloatValueVal))
	}
	Sfixed64ValueStr, ok := ctx.UserValue("Sfixed64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Sfixed64Value")
	}
	Sfixed64ValueStrs := strings.Split(Sfixed64ValueStr, ",")
	for _, str := range Sfixed64ValueStrs {
		Sfixed64ValueVal, convErr := strconv.ParseInt(str, 10, 64)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
			return nil, err
		}
		arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64ValueVal)
	}
	Fixed64ValueStr, ok := ctx.UserValue("Fixed64Value").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter Fixed64Value")
	}
	Fixed64ValueStrs := strings.Split(Fixed64ValueStr, ",")
	for _, str := range Fixed64ValueStrs {
		Fixed64ValueVal, convErr := strconv.ParseUint(str, 10, 64)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
			return nil, err
		}
		arg.Fixed64Value = append(arg.Fixed64Value, Fixed64ValueVal)
	}
	DoubleValueStr, ok := ctx.UserValue("DoubleValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter DoubleValue")
	}
	DoubleValueStrs := strings.Split(DoubleValueStr, ",")
	for _, str := range DoubleValueStrs {
		DoubleValueVal, convErr := strconv.ParseFloat(str, 64)
		if convErr != nil {
			err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
			return nil, err
		}
		arg.DoubleValue = append(arg.DoubleValue, DoubleValueVal)
	}
	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringValue")
	}
	arg.StringValue = strings.Split(StringValueStr, ",")
	BytesValueStr, ok := ctx.UserValue("BytesValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter BytesValue")
	}
	arg.BytesValue = append(arg.BytesValue, []byte(BytesValueStr))
	StringValueQueryStr, ok := ctx.UserValue("StringValueQuery").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringValueQuery")
	}
	arg.StringValueQuery = strings.Split(StringValueQueryStr, ",")
	return arg, err
}

func buildExampleServiceNameCheckRepeatedQueryRepeatedCheck(ctx *fasthttp.RequestCtx) (arg *RepeatedCheck, err error) {
	arg = &RepeatedCheck{}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		case "BoolValue[]":
			BoolValueStr := string(value)
			switch BoolValueStr {
			case "true", "t", "1":
				arg.BoolValue = append(arg.BoolValue, true)
			case "false", "f", "0":
				arg.BoolValue = append(arg.BoolValue, false)
			default:
				err = fmt.Errorf("unknown bool string value %s", BoolValueStr)
				return
			}
		case "EnumValue[]":
			EnumValueStr := string(value)
			if OptionsValue, optValueOk := Options_value[strings.ToUpper(EnumValueStr)]; optValueOk {
				arg.EnumValue = append(arg.EnumValue, Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
					if _, optIntValueOk := Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = append(arg.EnumValue, Options(intOptionValue))
					}
				}
			}
		case "Int32Value[]":
			Int32ValueStr := string(value)
			Int32ValueVal, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = append(arg.Int32Value, int32(Int32ValueVal))
		case "Sint32Value[]":
			Sint32ValueStr := string(value)
			Sint32ValueVal, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = append(arg.Sint32Value, int32(Sint32ValueVal))
		case "Uint32Value[]":
			Uint32ValueStr := string(value)
			Uint32ValueVal, convErr := strconv.ParseInt(Uint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32ValueVal))
		case "Int64Value[]":
			Int64ValueStr := string(value)
			Int64ValueVal, convErr := strconv.ParseInt(Int64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = append(arg.Int64Value, Int64ValueVal)
		case "Sint64Value[]":
			Sint64ValueStr := string(value)
			Sint64ValueVal, convErr := strconv.ParseInt(Sint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = append(arg.Sint64Value, Sint64ValueVal)
		case "Uint64Value[]":
			Uint64ValueStr := string(value)
			Uint64ValueVal, convErr := strconv.ParseUint(Uint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = append(arg.Uint64Value, Uint64ValueVal)
		case "Sfixed32Value[]":
			Sfixed32ValueStr := string(value)
			Sfixed32ValueVal, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32ValueVal))
		case "Fixed32Value[]":
			Fixed32ValueStr := string(value)
			Fixed32ValueVal, convErr := strconv.ParseInt(Fixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32ValueVal))
		case "FloatValue[]":
			FloatValueStr := string(value)
			FloatValueVal, convErr := strconv.ParseFloat(FloatValueStr, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = append(arg.FloatValue, float32(FloatValueVal))
		case "Sfixed64Value[]":
			Sfixed64ValueStr := string(value)
			Sfixed64ValueVal, convErr := strconv.ParseInt(Sfixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64ValueVal)
		case "Fixed64Value[]":
			Fixed64ValueStr := string(value)
			Fixed64ValueVal, convErr := strconv.ParseUint(Fixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = append(arg.Fixed64Value, Fixed64ValueVal)
		case "DoubleValue[]":
			DoubleValueStr := string(value)
			DoubleValueVal, convErr := strconv.ParseFloat(DoubleValueStr, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = append(arg.DoubleValue, DoubleValueVal)
		case "StringValue[]":
			StringValue := string(value)
			arg.StringValue = append(arg.StringValue, StringValue)
		case "BytesValue[]":
			BytesValue := value
			arg.BytesValue = append(arg.BytesValue, BytesValue)
		case "StringValueQuery[]":
			StringValueQuery := string(value)
			arg.StringValueQuery = append(arg.StringValueQuery, StringValueQuery)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringValue")
	}
	arg.StringValue = strings.Split(StringValueStr, ",")
	return arg, err
}

func buildExampleServiceNameCheckRepeatedPostRepeatedCheck(ctx *fasthttp.RequestCtx) (arg *RepeatedCheck, err error) {
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
		case "BoolValue[]":
			BoolValueStr := string(value)
			switch BoolValueStr {
			case "true", "t", "1":
				arg.BoolValue = append(arg.BoolValue, true)
			case "false", "f", "0":
				arg.BoolValue = append(arg.BoolValue, false)
			default:
				err = fmt.Errorf("unknown bool string value %s", BoolValueStr)
				return
			}
		case "EnumValue[]":
			EnumValueStr := string(value)
			if OptionsValue, optValueOk := Options_value[strings.ToUpper(EnumValueStr)]; optValueOk {
				arg.EnumValue = append(arg.EnumValue, Options(OptionsValue))
			} else {
				if intOptionValue, convErr := strconv.ParseInt(EnumValueStr, 10, 32); convErr == nil {
					if _, optIntValueOk := Options_name[int32(intOptionValue)]; optIntValueOk {
						arg.EnumValue = append(arg.EnumValue, Options(intOptionValue))
					}
				}
			}
		case "Int32Value[]":
			Int32ValueStr := string(value)
			Int32ValueVal, convErr := strconv.ParseInt(Int32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int32Value: %w", convErr)
				return
			}
			arg.Int32Value = append(arg.Int32Value, int32(Int32ValueVal))
		case "Sint32Value[]":
			Sint32ValueStr := string(value)
			Sint32ValueVal, convErr := strconv.ParseInt(Sint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint32Value: %w", convErr)
				return
			}
			arg.Sint32Value = append(arg.Sint32Value, int32(Sint32ValueVal))
		case "Uint32Value[]":
			Uint32ValueStr := string(value)
			Uint32ValueVal, convErr := strconv.ParseInt(Uint32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint32Value: %w", convErr)
				return
			}
			arg.Uint32Value = append(arg.Uint32Value, uint32(Uint32ValueVal))
		case "Int64Value[]":
			Int64ValueStr := string(value)
			Int64ValueVal, convErr := strconv.ParseInt(Int64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Int64Value: %w", convErr)
				return
			}
			arg.Int64Value = append(arg.Int64Value, Int64ValueVal)
		case "Sint64Value[]":
			Sint64ValueStr := string(value)
			Sint64ValueVal, convErr := strconv.ParseInt(Sint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sint64Value: %w", convErr)
				return
			}
			arg.Sint64Value = append(arg.Sint64Value, Sint64ValueVal)
		case "Uint64Value[]":
			Uint64ValueStr := string(value)
			Uint64ValueVal, convErr := strconv.ParseUint(Uint64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Uint64Value: %w", convErr)
				return
			}
			arg.Uint64Value = append(arg.Uint64Value, Uint64ValueVal)
		case "Sfixed32Value[]":
			Sfixed32ValueStr := string(value)
			Sfixed32ValueVal, convErr := strconv.ParseInt(Sfixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed32Value: %w", convErr)
				return
			}
			arg.Sfixed32Value = append(arg.Sfixed32Value, int32(Sfixed32ValueVal))
		case "Fixed32Value[]":
			Fixed32ValueStr := string(value)
			Fixed32ValueVal, convErr := strconv.ParseInt(Fixed32ValueStr, 10, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed32Value: %w", convErr)
				return
			}
			arg.Fixed32Value = append(arg.Fixed32Value, uint32(Fixed32ValueVal))
		case "FloatValue[]":
			FloatValueStr := string(value)
			FloatValueVal, convErr := strconv.ParseFloat(FloatValueStr, 32)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter FloatValue: %w", convErr)
				return
			}
			arg.FloatValue = append(arg.FloatValue, float32(FloatValueVal))
		case "Sfixed64Value[]":
			Sfixed64ValueStr := string(value)
			Sfixed64ValueVal, convErr := strconv.ParseInt(Sfixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Sfixed64Value: %w", convErr)
				return
			}
			arg.Sfixed64Value = append(arg.Sfixed64Value, Sfixed64ValueVal)
		case "Fixed64Value[]":
			Fixed64ValueStr := string(value)
			Fixed64ValueVal, convErr := strconv.ParseUint(Fixed64ValueStr, 10, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter Fixed64Value: %w", convErr)
				return
			}
			arg.Fixed64Value = append(arg.Fixed64Value, Fixed64ValueVal)
		case "DoubleValue[]":
			DoubleValueStr := string(value)
			DoubleValueVal, convErr := strconv.ParseFloat(DoubleValueStr, 64)
			if convErr != nil {
				err = fmt.Errorf("conversion failed for parameter DoubleValue: %w", convErr)
				return
			}
			arg.DoubleValue = append(arg.DoubleValue, DoubleValueVal)
		case "StringValue[]":
			StringValue := string(value)
			arg.StringValue = append(arg.StringValue, StringValue)
		case "BytesValue[]":
			BytesValue := value
			arg.BytesValue = append(arg.BytesValue, BytesValue)
		case "StringValueQuery[]":
			StringValueQuery := string(value)
			arg.StringValueQuery = append(arg.StringValueQuery, StringValueQuery)
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	StringValueStr, ok := ctx.UserValue("StringValue").(string)
	if !ok {
		return nil, errors.New("incorrect type for parameter StringValue")
	}
	arg.StringValue = strings.Split(StringValueStr, ",")
	return arg, err
}

func buildExampleServiceNameEmptyGetEmpty(ctx *fasthttp.RequestCtx) (arg *Empty, err error) {
	arg = &Empty{}
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		var strKey = string(key)
		switch strKey {
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func buildExampleServiceNameEmptyPostEmpty(ctx *fasthttp.RequestCtx) (arg *Empty, err error) {
	arg = &Empty{}
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
		default:
			err = fmt.Errorf("unknown query parameter %s", strKey)
			return
		}
	})
	return arg, err
}

func chainServerMiddlewaresExample(
	middlewares []func(ctx *fasthttp.RequestCtx, req interface{}, handler func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
) func(ctx *fasthttp.RequestCtx, req interface{}, handler func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
	switch len(middlewares) {
	case 0:
		return nil
	case 1:
		return middlewares[0]
	default:
		return func(ctx *fasthttp.RequestCtx, req interface{}, handler func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
			return middlewares[0](ctx, req, getChainServerMiddlewareHandlerExample(middlewares, 0, handler))
		}
	}
}

func getChainServerMiddlewareHandlerExample(
	middlewares []func(ctx *fasthttp.RequestCtx, req interface{}, handler func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error)) (resp interface{}, err error),
	curr int,
	finalHandler func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error),
) func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
	if curr == len(middlewares)-1 {
		return finalHandler
	}
	return func(ctx *fasthttp.RequestCtx, req interface{}) (resp interface{}, err error) {
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

func (p *ServiceNameHTTPGoClient) RPCName(ctx context.Context, request *InputMsgName) (resp *OutputMsgName, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/%s/%d%s", request.StringArgument, request.Int64Argument, queryArgs))
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
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/%t/%s/%d/%d/%d/%d/%d/%d/%d/%d/%f/%d/%d/%f/%s/%s%s", request.BoolValue, request.EnumValue, request.Int32Value, request.Sint32Value, request.Uint32Value, request.Int64Value, request.Sint64Value, request.Uint64Value, request.Sfixed32Value, request.Fixed32Value, request.FloatValue, request.Sfixed64Value, request.Fixed64Value, request.DoubleValue, request.StringValue, request.BytesValue, queryArgs))
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
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/commonTypes%s", queryArgs))
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
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/imports%s", queryArgs))
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
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/%s%s", request.StringArgument, queryArgs))
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

func (p *ServiceNameHTTPGoClient) Optional(ctx context.Context, request *OptionalField) (resp *OptionalField, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/optional%s", queryArgs))
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
	req := &fasthttp.Request{}
	var queryArgs string
	var parameters = []string{
		"Int64Argument=%d",
		"StringArgument=%s",
	}
	var values = []interface{}{
		request.Int64Argument,
		request.StringArgument,
	}
	queryArgs = fmt.Sprintf("?"+strings.Join(parameters, "&"), values...)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/test/get%s", queryArgs))
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

func (p *ServiceNameHTTPGoClient) CheckRepeatedPath(ctx context.Context, request *RepeatedCheck) (resp *RepeatedCheck, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	BoolValueStrs := make([]string, len(request.BoolValue))
	for i, v := range request.BoolValue {
		if v {
			BoolValueStrs[i] = "true"
		} else {
			BoolValueStrs[i] = "false"
		}
	}
	BoolValueRequest := strings.Join(BoolValueStrs, ",")
	EnumValueStrs := make([]string, len(request.EnumValue))
	for i, v := range request.EnumValue {
		EnumValueStrs[i] = strconv.FormatInt(int64(v), 10)
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
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/repeated/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s%s", BoolValueRequest, EnumValueRequest, Int32ValueRequest, Sint32ValueRequest, Uint32ValueRequest, Int64ValueRequest, Sint64ValueRequest, Uint64ValueRequest, Sfixed32ValueRequest, Fixed32ValueRequest, FloatValueRequest, Sfixed64ValueRequest, Fixed64ValueRequest, DoubleValueRequest, StringValueRequest, BytesValueRequest, StringValueQueryRequest, queryArgs))
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

func (p *ServiceNameHTTPGoClient) CheckRepeatedQuery(ctx context.Context, request *RepeatedCheck) (resp *RepeatedCheck, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	var parameters = []string{}
	var values = []interface{}{}
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
	StringValueRequest := strings.Join(request.StringValue, ",")
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/repeated/%s%s", StringValueRequest, queryArgs))
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

func (p *ServiceNameHTTPGoClient) CheckRepeatedPost(ctx context.Context, request *RepeatedCheck) (resp *RepeatedCheck, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	StringValueRequest := strings.Join(request.StringValue, ",")
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/repeated/%s%s", StringValueRequest, queryArgs))
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

func (p *ServiceNameHTTPGoClient) EmptyGet(ctx context.Context, request *Empty) (resp *Empty, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/emptyGet%s", queryArgs))
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
	resp = &Empty{}
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

func (p *ServiceNameHTTPGoClient) EmptyPost(ctx context.Context, request *Empty) (resp *Empty, err error) {
	req := &fasthttp.Request{}
	var queryArgs string
	var body []byte
	if rqEJ, ok := interface{}(request).(easyjson.Marshaler); ok {
		body, err = easyjson.Marshal(rqEJ)
	} else {
		body, err = json.Marshal(request)
	}
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	req.SetRequestURI(p.host + fmt.Sprintf("/v1/emptyPost%s", queryArgs))
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
	resp = &Empty{}
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
