package main

import (
	"context"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/fasthttp/router"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/MUlt1mate/protoc-gen-httpgo/benchmark/proto"
)

const (
	grpcServerEndpoint = "localhost:8001"
	httpServerEndpoint = "localhost:8002"
)

func main() {
	var (
		handler = Handler{}
		ctx     = context.Background()
	)

	switch os.Args[1] {
	case "grpc-gateway":
		go runGRPCServer(ctx, handler)
		go runGRPCGateway(ctx)
	case "httpgo":
		go runHTTPGo(ctx, handler)
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
	}()

	var c = make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c
	os.Exit(0)
}

func runGRPCGateway(ctx context.Context) {
	var err error
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = proto.RegisterAPIMeasureHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Fatal(err)
	}
	if err = http.ListenAndServe(httpServerEndpoint, mux); err != nil {
		log.Fatal(err)
	}
}

func runGRPCServer(ctx context.Context, handler Handler) {
	var (
		srv    = grpc.NewServer()
		config = net.ListenConfig{}
		listen net.Listener
		err    error
	)
	if listen, err = config.Listen(ctx, "tcp", grpcServerEndpoint); err != nil {
		log.Fatal(err)
	}
	if err = srv.Serve(listen); err != nil {
		log.Fatal(err)
	}
	proto.RegisterAPIMeasureServer(srv, &handler)
}

func runHTTPGo(ctx context.Context, handler Handler) {
	var fasthttpRouter = router.New()
	_ = proto.RegisterAPIMeasureHTTPGoServer(ctx, fasthttpRouter, &handler, nil)
	var srv = &fasthttp.Server{
		Handler: fasthttpRouter.Handler,
	}

	if err := srv.ListenAndServe(httpServerEndpoint); err != nil {
		log.Fatal(err)
	}
}
