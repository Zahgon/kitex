package server

import (
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/sep"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func WithStreamOptions(opts ...StreamOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStreamMiddleware(mw sep.StreamMiddleware) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamMiddlewareBuilder(mwb sep.StreamMiddlewareBuilder) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamRecvMiddleware(mw sep.StreamRecvMiddleware) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamRecvMiddlewareBuilder(mwb sep.StreamRecvMiddlewareBuilder) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamSendMiddleware(mw sep.StreamSendMiddleware) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamSendMiddlewareBuilder(mwb sep.StreamSendMiddlewareBuilder) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamEventHandler(hdl rpcinfo.ServerStreamEventHandler) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithRecvMiddleware(mw endpoint.RecvMiddleware) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRecvMiddlewareBuilder(mwb endpoint.RecvMiddlewareBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSendMiddleware(mw endpoint.SendMiddleware) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSendMiddlewareBuilder(mwb endpoint.SendMiddlewareBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCompatibleMiddlewareForUnary() Option { _ = "STUB: not implemented"; return *new(Option) }
