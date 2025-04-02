//go:generate protoc -I=. -I=./vendor --go_out=. --go_opt=paths=source_relative proto/somepackage/somepackage.proto
//go:generate protoc -I=. -I=./vendor -I=./proto --go_out=. --go_opt=paths=source_relative proto/example.proto proto/example2.proto
//go:generate protoc -I=. -I=./vendor -I=./proto --httpgo_out=. --httpgo_opt=paths=source_relative,marshaller=easyjson,context=native proto/example.proto
//go:generate protoc -I=. -I=./vendor -I=./proto --httpgo_out=. --httpgo_opt=only=client,paths=source_relative,context=native proto/example2.proto
//go:generate protoc -I=. -I=./vendor -I=./proto --httpgo_out=. --httpgo_opt=paths=source_relative,autoURI=true,context=native proto/no_url.proto
//go:generate easyjson -all proto/example.pb.go

package main

import (
	_ "github.com/mailru/easyjson/gen"
	_ "google.golang.org/protobuf/encoding/protojson"
)
