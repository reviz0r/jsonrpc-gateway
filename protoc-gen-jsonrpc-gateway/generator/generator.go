package generator

import (
	"go/format"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Generator .
type Generator interface {
	Generate(request *plugin.CodeGeneratorRequest) ([]*plugin.CodeGeneratorResponse_File, error)
}

type generator struct{}

// New .
func New() Generator {
	return new(generator)
}

func (g *generator) Generate(request *plugin.CodeGeneratorRequest) ([]*plugin.CodeGeneratorResponse_File, error) {
	var files []*plugin.CodeGeneratorResponse_File

	for _, fileToGenerate := range request.FileToGenerate {
		var file *descriptor.FileDescriptorProto
		for _, f := range request.ProtoFile {
			if f.GetName() == fileToGenerate {
				file = f
			}
		}

		glog.V(1).Infof("Processing %s", file.GetName())

		code, err := g.generate(file)
		if err != nil {
			glog.Errorf("error while generate file: %v", err)
			return nil, err
		}

		if code == nil {
			continue
		}

		formattedCode, err := format.Source([]byte(*code))
		if err != nil {
			glog.Errorf("error while format file: %v", err)
			return nil, err
		}

		fileName := g.fileName(file)
		files = append(files, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fileName),
			Content: proto.String(string(formattedCode)),
		})

		glog.V(1).Infof("Will emit %s", fileName)
	}

	return files, nil
}
