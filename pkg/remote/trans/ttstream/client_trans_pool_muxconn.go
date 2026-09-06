package ttstream

import (
	"runtime"
	"sync"
	"time"
)

var DefaultMuxConnConfig = MuxConnConfig{
	PoolSize:       runtime.GOMAXPROCS(0),
	MaxIdleTimeout: time.Minute,
}

type MuxConnConfig struct {
	PoolSize       int
	MaxIdleTimeout time.Duration
}

var _ transPool = (*muxConnTransPool)(nil)

type muxConnTransList struct {
	L          sync.RWMutex
	size       int
	cursor     uint32
	transports []*clientTransport
	pool       *muxConnTransPool
}

func newMuxConnTransList(size int, pool *muxConnTransPool) *muxConnTransList {
	_ = "STUB: not implemented"
	return nil
}

func (tl *muxConnTransList) Close() { _ = "STUB: not implemented"; return }

func (tl *muxConnTransList) Get(network, addr string) (*clientTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newMuxConnTransPool(config MuxConnConfig) transPool {
	_ = "STUB: not implemented"
	return *new(transPool)
}

type muxConnTransPool struct {
	config      MuxConnConfig
	pool        sync.Map
	activity    sync.Map
	cleanerOnce int32
}

func (p *muxConnTransPool) Get(network, addr string) (trans *clientTransport, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *muxConnTransPool) Put(trans *clientTransport) { _ = "STUB: not implemented"; return }
