//go:generate protoc -I=. -I=./vendor -I=/usr/local/include --go_out=. --go_opt=paths=source_relative proto/somepackage/somepackage.proto
//go:generate protoc -I=. -I=./vendor -I=/usr/local/include -I=./proto --go_out=. --go_opt=paths=source_relative proto/example.proto proto/example2.proto
//go:generate protoc -I=. -I=./vendor -I=/usr/local/include -I=./proto --httpgo_out=paths=source_relative,marshaller=easyjson,context=native:. proto/example.proto
//go:generate protoc -I=. -I=./vendor -I=/usr/local/include -I=./proto --httpgo_out=only=client,paths=source_relative,context=native:. proto/example2.proto
//go:generate protoc -I=. -I=./vendor -I=/usr/local/include -I=./proto --httpgo_out=paths=source_relative,autoURI=true,context=native:. proto/no_url.proto
//go:generate easyjson -all proto/example.pb.go

package main

import (
	_ "github.com/mailru/easyjson/gen"
	_ "google.golang.org/protobuf/encoding/protojson"
)
