package ttstream

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

var _ ClientStreamMeta = (*clientStream)(nil)

func newClientStream(ctx context.Context, writer streamWriter, smeta streamFrame) *clientStream {
	_ = "STUB: not implemented"
	return nil
}

type clientStream struct {
	*stream
	state            int32
	metaFrameHandler MetaFrameHandler

	closeStreamException atomic.Value
	storeExceptionOnce   sync.Once

	headerSig  chan int32
	trailerSig chan int32

	traceCtl *rpcinfo.TraceController

	streamTimeout time.Duration
}

func (s *clientStream) Header() (streaming.Header, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Header), nil
}

func (s *clientStream) Trailer() (streaming.Trailer, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Trailer), nil
}

func (s *clientStream) SendMsg(ctx context.Context, req any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStream) RecvMsg(ctx context.Context, resp any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStream) CloseSend(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *clientStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *clientStream) ctxDoneCallback(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStream) recvTimeoutCallback(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStream) parseCtxErr(ctx context.Context) (finalEx *Exception, noCancel bool, cancelPath string) {
	_ = "STUB: not implemented"
	return nil, false, ""
}

func (s *clientStream) close(exception error, sendRst bool, cancelPath string, trailer streaming.Trailer) {
	_ = "STUB: not implemented"
	return
}

func (s *clientStream) closeSignalMeta(trailer streaming.Trailer) {
	_ = "STUB: not implemented"
	return
}

func (s *clientStream) setMetaFrameHandler(metaHandler MetaFrameHandler) {
	_ = "STUB: not implemented"
	return
}

func (s *clientStream) setTraceController(traceCtl *rpcinfo.TraceController) {
	_ = "STUB: not implemented"
	return
}

func (s *clientStream) setStreamTimeout(tm time.Duration) { _ = "STUB: not implemented"; return }

func (s *clientStream) handleStreamStartEvent(event rpcinfo.StreamStartEvent) {
	_ = "STUB: not implemented"
	return
}

func (s *clientStream) handleStreamRecvHeaderEvent(event rpcinfo.StreamRecvHeaderEvent) {
	_ = "STUB: not implemented"
	return
}

func (s *clientStream) handleStreamFinishEvent(event rpcinfo.StreamFinishEvent) {
	_ = "STUB: not implemented"
	return
}

func (s *clientStream) onReadMetaFrame(fr *Frame) error { _ = "STUB: not implemented"; return nil }

func (s *clientStream) onReadHeaderFrame(fr *Frame) error { _ = "STUB: not implemented"; return nil }

func (s *clientStream) onReadTrailerFrame(fr *Frame) error { _ = "STUB: not implemented"; return nil }

func (s *clientStream) onReadRstFrame(fr *Frame) (err error) { _ = "STUB: not implemented"; return nil }

func (s *clientStream) cancelSignalMeta() { _ = "STUB: not implemented"; return }
