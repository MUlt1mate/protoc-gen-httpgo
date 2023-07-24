package generator

import (
	"errors"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Generator struct {
	methods []methodParams
}

type methodParams struct {
	httpMethodName string
	pathPattern    string
}

func getRuleMethodAndURI(protoMethod *protogen.Method) (methodParams, error) {
	m := methodParams{}
	options, ok := protoMethod.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return m, errors.New("empty option")
	}

	httpRule, ok := proto.GetExtension(options, annotations.E_Http).(*annotations.HttpRule)
	if !ok {
		return m, errors.New("empty rule")
	}

	switch httpRule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		m = methodParams{
			httpMethodName: "GET",
			pathPattern:    httpRule.GetGet(),
		}
	case *annotations.HttpRule_Put:
		m = methodParams{
			httpMethodName: "PUT",
			pathPattern:    httpRule.GetPut(),
		}
	case *annotations.HttpRule_Post:
		m = methodParams{
			httpMethodName: "POST",
			pathPattern:    httpRule.GetPost(),
		}
	case *annotations.HttpRule_Delete:
		m = methodParams{
			httpMethodName: "DELETE",
			pathPattern:    httpRule.GetDelete(),
		}
	case *annotations.HttpRule_Patch:
		m = methodParams{
			httpMethodName: "PATCH",
			pathPattern:    httpRule.GetPatch(),
		}
	}
	return m, nil
}
