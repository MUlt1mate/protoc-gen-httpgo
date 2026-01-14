package implementation

import (
	"context"
	"errors"
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

func (h *Handler) TopLevelArray(ctx context.Context, empty *proto.Empty) (*proto.Array, error) {
	return &proto.Array{Items: []*proto.ArrayItem{{Value: "a"}, {Value: "b"}}}, nil
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
