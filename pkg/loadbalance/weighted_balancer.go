package loadbalance

import (
	"sync"

	"golang.org/x/sync/singleflight"

	"github.com/cloudwego/kitex/pkg/discovery"
)

const (
	lbKindRoundRobin = iota
	lbKindInterleaved
	lbKindRandom
	lbKindRandomWithAliasMethod
)

type weightedBalancer struct {
	kind        int
	pickerCache sync.Map
	sfg         singleflight.Group
}

func NewWeightedBalancer() Loadbalancer { _ = "STUB: not implemented"; return *new(Loadbalancer) }

func NewWeightedRoundRobinBalancer() Loadbalancer {
	_ = "STUB: not implemented"
	return *new(Loadbalancer)
}

func NewInterleavedWeightedRoundRobinBalancer() Loadbalancer {
	_ = "STUB: not implemented"
	return *new(Loadbalancer)
}

func NewWeightedRandomBalancer() Loadbalancer { _ = "STUB: not implemented"; return *new(Loadbalancer) }

func NewWeightedRandomWithAliasMethodBalancer() Loadbalancer {
	_ = "STUB: not implemented"
	return *new(Loadbalancer)
}

func (wb *weightedBalancer) GetPicker(e discovery.Result) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (wb *weightedBalancer) createPicker(e discovery.Result) (picker Picker) {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (wb *weightedBalancer) Rebalance(change discovery.Change) { _ = "STUB: not implemented"; return }

func (wb *weightedBalancer) Delete(change discovery.Change) { _ = "STUB: not implemented"; return }

func (wb *weightedBalancer) Name() string { _ = "STUB: not implemented"; return "" }
