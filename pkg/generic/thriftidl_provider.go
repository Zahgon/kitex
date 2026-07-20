package generic

import (
	"os"
	"sync"

	"github.com/cloudwego/dynamicgo/meta"
	dthrift "github.com/cloudwego/dynamicgo/thrift"
	"github.com/cloudwego/thriftgo/parser"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
	"github.com/cloudwego/kitex/pkg/generic/thrift"
)

var (
	_ Closer = &ThriftContentProvider{}
	_ Closer = &ThriftContentWithAbsIncludePathProvider{}

	isGoTagAliasDisabled = os.Getenv("KITEX_GENERIC_GOTAG_ALIAS_DISABLED") == "True"
	goTagMapper          = dthrift.FindAnnotationMapper("go.tag", dthrift.AnnoScopeField)
)

type thriftFileProvider struct {
	closeOnce sync.Once
	svcs      chan *descriptor.ServiceDescriptor
	opts      *ProviderOption
}

func NewThriftFileProvider(path string, includeDirs ...string) (DescriptorProvider, error) {
	_ = "STUB: not implemented"
	return *new(DescriptorProvider), nil
}

func NewThriftFileProviderWithOption(path string, opts []ThriftIDLProviderOption, includeDirs ...string) (DescriptorProvider, error) {
	_ = "STUB: not implemented"
	return *new(DescriptorProvider), nil
}

func NewThriftFileProviderWithDynamicGo(path string, includeDirs ...string) (DescriptorProvider, error) {
	_ = "STUB: not implemented"
	return *new(DescriptorProvider), nil
}

func NewThriftFileProviderWithDynamicgoWithOption(path string, opts []ThriftIDLProviderOption, includeDirs ...string) (DescriptorProvider, error) {
	_ = "STUB: not implemented"
	return *new(DescriptorProvider), nil
}

func newServiceDescriptorFromPath(path string, parseMode thrift.ParseMode, goTagOpt *goTagOption, serviceName string, includeDirs ...string) (*descriptor.ServiceDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *thriftFileProvider) Provide() <-chan *descriptor.ServiceDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (p *thriftFileProvider) Close() error { _ = "STUB: not implemented"; return nil }

func (p *thriftFileProvider) Option() ProviderOption {
	_ = "STUB: not implemented"
	return *new(ProviderOption)
}

type ThriftContentProvider struct {
	closeOnce   sync.Once
	svcs        chan *descriptor.ServiceDescriptor
	opts        *ProviderOption
	parseMode   thrift.ParseMode
	goTagOpt    *goTagOption
	serviceName string
}

var _ DescriptorProvider = (*ThriftContentProvider)(nil)

const defaultMainIDLPath = "main.thrift"

func NewThriftContentProvider(mainIDLContent string, includes map[string]string, opts ...ThriftIDLProviderOption) (*ThriftContentProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewThriftContentProviderWithDynamicGo(mainIDLContent string, includes map[string]string, opts ...ThriftIDLProviderOption) (*ThriftContentProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ThriftContentProvider) UpdateIDL(main string, includes map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *ThriftContentProvider) Provide() <-chan *descriptor.ServiceDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (p *ThriftContentProvider) Close() error { _ = "STUB: not implemented"; return nil }

func (p *ThriftContentProvider) Option() ProviderOption {
	_ = "STUB: not implemented"
	return *new(ProviderOption)
}

func (p *ThriftContentProvider) newDynamicGoDsc(svc *descriptor.ServiceDescriptor, path, content string, includes map[string]string, parseMode thrift.ParseMode, goTag *goTagOption, serviceName string) {
	_ = "STUB: not implemented"
	return
}

func parseIncludes(tree *parser.Thrift, parsed map[string]*parser.Thrift, sources map[string]string, isAbsIncludePath bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func absPath(path, includePath string) string { _ = "STUB: not implemented"; return "" }

func ParseContent(path, content string, includes map[string]string, isAbsIncludePath bool) (*parser.Thrift, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ThriftContentWithAbsIncludePathProvider struct {
	closeOnce   sync.Once
	svcs        chan *descriptor.ServiceDescriptor
	opts        *ProviderOption
	parseMode   thrift.ParseMode
	goTagOpt    *goTagOption
	serviceName string
}

var _ DescriptorProvider = (*ThriftContentWithAbsIncludePathProvider)(nil)

func NewThriftContentWithAbsIncludePathProvider(mainIDLPath string, includes map[string]string, opts ...ThriftIDLProviderOption) (*ThriftContentWithAbsIncludePathProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewThriftContentWithAbsIncludePathProviderWithDynamicGo(mainIDLPath string, includes map[string]string, opts ...ThriftIDLProviderOption) (*ThriftContentWithAbsIncludePathProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ThriftContentWithAbsIncludePathProvider) UpdateIDL(mainIDLPath string, includes map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *ThriftContentWithAbsIncludePathProvider) Provide() <-chan *descriptor.ServiceDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (p *ThriftContentWithAbsIncludePathProvider) Close() error {
	_ = "STUB: not implemented"
	return nil
}

func (p *ThriftContentWithAbsIncludePathProvider) Option() ProviderOption {
	_ = "STUB: not implemented"
	return *new(ProviderOption)
}

func (p *ThriftContentWithAbsIncludePathProvider) newDynamicGoDsc(svc *descriptor.ServiceDescriptor, path, content string, includes map[string]string, parseMode thrift.ParseMode, goTag *goTagOption, serviceName string) {
	_ = "STUB: not implemented"
	return
}

func getParseMode(opt *thriftIDLProviderOptions) thrift.ParseMode {
	_ = "STUB: not implemented"
	return *new(thrift.ParseMode)
}

func getDynamicGoParseMode(parseMode thrift.ParseMode) (meta.ParseServiceMode, error) {
	_ = "STUB: not implemented"
	return *new(meta.ParseServiceMode), nil
}

func newServiceDescriptorFromContent(path, content string, includes map[string]string, isAbsIncludePath bool, parseMode thrift.ParseMode, goTagOpt *goTagOption, serviceName string) (*descriptor.ServiceDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeDynamicgoOptions(dOpts *dthrift.Options, dParseMode meta.ParseServiceMode, serviceName string) (ret dthrift.Options) {
	_ = "STUB: not implemented"
	return *new(dthrift.Options)
}

func newDynamicGoDscFromContent(svc *descriptor.ServiceDescriptor, path, content string, includes map[string]string, isAbsIncludePath bool, parseMode thrift.ParseMode, goTag *goTagOption, serviceName string, dopts *dthrift.Options) error {
	_ = "STUB: not implemented"
	return nil
}

func handleGoTagForDynamicGo(goTagOpt *goTagOption) { _ = "STUB: not implemented"; return }
