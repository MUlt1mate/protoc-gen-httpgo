package implementation

import (
	proto "github.com/MUlt1mate/protoc-gen-httpgo/example/proto/common"
)

var (
	dd              = "d d"
	hh              = []byte("h h")
	opt             = proto.Options_FIRST
	AllTextTypesMsg = proto.AllTextTypesMsg{
		String_:        "a a",
		RepeatedString: []string{"b b", "c c"},
		OptionalString: &dd,
		Bytes:          []byte("e e"),
		RepeatedBytes:  [][]byte{[]byte("f f"), []byte("g g")},
		OptionalBytes:  hh,
		Enum:           proto.Options_FIRST,
		RepeatedEnum:   []proto.Options{proto.Options_FIRST, proto.Options_SECOND},
		OptionalEnum:   &opt,
	}

	AllTypesMsg = proto.AllTypesMsg{
		SliceStringValue: []string{"a a", "b b"},
		BytesValue:       []byte("hello world"),
		StringValue:      "hello world",
	}

	MultipartFormRequestMsg = proto.MultipartFormRequest{
		Document: &proto.FileEx{
			File: []byte(`file content`),
			Name: "file.exe",
		},
		OtherField: "otherField",
	}

	MultipartFormRequestAllTypesMsg = proto.MultipartFormAllTypes{
		BoolValue:        true,
		EnumValue:        proto.Options_SECOND,
		Int32Value:       1,
		Sint32Value:      2,
		Uint32Value:      []uint32{3, 4},
		Int64Value:       5,
		Sint64Value:      Ptr(int64(6)),
		Uint64Value:      7,
		Sfixed32Value:    8,
		Fixed32Value:     []uint32{9, 10},
		FloatValue:       11.12,
		Sfixed64Value:    13,
		Fixed64Value:     Ptr(uint64(14)),
		DoubleValue:      15.16,
		StringValue:      "17 ",
		BytesValue:       []byte(" 18"),
		SliceStringValue: []string{"19 20", "21"},
		SliceInt32Value:  []int32{22, 23},
		Document: &proto.FileEx{
			File: []byte(`file content`),
			Name: "file.exe",
		},
		RepeatedStringValue: []string{"24"},
	}
)

func Ptr[T any](v T) *T {
	return &v
}
