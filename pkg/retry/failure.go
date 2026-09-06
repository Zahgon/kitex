package retry

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

const maxFailureRetryTimes = 5

func AllErrorRetry() *ShouldResultRetry { _ = "STUB: not implemented"; return nil }

func NewFailurePolicy() *FailurePolicy { _ = "STUB: not implemented"; return nil }

func NewFailurePolicyWithResultRetry(rr *ShouldResultRetry) *FailurePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *FailurePolicy) WithMaxRetryTimes(retryTimes int) { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) WithMaxDurationMS(maxMS uint32) { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) DisableChainRetryStop() { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) WithDDLStop() { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) WithFixedBackOff(fixMS int) { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) WithRandomBackOff(minMS, maxMS int) { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) WithRetryBreaker(errRate float64) { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) WithRetrySameNode() { _ = "STUB: not implemented"; return }

func (p *FailurePolicy) WithSpecifiedResultRetry(rr *ShouldResultRetry) {
	_ = "STUB: not implemented"
	return
}

func (p *FailurePolicy) String() string { _ = "STUB: not implemented"; return "" }

func (p *FailurePolicy) Equals(np *FailurePolicy) bool { _ = "STUB: not implemented"; return false }

func (p *FailurePolicy) DeepCopy() *FailurePolicy { _ = "STUB: not implemented"; return nil }

func (p *FailurePolicy) isRespRetryWithCtxNonNil() bool { _ = "STUB: not implemented"; return false }

func (p *FailurePolicy) isErrorRetryWithCtxNonNil() bool { _ = "STUB: not implemented"; return false }

func (p *FailurePolicy) isRespRetryNonNil() bool { _ = "STUB: not implemented"; return false }

func (p *FailurePolicy) isErrorRetryNonNil() bool { _ = "STUB: not implemented"; return false }

func (p *FailurePolicy) isRetryForTimeout() bool { _ = "STUB: not implemented"; return false }

func (p *FailurePolicy) isRespRetry(ctx context.Context, resp interface{}, ri rpcinfo.RPCInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *FailurePolicy) isErrorRetry(ctx context.Context, err error, ri rpcinfo.RPCInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *FailurePolicy) convertResultRetry() { _ = "STUB: not implemented"; return }

type BackOff interface {
	Wait(callTimes int)
}

var NoneBackOff = &noneBackOff{}

func newFixedBackOff(fixMS int) BackOff { _ = "STUB: not implemented"; return *new(BackOff) }

func newRandomBackOff(minMS, maxMS int) BackOff { _ = "STUB: not implemented"; return *new(BackOff) }

type noneBackOff struct{}

func (p noneBackOff) String() string { _ = "STUB: not implemented"; return "" }

func (p *noneBackOff) Wait(callTimes int) { _ = "STUB: not implemented"; return }

type fixedBackOff struct {
	fixMS int
}

func (p *fixedBackOff) Wait(callTimes int) { _ = "STUB: not implemented"; return }

func (p fixedBackOff) String() string { _ = "STUB: not implemented"; return "" }

type randomBackOff struct {
	minMS int
	maxMS int
}

func (p *randomBackOff) randomDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (p *randomBackOff) Wait(callTimes int) { _ = "STUB: not implemented"; return }

func (p randomBackOff) String() string { _ = "STUB: not implemented"; return "" }

func checkFixedBackOff(fixMS int) error { _ = "STUB: not implemented"; return nil }

func checkRandomBackOff(minMS, maxMS int) error { _ = "STUB: not implemented"; return nil }
