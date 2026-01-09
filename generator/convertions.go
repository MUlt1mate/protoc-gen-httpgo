package generator

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (g *generator) genFieldConvertor(
	f field,
	reference string,
	returnPrefix string,
	source string,
	destination string,
	repeatedDestination bool,
	nakedReturn bool,
) error {
	switch f.kind {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind, protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind:
		g.gf.P(f.goName, ", convErr := ", strconvPackage.Ident("ParseInt"), "(", source, ", 10, 32)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		g.gf.P(f.goName, ", convErr := ", strconvPackage.Ident("ParseInt"), "(", source, ", 10, 64)")
	case protoreflect.DoubleKind:
		g.gf.P(f.goName, ", convErr := ", strconvPackage.Ident("ParseFloat"), "(", source, ", 64)")
	case protoreflect.FloatKind:
		g.gf.P(f.goName, ", convErr := ", strconvPackage.Ident("ParseFloat"), "(", source, ", 32)")
	}
	g.gf.P("if convErr != nil {")
	errValues := []interface{}{fmtPackage.Ident("Errorf"), "(\"conversion failed for parameter ", f.protoName, ": %w\", convErr)"}
	if nakedReturn {
		g.gf.P(append([]interface{}{"	err = "}, errValues...)...)
		g.gf.P("	return")
	} else {
		g.gf.P(append([]interface{}{"	return ", returnPrefix}, errValues...)...)
	}
	g.gf.P("}")
	var nextValues []interface{}
	if f.needToBeConverted() {
		nextValues = []interface{}{f.getGolangTypeName(), "(", f.goName, ")"}
	} else {
		nextValues = []interface{}{f.goName}
	}
	if repeatedDestination {
		nextValues = append(nextValues, ")")
		g.gf.P(append([]interface{}{destination, " = append(", destination, ", "}, nextValues...)...)
	} else {
		if reference != "" && f.needToBeConverted() {
			g.gf.P(append([]interface{}{f.goName, "Value := "}, nextValues...)...)
			nextValues = []interface{}{f.goName, "Value"}
		}
		g.gf.P(append([]interface{}{destination, " = ", reference}, nextValues...)...)
	}
	return nil
}
