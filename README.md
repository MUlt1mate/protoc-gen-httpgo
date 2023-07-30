# protoc-gen-httpgo

protoc plugin that generates HTTP server and client

## Features

- server and client generation
- native go json Marshal/Unmarshal
- uses google.api.http for path
- supports many types in path parameters

## Usage

```bash
protoc -I=. --httpgo_out=. --httpgo_opt=paths=source_relative example/proto/example.proto
```

## TODO

- tests
