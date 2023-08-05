package generator

import (
	"regexp"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

var uriParametersRegexp = regexp.MustCompile(`(?mU){(.*)}`)

// GenerateClients generates HTTP clients for all services if the file
func (g *Generator) GenerateClients(gf *protogen.GeneratedFile) (err error) {
	if *g.cfg.Only == onlyServer {
		return nil
	}
	for _, service := range g.services {
		if err = g.genServiceClient(gf, service); err != nil {
			return err
		}
	}
	return nil
}

// genServiceClient generates HTTP client for serviceParams
func (g *Generator) genServiceClient(gf *protogen.GeneratedFile, service serviceParams) (err error) {
	gf.P("var _  ", service.name, "HTTPGoService = & ", service.name, "HTTPGoClient{}")
	gf.P("")
	gf.P("type  ", service.name, "HTTPGoClient struct {")
	gf.P("	cl   *", fasthttpPackage.Ident("Client"), "")
	gf.P("	host string")
	gf.P("}")
	gf.P()
	gf.P("func Get", service.name, "HTTPGoClient(_ ", contextPackage.Ident("Context"), ", cl *", fasthttpPackage.Ident("Client"), ", host string) (* ", service.name, "HTTPGoClient, error) {")
	gf.P("	return & ", service.name, "HTTPGoClient{")
	gf.P("		cl:   cl,")
	gf.P("		host: host,")
	gf.P("	}, nil")
	gf.P("}")
	gf.P()
	for _, method := range service.methods {
		if err = g.genClientMethod(gf, service.name, method); err != nil {
			return err
		}
	}
	return nil
}

// genClientMethod generates method for HTTP client
func (g *Generator) genClientMethod(
	gf *protogen.GeneratedFile,
	srvName string,
	method methodParams,
) (err error) {
	var (
		requestURI, paramsURI string
	)
	if requestURI, paramsURI, err = g.getRequestURIAndParams(method); err != nil {
		return err
	}
	gf.P("func (p * ", srvName, "HTTPGoClient) ", method.name, "(ctx ", contextPackage.Ident("Context"), ", request *", method.inputMsgName, ") (resp *", method.outputMsgName, ", err error) {")
	g.genMarshalRequestStruct(gf)
	gf.P("	req := &fasthttp.Request{}")
	gf.P("	req.SetBody(body)")
	gf.P("	req.SetRequestURI(p.host + ", fmtPackage.Ident("Sprintf"), "(\""+requestURI+"\""+paramsURI+"))")
	gf.P("	req.Header.SetMethod(\"", method.httpMethodName, "\")")
	gf.P("	reqResp := &fasthttp.Response{}")
	gf.P("	if err = p.cl.Do(req, reqResp); err != nil {")
	gf.P("		return nil, err")
	gf.P("	}")
	gf.P("	resp = &", method.outputMsgName, "{}")
	g.genUnmarshalResponseStruct(gf)
	gf.P("	return resp, err")
	gf.P("}")
	gf.P()
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
func (g *Generator) genMarshalRequestStruct(gf *protogen.GeneratedFile) {
	gf.P("	var body []byte")
	switch *g.cfg.Marshaller {
	case marshallerEasyJSON:
		gf.P("	if rqEJ, ok := interface{}(request).(", easyjsonPackage.Ident("Marshaler"), "); ok {")
		gf.P("		body, err = ", easyjsonPackage.Ident("Marshal"), "(rqEJ)")
		gf.P("	} else {")
		gf.P("		body, err = ", jsonPackage.Ident("Marshal"), "(request)")
		gf.P("	}")
	default:
		gf.P("	body, err = ", jsonPackage.Ident("Marshal"), "(request)")
	}
	gf.P("	if err != nil {")
	gf.P("		return")
	gf.P("	}")
}

// genUnmarshalResponseStruct generates unmarshalling from []byte to struct for response
func (g *Generator) genUnmarshalResponseStruct(gf *protogen.GeneratedFile) {
	switch *g.cfg.Marshaller {
	case marshallerEasyJSON:
		gf.P("	if respEJ, ok := interface{}(resp).(", easyjsonPackage.Ident("Unmarshaler"), "); ok {")
		gf.P("		if err = ", easyjsonPackage.Ident("Unmarshal"), "(reqResp.Body(), respEJ); err != nil {")
		gf.P("			return nil, err")
		gf.P("		}")
		gf.P("	} else {")
		gf.P("		if err = ", jsonPackage.Ident("Unmarshal"), "(reqResp.Body(), resp); err != nil {")
		gf.P("			return nil, err")
		gf.P("		}")
		gf.P("	}")
	default:
		gf.P("	err = ", jsonPackage.Ident("Unmarshal"), "(reqResp.Body(), resp)")
	}
}
