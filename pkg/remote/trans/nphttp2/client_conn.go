package nphttp2

import (
	"context"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

const (
	contentSubTypeThrift   = "thrift"
	contentSubTypeProtobuf = "protobuf"
)

type streamDesc struct {
	isStreaming bool
}

type clientConn struct {
	tr   grpc.ClientTransport
	s    *grpc.Stream
	desc *streamDesc
}

var _ GRPCConn = (*clientConn)(nil)

func (c *clientConn) ReadFrame() (hdr, data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getContentSubType(codec serviceinfo.PayloadCodec) string { _ = "STUB: not implemented"; return "" }

func newClientConn(ctx context.Context, tr grpc.ClientTransport, addr string) (*clientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fullMethodName(pkg, svc, method string) string { _ = "STUB: not implemented"; return "" }

func (c *clientConn) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *clientConn) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *clientConn) WriteFrame(hdr, data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *clientConn) LocalAddr() net.Addr                { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *clientConn) RemoteAddr() net.Addr               { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *clientConn) SetDeadline(t time.Time) error      { _ = "STUB: not implemented"; return nil }
func (c *clientConn) SetReadDeadline(t time.Time) error  { _ = "STUB: not implemented"; return nil }
func (c *clientConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
func (c *clientConn) Close() error                       { _ = "STUB: not implemented"; return nil }

func (c *clientConn) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}
func (c *clientConn) Trailer() metadata.MD    { _ = "STUB: not implemented"; return *new(metadata.MD) }
func (c *clientConn) GetRecvCompress() string { _ = "STUB: not implemented"; return "" }

func (c *clientConn) cancel(err error) { _ = "STUB: not implemented"; return }

type hasGetRecvCompress interface {
	GetRecvCompress() string
}

type hasHeader interface {
	Header() (metadata.MD, error)
}

type hasTrailer interface {
	Trailer() metadata.MD
}
