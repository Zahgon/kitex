package retry

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func newFailureRetryer(policy Policy, r *ShouldResultRetry, cbC *cbContainer) (Retryer, error) {
	_ = "STUB: not implemented"
	return *new(Retryer), nil
}

type failureRetryer struct {
	enable bool
	*failureCommon
	policy *FailurePolicy
	sync.RWMutex
	errMsg string
}

func (r *failureRetryer) ShouldRetry(ctx context.Context, err error, callTimes int, req interface{}, cbKey string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *failureRetryer) AllowRetry(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *failureRetryer) Do(ctx context.Context, rpcCall RPCCallFunc, firstRI rpcinfo.RPCInfo, req, resp interface{}) (lastRI rpcinfo.RPCInfo, recycleRI bool, err error) {
	_ = "STUB: not implemented"
	return *new(rpcinfo.RPCInfo), false, nil
}

func (r *failureRetryer) UpdatePolicy(rp Policy) (err error) { _ = "STUB: not implemented"; return nil }

func (r *failureRetryer) AppendErrMsgIfNeeded(ctx context.Context, err error, ri rpcinfo.RPCInfo, msg string) {
	_ = "STUB: not implemented"
	return
}

func (r *failureRetryer) Prepare(ctx context.Context, prevRI, retryRI rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return
}

func (r *failureRetryer) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (r *failureRetryer) Dump() map[string]interface{} { _ = "STUB: not implemented"; return nil }

type failureCommon struct {
	backOff              BackOff
	specifiedResultRetry *ShouldResultRetry
	cbContainer          *cbContainer
}

func (f *failureCommon) setSpecifiedResultRetryIfNeeded(rr *ShouldResultRetry, fp *FailurePolicy) {
	_ = "STUB: not implemented"
	return
}

func (r *failureCommon) isRetryErr(ctx context.Context, err error, ri rpcinfo.RPCInfo, fp *FailurePolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *failureCommon) shouldRetry(ctx context.Context, callTimes int, req interface{}, cbKey string, fp *FailurePolicy) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *failureCommon) isRetryResult(ctx context.Context, cRI rpcinfo.RPCInfo, resp interface{}, err error, fp *FailurePolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *failureCommon) dumpSpecifiedResultRetry(fp FailurePolicy) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func initBackOff(policy *BackOffPolicy) (bo BackOff, err error) {
	_ = "STUB: not implemented"
	return *new(BackOff), nil
}

func isExceedMaxDuration(ctx context.Context, start time.Time, maxDuration time.Duration, callTimes int32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
