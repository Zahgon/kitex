package genericserver

import (
	"context"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
)

func RegisterUnknownServiceOrMethodHandler(svr server.Server, unknownHandler *UnknownServiceOrMethodHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUnknownServiceOrMethodServer(unknownHandler *UnknownServiceOrMethodHandler, options ...server.Option) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

type UnknownServiceOrMethodHandler struct {
	DefaultHandler func(ctx context.Context, service, method string, request interface{}) (response interface{}, err error)

	StreamingHandler func(ctx context.Context, service, method string, stream generic.BidiStreamingServer) (err error)
}
