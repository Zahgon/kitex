package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift/base"
)

type RawReaderWriter struct {
	*RawReader
	*RawWriter
}

func NewRawReaderWriter() *RawReaderWriter { _ = "STUB: not implemented"; return nil }

func NewRawWriter() *RawWriter { _ = "STUB: not implemented"; return nil }

type RawWriter struct{}

var _ MessageWriter = (*RawWriter)(nil)

func (m *RawWriter) Write(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRawReader() *RawReader { _ = "STUB: not implemented"; return nil }

type RawReader struct{}

var _ MessageReader = (*RawReader)(nil)

func (m *RawReader) Read(ctx context.Context, method string, isClient bool, dataLen int, in bufiox.Reader) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
