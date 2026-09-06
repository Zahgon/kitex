package netpoll

import (
	"sync"

	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/remote"
)

var bytebufPool sync.Pool

func init() {
	bytebufPool.New = newNetpollByteBuffer
}

func NewReaderByteBuffer(r netpoll.Reader) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func NewWriterByteBuffer(w netpoll.Writer) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func NewReaderWriterByteBuffer(rw netpoll.ReadWriter) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func newNetpollByteBuffer() interface{} { _ = "STUB: not implemented"; return nil }

type netpollByteBuffer struct {
	writer   netpoll.Writer
	reader   netpoll.Reader
	status   int
	readSize int
}

var _ remote.ByteBuffer = &netpollByteBuffer{}

func (b *netpollByteBuffer) NetpollReader() netpoll.Reader {
	_ = "STUB: not implemented"
	return *new(netpoll.Reader)
}

func (b *netpollByteBuffer) Next(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *netpollByteBuffer) Peek(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *netpollByteBuffer) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

func (b *netpollByteBuffer) ReadableLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *netpollByteBuffer) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *netpollByteBuffer) ReadString(n int) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *netpollByteBuffer) ReadBinary(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *netpollByteBuffer) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *netpollByteBuffer) MallocAck(n int) (err error) { _ = "STUB: not implemented"; return nil }

func (b *netpollByteBuffer) WrittenLen() (length int) { _ = "STUB: not implemented"; return 0 }

func (b *netpollByteBuffer) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *netpollByteBuffer) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *netpollByteBuffer) WriteBinary(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *netpollByteBuffer) WriteDirect(p []byte, remainCap int) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *netpollByteBuffer) ReadLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *netpollByteBuffer) Flush() (err error) { _ = "STUB: not implemented"; return nil }

func (b *netpollByteBuffer) NewBuffer() remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (b *netpollByteBuffer) AppendBuffer(buf remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *netpollByteBuffer) Bytes() (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *netpollByteBuffer) Release(e error) (err error) { _ = "STUB: not implemented"; return nil }

func (b *netpollByteBuffer) zero() { _ = "STUB: not implemented"; return }

func GetWrittenBytes(lb *netpoll.LinkBuffer) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
