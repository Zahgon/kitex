package remote

import (
	"io"
	"net"
)

type ByteBufferFactory interface {
	NewByteBuffer(conn net.Conn) (ByteBuffer, error)
}

type NocopyWrite interface {
	WriteDirect(buf []byte, remainCap int) error

	MallocAck(n int) error
}

type FrameWrite interface {
	WriteHeader(buf []byte) (err error)

	WriteData(buf []byte) (err error)
}

type ByteBuffer interface {
	io.ReadWriter

	Next(n int) (p []byte, err error)

	Peek(n int) (buf []byte, err error)

	Skip(n int) (err error)

	Release(e error) (err error)

	ReadableLen() (n int)

	ReadLen() (n int)

	ReadString(n int) (s string, err error)

	ReadBinary(p []byte) (n int, err error)

	Malloc(n int) (buf []byte, err error)

	WrittenLen() (length int)

	WriteString(s string) (n int, err error)

	WriteBinary(b []byte) (n int, err error)

	Flush() (err error)

	NewBuffer() ByteBuffer

	AppendBuffer(buf ByteBuffer) (err error)

	Bytes() (buf []byte, err error)
}

type ByteBufferIO struct {
	buffer ByteBuffer
}

func NewByteBufferIO(buffer ByteBuffer) io.ReadWriter {
	_ = "STUB: not implemented"
	return *new(io.ReadWriter)
}

func (p *ByteBufferIO) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *ByteBufferIO) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
