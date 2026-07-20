package grpc

import (
	"io"
	"net"

	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc/grpcframe"
)

type framer struct {
	*grpcframe.Framer
	reader netpoll.Reader
	writer bufWriter
}

func newFramer(conn net.Conn, writeBufferSize, readBufferSize, maxHeaderListSize uint32, reuseCfg ReuseWriteBufferConfig) *framer {
	_ = "STUB: not implemented"
	return nil
}

type ReuseWriteBufferConfig struct {
	Enable bool

	EnableReuseHTTP2FramerBuffer bool
}

type bufWriter interface {
	Write(b []byte) (n int, err error)
	Flush() error
	GetOffset() int
}

func newBufWriter(writer io.Writer, batchSize int, reuseCfg ReuseWriteBufferConfig) bufWriter {
	_ = "STUB: not implemented"
	return *new(bufWriter)
}

type keepBufWriter struct {
	writer    io.Writer
	buf       []byte
	offset    int
	batchSize int
	err       error
}

func (w *keepBufWriter) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *keepBufWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (w *keepBufWriter) GetOffset() int { _ = "STUB: not implemented"; return 0 }

func (w *keepBufWriter) flushAllocatedBuffer() error { _ = "STUB: not implemented"; return nil }

type reuseBufWriter struct {
	*keepBufWriter
}

func (w *reuseBufWriter) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *reuseBufWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (w *reuseBufWriter) GetOffset() int { _ = "STUB: not implemented"; return 0 }
