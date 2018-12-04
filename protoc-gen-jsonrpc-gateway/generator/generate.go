package generator

import (
	"bytes"

	"github.com/reviz0r/jsonrpc-gateway/protoc-gen-jsonrpc-gateway/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func (g *generator) generate(file *descriptor.FileDescriptorProto) (*string, error) {
	if len(file.GetService()) == 0 {
		return nil, nil
	}

	buf := bytes.NewBufferString("")
	err := template.FileTmpl.Execute(buf, file)
	if err != nil {
		return nil, err
	}

	out := buf.String()
	return &out, nil
}
