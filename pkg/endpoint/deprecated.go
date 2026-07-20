package endpoint

import (
	"context"

	"github.com/cloudwego/kitex/pkg/streaming"
)

type RecvEndpoint func(stream streaming.Stream, message interface{}) (err error)

type RecvMiddleware func(next RecvEndpoint) RecvEndpoint

type RecvMiddlewareBuilder func(ctx context.Context) RecvMiddleware

func RecvChain(mws ...RecvMiddleware) RecvMiddleware {
	_ = "STUB: not implemented"
	return *new(RecvMiddleware)
}

type SendEndpoint func(stream streaming.Stream, message interface{}) (err error)

type SendMiddleware func(next SendEndpoint) SendEndpoint

type SendMiddlewareBuilder func(ctx context.Context) SendMiddleware

func SendChain(mws ...SendMiddleware) SendMiddleware {
	_ = "STUB: not implemented"
	return *new(SendMiddleware)
}
