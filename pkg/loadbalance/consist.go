package loadbalance

import (
	"context"
	"hash/maphash"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/cloudwego/kitex/pkg/discovery"
)

var hashSeed = maphash.MakeSeed()

type KeyFunc func(ctx context.Context, request interface{}) string

type ConsistentHashOption struct {
	GetKey KeyFunc

	Replica uint32

	VirtualFactor uint32

	Weighted bool

	ExpireDuration time.Duration
}

func NewConsistentHashOption(f KeyFunc) ConsistentHashOption {
	_ = "STUB: not implemented"
	return *new(ConsistentHashOption)
}

var consistPickerPool sync.Pool

func init() {
	consistPickerPool.New = newConsistPicker
}

type virtualNode struct {
	hash     uint64
	RealNode *realNode
}

type realNode struct {
	Ins discovery.Instance
}

type consistResult struct {
	Primary  discovery.Instance
	Replicas []discovery.Instance
}

type consistInfo struct {
	realNodes    []realNode
	virtualNodes []virtualNode
}

type vNodeType struct {
	s []virtualNode
}

func (v *vNodeType) Len() int { _ = "STUB: not implemented"; return 0 }

func (v *vNodeType) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (v *vNodeType) Swap(i, j int) { _ = "STUB: not implemented"; return }

type consistPicker struct {
	cb   *consistBalancer
	info *consistInfo
}

func newConsistPicker() interface{} { _ = "STUB: not implemented"; return nil }

func (cp *consistPicker) zero() { _ = "STUB: not implemented"; return }

func (cp *consistPicker) Recycle() { _ = "STUB: not implemented"; return }

func (cp *consistPicker) Next(ctx context.Context, request interface{}) discovery.Instance {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}

func buildConsistResult(info *consistInfo, key uint64) *consistResult {
	_ = "STUB: not implemented"
	return nil
}

type consistBalancer struct {
	cachedConsistInfo sync.Map
	opt               ConsistentHashOption
	sfg               singleflight.Group
}

func NewConsistBalancer(opt ConsistentHashOption) Loadbalancer {
	_ = "STUB: not implemented"
	return *new(Loadbalancer)
}

func (cb *consistBalancer) GetPicker(e discovery.Result) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (cb *consistBalancer) newConsistInfo(e discovery.Result) *consistInfo {
	_ = "STUB: not implemented"
	return nil
}

func (cb *consistBalancer) buildNodes(ins []discovery.Instance) ([]realNode, []virtualNode) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *consistBalancer) buildVirtualNodes(rNodes []realNode) []virtualNode {
	_ = "STUB: not implemented"
	return nil
}

func (cb *consistBalancer) getVirtualNodeLen(rNode realNode) int {
	_ = "STUB: not implemented"
	return 0
}

func (cb *consistBalancer) updateConsistInfo(e discovery.Result) { _ = "STUB: not implemented"; return }

func (cb *consistBalancer) Rebalance(change discovery.Change) { _ = "STUB: not implemented"; return }

func (cb *consistBalancer) Delete(change discovery.Change) { _ = "STUB: not implemented"; return }

func (cb *consistBalancer) Name() string { _ = "STUB: not implemented"; return "" }
