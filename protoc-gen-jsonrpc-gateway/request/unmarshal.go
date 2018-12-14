package request

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Unmarshal parses a protoc request to a proto Message
func Unmarshal(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read protoc request: %v", err)
	}
	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protoc request: %v", err)
	}
	return req, nil
}
