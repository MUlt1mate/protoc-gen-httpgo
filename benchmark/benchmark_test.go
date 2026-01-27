package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/MUlt1mate/protoc-gen-httpgo/benchmark/proto"
)

// Benchmark for the HTTP handler
func BenchmarkSampleHandler(b *testing.B) {
	// Reset the timer to exclude setup time from benchmark
	b.ResetTimer()

	client := &http.Client{Transport: http.DefaultTransport}
	arg := proto.MeasureRequest{
		BoolValue:     true,
		Int32Value:    1,
		Sint32Value:   2,
		Uint32Value:   3,
		Int64Value:    4,
		Sint64Value:   5,
		Uint64Value:   6,
		Sfixed32Value: 7,
		Fixed32Value:  8,
		FloatValue:    9,
		Sfixed64Value: 10,
		Fixed64Value:  11,
		DoubleValue:   12,
		StringValue:   "test",
		BytesValue:    []byte("test"),
	}
	jsonData, err := json.Marshal(&arg)
	if err != nil {
		b.Fatal(err)
	}

	b.N = 5000
	url := "http://" + httpServerEndpoint + "/v1/measure"
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		if err != nil {
			b.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Body.Close()
	}
}
