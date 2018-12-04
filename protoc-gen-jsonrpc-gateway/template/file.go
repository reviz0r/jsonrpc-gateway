package template

import (
	"strings"
	"text/template"
)

var tmplFuncs = map[string]interface{}{
	"trimMsgName": func(in string) string {
		splitedIn := strings.Split(in, ".")
		if len(splitedIn) == 0 {
			return ""
		}

		return splitedIn[len(splitedIn)-1]
	},
}

// FileTmpl Шаблон генерируемого файла
var FileTmpl = template.Must(template.New("").Funcs(tmplFuncs).Parse(`
// Code generated by protoc-gen-jsonrpc-gateway. DO NOT EDIT.
// source: {{ .Name }}

/*
Package {{ .Package }} is a reverse proxy.

It translates gRPC into JSON-RPC 2.0
*/
package {{ .Package }}

import (
	"context"
	"io"

	"google.golang.org/grpc"

	"github.com/reviz0r/jsonrpc-gateway/jsonrpc"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

{{ range $service := .Service }}
{{ $serviceName := $service.GetName | printf "%sService" }}
type {{ $serviceName }} struct {
	Address string
	Opts    []grpc.DialOption
	TagToString   func(jsonrpc.MethodTag) string
	TagFromString func(string) jsonrpc.MethodTag
}

func (*{{ $serviceName }}) JsonrpcService() {}

func (s *{{ $serviceName }}) Methods() []string {
	methods := s.methods()

	var names []string
	for key := range methods {
		var tag string
		if s.TagToString != nil {
			tag = s.TagToString(key)
		} else {
			tag = jsonrpc.MethodTagToString(key)
		}
		names = append(names, tag)
	}

	return names
}

func (s *{{ $serviceName }}) Exec(ctx context.Context, method string, params io.Reader) (proto.Message, error) {
	methods := s.methods()

	var tag jsonrpc.MethodTag
	if s.TagFromString != nil {
		tag = s.TagFromString(method)
	} else {
		tag = jsonrpc.MethodTagFromString(method)
	}
	handler, exist := methods[tag]
	if !exist {
		return nil, jsonrpc.ErrMethodNotFound(nil)
	}

	return handler(ctx, params)
}

func (s *{{ $serviceName }}) methods() map[jsonrpc.MethodTag]jsonrpc.Method {
	return map[jsonrpc.MethodTag]jsonrpc.Method{
		{{ range $method := $service.GetMethod }}
		jsonrpc.MethodTag{Service: "{{ $service.GetName }}", Method: "{{ $method.GetName }}"}: func(ctx context.Context, params io.Reader) (proto.Message, error) {
			conn, err := grpc.Dial(s.Address, s.Opts...)
			if err != nil {
				return nil, err
			}
			defer conn.Close()

			client := New{{ $service.GetName }}Client(conn)
			in := new({{ trimMsgName $method.GetInputType }})
			err = jsonpb.Unmarshal(params, in)
			if err != nil {
				return nil, jsonrpc.ErrInvalidRequest(err.Error())
			}
			return client.{{ $method.GetName }}(ctx, in)
		},
		{{ end }}
	}
} {{ end }}
`))