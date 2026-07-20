package sdk

import (
	"errors"
	"flag"
	"os/exec"

	"github.com/cloudwego/thriftgo/plugin"

	kargs "github.com/cloudwego/kitex/tool/cmd/kitex/args"
)

var args kargs.Arguments

var errExitZero = errors.New("os.Exit(0)")

func init() {
	var queryVersion bool
	args.AddExtraFlag(&kargs.ExtraFlag{
		Apply: func(f *flag.FlagSet) {
			f.BoolVar(&queryVersion, "version", false,
				"Show the version of kitex")
		},
		Check: func(a *kargs.Arguments) error {
			if queryVersion {
				println(a.Version)
				return errExitZero
			}
			return nil
		},
	})
}

func RunKitexTool(wd string, plugins []plugin.SDKPlugin, kitexArgs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetKiteXSDKPlugin(pwd string, rawKiteXArgs []string) (*KiteXSDKPlugin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InvokeThriftgoBySDK(pwd string, cmd *exec.Cmd) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type KiteXSDKPlugin struct {
	KitexParams    []string
	ThriftgoParams []string
	Pwd            string
}

func (k *KiteXSDKPlugin) Invoke(req *plugin.Request) (res *plugin.Response) {
	_ = "STUB: not implemented"
	return nil
}

func (k *KiteXSDKPlugin) GetName() string { _ = "STUB: not implemented"; return "" }

func (k *KiteXSDKPlugin) GetPluginParameters() []string { _ = "STUB: not implemented"; return nil }

func (k *KiteXSDKPlugin) GetThriftgoParameters() []string { _ = "STUB: not implemented"; return nil }

func ParseKitexCmd(cmd *exec.Cmd) (thriftgoParams, kitexParams []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
