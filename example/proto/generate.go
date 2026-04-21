//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --go_out=. --go_opt=paths=source_relative common/structs.proto example2.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,library=fasthttp:fasthttp example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,library=nethttp:nethttp example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,library=gin,only=server:gin example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,library=fiber,only=server:fiber example.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,marshaller=protojson:. example2.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,autoURI=true:. no_url.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,stream=sse,library=fasthttp,only=server:fasthttp sse.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,stream=sse,library=nethttp,only=server:nethttp sse.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,stream=sse,library=gin,only=server:gin sse.proto
//go:generate protoc -I=. -I=../vendor -I=/usr/local/include --httpgo_out=paths=source_relative,stream=sse,library=fiber,only=server:fiber sse.proto

package proto

import (
	_ "github.com/fasthttp/router" // for vendoring before generation
	_ "google.golang.org/protobuf/encoding/protojson"
)
