package generic

import (
	"github.com/cloudwego/kitex/pkg/generic/thrift"

	dthrift "github.com/cloudwego/dynamicgo/thrift"
)

type thriftIDLProviderOptions struct {
	parseMode    *thrift.ParseMode
	goTag        *goTagOption
	serviceName  string
	dynamicGoOpt *dthrift.Options
}

type goTagOption struct {
	isGoTagAliasDisabled bool
}

type ThriftIDLProviderOption struct {
	F func(opt *thriftIDLProviderOptions)
}

func (o *thriftIDLProviderOptions) apply(opts []ThriftIDLProviderOption) {
	_ = "STUB: not implemented"
	return
}

func WithParseMode(parseMode thrift.ParseMode) ThriftIDLProviderOption {
	_ = "STUB: not implemented"
	return *new(ThriftIDLProviderOption)
}

func WithGoTagDisabled(disable bool) ThriftIDLProviderOption {
	_ = "STUB: not implemented"
	return *new(ThriftIDLProviderOption)
}

func WithIDLServiceName(serviceName string) ThriftIDLProviderOption {
	_ = "STUB: not implemented"
	return *new(ThriftIDLProviderOption)
}

func WithDynamicGoOptions(opts *dthrift.Options) ThriftIDLProviderOption {
	_ = "STUB: not implemented"
	return *new(ThriftIDLProviderOption)
}
