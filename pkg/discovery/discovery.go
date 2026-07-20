package discovery

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

const DefaultWeight = 10

type Result struct {
	Cacheable bool
	CacheKey  string
	Instances []Instance
}

type Change struct {
	Result  Result
	Added   []Instance
	Updated []Instance
	Removed []Instance
}

type Resolver interface {
	Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string)

	Resolve(ctx context.Context, desc string) (Result, error)

	Diff(cacheKey string, prev, next Result) (Change, bool)

	Name() string
}

func DefaultDiff(cacheKey string, prev, next Result) (Change, bool) {
	_ = "STUB: not implemented"
	return *new(Change), false
}

type instance struct {
	addr   net.Addr
	weight int
	tags   map[string]string
}

func (i *instance) Address() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (i *instance) Weight() int { _ = "STUB: not implemented"; return 0 }

func (i *instance) Tag(key string) (value string, exist bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (i *instance) Tags() map[string]string { _ = "STUB: not implemented"; return nil }

func NewInstance(network, address string, weight int, tags map[string]string) Instance {
	_ = "STUB: not implemented"
	return *new(Instance)
}

type SynthesizedResolver struct {
	TargetFunc  func(ctx context.Context, target rpcinfo.EndpointInfo) string
	ResolveFunc func(ctx context.Context, key string) (Result, error)
	DiffFunc    func(key string, prev, next Result) (Change, bool)
	NameFunc    func() string
}

func (sr SynthesizedResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (sr SynthesizedResolver) Resolve(ctx context.Context, key string) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (sr SynthesizedResolver) Diff(key string, prev, next Result) (Change, bool) {
	_ = "STUB: not implemented"
	return *new(Change), false
}

func (sr SynthesizedResolver) Name() string { _ = "STUB: not implemented"; return "" }

type Instance interface {
	Address() net.Addr
	Weight() int
	Tag(key string) (value string, exist bool)
}
