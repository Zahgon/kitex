package ttstream

import (
	"time"

	"github.com/cloudwego/kitex/pkg/remote/trans/ttstream/internal/container"
)

var DefaultLongConnConfig = LongConnConfig{
	MaxIdleTimeout: time.Minute,
}

type LongConnConfig struct {
	MaxIdleTimeout time.Duration
}

func newLongConnTransPool(config LongConnConfig) transPool {
	_ = "STUB: not implemented"
	return *new(transPool)
}

type longConnTransPool struct {
	transPool *container.ObjectPool
	config    LongConnConfig
}

func (c *longConnTransPool) Get(network, addr string) (trans *clientTransport, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *longConnTransPool) Put(trans *clientTransport) { _ = "STUB: not implemented"; return }
