package fallback

import (
	"context"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
)

func ErrorFallback(ef Func) *Policy { _ = "STUB: not implemented"; return nil }

func TimeoutAndCBFallback(ef Func) *Policy { _ = "STUB: not implemented"; return nil }

func NewFallbackPolicy(fb Func) *Policy { _ = "STUB: not implemented"; return nil }

type Policy struct {
	fallbackFunc     Func
	reportAsFallback bool
}

func (p *Policy) EnableReportAsFallback() *Policy { _ = "STUB: not implemented"; return nil }

func IsPolicyValid(p *Policy) bool { _ = "STUB: not implemented"; return false }

func UnwrapHelper(userFB RealReqRespFunc) Func { _ = "STUB: not implemented"; return *new(Func) }

type Func func(ctx context.Context, args utils.KitexArgs, result utils.KitexResult, err error) (fbErr error)

type RealReqRespFunc func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error)

func (p *Policy) DoIfNeeded(ctx context.Context, ri rpcinfo.RPCInfo, args, result interface{}, err error) (fbErr error, reportAsFallback bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func getBizErrIfExist(ri rpcinfo.RPCInfo, err error) (error, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
