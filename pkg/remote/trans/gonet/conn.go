package gonet

import (
	"errors"
	"net"

	"github.com/cloudwego/gopkg/bufiox"
)

var (
	_             bufioxReadWriter = &cliConn{}
	_             bufioxReadWriter = &svrConn{}
	errConnClosed error            = errors.New("connection has been closed")
)

type bufioxReadWriter interface {
	Reader() *bufiox.DefaultReader
	Writer() *bufiox.DefaultWriter
}

type cliConn struct {
	net.Conn
	r      *bufiox.DefaultReader
	w      *bufiox.DefaultWriter
	closed uint32
}

func newCliConn(conn net.Conn) *cliConn { _ = "STUB: not implemented"; return nil }

func (c *cliConn) Reader() *bufiox.DefaultReader { _ = "STUB: not implemented"; return nil }

func (c *cliConn) Writer() *bufiox.DefaultWriter { _ = "STUB: not implemented"; return nil }

func (c *cliConn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *cliConn) Close() error { _ = "STUB: not implemented"; return nil }

type svrConn struct {
	net.Conn
	r      *bufiox.DefaultReader
	w      *bufiox.DefaultWriter
	closed uint32
}

func newSvrConn(conn net.Conn) *svrConn { _ = "STUB: not implemented"; return nil }

func (bc *svrConn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (bc *svrConn) Close() error { _ = "STUB: not implemented"; return nil }

func (bc *svrConn) Reader() *bufiox.DefaultReader { _ = "STUB: not implemented"; return nil }

func (bc *svrConn) Writer() *bufiox.DefaultWriter { _ = "STUB: not implemented"; return nil }
