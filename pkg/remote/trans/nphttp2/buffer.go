package nphttp2

import (
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"
)

const (
	minMallocSize = 4 * 1024
)

type GRPCConn interface {
	net.Conn

	WriteFrame(hdr, data []byte) (n int, err error)

	ReadFrame() (hdr, data []byte, err error)
}

type buffer struct {
	conn       GRPCConn
	rbuf       []byte
	whdr, wbuf []byte
}

var (
	_ remote.ByteBuffer = (*buffer)(nil)
	_ remote.FrameWrite = (*buffer)(nil)
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &buffer{}
	},
}

func newBuffer(conn GRPCConn) *buffer { _ = "STUB: not implemented"; return nil }

func (b *buffer) growRbuf(n int) { _ = "STUB: not implemented"; return }

func (b *buffer) Next(n int) (p []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (b *buffer) WriteHeader(buf []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (b *buffer) WriteData(buf []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (b *buffer) Flush() (err error) { _ = "STUB: not implemented"; return nil }

func (b *buffer) Release(e error) (err error) { _ = "STUB: not implemented"; return nil }

func (b *buffer) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *buffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *buffer) Peek(n int) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (b *buffer) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

func (b *buffer) ReadableLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *buffer) ReadLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *buffer) ReadString(n int) (s string, err error) { _ = "STUB: not implemented"; return "", nil }

func (b *buffer) ReadBinary(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *buffer) Malloc(n int) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (b *buffer) WrittenLen() (length int) { _ = "STUB: not implemented"; return 0 }

func (b *buffer) WriteString(s string) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *buffer) WriteBinary(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *buffer) NewBuffer() remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (b *buffer) AppendBuffer(buf remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *buffer) Bytes() (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }
