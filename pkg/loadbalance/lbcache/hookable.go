package lbcache

import (
	"sync"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

type hookableRebalancer struct {
	inner          loadbalance.Rebalancer
	rebalanceL     sync.Mutex
	rebalanceIndex int
	rebalanceHooks map[int]func(*discovery.Change)
	deleteL        sync.Mutex
	deleteIndex    int
	deleteHooks    map[int]func(*discovery.Change)
}

var (
	_ loadbalance.Rebalancer = (*hookableRebalancer)(nil)
	_ Hookable               = (*hookableRebalancer)(nil)
)

func newHookRebalancer(inner loadbalance.Rebalancer) *hookableRebalancer {
	_ = "STUB: not implemented"
	return nil
}

func (b *hookableRebalancer) Rebalance(ch discovery.Change) { _ = "STUB: not implemented"; return }

func (b *hookableRebalancer) Delete(ch discovery.Change) { _ = "STUB: not implemented"; return }

func (b *hookableRebalancer) RegisterRebalanceHook(f func(ch *discovery.Change)) int {
	_ = "STUB: not implemented"
	return 0
}

func (b *hookableRebalancer) DeregisterRebalanceHook(index int) { _ = "STUB: not implemented"; return }

func (b *hookableRebalancer) RegisterDeleteHook(f func(ch *discovery.Change)) int {
	_ = "STUB: not implemented"
	return 0
}

func (b *hookableRebalancer) DeregisterDeleteHook(index int) { _ = "STUB: not implemented"; return }

type noopHookRebalancer struct{}

var _ Hookable = (*noopHookRebalancer)(nil)

func (noopHookRebalancer) RegisterRebalanceHook(func(ch *discovery.Change)) int {
	_ = "STUB: not implemented"
	return 0
}
func (noopHookRebalancer) RegisterDeleteHook(func(ch *discovery.Change)) int {
	_ = "STUB: not implemented"
	return 0
}
func (noopHookRebalancer) DeregisterRebalanceHook(index int) { _ = "STUB: not implemented"; return }
func (noopHookRebalancer) DeregisterDeleteHook(index int)    { _ = "STUB: not implemented"; return }
