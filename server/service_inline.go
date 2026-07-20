package server

import (
	"context"
	"reflect"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
)

var localAddr = utils.NewNetAddr("tcp", "127.0.0.1")

var invocationType = reflect.TypeOf(rpcinfo.NewServerInvocation()).Elem()

func constructServerCtxWithMetadata(cliCtx context.Context) (serverCtx context.Context) {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *server) constructServerRPCInfo(svrCtx context.Context, cliRPCInfo rpcinfo.RPCInfo) (newServerCtx context.Context, svrRPCInfo rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(rpcinfo.RPCInfo)
}

func (s *server) BuildServiceInlineInvokeChain() endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}
