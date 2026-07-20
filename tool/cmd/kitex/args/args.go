package args

import (
	"flag"
	"io"
	"os/exec"

	"github.com/cloudwego/kitex/tool/internal_pkg/generator"
)

const EnvPluginMode = "KITEX_PLUGIN_MODE"

type ExtraFlag struct {
	Apply func(*flag.FlagSet)

	Check func(*Arguments) error
}

type Arguments struct {
	generator.Config
	extends []*ExtraFlag
}

const (
	Thrift   = "thrift"
	Protobuf = "protobuf"

	Unknown = "unknown"
)

const cmdExample = `  # Generate client codes or update kitex_gen codes when a project is in $GOPATH:
  kitex {{path/to/IDL_file.thrift}}

  # Generate client codes or update kitex_gen codes  when a project is not in $GOPATH:
  kitex -module {{github.com/xxx_org/xxx_name}} {{path/to/IDL_file.thrift}}

  # Generate server codes:
  kitex -service {{svc_name}} {{path/to/IDL_file.thrift}}
`

func (a *Arguments) AddExtraFlag(e *ExtraFlag) { _ = "STUB: not implemented"; return }

func (a *Arguments) buildFlags(version string) *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func (a *Arguments) ParseArgs(version, curpath string, kitexArgs []string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (a *Arguments) IsThrift() bool { _ = "STUB: not implemented"; return false }

func (a *Arguments) IsProtobuf() bool { _ = "STUB: not implemented"; return false }

func guessIDLType(idl string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (a *Arguments) checkIDL(files []string) error { _ = "STUB: not implemented"; return nil }

func (a *Arguments) checkServiceName() error { _ = "STUB: not implemented"; return nil }

func refGoSrcPath(curpath string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (a *Arguments) checkPath(curpath string) error { _ = "STUB: not implemented"; return nil }

func (a *Arguments) BuildCmd(out io.Writer) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateCMD(path, idlType string) error { _ = "STUB: not implemented"; return nil }

func LookupTool(idlType, compilerPath string) string { _ = "STUB: not implemented"; return "" }

func initGoMod(curpath, module string) error { _ = "STUB: not implemented"; return nil }
