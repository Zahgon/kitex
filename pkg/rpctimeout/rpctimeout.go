package rpctimeout

import (
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

type timeoutAdjustKeyType int

const TimeoutAdjustKey timeoutAdjustKeyType = 1

func MiddlewareBuilder(moreTimeout time.Duration) endpoint.MiddlewareBuilder {
	_ = "STUB: not implemented"
	return *new(endpoint.MiddlewareBuilder)
}

var globalNeedFineGrainedErrCode int32 = 0

func EnableGlobalNeedFineGrainedErrCode() { _ = "STUB: not implemented"; return }

func DisableGlobalNeedFineGrainedErrCode() { _ = "STUB: not implemented"; return }

func LoadGlobalNeedFineGrainedErrCode() bool { _ = "STUB: not implemented"; return false }

var defaultBusinessTimeoutThreshold = time.Millisecond * 50

func SetBusinessTimeoutThreshold(t time.Duration) { _ = "STUB: not implemented"; return }

func LoadBusinessTimeoutThreshold() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
