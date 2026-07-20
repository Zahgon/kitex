package client

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/event"
	"github.com/cloudwego/kitex/pkg/loadbalance/lbcache"
	"github.com/cloudwego/kitex/pkg/proxy"
)

const maxRetry = 6

func newProxyMW(prx proxy.ForwardProxy) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

type instSnapshot struct {
	network string
	address string
	weight  int
}

type discoveryEventExtra struct {
	cacheable bool
	cacheKey  string
	added     []instSnapshot
	updated   []instSnapshot
	removed   []instSnapshot
}

func (e *discoveryEventExtra) KitexDumpLazyExtra() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func discoveryEventHandler(name string, bus event.Bus, queue event.Queue) func(d *discovery.Change) {
	_ = "STUB: not implemented"
	return nil
}

func newResolveMWBuilder(lbf *lbcache.BalancerFactory) endpoint.MiddlewareBuilder {
	_ = "STUB: not implemented"
	return *new(endpoint.MiddlewareBuilder)
}

func newIOErrorHandleMW(errHandle func(context.Context, error) error) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func isRemoteErr(err error) bool { _ = "STUB: not implemented"; return false }

func DefaultClientErrorHandler(ctx context.Context, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func ClientErrorHandlerWithAddr(ctx context.Context, err error) error {
	_ = "STUB: not implemented"
	return nil
}

type instInfo struct {
	Address string
	Weight  int
}

func snapshotInstances(insts []discovery.Instance) []instSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func formatSnapshots(snapshots []instSnapshot) []*instInfo { _ = "STUB: not implemented"; return nil }

func retryable(err error) bool { _ = "STUB: not implemented"; return false }

func getRemoteAddr(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
