// Package main runs code generation with parameters passed by protogen
package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/MUlt1mate/protoc-gen-httpgo/generator"
)

var flags flag.FlagSet

func main() {
	cfg := generator.Config{
		Marshaller:         flags.String("marshaller", "", "custom structs marshaller"),
		Only:               flags.String("only", "", "generate only server or client"),
		AutoURI:            flags.Bool("autoURI", false, "create method URI if annotation is missing"),
		BodylessMethodsStr: flags.String("bodylessMethods", "", "list of semicolon separated http methods that should not have a body"),
		ContextStruct:      flags.String("context", "", "server context type (native|fasthttp)"),
		Library:            flags.String("library", "fasthttp", "implementation library (nethttp|fasthttp)"),
	}
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(gen *protogen.Plugin) (err error) {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generator.Run(gen, cfg)
	})
}
