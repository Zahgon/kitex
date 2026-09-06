package client

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

var workerPool = newTimeoutPool(128, time.Minute)

func makeTimeoutErr(ctx context.Context, start time.Time, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func isBusinessTimeout(start time.Time, kitexTimeout time.Duration, actualDDL time.Time, threshold time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func rpcTimeoutMW(mwCtx context.Context) endpoint.UnaryMiddleware {
	_ = "STUB: not implemented"
	return *new(endpoint.UnaryMiddleware)
}
