package lbcache

import (
	"sync"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/cloudwego/kitex/pkg/utils"
)

var (
	sharedTickers    sync.Map
	sharedTickersSfg singleflight.Group
)

func getSharedTicker(b *Balancer, refreshInterval time.Duration) *utils.SharedTicker {
	_ = "STUB: not implemented"
	return nil
}
