package generator

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// genFieldConvertor generates convertors from string to struct field
func (g *generator) genFieldConvertor(
	f field,
	source string,
	repeatedDestination bool,
	returnPrefix string,
	nakedReturn bool,
) error {
	var (
		destination = "arg." + f.goName
		reference   string
	)
	if f.optional {
		reference = "&"
	}

	switch f.kind {
	case protoreflect.StringKind:
		if repeatedDestination {
			g.gf.P(destination, " = append(", destination, ", ", source, ")")
		} else {
			g.gf.P(destination, " = ", reference, source)
		}
		return nil
	case protoreflect.BytesKind:
		if repeatedDestination {
			g.gf.P(destination, " = append(", destination, ", []byte(", source, "))")
		} else {
			g.gf.P(destination, " = []byte(", source, ")")
		}
		return nil
	case protoreflect.BoolKind:
		g.genBoolFieldConvertor(f, source, destination, repeatedDestination, returnPrefix, nakedReturn)
		return nil
	case protoreflect.EnumKind:
		g.genEnumFieldConverter(f, source, destination, repeatedDestination, returnPrefix, nakedReturn)
		return nil
	default:
		return g.genNumericFieldConvertor(f, source, destination, repeatedDestination, returnPrefix, nakedReturn)
	}
}

func (g *generator) genNumericFieldConvertor(
	f field,
	source string,
	destination string,
	repeatedDestination bool,
	returnPrefix string,
	nakedReturn bool,
) error {
	convertDestination := f.goName + ", convErr := "
	errorVar := "convErr"
	shortConversion := !fieldNeedToBeConverted(f) && !repeatedDestination && !f.optional
	if shortConversion {
		convertDestination = destination + ", err = "
		errorVar = "err"
	}
	switch f.kind {
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		g.gf.P(convertDestination, strconvPackage.Ident("ParseInt"), "(", source, ", 10, 64)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		g.gf.P(convertDestination, strconvPackage.Ident("ParseInt"), "(", source, ", 10, 32)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		g.gf.P(convertDestination, strconvPackage.Ident("ParseUint"), "(", source, ", 10, 64)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		g.gf.P(convertDestination, strconvPackage.Ident("ParseUint"), "(", source, ", 10, 32)")
	case protoreflect.DoubleKind:
		g.gf.P(convertDestination, strconvPackage.Ident("ParseFloat"), "(", source, ", 64)")
	case protoreflect.FloatKind:
		g.gf.P(convertDestination, strconvPackage.Ident("ParseFloat"), "(", source, ", 32)")
	default:
		return fmt.Errorf("unexpected type %s for numeric field ", f.kind)
	}
	g.gf.P("if ", errorVar, " != nil {")
	errValues := []interface{}{fmtPackage.Ident("Errorf"), "(\"conversion failed for parameter ", f.protoName, ": %w\", ", errorVar, ")"}
	if nakedReturn {
		g.gf.P(append([]interface{}{"	err = "}, errValues...)...)
		g.gf.P("	return")
	} else {
		g.gf.P(append([]interface{}{"	return ", returnPrefix}, errValues...)...)
	}
	g.gf.P("}")
	if shortConversion {
		return nil
	}
	var nextValues []interface{}
	if fieldNeedToBeConverted(f) {
		nextValues = []interface{}{getFieldConversionFuncName(f), "(", f.goName, ")"}
	} else {
		nextValues = []interface{}{f.goName}
	}
	if repeatedDestination {
		nextValues = append(nextValues, ")")
		g.gf.P(append([]interface{}{destination, " = append(", destination, ", "}, nextValues...)...)
	} else {
		var reference string
		if f.optional {
			reference = "&"
		}
		if f.optional && fieldNeedToBeConverted(f) {
			g.gf.P(append([]interface{}{f.goName, "Value := "}, nextValues...)...)
			nextValues = []interface{}{f.goName, "Value"}
		}
		g.gf.P(append([]interface{}{destination, " = ", reference}, nextValues...)...)
	}
	return nil
}

func (g *generator) genBoolFieldConvertor(
	f field,
	source string,
	destination string,
	repeatedDestination bool,
	returnPrefix string,
	nakedReturn bool,
) {
	g.gf.P("switch ", source, " {")
	g.gf.P("case \"true\", \"t\", \"1\":")
	if repeatedDestination {
		g.gf.P("	", destination, " = append(", destination, ", true)")
	} else {
		if !f.optional {
			g.gf.P("	", destination, " = true")
		} else {
			g.gf.P("	", f.goName, " := true")
			g.gf.P("	", destination, " = &", f.goName)
		}
	}
	g.gf.P("case \"false\", \"f\", \"0\":")
	if repeatedDestination {
		g.gf.P("	", destination, " = append(", destination, ", false)")
	} else {
		if !f.optional {
			g.gf.P("	", destination, " = false")
		} else {
			g.gf.P("	", f.goName, " := false")
			g.gf.P("	", destination, " = &", f.goName)
		}
	}
	g.gf.P("default:")
	errValues := []interface{}{fmtPackage.Ident("Errorf"), "(\"unknown bool string value %s\", ", source, ")"}
	if nakedReturn {
		g.gf.P(append([]interface{}{"	err = "}, errValues...)...)
		g.gf.P("	return")
	} else {
		g.gf.P(append([]interface{}{"	return ", returnPrefix}, errValues...)...)
	}
	g.gf.P("}")
}

func (g *generator) genEnumFieldConverter(
	f field,
	source string,
	destination string,
	repeatedDestination bool,
	returnPrefix string,
	nakedReturn bool,
) {
	g.gf.P("if ", f.enumName.GoName, "Value, optValueOk := ", f.enumName, "_value[", stringsPackage.Ident("ToUpper"), "(", source, ")]; optValueOk {")
	if repeatedDestination {
		g.gf.P("	", destination, " = append(", destination, ", ", f.enumName, "(", f.enumName.GoName, "Value))")
	} else {
		if !f.optional {
			g.gf.P("	", destination, " = ", f.enumName, "(", f.enumName.GoName, "Value)")
		} else {
			g.gf.P("	", f.goName, " := ", f.enumName, "(", f.enumName.GoName, "Value)")
			g.gf.P("	", destination, " = &", f.goName)
		}
	}
	g.gf.P("} else {")
	g.gf.P("	if intOptionValue, convErr := ", strconvPackage.Ident("ParseInt"), "(", source, ", 10, 32); convErr == nil {")
	g.gf.P("		if _, optIntValueOk := ", f.enumName, "_name[int32(intOptionValue)]; optIntValueOk {")
	if repeatedDestination {
		g.gf.P("			", destination, " = append(", destination, ", ", f.enumName, "(intOptionValue))")
	} else {
		if !f.optional {
			g.gf.P("			", destination, " = ", f.enumName, "(intOptionValue)")
		} else {
			g.gf.P("			", f.goName, " := ", f.enumName, "(intOptionValue)")
			g.gf.P("			", destination, " = &", f.goName)
		}
	}
	g.gf.P("		}")
	g.gf.P("	} else {")
	errValues := []interface{}{fmtPackage.Ident("Errorf"), "(\"conversion failed for parameter ", f.protoName, ": %w\", convErr)"}
	if nakedReturn {
		g.gf.P(append([]interface{}{"		err = "}, errValues...)...)
		g.gf.P("		return")
	} else {
		g.gf.P(append([]interface{}{"		return ", returnPrefix}, errValues...)...)
	}
	g.gf.P("	}")
	g.gf.P("}")
}

// fieldNeedToBeConverted check if we need to convert to base field type after converted to numeric value
func fieldNeedToBeConverted(f field) bool {
	switch f.kind {
	case protoreflect.Int64Kind, protoreflect.DoubleKind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return false
	default:
		return true
	}
}

// getFieldConversionFuncName we have to substitute some of the type names for go compiler
func getFieldConversionFuncName(f field) any {
	switch f.kind {
	case protoreflect.Fixed64Kind:
		return protoreflect.Uint64Kind.String()
	case protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.Int32Kind.String()
	case protoreflect.Fixed32Kind:
		return protoreflect.Uint32Kind.String()
	case protoreflect.Sfixed64Kind, protoreflect.Sint64Kind:
		return protoreflect.Int64Kind.String()
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.EnumKind:
		return f.enumName
	}

	return f.kind.String()
}

func (g *generator) convertFuncToString(f field, source string) (string, error) {
	switch f.kind {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind, protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind:
		return g.gf.QualifiedGoIdent(strconvPackage.Ident("FormatInt")) + "(int64(" + source + "), 10)", nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return g.gf.QualifiedGoIdent(strconvPackage.Ident("FormatInt")) + "(" + source + ", 10)", nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return g.gf.QualifiedGoIdent(strconvPackage.Ident("FormatUint")) + "(" + source + ", 10)", nil
	case protoreflect.FloatKind:
		return g.gf.QualifiedGoIdent(strconvPackage.Ident("FormatFloat")) + "(float64(" + source + "), 'f', -1, 64)", nil
	case protoreflect.DoubleKind:
		return g.gf.QualifiedGoIdent(strconvPackage.Ident("FormatFloat")) + "(" + source + ", 'f', -1, 64)", nil
	case protoreflect.BytesKind:
		return "string(" + source + ")", nil
	case protoreflect.BoolKind:
		return g.gf.QualifiedGoIdent(strconvPackage.Ident("FormatBool")) + "(" + source + ")", nil
	case protoreflect.EnumKind:
		return source + ".String()", nil
	default:
		return "", fmt.Errorf(`unsupported type %s for path variable: "%s"`, f.kind, f.goName)
	}
}
