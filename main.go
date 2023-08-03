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
	}
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(gen *protogen.Plugin) (err error) {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			g := generator.NewGenerator(f, cfg)
			if err = g.GenerateServers(gen, f); err != nil {
				return err
			}
			if err = g.GenerateClients(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}
