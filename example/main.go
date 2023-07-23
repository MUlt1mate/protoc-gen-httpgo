package main

import (
	"context"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/MUlt1mate/protoc-gen-httpgo/example/implementation"
	"github.com/MUlt1mate/protoc-gen-httpgo/example/proto"
)

func main() {
	r := router.New()
	_ = proto.RegisterServiceNameHTTPServer(context.TODO(), r, &implementation.Handler{})

	go func() { _ = fasthttp.ListenAndServe(":8080", r.Handler) }()

	client, _ := proto.GetServiceNameClient(context.TODO(), &fasthttp.Client{}, "http://localhost:8080")
	resp, err := client.RPCName(context.Background(), &proto.InputMsgName{Int64Argument: 999, StringArgument: "rand"})
	log.Println(resp, err)
	allTypesresp, err := client.AllTypesTest(context.Background(), &proto.AllTypesMsg{
		BoolValue:     true,
		EnumValue:     proto.Options_SECOND,
		Int32Value:    1,
		Sint32Value:   2,
		Uint32Value:   3,
		Int64Value:    4,
		Sint64Value:   5,
		Uint64Value:   6,
		Sfixed32Value: 7,
		Fixed32Value:  8,
		FloatValue:    9,
		Sfixed64Value: 10,
		Fixed64Value:  11,
		DoubleValue:   12,
		StringValue:   "13",
		BytesValue:    []byte("14"),
	})
	log.Println(allTypesresp, err)
}
