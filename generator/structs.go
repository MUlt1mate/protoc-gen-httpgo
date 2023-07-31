package generator

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Generator struct {
	services map[string][]methodParams
	filename string
}

func NewGenerator(file *protogen.File) Generator {
	g := Generator{
		services: make(map[string][]methodParams),
		filename: getFilename(file),
	}
	for _, srv := range file.Services {
		var methods []methodParams
		for _, protoMethod := range srv.Methods {
			// not supported
			if protoMethod.Desc.IsStreamingClient() || protoMethod.Desc.IsStreamingServer() {
				continue
			}
			method, err := getRuleMethodAndURI(protoMethod)
			if err != nil {
				// if there is an error, we can't use the method. skip it for now
				continue
			}

			method.name = protoMethod.GoName
			method.inputMsgName = protoMethod.Input.GoIdent
			method.outputMsgName = protoMethod.Output.GoIdent
			var fields = make(map[string]field)
			for _, protoField := range protoMethod.Input.Fields {
				f := field{
					goName:    protoField.GoName,
					protoName: protoField.Desc.JSONName(),
					kind:      protoField.Desc.Kind(),
				}
				if protoField.Desc.Kind() == protoreflect.EnumKind {
					f.enumName = protoField.Enum.GoIdent.GoName
				}
				fields[f.protoName] = f
			}

			method.fields = fields
			methods = append(methods, method)
		}
		if len(methods) != 0 {
			g.services[srv.GoName] = methods
		}
	}
	return g
}

// getFilename returns capitalized filename for generated method naming
func getFilename(file *protogen.File) string {
	fileName := file.GeneratedFilenamePrefix
	i := strings.LastIndex(fileName, "/")
	if i != -1 {
		fileName = fileName[i+1:]
	}

	fileName = strings.NewReplacer(".", "", "-", "", "_", "").Replace(fileName)

	return strings.ToUpper(fileName[:1]) + fileName[1:]
}

type methodParams struct {
	fields         map[string]field
	inputMsgName   protogen.GoIdent
	outputMsgName  protogen.GoIdent
	name           string
	httpMethodName string
	uri            string
}

type field struct {
	goName    string
	protoName string
	enumName  string
	kind      protoreflect.Kind
}

// getGolangTypeName we have to substitute some of the type names for go compiler
func (f field) getGolangTypeName() string {
	switch f.kind {
	case protoreflect.Fixed64Kind:
		return protoreflect.Uint64Kind.String()
	case protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.Int32Kind.String()
	case protoreflect.Fixed32Kind:
		return protoreflect.Uint32Kind.String()
	}

	return f.kind.String()
}

func (f field) getVariablePlaceholder() (string, error) {
	switch f.kind {
	case protoreflect.StringKind,
		protoreflect.EnumKind,
		protoreflect.BytesKind:
		return "%s", nil
	case protoreflect.Int32Kind,
		protoreflect.Sint32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sint64Kind,
		protoreflect.Uint64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Fixed64Kind:
		return "%d", nil
	case
		protoreflect.FloatKind,
		protoreflect.DoubleKind:
		return "%.0f", nil
	case protoreflect.BoolKind:
		return "%t", nil
	default:
		return "", fmt.Errorf("unsupported type %s for path variable", f.kind.String())
	}
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
			uri:            httpRule.GetGet(),
		}
	case *annotations.HttpRule_Put:
		m = methodParams{
			httpMethodName: "PUT",
			uri:            httpRule.GetPut(),
		}
	case *annotations.HttpRule_Post:
		m = methodParams{
			httpMethodName: "POST",
			uri:            httpRule.GetPost(),
		}
	case *annotations.HttpRule_Delete:
		m = methodParams{
			httpMethodName: "DELETE",
			uri:            httpRule.GetDelete(),
		}
	case *annotations.HttpRule_Patch:
		m = methodParams{
			httpMethodName: "PATCH",
			uri:            httpRule.GetPatch(),
		}
	default:
		return m, fmt.Errorf("unknown method type %T", httpRule.GetPattern())
	}
	return m, nil
}
