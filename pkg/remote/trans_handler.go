package remote

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type ClientTransHandlerFactory interface {
	NewTransHandler(opt *ClientOption) (ClientTransHandler, error)
}

type ServerTransHandlerFactory interface {
	NewTransHandler(opt *ServerOption) (ServerTransHandler, error)
}

type TransReadWriter interface {
	Write(ctx context.Context, conn net.Conn, send Message) (nctx context.Context, err error)
	Read(ctx context.Context, conn net.Conn, msg Message) (nctx context.Context, err error)
}

type TransHandler interface {
	TransReadWriter
	OnInactive(ctx context.Context, conn net.Conn)
	OnError(ctx context.Context, err error, conn net.Conn)
	OnMessage(ctx context.Context, args, result Message) (context.Context, error)
	SetPipeline(pipeline *TransPipeline)
}

type ClientTransHandler interface {
	TransHandler
}

type ServerTransHandler interface {
	TransHandler
	OnActive(ctx context.Context, conn net.Conn) (context.Context, error)
	OnRead(ctx context.Context, conn net.Conn) error
}

type InvokeHandleFuncSetter interface {
	SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint)
}

type GracefulShutdown interface {
	GracefulShutdown(ctx context.Context) error
}

type ClientStreamFactory interface {
	NewStream(ctx context.Context, ri rpcinfo.RPCInfo) (streaming.ClientStream, error)
}
