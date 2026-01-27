## Compare performance between gRPC + grpc-gateway vs. httpgo

#### Host info

- Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
- 32 GB RAM

#### Test suite

- 5000 sequential POST requests with body
- default settings
- no middlewares
- fasthttp library for httpgo

#### How to run benchmark

- `make run`
- check output files

#### Results

|                   | httpgo       | grpc-gateway |
|-------------------|--------------|--------------|
| Time all          | 1.026s       | 2.271s       |
| Time/op           | 203982 ns/op | 453008 ns/op |
| Memory allocation | 3.5 MB       | 91.5 MB      |

##### Benchstat

```text
goos: linux
goarch: amd64
pkg: github.com/MUlt1mate/protoc-gen-httpgo/benchmark
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
                │ bench_grpcgateway.txt │          bench_httpgo.txt           │
                │        sec/op         │   sec/op     vs base                │
SampleHandler-8             368.4µ ± 7%   246.7µ ± 8%  -33.03% (p=0.000 n=10)


```

#### Why the difference?

grpc-gateway acts as a proxy that marshals/unmarshals every request twice (JSON → Struct → Protobuf) and communicates
over local network sockets.  
httpgo generates direct handlers, eliminating the proxy layer and reducing heap allocations
significantly.

