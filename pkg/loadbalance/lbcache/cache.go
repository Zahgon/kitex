package lbcache

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/cloudwego/kitex/pkg/diagnosis"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
)

const (
	defaultRefreshInterval = 5 * time.Second
	defaultExpireInterval  = 15 * time.Second
)

var (
	balancerFactories    sync.Map
	balancerFactoriesSfg singleflight.Group
)

type Options struct {
	RefreshInterval time.Duration

	ExpireInterval time.Duration

	DiagnosisService diagnosis.Service

	Cacheable bool
}

func (v *Options) check() { _ = "STUB: not implemented"; return }

type Hookable interface {
	RegisterRebalanceHook(func(ch *discovery.Change)) (index int)
	DeregisterRebalanceHook(index int)

	RegisterDeleteHook(func(ch *discovery.Change)) (index int)
	DeregisterDeleteHook(index int)
}

type BalancerFactory struct {
	Hookable
	opts       Options
	cache      sync.Map
	resolver   discovery.Resolver
	balancer   loadbalance.Loadbalancer
	rebalancer loadbalance.Rebalancer
	sfg        singleflight.Group
}

func cacheKey(resolver, balancer string, opts Options) string { _ = "STUB: not implemented"; return "" }

func newBalancerFactory(resolver discovery.Resolver, balancer loadbalance.Loadbalancer, opts Options) *BalancerFactory {
	_ = "STUB: not implemented"
	return nil
}

func NewBalancerFactory(resolver discovery.Resolver, balancer loadbalance.Loadbalancer, opts Options) *BalancerFactory {
	_ = "STUB: not implemented"
	return nil
}

func (b *BalancerFactory) watcher() { _ = "STUB: not implemented"; return }

func renameResultCacheKey(res *discovery.Result, resolverName string) {
	_ = "STUB: not implemented"
	return
}

func (b *BalancerFactory) Get(ctx context.Context, target rpcinfo.EndpointInfo) (*Balancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Balancer struct {
	b            *BalancerFactory
	target       string
	res          atomic.Value
	expire       int32
	sharedTicker *utils.SharedTicker
}

func (bl *Balancer) Refresh() { _ = "STUB: not implemented"; return }

func (bl *Balancer) Tick() { _ = "STUB: not implemented"; return }

func (bl *Balancer) GetResult() (res discovery.Result, ok bool) {
	_ = "STUB: not implemented"
	return *new(discovery.Result), false
}

func (bl *Balancer) GetPicker() loadbalance.Picker {
	_ = "STUB: not implemented"
	return *new(loadbalance.Picker)
}

func (bl *Balancer) close() { _ = "STUB: not implemented"; return }

const unknown = "unknown"

func Dump() interface{} { _ = "STUB: not implemented"; return nil }
