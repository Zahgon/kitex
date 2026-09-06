package retry

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
)

var errUnexpectedFinish = errors.New("backup request: all retries finished unexpectedly, " +
	"please submit an issue to https://github.com/cloudwego/kitex/issues")

func newBackupRetryer(policy Policy, cbC *cbContainer) (Retryer, error) {
	_ = "STUB: not implemented"
	return *new(Retryer), nil
}

type backupRetryer struct {
	enable      bool
	retryDelay  time.Duration
	policy      *BackupPolicy
	cbContainer *cbContainer
	sync.RWMutex
	errMsg string
}

type resultWrapper struct {
	ri   rpcinfo.RPCInfo
	resp interface{}
	err  error
}

func (r *backupRetryer) ShouldRetry(ctx context.Context, err error, callTimes int, req interface{}, cbKey string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *backupRetryer) AllowRetry(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *backupRetryer) Do(ctx context.Context, rpcCall RPCCallFunc, firstRI rpcinfo.RPCInfo, req, resp interface{}) (lastRI rpcinfo.RPCInfo, recycleRI bool, err error) {
	_ = "STUB: not implemented"
	return *new(rpcinfo.RPCInfo), false, nil
}

func (r *backupRetryer) Prepare(ctx context.Context, prevRI, retryRI rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return
}

func (r *backupRetryer) UpdatePolicy(rp Policy) (err error) { _ = "STUB: not implemented"; return nil }

func (r *backupRetryer) AppendErrMsgIfNeeded(ctx context.Context, err error, ri rpcinfo.RPCInfo, msg string) {
	_ = "STUB: not implemented"
	return
}

func (r *backupRetryer) Dump() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (r *backupRetryer) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func recordCost(ct int32, start time.Time, recordCostDoing *int32, sb *utils.StringBuilder, abort *int32, err error) {
	_ = "STUB: not implemented"
	return
}
