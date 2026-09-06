package client

import (
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint/cep"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

func WithStreamOptions(opts ...StreamOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStreamRecvTimeout(d time.Duration) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamRecvTimeoutConfig(cfg streaming.TimeoutConfig) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamMiddleware(mw cep.StreamMiddleware) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamMiddlewareBuilder(mwb cep.StreamMiddlewareBuilder) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamRecvMiddleware(mw cep.StreamRecvMiddleware) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamRecvMiddlewareBuilder(mwb cep.StreamRecvMiddlewareBuilder) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamSendMiddleware(mw cep.StreamSendMiddleware) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamSendMiddlewareBuilder(mwb cep.StreamSendMiddlewareBuilder) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

func WithStreamEventHandler(hdl rpcinfo.ClientStreamEventHandler) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}
