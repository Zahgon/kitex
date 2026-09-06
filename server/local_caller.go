package server

import (
	"context"
	"reflect"
	"sync"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/utils"
)

var localCallerAddr = utils.NewNetAddr("tcp", "127.0.0.1:0")

type LocalCaller interface {
	Call(ctx context.Context, method string, args, result any) error

	ResolveMethod(ctx context.Context, method string) (serviceinfo.MethodInfo, error)
}

type cachedResult struct {
	methodName string
	svcInfo    *serviceinfo.ServiceInfo
	mi         serviceinfo.MethodInfo
	argsType   reflect.Type
	resultType reflect.Type
}

func newCachedResult(methodName string, svcInfo *serviceinfo.ServiceInfo, mi serviceinfo.MethodInfo) *cachedResult {
	_ = "STUB: not implemented"
	return nil
}

type localCaller struct {
	svr      *server
	traceCtl *rpcinfo.TraceController
	caller   string
	cache    sync.Map
}

func NewLocalCaller(caller string, svr Server) (LocalCaller, error) {
	_ = "STUB: not implemented"
	return *new(LocalCaller), nil
}

func (lc *localCaller) Call(ctx context.Context, method string, args, result any) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func (lc *localCaller) ResolveMethod(_ context.Context, method string) (serviceinfo.MethodInfo, error) {
	_ = "STUB: not implemented"
	return *new(serviceinfo.MethodInfo), nil
}

func (lc *localCaller) resolve(method string) (*cachedResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lc *localCaller) doResolve(method string) (*cachedResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lc *localCaller) newLocalRPCInfo(svcInfo *serviceinfo.ServiceInfo, mi serviceinfo.MethodInfo, methodName, callerMethod string) rpcinfo.RPCInfo {
	_ = "STUB: not implemented"
	return *new(rpcinfo.RPCInfo)
}
