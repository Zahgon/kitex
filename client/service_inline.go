package client

import (
	"context"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/internal/client"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

type ContextServiceInlineHandler interface {
	WriteMeta(cliCtx context.Context, req interface{}) (newCliCtx context.Context, err error)
	ReadMeta(cliCtx context.Context, resp interface{}) (err error)
}

type serviceInlineClient struct {
	svcInfo *serviceinfo.ServiceInfo
	mws     []endpoint.Middleware
	eps     endpoint.Endpoint
	opt     *client.Options

	inited bool
	closed bool

	serverEps endpoint.Endpoint

	contextServiceInlineHandler ContextServiceInlineHandler
}

type ServerInitialInfo interface {
	BuildServiceInlineInvokeChain() endpoint.Endpoint
}

func NewServiceInlineClient(svcInfo *serviceinfo.ServiceInfo, s ServerInitialInfo, opts ...Option) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (kc *serviceInlineClient) SetContextServiceInlineHandler(simh ContextServiceInlineHandler) {
	_ = "STUB: not implemented"
	return
}

func (kc *serviceInlineClient) init() (err error) {
	if err = kc.checkOptions(); err != nil {
		return err
	}
	ctx := kc.initContext()
	kc.initMiddlewares(ctx)
	kc.richRemoteOption()
	if err = kc.buildInvokeChain(); err != nil {
		return err
	}
	kc.inited = true
	return nil
}

func (kc *serviceInlineClient) checkOptions() (err error) { _ = "STUB: not implemented"; return nil }

func (kc *serviceInlineClient) initContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (kc *serviceInlineClient) initMiddlewares(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (kc *serviceInlineClient) initRPCInfo(ctx context.Context, method string) (context.Context, rpcinfo.RPCInfo, *callopt.CallOptions) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(rpcinfo.RPCInfo), nil
}

func (kc *serviceInlineClient) Call(ctx context.Context, method string, request, response interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (kc *serviceInlineClient) richRemoteOption() { _ = "STUB: not implemented"; return }

func (kc *serviceInlineClient) buildInvokeChain() error { _ = "STUB: not implemented"; return nil }

func (kc *serviceInlineClient) invokeHandleEndpoint() (endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint), nil
}

func (kc *serviceInlineClient) Close() error { _ = "STUB: not implemented"; return nil }
