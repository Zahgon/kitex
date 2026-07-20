package gonet

import (
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/remote"
)

func NewDialer() remote.Dialer { _ = "STUB: not implemented"; return *new(remote.Dialer) }

type dialer struct {
	net.Dialer
}

func (d *dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
