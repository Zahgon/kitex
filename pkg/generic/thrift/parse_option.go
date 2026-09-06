package thrift

import "github.com/cloudwego/kitex/pkg/generic/descriptor"

type parseOptions struct {
	goTag       *descriptor.GoTagOption
	serviceName string
}

type ParseOption struct {
	F func(opt *parseOptions)
}

func (o *parseOptions) apply(opts []ParseOption) { _ = "STUB: not implemented"; return }

func WithGoTagDisabled(disable bool) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithIDLServiceName(serviceName string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}
