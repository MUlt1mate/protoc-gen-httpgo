// This is a protoc plugin that generates HTTP server and client code from proto files.
// It is supposed to be installed as a binary file and invoked by protoc when set as an argument
//
// Example:
// protoc -I=. --httpgo_out=paths=source_relative,context=native:. example/proto/example.proto
//
// For using reference check README file
// https://github.com/MUlt1mate/protoc-gen-httpgo/blob/main/README.md
// For development reference check generator package
package main
