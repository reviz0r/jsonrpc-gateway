package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func (g *generator) fileName(file *descriptor.FileDescriptorProto) string {
	name := file.GetName()
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	output := fmt.Sprintf("%s.pb.jgw.go", base)
	return output
}
