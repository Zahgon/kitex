package kitexutil

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func GetCaller(ctx context.Context) (string, bool) { _ = "STUB: not implemented"; return "", false }

func GetCallee(ctx context.Context) (string, bool) { _ = "STUB: not implemented"; return "", false }

func GetMethod(ctx context.Context) (string, bool) { _ = "STUB: not implemented"; return "", false }

func GetCallerHandlerMethod(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func GetIDLServiceName(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func GetCallerAddr(ctx context.Context) (net.Addr, bool) {
	_ = "STUB: not implemented"
	return *new(net.Addr), false
}

func GetCallerIP(ctx context.Context) (string, bool) { _ = "STUB: not implemented"; return "", false }

func GetTransportProtocol(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func GetRPCInfo(ctx context.Context) (rpcinfo.RPCInfo, bool) {
	_ = "STUB: not implemented"
	return *new(rpcinfo.RPCInfo), false
}

func GetRealReqFromKitexArgs(req interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func GetRealRespFromKitexResult(resp interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}
