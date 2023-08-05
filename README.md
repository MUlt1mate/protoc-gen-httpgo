# protoc-gen-httpgo

This is a protoc plugin that generates HTTP server and client code from proto files.

## Features

- Generation of both server and client code
- Provides multiple options for Marshaling/Unmarshaling:
    - Uses the native `encoding/json` by default
    - Optional usage of [easyjson](https://github.com/mailru/easyjson) for performance
- Utilizes google.api.http for defining HTTP paths
- Supports a wide range of data types in path parameters

## Usage

```bash  
protoc -I=. --httpgo_out=. --httpgo_opt=paths=source_relative example/proto/example.proto
```  

### Parameters

| Name       | Values                 | Description                                                                                                                                  |
|------------|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------|
| paths      | source_relative,import | Inherited from protogen, see [docs](https://protobuf.dev/reference/go/go-generated/#invocation) for more details                             |
| marshaller | easyjson               | Specifies the data marshaling/unmarshaling package. Uses `encoding/json` by default. Can be set to easyjson with fallback to `encoding/json` |
| only       | server,client          | Use to generate either the server or client code exclusively                                                                                 |

Example of usage:

```
--httpgo_opt=paths=source_relative,marshaller=easyjson,only=server
```

## TODO

- Implement comprehensive test cases
