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
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			g := generator.NewGenerator(
				f,
				cfg,
				gen.NewGeneratedFile(f.GeneratedFilenamePrefix+".httpgo.go", f.GoImportPath),
			)
			if err = g.GenerateServers(f); err != nil {
				return err
			}
			if err = g.GenerateClients(); err != nil {
				return err
			}
		}
		return nil
	})
}
