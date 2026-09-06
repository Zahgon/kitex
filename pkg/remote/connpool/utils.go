package connpool

import (
	"net"
)

type ConnWithPkgSize struct {
	net.Conn
	Written int32
	Readn   int32
}

func (c *ConnWithPkgSize) Read(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *ConnWithPkgSize) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *ConnWithPkgSize) Close() error { _ = "STUB: not implemented"; return nil }
