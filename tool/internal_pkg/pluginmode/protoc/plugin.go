package protoc

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/cloudwego/kitex/tool/internal_pkg/generator"
	"github.com/cloudwego/kitex/tool/internal_pkg/tpl/pbtpl"
)

type protocPlugin struct {
	generator.Config
	generator.PackageInfo
	Services    []*generator.ServiceInfo
	kg          generator.Generator
	err         error
	importPaths map[string]string
}

func (pp *protocPlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (pp *protocPlugin) init() {
	pp.Dependencies = map[string]string{
		"proto": "google.golang.org/protobuf/proto",
	}
}

func (pp *protocPlugin) parseM() { _ = "STUB: not implemented"; return }

func (pp *protocPlugin) GenerateFile(gen *protogen.Plugin, file *protogen.File) {
	_ = "STUB: not implemented"
	return
}

func (pp *protocPlugin) process(gen *protogen.Plugin) { _ = "STUB: not implemented"; return }

func (pp *protocPlugin) convertTypes(file *protogen.File) (ss []*generator.ServiceInfo) {
	_ = "STUB: not implemented"
	return nil
}

func (pp *protocPlugin) getCombineServiceName(name string, svcs []*generator.ServiceInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (pp *protocPlugin) convertParameter(msg *protogen.Message, paramName string) *generator.Parameter {
	_ = "STUB: not implemented"
	return nil
}

func (pp *protocPlugin) makeRenderArgs(gf *protogen.GeneratedFile, file *protogen.File, streamx bool) *pbtpl.Args {
	_ = "STUB: not implemented"
	return nil
}

func (pp *protocPlugin) adjustPath(path string) (ret string) { _ = "STUB: not implemented"; return "" }

func (pp *protocPlugin) fixImport(path string) string { _ = "STUB: not implemented"; return "" }
