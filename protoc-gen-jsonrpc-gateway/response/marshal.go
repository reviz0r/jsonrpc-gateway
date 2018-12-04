package response

import (
	"io"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Marshal .
func Marshal(w io.Writer, out []*plugin.CodeGeneratorResponse_File, err error) {
	var response = new(plugin.CodeGeneratorResponse)

	if err != nil {
		response.Error = proto.String(err.Error())
	} else {
		response.File = out
	}

	buf, err := proto.Marshal(response)
	if err != nil {
		glog.Fatal(err)
	}

	if _, err := w.Write(buf); err != nil {
		glog.Fatal(err)
	}
}
