package thriftgo

import (
	"fmt"

	"github.com/cloudwego/thriftgo/generator/backend"
	"github.com/cloudwego/thriftgo/generator/golang"
	"github.com/cloudwego/thriftgo/parser"
	"github.com/cloudwego/thriftgo/plugin"

	"github.com/cloudwego/kitex/tool/internal_pkg/generator"
)

var (
	prelude  = map[string]bool{"client": true, "server": true, "callopt": true, "context": true, "thrift": true, "kitex": true}
	keyWords = []string{"client", "server", "callopt", "context", "thrift", "kitex"}
)

type converter struct {
	Warnings []string
	Utils    *golang.CodeUtils
	Config   generator.Config
	Package  generator.PackageInfo
	Services []*generator.ServiceInfo
	svc2ast  map[*generator.ServiceInfo]*parser.Thrift
}

func (c *converter) init(req *plugin.Request) error {
	if req.Language != "go" {
		return fmt.Errorf("expect language to be 'go'. Encountered '%s'", req.Language)
	}

	if err := c.Config.Unpack(req.PluginParameters); err != nil {
		return err
	}

	c.Utils = golang.NewCodeUtils(c.initLogs())
	c.Utils.HandleOptions(req.GeneratorParameters)

	return nil
}

func (c *converter) initLogs() backend.LogFunc {
	_ = "STUB: not implemented"
	return *new(backend.LogFunc)
}

func (c *converter) avoidIncludeConflict(ast *parser.Thrift, ref string) (*parser.Thrift, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func (c *converter) copyTreeWithRef(ast *parser.Thrift, ref string) *parser.Thrift {
	_ = "STUB: not implemented"
	return nil
}

func (c *converter) copyFunctionWithRef(f *parser.Function, ref string) *parser.Function {
	_ = "STUB: not implemented"
	return nil
}

func (c *converter) copyTypeWithRef(t *parser.Type, ref string) (res *parser.Type) {
	_ = "STUB: not implemented"
	return nil
}

func (c *converter) getImports(t *parser.Type) (res []generator.PkgInfo) {
	_ = "STUB: not implemented"
	return nil
}

func (c *converter) fixImportConflicts() { _ = "STUB: not implemented"; return }

type ast2svc map[string][]*generator.ServiceInfo

func (t ast2svc) findService(ast *parser.Thrift, name string) *generator.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func (c *converter) convertTypes(req *plugin.Request) error { _ = "STUB: not implemented"; return nil }

func (c *converter) fixStreamingForExtendedServices(ast *parser.Thrift, all ast2svc) {
	_ = "STUB: not implemented"
	return
}

func (c *converter) makeService(pkg generator.PkgInfo, svc *golang.Service) (*generator.ServiceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *converter) makeMethod(si *generator.ServiceInfo, f *golang.Function) (*generator.MethodInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *converter) persist(res *plugin.Response) error { _ = "STUB: not implemented"; return nil }

func (c *converter) getCombineServiceName(name string, svcs []*generator.ServiceInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *converter) IsHessian2() bool { _ = "STUB: not implemented"; return false }

func (c *converter) copyAnnotations(annotations parser.Annotations) parser.Annotations {
	_ = "STUB: not implemented"
	return *new(parser.Annotations)
}
