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

const (
	marshallerEasyJSON       = "easyjson"
	marshallerProtoJSON      = "protojson"
	onlyServer               = "server"
	onlyClient               = "client"
	pathRepeatedArgDelimiter = ","
	libraryNetHTTP           = "net/http"
	libraryFastHTTP          = "fasthttp"
)

type (
	generator struct {
		gf              *protogen.GeneratedFile
		cfg             Config
		services        []serviceParams
		bodylessMethods map[string]struct{}
		filename        string
		clientInput     string
		clientOutput    string
		serverInput     string
		serverOutput    string
		lib             protogen.GoImportPath
	}

	serviceParams struct {
		name    string
		methods []methodParams
	}

	methodParams struct {
		inputFields    map[string]field
		outputFields   map[string]field
		inputMsgName   protogen.GoIdent
		outputMsgName  protogen.GoIdent
		name           string
		httpMethodName string
		uri            string
		comment        string
		responseBody   string
		inputFieldList []string // slice for constant sorting
		hasBody        bool
	}

	field struct {
		goName      string // name in go generated files
		protoName   string // name in proto file and http requests
		enumName    string
		kind        protoreflect.Kind
		cardinality protoreflect.Cardinality
		optional    bool
	}
	Config struct {
		Marshaller         *string
		Only               *string
		AutoURI            *bool
		BodylessMethodsStr *string
		ContextStruct      *string
		Library            *string
	}
)

func newGenerator(
	file *protogen.File,
	cfg Config,
	gf *protogen.GeneratedFile,
) (generator, error) {
	var bodylessMethods = make(map[string]struct{})
	if cfg.BodylessMethodsStr == nil || *cfg.BodylessMethodsStr == "" {
		bodylessMethods = map[string]struct{}{"GET": {}, "DELETE": {}}
	} else {
		list := strings.Split(*cfg.BodylessMethodsStr, ";")
		for _, l := range list {
			bodylessMethods[strings.TrimSpace(l)] = struct{}{}
		}
	}

	g := generator{
		filename:        getFilename(file),
		cfg:             cfg,
		gf:              gf,
		bodylessMethods: bodylessMethods,
	}
	switch *cfg.Library {
	case libraryNetHTTP:
		g.lib = httpPackage
	case libraryFastHTTP:
		g.lib = fasthttpPackage
	default:
		return g, errors.New("unsupported library type: " + *cfg.Library)
	}
	g.initTemplates(gf)
	g.fillServices(file)
	return g, nil
}

// Run is the main function that start generation with given parameters
func Run(gen *protogen.Plugin, cfg Config) (err error) {
	var g generator
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		if g, err = newGenerator(
			f,
			cfg,
			gen.NewGeneratedFile(f.GeneratedFilenamePrefix+".httpgo.go", f.GoImportPath),
		); err != nil {
			return err
		}
		if err = g.GenerateServers(f); err != nil {
			return err
		}
		if err = g.GenerateClients(); err != nil {
			return err
		}
	}
	return nil
}

// fillServices scans services and methods from file for further generation
func (g *generator) fillServices(file *protogen.File) {
	for _, srv := range file.Services {
		var methods []methodParams
		for _, protoMethod := range srv.Methods {
			// not supported
			if protoMethod.Desc.IsStreamingClient() || protoMethod.Desc.IsStreamingServer() {
				continue
			}
			method, err := g.getRuleMethodAndURI(protoMethod, srv.GoName)
			if err != nil {
				// if there is an error, we can't use the method. skip it for now
				continue
			}

			fillMethod(&method, protoMethod)
			methods = append(methods, method)
		}
		if len(methods) != 0 {
			g.services = append(g.services, serviceParams{name: srv.GoName, methods: methods})
		}
	}
}

func fillMethod(method *methodParams, protoMethod *protogen.Method) {
	method.name = protoMethod.GoName
	method.inputMsgName = protoMethod.Input.GoIdent
	method.outputMsgName = protoMethod.Output.GoIdent
	var fields = make(map[string]field)
	for _, protoField := range protoMethod.Input.Fields {
		f := field{
			goName:      protoField.GoName,
			protoName:   protoField.Desc.JSONName(),
			kind:        protoField.Desc.Kind(),
			cardinality: protoField.Desc.Cardinality(),
			optional:    protoField.Desc.HasOptionalKeyword(),
		}
		if protoField.Desc.Kind() == protoreflect.EnumKind {
			f.enumName = protoField.Enum.GoIdent.GoName
		}
		fields[f.protoName] = f
		method.inputFieldList = append(method.inputFieldList, f.protoName)
	}
	method.inputFields = fields
	fields = make(map[string]field)
	for _, protoField := range protoMethod.Output.Fields {
		f := field{
			goName:      protoField.GoName,
			protoName:   protoField.Desc.JSONName(),
			kind:        protoField.Desc.Kind(),
			cardinality: protoField.Desc.Cardinality(),
			optional:    protoField.Desc.HasOptionalKeyword(),
		}
		if protoField.Desc.Kind() == protoreflect.EnumKind {
			f.enumName = protoField.Enum.GoIdent.GoName
		}
		fields[f.protoName] = f
	}
	method.outputFields = fields
	method.comment = strings.TrimSuffix(protoMethod.Comments.Leading.String(), "\n")
}

// initTemplates fill predefined templates
// we have to convert to strings here, because we can't pass other types like slices to protogen.P()
func (g *generator) initTemplates(gf *protogen.GeneratedFile) {
	if g.cfg.ContextStruct != nil && *g.cfg.ContextStruct == "native" {
		g.serverInput = "ctx " + gf.QualifiedGoIdent(contextPackage.Ident("Context")) + ", req interface{}"
		g.clientInput = "ctx " + gf.QualifiedGoIdent(contextPackage.Ident("Context")) + ", req interface{}"
		g.clientOutput = "resp interface{}, err error"
	} else {
		// this is default behavior for backward compatibility
		g.serverInput = "ctx *" + gf.QualifiedGoIdent(fasthttpPackage.Ident("RequestCtx")) + ", req interface{}"
		g.clientInput = "ctx " + gf.QualifiedGoIdent(contextPackage.Ident("Context")) + ", req *" + gf.QualifiedGoIdent(fasthttpPackage.Ident("Request"))
		g.clientOutput = "resp *" + gf.QualifiedGoIdent(fasthttpPackage.Ident("Response")) + ", err error"
	}

	g.serverOutput = "resp interface{}, err error"
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
		return "%f", nil
	case protoreflect.BoolKind:
		return "%t", nil
	default:
		return "", fmt.Errorf(`unsupported type %s for path variable: "%s"`, f.kind, f.goName)
	}
}

func (g *generator) getRuleMethodAndURI(protoMethod *protogen.Method, serviceName string) (methodParams, error) {
	m := methodParams{}
	options, ok := protoMethod.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return m, errors.New("empty option")
	}

	httpRule, ok := proto.GetExtension(options, annotations.E_Http).(*annotations.HttpRule)
	if !ok && !*g.cfg.AutoURI {
		return m, errors.New("empty rule")
	}

	if *g.cfg.AutoURI {
		return methodParams{
			httpMethodName: "POST",
			uri:            serviceName + "/" + protoMethod.GoName,
			hasBody:        true,
		}, nil
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
	m.hasBody = g.MethodShouldHasBody(m.httpMethodName)
	m.responseBody = httpRule.GetResponseBody()
	return m, nil
}

// HasBody checks if method may have a body
func (m methodParams) HasBody() bool {
	return m.hasBody
}

// MethodShouldHasBody checks if method may have a body
func (g *generator) MethodShouldHasBody(method string) bool {
	_, ok := g.bodylessMethods[method]
	return !ok
}

func titleString(input string) string {
	input = strings.ToLower(input)
	return strings.ToUpper(string(input[0])) + input[1:]
}
