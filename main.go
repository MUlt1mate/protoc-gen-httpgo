package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/MUlt1mate/protoc-gen-httpgo/generator"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) (err error) {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err = generator.GenerateServers(gen, f); err != nil {
				return err
			}
			if err = generator.GenerateClients(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}
