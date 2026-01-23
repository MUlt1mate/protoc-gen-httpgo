package implementation

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	proto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
	protofasthttp "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/fasthttp"
)

type Handler struct {
}

var _ protofasthttp.ServiceNameHTTPGoService = &Handler{}

func (h *Handler) RPCName(_ context.Context, request *proto.InputMsgName) (*proto.OutputMsgName, error) {
	p := &proto.OutputMsgName{
		StringValue: request.StringArgument,
		IntValue:    request.Int64Argument,
	}
	return p, nil
}

func (h *Handler) AllTypesTest(_ context.Context, msg *proto.AllTypesMsg) (*proto.AllTypesMsg, error) {
	return msg, nil
}

func (h *Handler) CommonTypes(_ context.Context, _ *anypb.Any) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h *Handler) SameInputAndOutput(_ context.Context, req *proto.InputMsgName) (*proto.OutputMsgName, error) {
	return &proto.OutputMsgName{
		StringValue: req.StringArgument,
		IntValue:    req.Int64Argument,
	}, nil
}

func (h *Handler) Optional(_ context.Context, req *proto.OptionalField) (*proto.OptionalField, error) {
	return req, nil
}

func (h *Handler) GetMethod(_ context.Context, req *proto.InputMsgName) (*proto.OutputMsgName, error) {
	return &proto.OutputMsgName{
		StringValue: req.StringArgument,
		IntValue:    req.Int64Argument,
	}, nil
}

func (h *Handler) CheckRepeatedPath(_ context.Context, req *proto.RepeatedCheck) (*proto.RepeatedCheck, error) {
	return req, nil
}

func (h *Handler) CheckRepeatedQuery(_ context.Context, req *proto.RepeatedCheck) (*proto.RepeatedCheck, error) {
	return req, nil
}

func (h *Handler) CheckRepeatedPost(_ context.Context, req *proto.RepeatedCheck) (*proto.RepeatedCheck, error) {
	return req, nil
}

func (h *Handler) EmptyGet(_ context.Context, _ *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (h *Handler) EmptyPost(_ context.Context, _ *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (h *Handler) TopLevelArray(ctx context.Context, req *proto.Array) (*proto.Array, error) {
	return req, nil
}

func (h *Handler) OnlyStructInGet(ctx context.Context, onlyStruct *proto.OnlyStruct) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (h *Handler) MultipartForm(ctx context.Context, request *proto.MultipartFormRequest) (*proto.Empty, error) {
	if request == nil || request.Document == nil {
		return nil, errors.New("empty request")
	}
	if diff := cmp.Diff(
		&MultipartFormRequestMsg,
		request,
		cmpopts.IgnoreUnexported(MultipartFormRequestMsg, proto.FileEx{}),
		cmpopts.IgnoreFields(proto.FileEx{}, "Headers"),
	); diff != "" {
		log.Println(diff)
	}
	return &proto.Empty{}, nil
}

func (h *Handler) MultipartFormAllTypes(ctx context.Context, request *proto.MultipartFormAllTypes) (*proto.Empty, error) {
	if request == nil || request.Document == nil {
		return nil, errors.New("empty request")
	}
	if diff := cmp.Diff(
		&MultipartFormRequestAllTypesMsg,
		request,
		cmpopts.IgnoreUnexported(MultipartFormRequestAllTypesMsg, proto.FileEx{}),
		cmpopts.IgnoreFields(proto.FileEx{}, "Headers"),
	); diff != "" {
		log.Println(diff)
	}
	return &proto.Empty{}, nil
}

func (h *Handler) AllTextTypesPost(ctx context.Context, msg *proto.AllTextTypesMsg) (*proto.AllTextTypesMsg, error) {
	if diff := cmp.Diff(&AllTextTypesMsg, msg, cmpopts.IgnoreUnexported(AllTextTypesMsg)); diff != "" {
		log.Println(diff)
	}
	return msg, nil
}

func (h *Handler) AllTextTypesGet(ctx context.Context, msg *proto.AllTextTypesMsg) (*proto.AllTextTypesMsg, error) {
	if diff := cmp.Diff(&AllTextTypesMsg, msg, cmpopts.IgnoreUnexported(AllTextTypesMsg)); diff != "" {
		log.Println(diff)
	}
	return msg, nil
}

func (h *Handler) AllTypesMaxTest(ctx context.Context, msg *proto.AllNumberTypesMsg) (*proto.AllNumberTypesMsg, error) {
	return msg, nil
}

func (h *Handler) AllTypesMaxQueryTest(ctx context.Context, msg *proto.AllNumberTypesMsg) (*proto.AllNumberTypesMsg, error) {
	return msg, nil
}

func (h *Handler) GetMessage(ctx context.Context, request *proto.GetMessageRequest) (*proto.Message, error) {
	if request.Name != "messages/123456" {
		return nil, fmt.Errorf("expected Name messages/123456, but got %s", request.Name)
	}
	return &proto.Message{}, nil
}

func (h *Handler) GetMessageV2(ctx context.Context, request *proto.GetMessageRequestV2) (*proto.Message, error) {
	if request.MessageId != "123456" {
		return nil, fmt.Errorf("expected MessageId 123456, but got %s", request.MessageId)
	}
	if request.Sub == nil {
		return nil, fmt.Errorf("empty sub")
	}
	if request.Sub.Subfield != "foo" {
		return nil, fmt.Errorf("expected subfield foo, but got %s", request.Sub.Subfield)
	}
	return &proto.Message{}, nil
}

func (h *Handler) UpdateMessage(ctx context.Context, request *proto.UpdateMessageRequest) (*proto.Message, error) {
	if request.MessageId != "123456" {
		return nil, fmt.Errorf("expected MessageId 123456, but got %s", request.MessageId)
	}
	if request.Message == nil {
		return nil, fmt.Errorf("empty sub")
	}
	if request.Message.Text != "Hi!" {
		return nil, fmt.Errorf("expected text Hi!, but got %s", request.Message.Text)
	}
	return &proto.Message{}, nil
}

func (h *Handler) UpdateMessageV2(ctx context.Context, request *proto.MessageV2) (*proto.MessageV2, error) {
	if request.MessageId != "123456" {
		return nil, fmt.Errorf("expected MessageId 123456, but got %s", request.MessageId)
	}
	if request.Text != "Hi!" {
		return nil, fmt.Errorf("expected text Hi!, but got %s", request.Text)
	}
	return &proto.MessageV2{}, nil
}

func (h *Handler) GetMessageV3(ctx context.Context, request *proto.GetMessageRequestV3) (*proto.MessageV2, error) {
	switch request.MessageId {
	case "123456":
		if request.UserId != "me" {
			return nil, fmt.Errorf("expected UserId me, but got %s", request.UserId)
		}
	case "234567":
		if request.UserId != "" {
			return nil, fmt.Errorf("expected UserId '', but got %s", request.UserId)
		}
	default:
		return nil, fmt.Errorf("unexpected MessageId: %s", request.MessageId)
	}
	return &proto.MessageV2{}, nil
}

func (h *Handler) GetMessageV4(ctx context.Context, request *proto.GetMessageRequestV3) (*proto.MessageV2, error) {
	if request.MessageId != "base/seg1/seg2.ext" {
		return nil, fmt.Errorf("unexpected MessageId: %s", request.MessageId)
	}
	return &proto.MessageV2{MessageId: request.MessageId}, nil
}

func (h *Handler) UpdateMessageV3(ctx context.Context, request *proto.UpdateMessageRequest) (*proto.UpdateMessageRequest, error) {
	return request, nil
}
