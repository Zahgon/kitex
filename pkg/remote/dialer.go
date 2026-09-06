package remote

import (
	"net"
	"time"
)

type Dialer interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

func NewDefaultDialer() Dialer { _ = "STUB: not implemented"; return *new(Dialer) }

type SynthesizedDialer struct {
	DialFunc func(network, address string, timeout time.Duration) (net.Conn, error)
}

func (sd SynthesizedDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
