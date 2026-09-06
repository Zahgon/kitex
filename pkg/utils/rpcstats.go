package utils

import (
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
)

func CalculateEventCost(rpcstats rpcinfo.RPCStats, start, end stats.Event) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
