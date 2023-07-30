//go:generate protoc -I=. -I=./vendor --go_out=. --go_opt=paths=source_relative proto/somepackage/somepackage.proto
//go:generate protoc -I=. -I=./vendor -I=./proto/somepackage --go_out=. --go_opt=paths=source_relative proto/example.proto proto/example2.proto
//go:generate protoc -I=. -I=./vendor -I=./proto/somepackage --httpgo_out=. --httpgo_opt=paths=source_relative proto/example.proto proto/example2.proto

package main
