package generator

import "google.golang.org/protobuf/compiler/protogen"

var (
	errorsPackage  = protogen.GoImportPath("errors")
	contextPackage = protogen.GoImportPath("context")
	jsonPackage    = protogen.GoImportPath("encoding/json")
	fmtPackage     = protogen.GoImportPath("fmt")
	strconvPackage = protogen.GoImportPath("strconv")
	stringsPackage = protogen.GoImportPath("strings")
	httpPackage    = protogen.GoImportPath("net/http")
	ioPackage      = protogen.GoImportPath("io")
	bytesPackage   = protogen.GoImportPath("bytes")
	urlPackage     = protogen.GoImportPath("net/url")

	fasthttpPackage  = protogen.GoImportPath("github.com/valyala/fasthttp")
	routerPackage    = protogen.GoImportPath("github.com/fasthttp/router")
	easyjsonPackage  = protogen.GoImportPath("github.com/mailru/easyjson")
	protojsonPackage = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
)
