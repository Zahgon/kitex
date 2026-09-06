package connpool

import (
	"net"
	"sync"
)

var (
	isSet          bool
	rwLock         sync.RWMutex
	commonReporter Reporter = &DummyReporter{}
)

type ConnectionPoolType int8

const (
	Short ConnectionPoolType = iota
	Long
)

type Reporter interface {
	ConnSucceed(poolType ConnectionPoolType, serviceName string, addr net.Addr)
	ConnFailed(poolType ConnectionPoolType, serviceName string, addr net.Addr)
	ReuseSucceed(poolType ConnectionPoolType, serviceName string, addr net.Addr)
}

func SetReporter(r Reporter) { _ = "STUB: not implemented"; return }

func GetCommonReporter() Reporter { _ = "STUB: not implemented"; return *new(Reporter) }
