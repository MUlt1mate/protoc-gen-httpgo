//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --go_out=. --go_opt=paths=source_relative common/structs.proto example2.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,library=fasthttp:fasthttp example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,library=nethttp:nethttp example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,marshaller=protojson:. example2.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,autoURI=true:. no_url.proto

package proto

import (
	_ "github.com/fasthttp/router" // for vendoring before generation
	_ "google.golang.org/protobuf/encoding/protojson"
)
