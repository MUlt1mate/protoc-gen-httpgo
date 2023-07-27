package implementation

import (
	"context"
	"log"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
)

type Handler struct {
}

var _ proto.ServiceNameHTTPGoService = &Handler{}

func (h *Handler) RPCName(_ context.Context, request *proto.InputMsgName) (*proto.OutputMsgName, error) {
	log.Println("request", request)
	p := &proto.OutputMsgName{
		StringValue: request.StringArgument,
		IntValue:    request.Int64Argument,
	}
	return p, nil
}

func (h *Handler) AllTypesTest(_ context.Context, msg *proto.AllTypesMsg) (*proto.AllTypesMsg, error) {
	log.Println("request", msg)
	p := &proto.AllTypesMsg{
		BoolValue:     msg.BoolValue,
		EnumValue:     msg.EnumValue,
		Int32Value:    msg.Int32Value,
		Sint32Value:   msg.Sint32Value,
		Uint32Value:   msg.Uint32Value,
		Int64Value:    msg.Int64Value,
		Sint64Value:   msg.Sint64Value,
		Uint64Value:   msg.Uint64Value,
		Sfixed32Value: msg.Sfixed32Value,
		Fixed32Value:  msg.Fixed32Value,
		FloatValue:    msg.FloatValue,
		Sfixed64Value: msg.Sfixed64Value,
		Fixed64Value:  msg.Fixed64Value,
		DoubleValue:   msg.DoubleValue,
		StringValue:   msg.StringValue,
		BytesValue:    msg.BytesValue,
	}
	return p, nil
}
