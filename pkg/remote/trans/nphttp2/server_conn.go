package nphttp2

import (
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type serverConnKey struct{}

type serverConn struct {
	tr grpc.ServerTransport
	s  *grpc.Stream
}

var _ GRPCConn = (*serverConn)(nil)

func newServerConn(tr grpc.ServerTransport, s *grpc.Stream) *serverConn {
	_ = "STUB: not implemented"
	return nil
}

func (c *serverConn) ReadFrame() (hdr, data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GetServerConn(st streaming.Stream) (GRPCConn, error) {
	_ = "STUB: not implemented"
	return *new(GRPCConn), nil
}

func (c *serverConn) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *serverConn) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *serverConn) WriteFrame(hdr, data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *serverConn) LocalAddr() net.Addr                { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *serverConn) RemoteAddr() net.Addr               { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *serverConn) SetDeadline(t time.Time) error      { _ = "STUB: not implemented"; return nil }
func (c *serverConn) SetReadDeadline(t time.Time) error  { _ = "STUB: not implemented"; return nil }
func (c *serverConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
func (c *serverConn) Close() error                       { _ = "STUB: not implemented"; return nil }
