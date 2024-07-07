package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/MUlt1mate/protoc-gen-httpgo/generator"
)

var flags flag.FlagSet

func main() {
	cfg := generator.Config{
		Marshaller: flags.String("marshaller", "", "custom structs marshaller"),
		Only:       flags.String("only", "", "generate only server or client"),
	}
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(gen *protogen.Plugin) (err error) {
		return generator.Run(gen, cfg)
	})
}
