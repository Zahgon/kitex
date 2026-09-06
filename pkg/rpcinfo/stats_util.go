package rpcinfo

import (
	"context"

	"github.com/cloudwego/kitex/pkg/stats"
)

func Record(ctx context.Context, ri RPCInfo, event stats.Event, err error) {
	_ = "STUB: not implemented"
	return
}

func CalcEventCostUs(start, end Event) uint64 { _ = "STUB: not implemented"; return 0 }

func ClientPanicToErr(ctx context.Context, panicInfo interface{}, ri RPCInfo, logErr bool) error {
	_ = "STUB: not implemented"
	return nil
}
