package client

import (
	"context"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/internal/client"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/cep"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/loadbalance/lbcache"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/rpcinfo/remoteinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/transport"
)

type Client interface {
	Call(ctx context.Context, method string, request, response interface{}) error
}

type kClient struct {
	svcInfo *serviceinfo.ServiceInfo
	eps     endpoint.UnaryEndpoint
	sEps    cep.StreamEndpoint

	opt *client.Options
	lbf *lbcache.BalancerFactory

	inited bool
	closed bool
}

type kcFinalizerClient struct {
	*kClient
}

func (kf *kcFinalizerClient) Call(ctx context.Context, method string, request, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClient(svcInfo *serviceinfo.ServiceInfo, opts ...Option) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (kc *kClient) init() (err error) {
	initTransportProtocol(kc.svcInfo, kc.opt.Configs)
	if err = kc.checkOptions(); err != nil {
		return err
	}
	if err = kc.initCircuitBreaker(); err != nil {
		return err
	}
	if err = kc.initRetryer(); err != nil {
		return err
	}
	if err = kc.initProxy(); err != nil {
		return err
	}
	if err = kc.initConnPool(); err != nil {
		return err
	}
	if err = kc.initLBCache(); err != nil {
		return err
	}
	ctx := kc.initContext()
	mw := kc.initMiddlewares(ctx)
	kc.initDebugService()
	kc.richRemoteOption()
	if err = kc.buildInvokeChain(mw); err != nil {
		return err
	}
	if err = kc.warmingUp(); err != nil {
		return err
	}
	kc.inited = true
	return nil
}

func (kc *kClient) checkOptions() (err error) { _ = "STUB: not implemented"; return nil }

func (kc *kClient) initCircuitBreaker() error { _ = "STUB: not implemented"; return nil }

func (kc *kClient) initRetryer() error { _ = "STUB: not implemented"; return nil }

func (kc *kClient) initContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (kc *kClient) initProxy() error { _ = "STUB: not implemented"; return nil }

func (kc *kClient) initConnPool() error { _ = "STUB: not implemented"; return nil }

func (kc *kClient) initLBCache() error { _ = "STUB: not implemented"; return nil }

type middleware struct {
	mws  []endpoint.Middleware
	uMws []endpoint.UnaryMiddleware

	smws []endpoint.Middleware
	sMws []cep.StreamMiddleware
}

func (kc *kClient) initMiddlewares(ctx context.Context) (mw middleware) {
	_ = "STUB: not implemented"
	return *new(middleware)
}

func richMWsWithBuilder(ctx context.Context, mwBs []endpoint.MiddlewareBuilder) (mws []endpoint.Middleware) {
	_ = "STUB: not implemented"
	return nil
}

func (kc *kClient) initRPCInfo(ctx context.Context, method string, retryTimes int,
	firstRI rpcinfo.RPCInfo, streamCall bool,
) (context.Context, rpcinfo.RPCInfo, *callopt.CallOptions) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(rpcinfo.RPCInfo), nil
}

func applyCallOptions(ctx context.Context, cfg rpcinfo.MutableRPCConfig, svr remoteinfo.RemoteInfo, opt *client.Options) (context.Context, *callopt.CallOptions) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (kc *kClient) Call(ctx context.Context, method string, request, response interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (kc *kClient) rpcCallWithRetry(ri rpcinfo.RPCInfo, method string) retry.RPCCallFunc {
	_ = "STUB: not implemented"
	return *new(retry.RPCCallFunc)
}

func (kc *kClient) initDebugService() { _ = "STUB: not implemented"; return }

func (kc *kClient) richRemoteOption() { _ = "STUB: not implemented"; return }

func (kc *kClient) buildInvokeChain(mw middleware) error { _ = "STUB: not implemented"; return nil }

func (kc *kClient) invokeHandleEndpoint() (endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint), nil
}

func (kc *kClient) Close() error { _ = "STUB: not implemented"; return nil }

func newCliTransHandler(opt *remote.ClientOption) (remote.ClientTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandler), nil
}

func initTransportProtocol(svcInfo *serviceinfo.ServiceInfo, cfg rpcinfo.RPCConfig) {
	_ = "STUB: not implemented"
	return
}

func (kc *kClient) warmingUp() error { _ = "STUB: not implemented"; return nil }

func validateForCall(ctx context.Context, inited, closed bool) { _ = "STUB: not implemented"; return }

func getCalloptRetryPolicy(callOpts *callopt.CallOptions) (callOptRetry *retry.Policy) {
	_ = "STUB: not implemented"
	return nil
}

func doFallbackIfNeeded(ctx context.Context, ri rpcinfo.RPCInfo, request, response interface{}, oriErr error, cliFallback *fallback.Policy, callOpts *callopt.CallOptions) (err, reportErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFallbackPolicy(cliOptFB *fallback.Policy, callOpts *callopt.CallOptions) (fb *fallback.Policy, hasFallback bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func purifyProtocol(cfg rpcinfo.MutableRPCConfig, tp transport.Protocol, streamCall bool) {
	_ = "STUB: not implemented"
	return
}

func initRPCInfo(ctx context.Context, method string, opt *client.Options, svcInfo *serviceinfo.ServiceInfo,
	retryTimes int, firstRI rpcinfo.RPCInfo, streamCall bool,
) (context.Context, rpcinfo.RPCInfo, *callopt.CallOptions) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(rpcinfo.RPCInfo), nil
}
