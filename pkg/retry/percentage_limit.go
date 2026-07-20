package retry

import (
	"github.com/bytedance/gopkg/cloud/circuitbreaker"
)

func recordRetryStat(cbKey string, panel circuitbreaker.Panel, callTimes int32) {
	_ = "STUB: not implemented"
	return
}
