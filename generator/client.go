package generator

import (
	"regexp"
	"strings"
)

var uriParametersRegexp = regexp.MustCompile(`(?mU){(.*)}`)

// GenerateClients generates HTTP clients for all services if the file
func (g *Generator) GenerateClients() (err error) {
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
func (g *Generator) genServiceClient(service serviceParams) (err error) {
	g.gf.P("var _  ", service.name, "HTTPGoService = & ", service.name, "HTTPGoClient{}")
	g.gf.P("")
	g.gf.P("type ", service.name, "HTTPGoClient struct {")
	g.gf.P("	cl          *", fasthttpPackage.Ident("Client"), "")
	g.gf.P("	host        string")
	g.gf.P("	middlewares []func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, ")")
	g.gf.P("	middleware  func(", g.clientInput, ", handler func(", g.clientInput, ") (", g.clientOutput, ")) (", g.clientOutput, ")")
	g.gf.P("}")
	g.gf.P("")
	g.gf.P("func Get", service.name, "HTTPGoClient(")
	g.gf.P("	_ ", contextPackage.Ident("Context"), ",")
	g.gf.P("	cl *", fasthttpPackage.Ident("Client"), ",")
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
func (g *Generator) genClientMethod(
	srvName string,
	method methodParams,
) (err error) {
	var (
		requestURI, paramsURI string
	)
	if requestURI, paramsURI, err = g.getRequestURIAndParams(method); err != nil {
		return err
	}
	g.gf.P("func (p * ", srvName, "HTTPGoClient) ", method.name, "(ctx ", contextPackage.Ident("Context"), ", request *", method.inputMsgName, ") (resp *", method.outputMsgName, ", err error) {")
	g.genMarshalRequestStruct()
	g.gf.P("	req := &fasthttp.Request{}")
	g.gf.P("	req.SetBody(body)")
	g.gf.P("	req.SetRequestURI(p.host + ", fmtPackage.Ident("Sprintf"), "(\""+requestURI+"\""+paramsURI+"))")
	g.gf.P("	req.Header.SetMethod(\"", method.httpMethodName, "\")")
	g.gf.P("	reqResp := &fasthttp.Response{}")
	g.gf.P("	var handler = func(", g.clientInput, ") (", g.clientOutput, ") {")
	g.gf.P("		resp = &fasthttp.Response{}")
	g.gf.P("		err = p.cl.Do(req, resp)")
	g.gf.P("		return resp, err")
	g.gf.P("	}")
	g.gf.P("	if p.middleware == nil {")
	g.gf.P("		if reqResp, err = handler(req); err != nil {")
	g.gf.P("			return nil, err")
	g.gf.P("		}")
	g.gf.P("	} else {")
	g.gf.P("		if reqResp, err = p.middleware(req, handler); err != nil {")
	g.gf.P("			return nil, err")
	g.gf.P("		}")
	g.gf.P("	}")
	g.gf.P("	resp = &", method.outputMsgName, "{}")
	g.genUnmarshalResponseStruct()
	g.gf.P("	return resp, err")
	g.gf.P("}")
	g.gf.P()
	return nil
}

// getRequestURIAndParams returns the request URI and parameters for the HTTP client method
func (g *Generator) getRequestURIAndParams(method methodParams) (requestURI, paramsURI string, err error) {
	requestURI = method.uri
	var placeholder string
	for _, match := range uriParametersRegexp.FindAllStringSubmatch(method.uri, -1) {
		if f, ok := method.fields[match[1]]; ok {
			if placeholder, err = f.getVariablePlaceholder(); err != nil {
				return "", "", err
			}
			requestURI = strings.ReplaceAll(requestURI, match[0], placeholder)
			paramsURI += ", request." + f.goName
		}
	}
	return requestURI, paramsURI, nil
}

// genMarshalRequestStruct generates marshalling from struct to []byte for request
func (g *Generator) genMarshalRequestStruct() {
	g.gf.P("	var body []byte")
	switch *g.cfg.Marshaller {
	case marshallerEasyJSON:
		g.gf.P("	if rqEJ, ok := interface{}(request).(", easyjsonPackage.Ident("Marshaler"), "); ok {")
		g.gf.P("		body, err = ", easyjsonPackage.Ident("Marshal"), "(rqEJ)")
		g.gf.P("	} else {")
		g.gf.P("		body, err = ", jsonPackage.Ident("Marshal"), "(request)")
		g.gf.P("	}")
	default:
		g.gf.P("	body, err = ", jsonPackage.Ident("Marshal"), "(request)")
	}
	g.gf.P("	if err != nil {")
	g.gf.P("		return nil, err")
	g.gf.P("	}")
}

// genUnmarshalResponseStruct generates unmarshalling from []byte to struct for response
func (g *Generator) genUnmarshalResponseStruct() {
	switch *g.cfg.Marshaller {
	case marshallerEasyJSON:
		g.gf.P("	if respEJ, ok := interface{}(resp).(", easyjsonPackage.Ident("Unmarshaler"), "); ok {")
		g.gf.P("		if err = ", easyjsonPackage.Ident("Unmarshal"), "(reqResp.Body(), respEJ); err != nil {")
		g.gf.P("			return nil, err")
		g.gf.P("		}")
		g.gf.P("	} else {")
		g.gf.P("		if err = ", jsonPackage.Ident("Unmarshal"), "(reqResp.Body(), resp); err != nil {")
		g.gf.P("			return nil, err")
		g.gf.P("		}")
		g.gf.P("	}")
	default:
		g.gf.P("	err = ", jsonPackage.Ident("Unmarshal"), "(reqResp.Body(), resp)")
	}
}

// genChainClientMiddlewares generates client middleware chain functions
func (g *Generator) genChainClientMiddlewares() {
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
	g.gf.P("			return middlewares[0](req, getChainClientMiddlewareHandler", g.filename, "(middlewares, 0, handler))")
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
	g.gf.P("		return middlewares[curr+1](req, getChainClientMiddlewareHandler", g.filename, "(middlewares, curr+1, finalHandler))")
	g.gf.P("	}")
	g.gf.P("}")
	g.gf.P()
}
