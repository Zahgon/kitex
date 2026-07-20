package retry

import (
	"context"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

type ctxKey string

const (
	TransitKey = "RetryReq"

	CtxReqOp ctxKey = "K_REQ_OP"

	Wildcard = "*"
)

const (
	OpNo int32 = iota
	OpDoing
	OpDone
)

var tagValueFirstTry = "0"

type DDLStopFunc func(ctx context.Context, policy StopPolicy) (bool, string)

var ddlStopFunc DDLStopFunc

func RegisterDDLStop(f DDLStopFunc) { _ = "STUB: not implemented"; return }

func ddlStop(ctx context.Context, policy StopPolicy) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func chainStop(ctx context.Context, policy StopPolicy) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func circuitBreakerStop(ctx context.Context, policy StopPolicy, cbC *cbContainer, request interface{}, cbKey string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func handleRetryInstance(retrySameNode bool, prevRI, retryRI rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return
}

func makeRetryErr(ctx context.Context, msg string, callTimes int32) error {
	_ = "STUB: not implemented"
	return nil
}

func panicToErr(ctx context.Context, panicInfo interface{}, ri rpcinfo.RPCInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func appendErrMsg(err error, msg string) { _ = "STUB: not implemented"; return }

func recordRetryInfo(ri rpcinfo.RPCInfo, callTimes int32, lastCosts string) {
	_ = "STUB: not implemented"
	return
}

func IsLocalRetryRequest(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func IsRemoteRetryRequest(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func getNewRespFunc(mi serviceinfo.MethodInfo) func() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func shallowCopyResults(src, dst interface{}) { _ = "STUB: not implemented"; return }
