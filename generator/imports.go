package generator

import "google.golang.org/protobuf/compiler/protogen"

var (
	errorsPackage  = protogen.GoImportPath("errors")
	logPackage     = protogen.GoImportPath("log")
	contextPackage = protogen.GoImportPath("context")
	jsonPackage    = protogen.GoImportPath("encoding/json")
	fmtPackage     = protogen.GoImportPath("fmt")
	strconvPackage = protogen.GoImportPath("strconv")

	fasthttpPackage = protogen.GoImportPath("github.com/valyala/fasthttp")
	routerPackage   = protogen.GoImportPath("github.com/fasthttp/router")
)
