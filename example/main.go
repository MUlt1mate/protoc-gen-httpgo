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
}
