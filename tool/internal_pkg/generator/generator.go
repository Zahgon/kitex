package generator

import (
	"time"

	"github.com/cloudwego/kitex/tool/internal_pkg/util"
)

const (
	KitexGenPath = "kitex_gen"
	DefaultCodec = "thrift"

	BuildFileName       = "build.sh"
	BootstrapFileName   = "bootstrap.sh"
	ToolVersionFileName = "kitex_info.yaml"
	HandlerFileName     = "handler.go"
	MainFileName        = "main.go"
	ClientFileName      = "client.go"
	ServerFileName      = "server.go"
	InvokerFileName     = "invoker.go"
	ServiceFileName     = "*service.go"
	ExtensionFilename   = "extensions.yaml"

	DefaultThriftPluginTimeLimit = time.Minute

	MultipleServicesTpl = "multiple_services"
)

var (
	kitexImportPath = "github.com/cloudwego/kitex"

	globalMiddlewares  []Middleware
	globalDependencies = map[string]string{
		"kitex":     kitexImportPath,
		"client":    ImportPathTo("client"),
		"server":    ImportPathTo("server"),
		"callopt":   ImportPathTo("client/callopt"),
		"frugal":    "github.com/cloudwego/frugal",
		"fieldmask": "github.com/cloudwego/thriftgo/fieldmask",
	}
)

func SetKitexImportPath(path string) { _ = "STUB: not implemented"; return }

func ImportPathTo(pkg string) string { _ = "STUB: not implemented"; return "" }

func AddGlobalMiddleware(mw Middleware) { _ = "STUB: not implemented"; return }

func AddGlobalDependency(ref, path string) bool { _ = "STUB: not implemented"; return false }

type Generator interface {
	GenerateService(pkg *PackageInfo) ([]*File, error)
	GenerateMainPackage(pkg *PackageInfo) ([]*File, error)
	GenerateCustomPackage(pkg *PackageInfo) ([]*File, error)
}

type Config struct {
	Verbose               bool
	GenerateMain          bool
	GenerateInvoker       bool
	Version               string
	NoFastAPI             bool
	ModuleName            string
	ServiceName           string
	Use                   string
	IDLType               string
	Includes              util.StringSlice
	ThriftOptions         util.StringSlice
	ProtobufOptions       util.StringSlice
	Hessian2Options       util.StringSlice
	IDL                   string
	OutputPath            string
	PackagePrefix         string
	CombineService        bool
	CopyIDL               bool
	ThriftPlugins         util.StringSlice
	ProtobufPlugins       util.StringSlice
	Features              []feature
	FrugalPretouch        bool
	ThriftPluginTimeLimit time.Duration
	CompilerPath          string

	ExtensionFile string
	tmplExt       *TemplateExtension

	Record    bool
	RecordCmd []string

	TemplateDir string

	GenPath string

	DeepCopyAPI           bool
	Protocol              string
	HandlerReturnKeepResp bool

	NoDependencyCheck bool
	Rapid             bool
	LocalThriftgo     bool

	FrugalStruct util.StringSlice
	NoRecurse    bool

	BuiltinTpl util.StringSlice

	StreamX bool
}

func (c *Config) Pack() (res []string) { _ = "STUB: not implemented"; return nil }

func (c *Config) Unpack(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *Config) AddFeature(key string) bool { _ = "STUB: not implemented"; return false }

func (c *Config) ApplyExtension() error { _ = "STUB: not implemented"; return nil }

func (c *Config) IsUsingMultipleServicesTpl() bool { _ = "STUB: not implemented"; return false }

func (c *Config) PkgOutputPath(pkg string) string { _ = "STUB: not implemented"; return "" }

func NewGenerator(config *Config, middlewares []Middleware) Generator {
	_ = "STUB: not implemented"
	return *new(Generator)
}

type Middleware func(HandleFunc) HandleFunc

type HandleFunc func(*Task, *PackageInfo) (*File, error)

type generator struct {
	*Config
	middlewares []Middleware
}

func (g *generator) chainMWs(handle HandleFunc) HandleFunc {
	_ = "STUB: not implemented"
	return *new(HandleFunc)
}

func (g *generator) GenerateMainPackage(pkg *PackageInfo) (fs []*File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *generator) generateHandler(pkg *PackageInfo, svc *ServiceInfo, handlerFileName string) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *generator) GenerateService(pkg *PackageInfo) ([]*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *generator) updatePackageInfo(pkg *PackageInfo) { _ = "STUB: not implemented"; return }

func (g *generator) setImports(name string, pkg *PackageInfo) { _ = "STUB: not implemented"; return }

func needCallOpt(pkg *PackageInfo) bool { _ = "STUB: not implemented"; return false }
