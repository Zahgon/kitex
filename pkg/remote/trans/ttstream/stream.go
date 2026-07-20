package ttstream

import (
	"context"

	"github.com/cloudwego/gopkg/protocol/thrift"

	"github.com/cloudwego/kitex/pkg/remote/trans/ttstream/internal/container"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	ktransport "github.com/cloudwego/kitex/transport"
)

var (
	_ streaming.ClientStream          = (*clientStream)(nil)
	_ streaming.ServerStream          = (*serverStream)(nil)
	_ streaming.CloseCallbackRegister = (*stream)(nil)
)

var defaultRstException = thrift.NewApplicationException(13, "rst")

func newBasicStream(ctx context.Context, writer streamWriter, smeta streamFrame) *stream {
	_ = "STUB: not implemented"
	return nil
}

type streamFrame struct {
	sid     int32
	method  string
	meta    IntHeader
	header  streaming.Header
	trailer streaming.Trailer
}

const (
	streamSigNone     int32 = 0
	streamSigActive   int32 = 1
	streamSigInactive int32 = -1
	streamSigCancel   int32 = -2
)

const (
	streamStateActive          int32 = 0
	streamStateHalfCloseLocal  int32 = 1
	streamStateHalfCloseRemote int32 = 2
	streamStateInactive        int32 = 3
)

type stream struct {
	streamFrame
	ctx      context.Context
	rpcInfo  rpcinfo.RPCInfo
	reader   *streamReader
	writer   streamWriter
	wheader  streaming.Header
	wtrailer streaming.Trailer

	recvTimeoutConfig   streaming.TimeoutConfig
	recvTimeoutCallback container.CtxDoneCallback
	closeCallback       []func(error)
}

func (s *stream) Service() string { _ = "STUB: not implemented"; return "" }

func (s *stream) fromService() string { _ = "STUB: not implemented"; return "" }

func (s *stream) Method() string { _ = "STUB: not implemented"; return "" }

func (s *stream) TransportProtocol() ktransport.Protocol {
	_ = "STUB: not implemented"
	return *new(ktransport.Protocol)
}

func (s *stream) SendMsg(ctx context.Context, msg any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) RecvMsg(ctx context.Context, data any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) RegisterCloseCallback(cb func(error)) { _ = "STUB: not implemented"; return }

func (s *stream) setRecvTimeoutConfig(cfg rpcinfo.RPCConfig, cb container.CtxDoneCallback) {
	_ = "STUB: not implemented"
	return
}

func (s *stream) runCloseCallback(exception error) { _ = "STUB: not implemented"; return }

func (s *stream) writeFrame(ftype int32, header streaming.Header, trailer streaming.Trailer, payload []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) sendTrailer(exception error) (err error) { _ = "STUB: not implemented"; return nil }

func (s *stream) sendRst(exception error, cancelPath string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) onReadDataFrame(fr *Frame) (err error) { _ = "STUB: not implemented"; return nil }
