package generic

import (
	"github.com/cloudwego/dynamicgo/conv"

	"github.com/cloudwego/kitex/pkg/generic/thrift"
)

var (
	DefaultHTTPDynamicGoConvOpts = conv.Options{
		EnableHttpMapping:      true,
		EnableValueMapping:     true,
		WriteRequireField:      true,
		WriteDefaultField:      true,
		OmitHttpMappingErrors:  true,
		NoBase64Binary:         true,
		UseKitexHttpEncoding:   true,
		WriteHttpValueFallback: true,
		ReadHttpValueFallback:  true,
		MergeBaseFunc:          thrift.MergeBase,
	}
	DefaultJSONDynamicGoConvOpts = conv.Options{
		WriteRequireField:  true,
		WriteDefaultField:  true,
		EnableValueMapping: true,
		String2Int64:       true,
		MergeBaseFunc:      thrift.MergeBase,
	}
)

type Options struct {
	dynamicgoConvOpts conv.Options

	useRawBodyForHTTPResp bool
}

type Option struct {
	F func(opt *Options)
}

func (o *Options) apply(opts []Option) { _ = "STUB: not implemented"; return }

func WithCustomDynamicGoConvOpts(opts *conv.Options) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func UseRawBodyForHTTPResp(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }
