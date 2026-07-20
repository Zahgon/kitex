package remote

import (
	"context"
	"net"
	"time"
)

type ConnOption struct {
	Dialer         Dialer
	ConnectTimeout time.Duration
}

type ConnPool interface {
	Get(ctx context.Context, network, address string, opt ConnOption) (net.Conn, error)

	Put(conn net.Conn) error

	Discard(conn net.Conn) error

	Close() error
}

type LongConnPool interface {
	ConnPool

	Clean(network, address string)
}

type ConnPoolReporter interface {
	EnableReporter()
}

type RawConn interface {
	RawConn() net.Conn
}

type IsActive interface {
	IsActive() bool
}
