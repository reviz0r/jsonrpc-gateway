package main

import (
	"os"

	"github.com/reviz0r/jsonrpc-gateway/protoc-gen-jsonrpc-gateway/generator"
	"github.com/reviz0r/jsonrpc-gateway/protoc-gen-jsonrpc-gateway/request"
	"github.com/reviz0r/jsonrpc-gateway/protoc-gen-jsonrpc-gateway/response"

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()

	in, err := request.Unmarshal(os.Stdin)
	if err != nil {
		response.Marshal(os.Stdout, nil, err)
		return
	}

	gen := generator.New()
	out, err := gen.Generate(in)
	if err != nil {
		response.Marshal(os.Stdout, nil, err)
		return
	}

	response.Marshal(os.Stdout, out, nil)
}
