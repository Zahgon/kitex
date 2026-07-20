package cep

import (
	"context"

	"github.com/cloudwego/kitex/pkg/streaming"
)

type StreamEndpoint func(ctx context.Context) (st streaming.ClientStream, err error)

type StreamMiddleware func(next StreamEndpoint) StreamEndpoint

type StreamMiddlewareBuilder func(ctx context.Context) StreamMiddleware

type StreamRecvEndpoint func(ctx context.Context, stream streaming.ClientStream, message interface{}) (err error)

func (e StreamRecvEndpoint) EqualsTo(e2 StreamRecvEndpoint) bool {
	_ = "STUB: not implemented"
	return false
}

type StreamRecvMiddleware func(next StreamRecvEndpoint) StreamRecvEndpoint

type StreamRecvMiddlewareBuilder func(ctx context.Context) StreamRecvMiddleware

type StreamSendEndpoint func(ctx context.Context, stream streaming.ClientStream, message interface{}) (err error)

func (e StreamSendEndpoint) EqualsTo(e2 StreamSendEndpoint) bool {
	_ = "STUB: not implemented"
	return false
}

type StreamSendMiddleware func(next StreamSendEndpoint) StreamSendEndpoint

type StreamSendMiddlewareBuilder func(ctx context.Context) StreamSendMiddleware

func StreamChain(mws ...StreamMiddleware) StreamMiddleware {
	_ = "STUB: not implemented"
	return *new(StreamMiddleware)
}

func StreamRecvChain(mws ...StreamRecvMiddleware) StreamRecvMiddleware {
	_ = "STUB: not implemented"
	return *new(StreamRecvMiddleware)
}

func StreamSendChain(mws ...StreamSendMiddleware) StreamSendMiddleware {
	_ = "STUB: not implemented"
	return *new(StreamSendMiddleware)
}

func DummyDummyMiddleware(next StreamEndpoint) StreamEndpoint {
	_ = "STUB: not implemented"
	return *new(StreamEndpoint)
}
