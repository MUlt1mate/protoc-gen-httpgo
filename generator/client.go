package generator

import (
	"regexp"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var uriParametersRegexp = regexp.MustCompile(`(?mU){(.*)}`)

// GenerateClients generates HTTP clients for all services if the file
func GenerateClients(gen *protogen.Plugin, file *protogen.File) (err error) {
	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+".httpgo.cl.go", file.GoImportPath)
	g.P("// Code generated by protoc-gen-httpgo. DO NOT EDIT.")
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)
	for _, srv := range file.Services {
		if err = genServiceClient(g, srv); err != nil {
			return err
		}
	}
	return nil
}

// genServiceClient generates HTTP client for service
func genServiceClient(g *protogen.GeneratedFile, srv *protogen.Service) (err error) {
	g.P("var _  ", srv.GoName, "HTTPService = & ", srv.GoName, "Client{}")
	g.P("")
	g.P("type  ", srv.GoName, "Client struct {")
	g.P("	cl   *", fasthttpPackage.Ident("Client"), "")
	g.P("	host string")
	g.P("}")
	g.P("")
	g.P("func Get", srv.GoName, "Client(_ ", contextPackage.Ident("Context"), ", cl *", fasthttpPackage.Ident("Client"), ", host string) (* ", srv.GoName, "Client, error) {")
	g.P("	return & ", srv.GoName, "Client{")
	g.P("		cl:   cl,")
	g.P("		host: host,")
	g.P("	}, nil")
	g.P("}")
	g.P("")
	for _, method := range srv.Methods {
		if err = genClientMethod(g, srv.GoName, method); err != nil {
			return err
		}
	}
	return nil
}

// genClientMethod generates method for HTTP client
func genClientMethod(
	g *protogen.GeneratedFile,
	srvName string,
	method *protogen.Method,
) (err error) {
	var params methodParams
	if params, err = getRuleMethodAndURI(method); err != nil {
		return err
	}
	requestURI, paramsURI := getRequestURIAndParams(params.pattern, method)
	g.P("func (p * ", srvName, "Client) ", method.GoName, "(ctx ", contextPackage.Ident("Context"), ", request *", method.Input.GoIdent, ") (resp *", method.Output.GoIdent, ", err error) {")
	g.P("    body, _ := ", jsonPackage.Ident("Marshal"), "(request)")
	g.P("    req := &fasthttp.Request{}")
	g.P("    req.SetBody(body)")
	g.P("    req.SetRequestURI(p.host + ", fmtPackage.Ident("Sprintf"), "(\""+requestURI+"\""+paramsURI+"))")
	g.P("    req.Header.SetMethod(\"", params.serverMethod, "\")")
	g.P("    reqResp := &fasthttp.Response{}")
	g.P("    if err = p.cl.Do(req, reqResp); err != nil {")
	g.P("        return nil, err")
	g.P("    }")
	g.P("    resp = &", method.Output.GoIdent, "{}")
	g.P("    err = ", jsonPackage.Ident("Unmarshal"), "(reqResp.Body(), resp)")
	g.P("    return resp, err")
	g.P("}")
	return nil
}

// getRequestURIAndParams returns the request URI and parameters for the HTTP client method
func getRequestURIAndParams(pattern string, method *protogen.Method) (requestURI, paramsURI string) {
	requestURI = pattern
	for _, match := range uriParametersRegexp.FindAllStringSubmatch(pattern, -1) {
		for _, f := range method.Input.Fields {
			if f.GoName == strings.Title(match[1]) {
				requestURI = strings.Replace(requestURI, match[0], getVariablePlaceholder(f.Desc.Kind()), -1)
				paramsURI += ", request." + strings.Title(match[1])
			}
		}
	}
	return requestURI, paramsURI
}

func getVariablePlaceholder(parameterKind protoreflect.Kind) string {
	switch parameterKind {
	case protoreflect.StringKind:
		return "%s"
	default:
		return "%d"
	}
}