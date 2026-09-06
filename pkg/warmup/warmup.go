package warmup

import (
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type ErrorHandling int

const (
	IgnoreError ErrorHandling = iota
	WarningLog
	ErrorLog
	FailFast
)

type ClientOption struct {
	ErrorHandling
	ResolverOption *ResolverOption
	PoolOption     *PoolOption
}

type ResolverOption struct {
	Dests []*rpcinfo.EndpointBasicInfo
}

type PoolOption struct {
	Targets  map[string][]string
	ConnNum  int
	Parallel int
}

type Pool interface {
	WarmUp(eh ErrorHandling, wuo *PoolOption, co remote.ConnOption) error
}
