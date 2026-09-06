package protoc

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const PluginName = "protoc-gen-kitex"

func Run() int { _ = "STUB: not implemented"; return 0 }

func DoRun(opts protogen.Options) error { _ = "STUB: not implemented"; return nil }

func GenKitex(req *pluginpb.CodeGeneratorRequest, opts protogen.Options) (*pluginpb.CodeGeneratorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
