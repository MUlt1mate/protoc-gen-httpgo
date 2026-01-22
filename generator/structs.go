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
	marshallerProtoJSON      = "protojson"
	onlyServer               = "server"
	onlyClient               = "client"
	pathRepeatedArgDelimiter = ","
	libraryNetHTTP           = "nethttp"
	libraryFastHTTP          = "fasthttp"
	contentTypeJSONApp       = "application/json"
	contentTypeMultipart     = "multipart/form-data"
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
		marshaller      protogen.GoImportPath
	}

	serviceParams struct {
		name    string
		methods []methodParams
	}

	methodParams struct {
		rule           *annotations.HttpRule
		inputFields    map[string]field
		outputFields   map[string]field
		inputMsgName   protogen.GoIdent
		outputMsgName  protogen.GoIdent
		name           string
		httpMethodName string
		uri            methodURI
		comment        string
		responseBody   string
		inputFieldList []string // slice for constant sorting
		hasBody        bool
		withFiles      bool
	}

	methodURI struct {
		protoURI string
		argList  []string
		args     map[string]methodURIArg
	}

	methodURIArg struct {
		PathTpl        string // how argument represented in path
		DestinationTpl string // patter for fmt.Sprintf
	}

	field struct {
		goName          string // name in go generated files
		protoName       string // name in proto file and http requests
		jsonName        string // name in json tag and http requests
		structTypeIdent protogen.GoIdent
		kind            protoreflect.Kind
		cardinality     protoreflect.Cardinality
		optional        bool
		isFile          bool
		fileStruct      struct {
			Name string
			Path string
		}
		structFields []*protogen.Field
	}
	Config struct {
		Marshaller         *string
		Only               *string
		AutoURI            *bool
		BodylessMethodsStr *string
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
	switch *cfg.Marshaller {
	case marshallerProtoJSON:
		g.marshaller = protojsonPackage
	default:
		g.marshaller = jsonPackage
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
	var (
		fields = make(map[string]field)
	)
	for _, protoField := range protoMethod.Input.Fields {
		f := convertField(protoField)
		if f.isFile {
			method.withFiles = true
		}
		fields[f.protoName] = f
		method.inputFieldList = append(method.inputFieldList, f.protoName)
	}
	method.inputFields = fields
	fields = make(map[string]field)
	for _, protoField := range protoMethod.Output.Fields {
		f := field{
			goName:      protoField.GoName,
			protoName:   protoField.Desc.TextName(),
			jsonName:    protoField.Desc.JSONName(),
			kind:        protoField.Desc.Kind(),
			cardinality: protoField.Desc.Cardinality(),
			optional:    protoField.Desc.HasOptionalKeyword(),
		}
		if protoField.Desc.Kind() == protoreflect.EnumKind {
			f.structTypeIdent = protoField.Enum.GoIdent
		}
		fields[f.protoName] = f
	}
	method.outputFields = fields
	method.comment = strings.TrimSuffix(protoMethod.Comments.Leading.String(), "\n")
}

func convertField(protoField *protogen.Field) field {
	f := field{
		goName:      protoField.GoName,
		protoName:   protoField.Desc.TextName(),
		jsonName:    protoField.Desc.JSONName(),
		kind:        protoField.Desc.Kind(),
		cardinality: protoField.Desc.Cardinality(),
		optional:    protoField.Desc.HasOptionalKeyword(),
		isFile:      isFileField(protoField),
	}
	if f.isFile {
		f.fileStruct.Path = string(protoField.Message.GoIdent.GoImportPath)
		f.fileStruct.Name = protoField.Message.GoIdent.GoName
	}
	if protoField.Desc.Kind() == protoreflect.EnumKind {
		f.structTypeIdent = protoField.Enum.GoIdent
	}
	if protoField.Desc.Kind() == protoreflect.MessageKind {
		f.structFields = protoField.Message.Fields
		f.structTypeIdent = protoField.Message.GoIdent
	}
	return f
}

// initTemplates fill predefined templates
func (g *generator) initTemplates(gf *protogen.GeneratedFile) {
	g.serverInput = "ctx " + gf.QualifiedGoIdent(contextPackage.Ident("Context")) + ", req any"
	g.serverOutput = "resp any, err error"
	g.clientInput = "ctx " + gf.QualifiedGoIdent(contextPackage.Ident("Context")) + ", req *" + gf.QualifiedGoIdent(g.lib.Ident("Request"))
	g.clientOutput = "resp *" + gf.QualifiedGoIdent(g.lib.Ident("Response")) + ", err error"
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

func getVariablePlaceholder(kind protoreflect.Kind) (string, error) {
	switch kind {
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
		return "", fmt.Errorf(`unsupported type %s for path variable`, kind)
	}
}

func (f field) fileStructIdent() protogen.GoIdent {
	return protogen.GoImportPath(f.fileStruct.Path).Ident(f.fileStruct.Name)
}

func (g *generator) getRuleMethodAndURI(protoMethod *protogen.Method, serviceName string) (methodParams, error) {
	m := methodParams{}
	options, ok := protoMethod.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return m, errors.New("empty option")
	}

	httpRule, _ := proto.GetExtension(options, annotations.E_Http).(*annotations.HttpRule)
	if httpRule == nil && !*g.cfg.AutoURI {
		return m, errors.New("empty rule")
	}

	if *g.cfg.AutoURI {
		return methodParams{
			httpMethodName: "POST",
			uri:            methodURI{protoURI: "/" + serviceName + "/" + protoMethod.GoName},
			hasBody:        true,
		}, nil
	}

	m.rule = httpRule
	m.httpMethodName, m.uri.protoURI = getRuleMethodAndURI(httpRule)
	m.uri.parseURI(*g.cfg.Library)
	m.hasBody = g.MethodShouldHasBody(m.httpMethodName)
	m.responseBody = httpRule.GetResponseBody()
	return m, nil
}

func getRuleMethodAndURI(httpRule *annotations.HttpRule) (string, string) {
	switch httpRule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		return "GET", httpRule.GetGet()
	case *annotations.HttpRule_Put:
		return "PUT", httpRule.GetPut()
	case *annotations.HttpRule_Post:
		return "POST", httpRule.GetPost()
	case *annotations.HttpRule_Delete:
		return "DELETE", httpRule.GetDelete()
	case *annotations.HttpRule_Patch:
		return "PATCH", httpRule.GetPatch()
	case *annotations.HttpRule_Custom:
		return httpRule.GetCustom().Kind, httpRule.GetCustom().Path
	default:
		// doesn't return error here, because it won't compile with other value anyway
		return "", ""
	}
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

func (m methodParams) GetContentType() string {
	if m.withFiles {
		return contentTypeMultipart
	}
	return contentTypeJSONApp
}

func isFileField(field *protogen.Field) bool {
	if field.Desc.Kind() != protoreflect.MessageKind {
		return false
	}

	msg := field.Message
	if msg == nil {
		return false
	}

	var (
		matchedFields int
		totalFields   = 3
	)

	for _, nf := range msg.Fields {
		switch nf.Desc.JSONName() {
		case "file":
			if nf.Desc.Kind() == protoreflect.BytesKind {
				matchedFields++
			}
		case "name":
			if nf.Desc.Kind() == protoreflect.StringKind {
				matchedFields++
			}
		case "headers":
			matchedFields++
		}
	}

	return matchedFields == totalFields
}

func (m *methodURI) parseURI(library string) {
	m.args = make(map[string]methodURIArg)
	var path string
	for _, match := range uriParametersRegexp.FindAllStringSubmatch(m.protoURI, -1) {
		fieldName := match[1]
		arg := methodURIArg{}
		if i := strings.Index(fieldName, "="); i != -1 {
			// split variable name and segments name=messages/*
			pattern := fieldName[i+1:]
			fieldName = fieldName[:i]
			// and change uri to messages/{name}
			if strings.Contains(pattern, "**") {
				arg.PathTpl = "{" + fieldName + ":*}"
				if library == libraryNetHTTP {
					arg.PathTpl = "{" + fieldName + "...}"
				}
				arg.DestinationTpl = strings.Replace(pattern, "**", "%s", 1)
				path = strings.Replace(pattern, "**", arg.PathTpl, 1)
			} else {
				arg.PathTpl = "{" + fieldName + "}"
				arg.DestinationTpl = strings.Replace(pattern, "*", "%s", 1)
				path = strings.Replace(pattern, "*", arg.PathTpl, 1)
			}
			m.protoURI = strings.Replace(m.protoURI, match[0], path, 1)
		}
		m.argList = append(m.argList, fieldName)
		m.args[fieldName] = arg
	}
}

func (m methodParams) Copy() methodParams {
	cp := m
	if m.rule != nil {
		ruleCopy := *m.rule //nolint:govet // we are copying mutex, but it's not used
		cp.rule = &ruleCopy
	}

	if m.inputFields != nil {
		cp.inputFields = make(map[string]field, len(m.inputFields))
		for k, v := range m.inputFields {
			cp.inputFields[k] = v
		}
	}

	if m.outputFields != nil {
		cp.outputFields = make(map[string]field, len(m.outputFields))
		for k, v := range m.outputFields {
			cp.outputFields[k] = v
		}
	}

	if m.inputFieldList != nil {
		cp.inputFieldList = make([]string, len(m.inputFieldList))
		copy(cp.inputFieldList, m.inputFieldList)
	}

	return cp
}
