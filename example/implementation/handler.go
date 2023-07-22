package implementation

import (
	"context"
	"log"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
)

type Handler struct {
}

func (h *Handler) RPCName(_ context.Context, request *proto.InputMsgName) (*proto.OutputMsgName, error) {
	log.Println("request", request)
	p := &proto.OutputMsgName{
		StringValue: request.StringArgument,
		IntValue:    request.Int64Argument,
	}
	return p, nil
}
