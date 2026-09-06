package loadbalance

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/pkg/discovery"
)

var (
	_ Picker = &WeightedRoundRobinPicker{}
	_ Picker = &RoundRobinPicker{}
)

const wrrVNodesBatchSize = 500

type wrrNode struct {
	discovery.Instance
	current int
}

func newWeightedRoundRobinPicker(instances []discovery.Instance) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

type WeightedRoundRobinPicker struct {
	nodes []*wrrNode
	size  uint64

	iterator  *round
	vsize     uint64
	vcapacity uint64
	vnodes    []discovery.Instance
	vlock     sync.RWMutex
}

func (wp *WeightedRoundRobinPicker) Next(ctx context.Context, request interface{}) (ins discovery.Instance) {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}

func (wp *WeightedRoundRobinPicker) buildVirtualWrrNodes(vtarget uint64) {
	_ = "STUB: not implemented"
	return
}

func nextWrrNode(nodes []*wrrNode) (selected *wrrNode) { _ = "STUB: not implemented"; return nil }

type RoundRobinPicker struct {
	size      uint64
	instances []discovery.Instance
	iterator  *round
}

func newRoundRobinPicker(instances []discovery.Instance) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (rp *RoundRobinPicker) Next(ctx context.Context, request interface{}) (ins discovery.Instance) {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}

func gcdInt(a, b int) int { _ = "STUB: not implemented"; return 0 }
