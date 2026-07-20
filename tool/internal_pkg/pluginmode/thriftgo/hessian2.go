package thriftgo

import (
	"io"
	"regexp"

	"github.com/cloudwego/thriftgo/config"

	"github.com/cloudwego/kitex/tool/internal_pkg/generator"
)

const (
	JavaExtensionOption = "java_extension"

	DubboCodec        = "github.com/kitex-contrib/codec-dubbo"
	JavaThrift        = "java.thrift"
	JavaThriftAddress = "https://raw.githubusercontent.com/kitex-contrib/codec-dubbo/main/java/java.thrift"
)

func Hessian2PreHook(cfg *generator.Config) error { _ = "STUB: not implemented"; return nil }

func IsHessian2(a generator.Config) bool { _ = "STUB: not implemented"; return false }

func EnableJavaExtension(a generator.Config) bool { _ = "STUB: not implemented"; return false }

func runOption(cfg *generator.Config, opt string) error { _ = "STUB: not implemented"; return nil }

func runJavaExtensionOption(cfg *generator.Config) error { _ = "STUB: not implemented"; return nil }

func patchIDLRefConfig(cfg *generator.Config) error { _ = "STUB: not implemented"; return nil }

func loadIDLRefConfig(fileName string, reader io.Reader) (*config.RawConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	javaObjectRe                     = regexp.MustCompile(`\*java\.Object\b`)
	javaExceptionRe                  = regexp.MustCompile(`\*java\.Exception\b`)
	javaExceptionEmptyVerificationRe = regexp.MustCompile(`return p\.Exception != nil\b`)
)

func Hessian2PatchByReplace(args generator.Config, subDirPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func replaceJavaObject(content []byte) []byte { _ = "STUB: not implemented"; return nil }

func replaceJavaException(content []byte) []byte { _ = "STUB: not implemented"; return nil }

func replaceJavaExceptionEmptyVerification(content []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
