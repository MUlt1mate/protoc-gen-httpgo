//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --go_out=. --go_opt=paths=source_relative common/structs.proto example2.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,marshaller=easyjson,context=native,library=fasthttp:fasthttp example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,marshaller=easyjson,context=native,library=nethttp:nethttp example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=only=client,paths=source_relative,context=native:. example2.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,autoURI=true,context=native:. no_url.proto
//go:generate easyjson -all common/structs.pb.go

package proto

import (
	_ "github.com/fasthttp/router" // for vendoring before generation
	_ "github.com/mailru/easyjson/gen"
	_ "google.golang.org/protobuf/encoding/protojson"
)
