package proto

import (
	"context"
)

type RawReaderWriter struct {
	*RawReader
	*RawWriter
}

func NewRawReaderWriter() *RawReaderWriter { _ = "STUB: not implemented"; return nil }

func NewRawWriter() *RawWriter { _ = "STUB: not implemented"; return nil }

type RawWriter struct{}

var _ MessageWriter = (*RawWriter)(nil)

func (m *RawWriter) Write(ctx context.Context, msg interface{}, method string, isClient bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRawReader() *RawReader { _ = "STUB: not implemented"; return nil }

type RawReader struct{}

var _ MessageReader = (*RawReader)(nil)

func (m *RawReader) Read(ctx context.Context, method string, isClient bool, actualMsgBuf []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
