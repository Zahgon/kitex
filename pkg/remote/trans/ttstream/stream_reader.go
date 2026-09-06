package ttstream

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote/trans/ttstream/internal/container"
)

type streamReader struct {
	pipe      *container.Pipe[streamMsg]
	cache     [1]streamMsg
	exception error
}

type streamMsg struct {
	payload   []byte
	exception error
}

func newStreamReader(ctx context.Context, callback container.CtxDoneCallback) *streamReader {
	_ = "STUB: not implemented"
	return nil
}

func (s *streamReader) input(payload []byte) { _ = "STUB: not implemented"; return }

func (s *streamReader) output() (payload []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *streamReader) outputWithCtx(ctx context.Context, perReadCallback container.CtxDoneCallback) (payload []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *streamReader) handleReadResult(n int, err error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *streamReader) close(exception error) { _ = "STUB: not implemented"; return }
