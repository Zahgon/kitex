package prutal

import (
	"github.com/cloudwego/kitex/tool/internal_pkg/generator"

	"github.com/cloudwego/prutal/prutalgen/pkg/prutalgen"
)

type PrutalGen struct {
	g generator.Generator
	c *generator.Config

	pp []*generator.PackageInfo
}

func NewPrutalGen(c generator.Config) *PrutalGen { _ = "STUB: not implemented"; return nil }

func (pg *PrutalGen) initPackageInfo(f *prutalgen.Proto) *generator.PackageInfo {
	_ = "STUB: not implemented"
	return nil
}

func (pg *PrutalGen) generateClientServerFiles(f *prutalgen.Proto, p *generator.PackageInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (pg *PrutalGen) Process() error { _ = "STUB: not implemented"; return nil }

func (pg *PrutalGen) convertTypes(f *prutalgen.Proto) (ss []*generator.ServiceInfo) {
	_ = "STUB: not implemented"
	return nil
}

func getCombineServiceName(name string, svcs []*generator.ServiceInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (pg *PrutalGen) convertParameter(f *prutalgen.Proto, t *prutalgen.Type, paramName string) *generator.Parameter {
	_ = "STUB: not implemented"
	return nil
}

func genStructsAndKitexInterfaces(f *prutalgen.Proto, c *generator.Config, dir string) {
	_ = "STUB: not implemented"
	return
}

func genKitexServiceInterface(f *prutalgen.Proto, w *prutalgen.CodeWriter, streamx bool) {
	_ = "STUB: not implemented"
	return
}

func getImportPath(pkg, module, prefix string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func writeFile(fn string, data []byte) { _ = "STUB: not implemented"; return }
