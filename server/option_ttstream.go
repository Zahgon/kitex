package server

import (
	"github.com/cloudwego/kitex/pkg/remote/trans/ttstream"
)

func WithTTHeaderStreamingOptions(opts ...TTHeaderStreamingOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTTHeaderStreamingTransportOptions(opts ...ttstream.ServerHandlerOption) TTHeaderStreamingOption {
	_ = "STUB: not implemented"
	return *new(TTHeaderStreamingOption)
}
