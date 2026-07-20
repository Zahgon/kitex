package ttstream

import (
	"context"

	"github.com/cloudwego/kitex/pkg/streaming"
)

var _ ServerStreamMeta = (*serverStream)(nil)

func newServerStream(ctx context.Context, writer streamWriter, smeta streamFrame) *serverStream {
	_ = "STUB: not implemented"
	return nil
}

type serverStream struct {
	*stream
	state      int32
	cancelFunc cancelWithReason
}

func (s *serverStream) SetHeader(hd streaming.Header) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SendHeader(hd streaming.Header) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) writeHeader(hd streaming.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) sendHeader() (err error) { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SetTrailer(tl streaming.Trailer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) writeTrailer(tl streaming.Trailer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) RecvMsg(ctx context.Context, req any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) SendMsg(ctx context.Context, res any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) CloseSend(exception error) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) closeRecv(exception error) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) close(exception *Exception) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) onReadTrailerFrame(fr *Frame) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) onReadRstFrame(fr *Frame) (err error) { _ = "STUB: not implemented"; return nil }

func (s *serverStream) closeTest(exception error, cancelPath string) error {
	_ = "STUB: not implemented"
	return nil
}
