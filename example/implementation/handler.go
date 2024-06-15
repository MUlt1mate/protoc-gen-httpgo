package implementation

import (
	"context"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto/somepackage"
)

type Handler struct {
}

var _ proto.ServiceNameHTTPGoService = &Handler{}

func (h *Handler) RPCName(_ context.Context, request *proto.InputMsgName) (*proto.OutputMsgName, error) {
	p := &proto.OutputMsgName{
		StringValue: request.StringArgument,
		IntValue:    request.Int64Argument,
	}
	return p, nil
}

func (h *Handler) AllTypesTest(_ context.Context, msg *proto.AllTypesMsg) (*proto.AllTypesMsg, error) {
	p := &proto.AllTypesMsg{
		BoolValue:        msg.BoolValue,
		EnumValue:        msg.EnumValue,
		Int32Value:       msg.Int32Value,
		Sint32Value:      msg.Sint32Value,
		Uint32Value:      msg.Uint32Value,
		Int64Value:       msg.Int64Value,
		Sint64Value:      msg.Sint64Value,
		Uint64Value:      msg.Uint64Value,
		Sfixed32Value:    msg.Sfixed32Value,
		Fixed32Value:     msg.Fixed32Value,
		FloatValue:       msg.FloatValue,
		Sfixed64Value:    msg.Sfixed64Value,
		Fixed64Value:     msg.Fixed64Value,
		DoubleValue:      msg.DoubleValue,
		StringValue:      msg.StringValue,
		BytesValue:       msg.BytesValue,
		SliceStringValue: msg.SliceStringValue,
	}
	return p, nil
}

func (h *Handler) CommonTypes(_ context.Context, _ *anypb.Any) (*emptypb.Empty, error) {
	panic("implement me")
}

func (h *Handler) Imports(_ context.Context, _ *somepackage.SomeCustomMsg1) (*somepackage.SomeCustomMsg2, error) {
	panic("implement me")
}

func (h *Handler) SameInputAndOutput(_ context.Context, _ *proto.InputMsgName) (*proto.OutputMsgName, error) {
	panic("implement me")
}

func (h *Handler) Optional(_ context.Context, _ *proto.InputMsgName) (*proto.OptionalField, error) {
	panic("implement me")
}

func (h *Handler) GetMethod(_ context.Context, req *proto.InputMsgName) (*proto.OutputMsgName, error) {
	return &proto.OutputMsgName{
		StringValue: req.StringArgument,
		IntValue:    req.Int64Argument,
	}, nil
}
