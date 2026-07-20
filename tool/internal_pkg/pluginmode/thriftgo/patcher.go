package thriftgo

import (
	"text/template"

	"github.com/cloudwego/thriftgo/generator/golang"
	"github.com/cloudwego/thriftgo/parser"
	"github.com/cloudwego/thriftgo/plugin"

	"github.com/cloudwego/kitex/tool/internal_pkg/util"
)

var extraTemplates []string

type PatchFunc func(req *plugin.Request, p *Patcher) ([]*plugin.Generated, error)

var extraPatchFunc PatchFunc

func AppendToTemplate(text string) { _ = "STUB: not implemented"; return }

func AppendExtensionPatches(f PatchFunc) { _ = "STUB: not implemented"; return }

const kitexUnusedProtection = `
// KitexUnusedProtection is used to prevent 'imported and not used' error.
var KitexUnusedProtection = struct{}{}
`

//lint:ignore U1000 until protectionInsertionPoint is used
var protectionInsertionPoint = "KitexUnusedProtection"

type Patcher struct {
	noFastAPI             bool
	utils                 *golang.CodeUtils
	module                string
	copyIDL               bool
	version               string
	record                bool
	recordCmd             []string
	deepCopyAPI           bool
	protocol              string
	handlerReturnKeepResp bool

	frugalStruct []string

	fileTpl *template.Template
	libs    map[string]string
}

func (p *Patcher) GetModule() string { _ = "STUB: not implemented"; return "" }

func (p *Patcher) GetUtils() *golang.CodeUtils { _ = "STUB: not implemented"; return nil }

func (p *Patcher) UseFrugalForStruct(st *golang.StructLike) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Patcher) UseLib(path, alias string) string { _ = "STUB: not implemented"; return "" }

func (p *Patcher) buildTemplates() (err error) { _ = "STUB: not implemented"; return nil }

const ImportInsertPoint = "// imports insert-point"

func (p *Patcher) patch(req *plugin.Request) (patches []*plugin.Generated, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Patcher) patchHessian(path string, scope *golang.Scope, pkgName, base string) (patch *plugin.Generated, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBashPath() string { _ = "STUB: not implemented"; return "" }

func (p *Patcher) extractLocalLibs(imports []util.Import) []util.Import {
	_ = "STUB: not implemented"
	return nil
}

func doRecord(recordCmd []string) string { _ = "STUB: not implemented"; return "" }

func (p *Patcher) reorderStructFields(fields []*golang.Field) ([]*golang.Field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Patcher) isBinaryOrStringType(t *parser.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Patcher) IsHessian2() bool { _ = "STUB: not implemented"; return false }

var typeIDToGoType = map[string]string{
	"Bool":   "bool",
	"Byte":   "int8",
	"I16":    "int16",
	"I32":    "int32",
	"I64":    "int64",
	"Double": "float64",
	"String": "string",
	"Binary": "[]byte",
}
