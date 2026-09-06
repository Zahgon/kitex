package gonet

import (
	"sync"

	"github.com/cloudwego/gopkg/bufiox"

	"github.com/cloudwego/kitex/pkg/remote"
)

var rwPool = sync.Pool{New: func() any { return &bufferReadWriter{} }}

var _ remote.ByteBuffer = &bufferReadWriter{}

type bufferReadWriter struct {
	reader *bufiox.DefaultReader
	writer *bufiox.DefaultWriter
	status int
}

func newBufferReader(reader *bufiox.DefaultReader) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func newBufferWriter(writer *bufiox.DefaultWriter) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func newBufferReadWriter(irw bufioxReadWriter) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (rw *bufferReadWriter) readable() bool { _ = "STUB: not implemented"; return false }

func (rw *bufferReadWriter) writable() bool { _ = "STUB: not implemented"; return false }

func (rw *bufferReadWriter) Next(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rw *bufferReadWriter) Peek(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rw *bufferReadWriter) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

func (rw *bufferReadWriter) ReadableLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (rw *bufferReadWriter) ReadString(n int) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rw *bufferReadWriter) ReadBinary(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *bufferReadWriter) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *bufferReadWriter) ReadLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (rw *bufferReadWriter) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rw *bufferReadWriter) WrittenLen() (length int) { _ = "STUB: not implemented"; return 0 }

func (rw *bufferReadWriter) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *bufferReadWriter) WriteBinary(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *bufferReadWriter) Flush() (err error) { _ = "STUB: not implemented"; return nil }

func (rw *bufferReadWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *bufferReadWriter) Release(e error) (err error) { _ = "STUB: not implemented"; return nil }

func (rw *bufferReadWriter) AppendBuffer(buf remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rw *bufferReadWriter) NewBuffer() remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (rw *bufferReadWriter) Bytes() (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rw *bufferReadWriter) zero() { _ = "STUB: not implemented"; return }
