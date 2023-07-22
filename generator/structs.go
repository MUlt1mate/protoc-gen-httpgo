package generator

import (
	"errors"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type methodParams struct {
	serverMethod string
	pattern      string
}

func getRuleMethodAndURI(protoMethod *protogen.Method) (methodParams, error) {
	options, ok := protoMethod.Desc.Options().(*descriptorpb.MethodOptions)
	m := methodParams{}
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
			serverMethod: "GET",
			pattern:      httpRule.GetGet(),
		}
	case *annotations.HttpRule_Put:
		m = methodParams{
			serverMethod: "PUT",
			pattern:      httpRule.GetPut(),
		}
	case *annotations.HttpRule_Post:
		m = methodParams{
			serverMethod: "POST",
			pattern:      httpRule.GetPost(),
		}
	case *annotations.HttpRule_Delete:
		m = methodParams{
			serverMethod: "DELETE",
			pattern:      httpRule.GetDelete(),
		}
	case *annotations.HttpRule_Patch:
		m = methodParams{
			serverMethod: "PATCH",
			pattern:      httpRule.GetPatch(),
		}
	}
	return m, nil
}
