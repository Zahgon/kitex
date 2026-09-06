//go:build !windows

package ttstream

import (
	"testing"

	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

type mockStreamWriter struct {
	writeFrameFunc  func(f *Frame) error
	closeStreamFunc func(sid int32) error
}

func (m mockStreamWriter) WriteFrame(f *Frame) error { _ = "STUB: not implemented"; return nil }

func (m mockStreamWriter) CloseStream(sid int32) error { _ = "STUB: not implemented"; return nil }

func newTestConnectionPipe(t *testing.T) (netpoll.Connection, netpoll.Connection) {
	_ = "STUB: not implemented"
	return *new(netpoll.Connection), *new(netpoll.Connection)
}

func newTestStreamPipe(sinfo *serviceinfo.ServiceInfo, method string) (*clientStream, *serverStream, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
