//go:generate protoc -I=. -I=./vendor --go_out=. --go_opt=paths=source_relative proto/example.proto
//go:generate protoc --httpgo_out=. --httpgo_opt=paths=source_relative proto/example.proto

package main
