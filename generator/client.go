package generator

import (
	"fmt"
	"regexp"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var uriParametersRegexp = regexp.MustCompile(`(?mU){(.*)}`)

// GenerateClients generates HTTP clients for all services in the file
func (g *generator) GenerateClients() (err error) {
	if *g.cfg.Only == onlyServer {
		return nil
	}
	for _, service := range g.services {
		if err = g.genServiceClient(service); err != nil {
			return err
		}
	}

	g.genChainClientMiddlewares()
	return nil
}

// genServiceClient generates HTTP client for serviceParams
func (g *generator) genServiceClient(service serviceParams) (err error) {
	g.gf.P("var _  ", service.name, "HTTPGoService = & ", service.name, "HTTPGoClient{}")
	g.gf.P("")
	g.gf.P("type ", service.name, "HTTPGoClient struct {")
	g.gf.P("	cl          *", g.lib.Ident("Client"), "")
	g.gf.P("	host        string")
	g.gf.P("	middlewares []func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, ")")
	g.gf.P("	middleware  func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, ")")
	g.gf.P("}")
	g.gf.P("")
	g.gf.P("func Get", service.name, "HTTPGoClient(")
	g.gf.P("	_ ", contextPackage.Ident("Context"), ",")
	g.gf.P("	cl *", g.lib.Ident("Client"), ",")
	g.gf.P("	host string,")
	g.gf.P("	middlewares []func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, "),")
	g.gf.P(") (*", service.name, "HTTPGoClient, error) {")
	g.gf.P("	return &", service.name, "HTTPGoClient{")
	g.gf.P("		cl:          cl,")
	g.gf.P("		host:        host,")
	g.gf.P("		middlewares: middlewares,")
	g.gf.P("		middleware:  chainClientMiddlewares", g.filename, "(middlewares),")
	g.gf.P("	}, nil")
	g.gf.P("}")
	g.gf.P()
	for _, method := range service.methods {
		if err = g.genClientMethod(service.name, method); err != nil {
			return err
		}
	}
	return nil
}

// genClientMethod generates method for HTTP client
func (g *generator) genClientMethod(
	srvName string,
	method methodParams,
) (err error) {
	if method.comment != "" {
		comment := method.comment
		if !strings.HasPrefix(comment, "// "+method.name) {
			comment = "// " + method.name + strings.TrimLeft(comment, "/")
		}
		g.gf.P(comment)
	}
	g.gf.P("func (p * ", srvName, "HTTPGoClient) ", method.name, "(ctx ", contextPackage.Ident("Context"), ", request *", method.inputMsgName, ") (resp *", method.outputMsgName, ", err error) {")
	switch *g.cfg.Library {
	case libraryNetHTTP:
		g.gf.P("	req := &", g.lib.Ident("Request"), "{Header: make(", g.lib.Ident("Header"), ")}")
	case libraryFastHTTP:
		g.gf.P("	req := &", g.lib.Ident("Request"), "{}")
	}
	g.gf.P("	var queryArgs string")
	if method.HasBody() {
		g.genMarshalRequestStruct()
	} else {
		if err = g.genQueryRequestParameters(method); err != nil {
			return err
		}
	}
	var (
		requestURI, paramsURI string
		params                []string
	)
	if requestURI, params, err = g.getRequestURIAndParams(method); err != nil {
		return err
	}
	if len(params) > 0 {
		paramsURI = "," + strings.Join(params, ", ")
	}
	switch *g.cfg.Library {
	case libraryNetHTTP:
		g.gf.P("	u, err := ", urlPackage.Ident("Parse"), "(", fmtPackage.Ident("Sprintf"), "(\"%s"+requestURI+"%s\",p.host"+paramsURI+",queryArgs))")
		g.gf.P("	if err != nil {")
		g.gf.P("		return nil, err")
		g.gf.P("	}")
		g.gf.P("	req.URL = u")
		g.gf.P("	req.Method = ", g.lib.Ident("Method"+titleString(method.httpMethodName)))
	case libraryFastHTTP:
		g.gf.P("	req.SetRequestURI(p.host + ", fmtPackage.Ident("Sprintf"), "(\""+requestURI+"%s\""+paramsURI+",queryArgs))")
		g.gf.P("	req.Header.SetMethod(\"", method.httpMethodName, "\")")
	}
	g.gf.P("	var reqResp interface{}")
	g.gf.P("	ctx = context.WithValue(ctx, \"proto_service\", \"" + srvName + "\")")
	g.gf.P("	ctx = context.WithValue(ctx, \"proto_method\", \"" + method.name + "\")")
	g.gf.P("	var handler = func(", g.clientInput, ") (", g.clientOutput, ") {")
	switch *g.cfg.Library {
	case libraryNetHTTP:
		g.gf.P("		resp, err = p.cl.Do(req.(*", g.lib.Ident("Request"), "))")
		g.gf.P("		return resp, err")
	case libraryFastHTTP:
		g.gf.P("		resp = &", g.lib.Ident("Response"), "{}")
		g.gf.P("		err = p.cl.Do(req.(*", g.lib.Ident("Request"), "), resp.(*", g.lib.Ident("Response"), "))")
		g.gf.P("		return resp, err")
	}
	g.gf.P("	}")
	g.gf.P("	if p.middleware == nil {")
	g.gf.P("		if reqResp, err = handler(ctx, req); err != nil {")
	g.gf.P("			return nil, err")
	g.gf.P("		}")
	g.gf.P("	} else {")
	g.gf.P("		if reqResp, err = p.middleware(ctx, req, handler); err != nil {")
	g.gf.P("			return nil, err")
	g.gf.P("		}")
	g.gf.P("	}")
	g.gf.P("	resp = &", method.outputMsgName, "{}")
	if err = g.genUnmarshalResponseStruct(method); err != nil {
		return err
	}
	g.gf.P("	return resp, err")
	g.gf.P("}")
	g.gf.P()
	return nil
}

// getRequestURIAndParams returns the request URI and parameters for the HTTP client method
func (g *generator) getRequestURIAndParams(method methodParams) (requestURI string, params []string, err error) {
	requestURI = method.uri
	var placeholder string
	for _, match := range uriParametersRegexp.FindAllStringSubmatch(method.uri, -1) {
		if f, ok := method.inputFields[match[1]]; ok {
			if placeholder, err = f.getVariablePlaceholder(); err != nil {
				return "", nil, err
			}
			parameterName := "request." + f.goName
			if f.cardinality == protoreflect.Repeated {
				parameterName = f.goName + "Request"
				placeholder = "%s"
				if err = g.genClientRepeatedFieldRequestValues(f); err != nil {
					return "", nil, err
				}
			}
			requestURI = strings.ReplaceAll(requestURI, match[0], placeholder)
			params = append(params, parameterName)
		}
	}
	return requestURI, params, nil
}

func (g *generator) genClientRepeatedFieldRequestValues(f field) (err error) {
	switch f.kind {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind, protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P(f.goName, "Strs[i] = ", strconvPackage.Ident("FormatInt"), "(int64(v), 10)")
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P(f.goName, "Strs[i] = ", strconvPackage.Ident("FormatInt"), "(v, 10)")
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P(f.goName, "Strs[i] = ", strconvPackage.Ident("FormatUint"), "(v, 10)")
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.FloatKind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P(f.goName, "Strs[i] = ", strconvPackage.Ident("FormatFloat"), "(float64(v), 'f', -1, 64)")
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.DoubleKind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P(f.goName, "Strs[i] = ", strconvPackage.Ident("FormatFloat"), "(v, 'f', -1, 64)")
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.StringKind:
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(request.", f.goName, ", \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.BytesKind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P(f.goName, "Strs[i] = string(v)")
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.BoolKind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P("if v {")
		g.gf.P(f.goName, "Strs[i] = \"true\"")
		g.gf.P("} else {")
		g.gf.P(f.goName, "Strs[i] = \"false\"")
		g.gf.P("}")
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	case protoreflect.EnumKind:
		g.gf.P(f.goName, "Strs := make([]string, len(request.", f.goName, "))")
		g.gf.P("for i, v := range request.", f.goName, " {")
		g.gf.P(f.goName, "Strs[i] = ", strconvPackage.Ident("FormatInt"), "(int64(v), 10)") // Assuming Enum is represented as int
		g.gf.P("}")
		g.gf.P(f.goName, "Request := ", stringsPackage.Ident("Join"), "(", f.goName, "Strs, \""+pathRepeatedArgDelimiter+"\")")
	default:
		err = fmt.Errorf(`unsupported type %s for path variable: "%s"`, f.kind, f.goName)
		return err
	}

	return nil
}

// genMarshalRequestStruct generates marshalling from struct to []byte for request
func (g *generator) genMarshalRequestStruct() {
	g.gf.P("	var body []byte")
	switch *g.cfg.Marshaller {
	case marshallerEasyJSON:
		g.gf.P("	if rqEJ, ok := interface{}(request).(", easyjsonPackage.Ident("Marshaler"), "); ok {")
		g.gf.P("		body, err = ", easyjsonPackage.Ident("Marshal"), "(rqEJ)")
		g.gf.P("	} else {")
		g.gf.P("		body, err = ", jsonPackage.Ident("Marshal"), "(request)")
		g.gf.P("	}")
	case marshallerProtoJSON:
		g.gf.P("	body, err = ", protojsonPackage.Ident("Marshal"), "(request)")
	default:
		g.gf.P("	body, err = ", jsonPackage.Ident("Marshal"), "(request)")
	}
	g.gf.P("	if err != nil {")
	g.gf.P("		return nil, err")
	g.gf.P("	}")
	switch *g.cfg.Library {
	case libraryNetHTTP:
		g.gf.P("	req.Body = ", ioPackage.Ident("NopCloser"), "(", bytesPackage.Ident("NewBuffer"), "(body))")
	case libraryFastHTTP:
		g.gf.P("	req.SetBody(body)")
	}
}

// genQueryRequestParameters
//
//nolint:prealloc // false positive
func (g *generator) genQueryRequestParameters(method methodParams) (err error) {
	pathParams := make(map[string]struct{})
	for _, match := range uriParametersRegexp.FindAllStringSubmatch(method.uri, -1) {
		pathParams[match[1]] = struct{}{}
	}
	if len(pathParams) == len(method.inputFieldList) {
		return nil
	}
	var (
		parameters, values []string
		placeholder        string
	)
	for _, f := range method.inputFieldList {
		if _, ok := pathParams[f]; ok {
			continue
		}
		methodField := method.inputFields[f]
		if methodField.cardinality == protoreflect.Repeated {
			continue
		}
		if placeholder, err = methodField.getVariablePlaceholder(); err != nil {
			return err
		}
		parameters = append(parameters, methodField.protoName+"="+placeholder)
		values = append(values, "request."+methodField.goName)
	}
	g.gf.P("var parameters = []string{")
	for _, q := range parameters {
		g.gf.P("\"", q, "\",")
	}
	g.gf.P("}")
	g.gf.P("var values = []interface{}{")
	for _, q := range values {
		g.gf.P(q, ",")
	}
	g.gf.P("}")
	for _, f := range method.inputFieldList {
		if _, ok := pathParams[f]; ok {
			continue
		}
		methodField := method.inputFields[f]
		if methodField.cardinality != protoreflect.Repeated {
			continue
		}
		if placeholder, err = methodField.getVariablePlaceholder(); err != nil {
			return err
		}
		uriName := methodField.protoName + "[]"
		g.gf.P("for _,v:= range request.", methodField.goName, " {")
		g.gf.P("	parameters = append(parameters, \"", uriName, "=", placeholder, "\")")
		g.gf.P("	values = append(values, v)")
		g.gf.P("}")
	}
	g.gf.P("queryArgs=", fmtPackage.Ident("Sprintf"), "(\"?\"+", stringsPackage.Ident("Join"), "(parameters, \"&\"),values...)")
	return nil
}

// genUnmarshalResponseStruct generates unmarshalling from []byte to struct for response
func (g *generator) genUnmarshalResponseStruct(method methodParams) error {
	respStruct := "resp"
	respStructPointer := respStruct
	switch *g.cfg.Library {
	case libraryNetHTTP:
		g.gf.P("	var respBody []byte")
		g.gf.P("	if respBody, err = ", ioPackage.Ident("ReadAll"), "(reqResp.(*", g.lib.Ident("Response"), ").Body); err != nil {")
		g.gf.P("		return nil, err")
		g.gf.P("	}")
		g.gf.P("	_ = reqResp.(*", g.lib.Ident("Response"), ").Body.Close()")
	case libraryFastHTTP:
		g.gf.P("	var respBody = reqResp.(*", g.lib.Ident("Response"), ").Body()")
	}

	if method.responseBody != "" {
		respField, ok := method.outputFields[method.responseBody]
		if !ok {
			return fmt.Errorf("field %s not found in struct %s for method %s", method.responseBody, method.outputMsgName.String(), method.name)
		}
		respStruct = "resp." + respField.goName
		respStructPointer = "&" + respStruct
	}
	switch *g.cfg.Marshaller {
	case marshallerEasyJSON:
		g.gf.P("	if respEJ, ok := interface{}(", respStruct, ").(", easyjsonPackage.Ident("Unmarshaler"), "); ok {")
		g.gf.P("		if err = ", easyjsonPackage.Ident("Unmarshal"), "(respBody, respEJ); err != nil {")
		g.gf.P("			return nil, err")
		g.gf.P("		}")
		g.gf.P("	} else {")
		g.gf.P("		if err = ", jsonPackage.Ident("Unmarshal"), "(respBody, ", respStructPointer, "); err != nil {")
		g.gf.P("			return nil, err")
		g.gf.P("		}")
		g.gf.P("	}")
	case marshallerProtoJSON:
		g.gf.P("	err = ", protojsonPackage.Ident("Unmarshal"), "(respBody, ", respStructPointer, ")")
	default:
		g.gf.P("	err = ", jsonPackage.Ident("Unmarshal"), "(respBody, ", respStructPointer, ")")
	}
	return nil
}

// genChainClientMiddlewares generates client middleware chain functions
func (g *generator) genChainClientMiddlewares() {
	g.gf.P("func chainClientMiddlewares", g.filename, "(")
	g.gf.P("	middlewares []func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, "),")
	g.gf.P(") func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, ") {")
	g.gf.P("	switch len(middlewares) {")
	g.gf.P("	case 0:")
	g.gf.P("		return nil")
	g.gf.P("	case 1:")
	g.gf.P("		return middlewares[0]")
	g.gf.P("	default:")
	g.gf.P("		return func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, ") {")
	g.gf.P("			return middlewares[0](ctx, req, getChainClientMiddlewareHandler", g.filename, "(middlewares, 0, handler))")
	g.gf.P("		}")
	g.gf.P("	}")
	g.gf.P("}")
	g.gf.P()
	g.gf.P("func getChainClientMiddlewareHandler", g.filename, "(")
	g.gf.P("	middlewares []func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, "),")
	g.gf.P("	curr int,")
	g.gf.P("	finalHandler func(", g.clientInput, ") (", g.clientOutput, "),")
	g.gf.P(") func(", g.clientInput, ") (", g.clientOutput, ") {")
	g.gf.P("	if curr == len(middlewares)-1 {")
	g.gf.P("		return finalHandler")
	g.gf.P("	}")
	g.gf.P("	return func(", g.clientInput, ") (", g.clientOutput, ") {")
	g.gf.P("		return middlewares[curr+1](ctx, req, getChainClientMiddlewareHandler", g.filename, "(middlewares, curr+1, finalHandler))")
	g.gf.P("	}")
	g.gf.P("}")
	g.gf.P()
}
