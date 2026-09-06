package retry

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func newMixedRetryer(policy Policy, r *ShouldResultRetry, cbC *cbContainer) (Retryer, error) {
	_ = "STUB: not implemented"
	return *new(Retryer), nil
}

type mixedRetryer struct {
	enable bool
	*failureCommon
	policy     *MixedPolicy
	retryDelay time.Duration
	sync.RWMutex
	errMsg string
}

func (r *mixedRetryer) ShouldRetry(ctx context.Context, err error, callTimes int, req interface{}, cbKey string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *mixedRetryer) AllowRetry(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *mixedRetryer) Do(ctx context.Context, rpcCall RPCCallFunc, firstRI rpcinfo.RPCInfo, req, resp interface{}) (lastRI rpcinfo.RPCInfo, recycleRI bool, err error) {
	_ = "STUB: not implemented"
	return *new(rpcinfo.RPCInfo), false, nil
}

func (r *mixedRetryer) UpdatePolicy(rp Policy) (err error) { _ = "STUB: not implemented"; return nil }

func (r *mixedRetryer) AppendErrMsgIfNeeded(ctx context.Context, err error, ri rpcinfo.RPCInfo, msg string) {
	_ = "STUB: not implemented"
	return
}

func (r *mixedRetryer) Prepare(ctx context.Context, prevRI, retryRI rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return
}

func (r *mixedRetryer) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (r *mixedRetryer) Dump() map[string]interface{} { _ = "STUB: not implemented"; return nil }
