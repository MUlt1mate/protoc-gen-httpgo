package generator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_methodURI_parseURI(t *testing.T) {
	tests := []struct {
		protoURI         string
		expectedArgList  []string
		expectedArgs     map[string]methodURIArg
		expectedProtoURI string
	}{
		{
			protoURI:         "/v1/{name}",
			expectedArgList:  []string{"name"},
			expectedArgs:     map[string]methodURIArg{"name": {}},
			expectedProtoURI: "/v1/{name}",
		},
		{
			protoURI:         "/v1/{name=messages/*}",
			expectedArgList:  []string{"name"},
			expectedArgs:     map[string]methodURIArg{"name": {PathTpl: "{name}", DestinationTpl: "messages/%s"}},
			expectedProtoURI: "/v1/messages/{name}",
		},
		{
			protoURI:         "{var=*}",
			expectedArgList:  []string{"var"},
			expectedArgs:     map[string]methodURIArg{"var": {PathTpl: "{var}", DestinationTpl: "%s"}},
			expectedProtoURI: "{var}",
		},
		{
			protoURI:         "/v1/{name=seg1/*/seg3}",
			expectedArgList:  []string{"name"},
			expectedArgs:     map[string]methodURIArg{"name": {PathTpl: "{name}", DestinationTpl: "seg1/%s/seg3"}},
			expectedProtoURI: "/v1/seg1/{name}/seg3",
		},
		{
			protoURI:         "/v1/{name=**}",
			expectedArgList:  []string{"name"},
			expectedArgs:     map[string]methodURIArg{"name": {PathTpl: "{name...}", DestinationTpl: "%s"}},
			expectedProtoURI: "/v1/{name...}",
		},
		{
			protoURI:         "/v1/{name=fixed/**}",
			expectedArgList:  []string{"name"},
			expectedArgs:     map[string]methodURIArg{"name": {PathTpl: "{name...}", DestinationTpl: "fixed/%s"}},
			expectedProtoURI: "/v1/fixed/{name...}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.protoURI, func(t *testing.T) {
			m := &methodURI{protoURI: tt.protoURI}
			m.parseURI(libraryNetHTTP)
			if m.protoURI != tt.expectedProtoURI {
				t.Errorf("got protoURI = %v, want %v", m.protoURI, tt.expectedProtoURI)
			}
			if diff := cmp.Diff(tt.expectedArgList, m.argList); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(tt.expectedArgs, m.args); diff != "" {
				t.Error(diff)
			}
		})
	}
}
