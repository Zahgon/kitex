package remote

import (
	"context"
	"net"
)

type BoundHandler interface{}

type OutboundHandler interface {
	BoundHandler
	Write(ctx context.Context, conn net.Conn, send Message) (context.Context, error)
}

type InboundHandler interface {
	BoundHandler
	OnActive(ctx context.Context, conn net.Conn) (context.Context, error)
	OnInactive(ctx context.Context, conn net.Conn) context.Context
	OnRead(ctx context.Context, conn net.Conn) (context.Context, error)
	OnMessage(ctx context.Context, args, result Message) (context.Context, error)
}

type DuplexBoundHandler interface {
	OutboundHandler
	InboundHandler
}

type TransPipeline struct {
	netHdlr TransHandler

	inboundHdrls  []InboundHandler
	outboundHdrls []OutboundHandler
}

var (
	_ TransHandler       = &TransPipeline{}
	_ ServerTransHandler = &TransPipeline{}
)

func newTransPipeline() *TransPipeline { _ = "STUB: not implemented"; return nil }

func NewTransPipeline(netHdlr TransHandler) *TransPipeline { _ = "STUB: not implemented"; return nil }

func (p *TransPipeline) AddInboundHandler(hdlr InboundHandler) *TransPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *TransPipeline) AddOutboundHandler(hdlr OutboundHandler) *TransPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *TransPipeline) Write(ctx context.Context, conn net.Conn, sendMsg Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (p *TransPipeline) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (p *TransPipeline) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (p *TransPipeline) OnRead(ctx context.Context, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TransPipeline) Read(ctx context.Context, conn net.Conn, msg Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (p *TransPipeline) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (p *TransPipeline) OnMessage(ctx context.Context, args, result Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (p *TransPipeline) SetPipeline(transPipe *TransPipeline) { _ = "STUB: not implemented"; return }

func (p *TransPipeline) GracefulShutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
