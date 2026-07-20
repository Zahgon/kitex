package invoke

import (
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/remote"
)

var _ Message = &message{}

type PayloadHandler interface {
	SetRequestBytes(buf []byte) error
	GetResponseBytes() ([]byte, error)
	GetRequestReaderByteBuffer() remote.ByteBuffer
	GetResponseWriterByteBuffer() remote.ByteBuffer
	Release() error
}

type Message interface {
	net.Conn
	PayloadHandler
}

type message struct {
	localAddr  net.Addr
	remoteAddr net.Addr
	request    remote.ByteBuffer
	response   remote.ByteBuffer
}

func NewMessage(local, remote net.Addr) Message { _ = "STUB: not implemented"; return *new(Message) }

func (f *message) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (f *message) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (f *message) Close() error { _ = "STUB: not implemented"; return nil }

func (f *message) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (f *message) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (f *message) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (f *message) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (f *message) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (f *message) SetRequestBytes(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (f *message) GetResponseBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *message) GetRequestReaderByteBuffer() remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (f *message) GetResponseWriterByteBuffer() remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (f *message) Release() error { _ = "STUB: not implemented"; return nil }
