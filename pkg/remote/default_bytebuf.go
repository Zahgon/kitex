package remote

import (
	"sync"
)

const (
	BitReadable = 1 << iota
	BitWritable
)

var bytebufPool sync.Pool

func init() {
	bytebufPool.New = newDefaultByteBuffer
}

func NewWriterBuffer(size int) ByteBuffer { _ = "STUB: not implemented"; return *new(ByteBuffer) }

func NewReaderBuffer(buf []byte) ByteBuffer { _ = "STUB: not implemented"; return *new(ByteBuffer) }

func NewReaderWriterBuffer(size int) ByteBuffer { _ = "STUB: not implemented"; return *new(ByteBuffer) }

type defaultByteBuffer struct {
	buff     []byte
	readIdx  int
	writeIdx int
	status   int
}

var _ ByteBuffer = &defaultByteBuffer{}

func newDefaultByteBuffer() interface{} { _ = "STUB: not implemented"; return nil }

func newReaderByteBuffer(buf []byte) ByteBuffer { _ = "STUB: not implemented"; return *new(ByteBuffer) }

func newWriterByteBuffer(estimatedLength int) ByteBuffer {
	_ = "STUB: not implemented"
	return *new(ByteBuffer)
}

func newReaderWriterByteBuffer(estimatedLength int) ByteBuffer {
	_ = "STUB: not implemented"
	return *new(ByteBuffer)
}

func (b *defaultByteBuffer) Next(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *defaultByteBuffer) Peek(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *defaultByteBuffer) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

func (b *defaultByteBuffer) ReadableLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *defaultByteBuffer) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *defaultByteBuffer) ReadString(n int) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *defaultByteBuffer) ReadBinary(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *defaultByteBuffer) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *defaultByteBuffer) WrittenLen() (length int) { _ = "STUB: not implemented"; return 0 }

func (b *defaultByteBuffer) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *defaultByteBuffer) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *defaultByteBuffer) WriteBinary(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *defaultByteBuffer) ReadLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *defaultByteBuffer) Flush() (err error) { _ = "STUB: not implemented"; return nil }

func (b *defaultByteBuffer) AppendBuffer(buf ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *defaultByteBuffer) Bytes() (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *defaultByteBuffer) NewBuffer() ByteBuffer {
	_ = "STUB: not implemented"
	return *new(ByteBuffer)
}

func (b *defaultByteBuffer) Release(e error) (err error) { _ = "STUB: not implemented"; return nil }

func (b *defaultByteBuffer) zero() { _ = "STUB: not implemented"; return }

func (b *defaultByteBuffer) readableCheck(n int) error { _ = "STUB: not implemented"; return nil }

func (b *defaultByteBuffer) writableLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *defaultByteBuffer) ensureWritable(minWritableBytes int) { _ = "STUB: not implemented"; return }
