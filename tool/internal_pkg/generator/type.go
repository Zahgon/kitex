package generator

import (
	"strings"
	"text/template"

	"github.com/cloudwego/kitex/tool/internal_pkg/util"
	"github.com/cloudwego/kitex/transport"
)

type File struct {
	Name    string
	Content string
}

type PackageInfo struct {
	Namespace    string
	Dependencies map[string]string
	*ServiceInfo
	Services []*ServiceInfo

	Codec            string
	NoFastAPI        bool
	Version          string
	RealServiceName  string
	Imports          map[string]map[string]bool
	ExternalKitexGen string
	Features         []feature
	FrugalPretouch   bool
	Module           string
	Protocol         transport.Protocol
	IDLName          string
	ServerPkg        string
	StreamX          bool
}

func (p *PackageInfo) AddImport(pkg, path string) { _ = "STUB: not implemented"; return }

func (p *PackageInfo) AddImports(pkgs ...string) { _ = "STUB: not implemented"; return }

func (p *PackageInfo) UpdateImportPath(pkg, newPath string) { _ = "STUB: not implemented"; return }

func (p *PackageInfo) toExternalGenPath(path string) string { _ = "STUB: not implemented"; return "" }

type PkgInfo struct {
	PkgName    string
	PkgRefName string
	ImportPath string
}

type ServiceInfo struct {
	PkgInfo
	ServiceName           string
	RawServiceName        string
	ServiceTypeName       func() string
	Base                  *ServiceInfo
	Methods               []*MethodInfo
	CombineServices       []*ServiceInfo
	HasStreaming          bool
	ServiceFilePath       string
	Protocol              string
	HandlerReturnKeepResp bool
	UseThriftReflection   bool

	RefName string

	GenerateHandler bool
}

func (s *ServiceInfo) AllMethods() (ms []*MethodInfo) { _ = "STUB: not implemented"; return nil }

func (s *ServiceInfo) FixHasStreamingForExtendedService() { _ = "STUB: not implemented"; return }

func (s *ServiceInfo) HasStreamingRecursive() bool { _ = "STUB: not implemented"; return false }

type MethodInfo struct {
	PkgInfo                `json:"pkg_info"`
	ServiceName            string       `json:"service_name,omitempty"`
	Name                   string       `json:"name,omitempty"`
	RawName                string       `json:"raw_name,omitempty"`
	Oneway                 bool         `json:"oneway,omitempty"`
	Void                   bool         `json:"void,omitempty"`
	Args                   []*Parameter `json:"args,omitempty"`
	ArgsLength             int          `json:"args_length,omitempty"`
	Resp                   *Parameter   `json:"resp,omitempty"`
	Exceptions             []*Parameter `json:"exceptions,omitempty"`
	ArgStructName          string       `json:"arg_struct_name,omitempty"`
	ResStructName          string       `json:"res_struct_name,omitempty"`
	IsResponseNeedRedirect bool         `json:"is_response_need_redirect,omitempty"`
	GenArgResultStruct     bool         `json:"gen_arg_result_struct,omitempty"`
	IsStreaming            bool         `json:"is_streaming,omitempty"`
	ClientStreaming        bool         `json:"client_streaming,omitempty"`
	ServerStreaming        bool         `json:"server_streaming,omitempty"`
}

func (m *MethodInfo) StreamingMode() string { _ = "STUB: not implemented"; return "" }

type Parameter struct {
	Deps    []PkgInfo
	Name    string
	RawName string
	Type    string
}

var funcs = map[string]interface{}{
	"ToLower":       strings.ToLower,
	"LowerFirst":    util.LowerFirst,
	"UpperFirst":    util.UpperFirst,
	"NotPtr":        util.NotPtr,
	"ReplaceString": util.ReplaceString,
	"SnakeString":   util.SnakeString,
	"HasFeature":    HasFeature,
	"FilterImports": FilterImports,
	"backquoted":    BackQuoted,
}

func AddTemplateFunc(key string, f interface{}) { _ = "STUB: not implemented"; return }

var templateNames = []string{
	"@client.go-NewClient-option",
	"@client.go-NewStreamClient-option",
	"@client.go-EOF",
	"@server.go-NewServer-option",
	"@server.go-EOF",
	"@invoker.go-NewInvoker-option",
	"@invoker.go-EOF",
}

func wrapTemplate(point, content string) string { _ = "STUB: not implemented"; return "" }

var templateExtensions = (func() map[string]string {
	m := make(map[string]string)
	for _, name := range templateNames {

		m[name] = wrapTemplate(name, "")
	}
	return m
})()

func SetTemplateExtension(name, text string) { _ = "STUB: not implemented"; return }

func applyExtension(name string, x *template.Template) (*template.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Task struct {
	Name string
	Path string
	Text string
	*template.Template
	Ext *APIExtension
}

func (t *Task) Build() error { _ = "STUB: not implemented"; return nil }

func fileTemplateExtension(name string) (option, eof string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (t *Task) makeExtension() (res []string) { _ = "STUB: not implemented"; return nil }

func (t *Task) Render(data interface{}) (*File, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Task) RenderString(data interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func FilterImports(Imports map[string]map[string]bool, ms []*MethodInfo) map[string]map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func BackQuoted(s string) string { _ = "STUB: not implemented"; return "" }
