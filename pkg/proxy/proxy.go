package proxy

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type Config struct {
	ServerInfo   *rpcinfo.EndpointBasicInfo
	Resolver     discovery.Resolver
	Balancer     loadbalance.Loadbalancer
	Pool         remote.ConnPool
	FixedTargets string
	RPCConfig    rpcinfo.RPCConfig
}

type ForwardProxy interface {
	Configure(*Config) error

	ResolveProxyInstance(ctx context.Context) error
}

type WithMiddleware interface {
	ProxyMiddleware() endpoint.Middleware
}

type ReverseProxy interface {
	Replace(net.Addr) (net.Addr, error)
}

type BackwardProxy = ReverseProxy

type ContextHandler interface {
	HandleContext(context.Context) context.Context
}
