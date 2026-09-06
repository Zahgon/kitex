package loadbalance

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/pkg/discovery"
)

type iwrrNode struct {
	discovery.Instance
	remainder int

	next *iwrrNode
}

type iwrrQueue struct {
	head *iwrrNode
	tail *iwrrNode
}

type InterleavedWeightedRoundRobinPicker struct {
	current *iwrrQueue
	next    *iwrrQueue
	gcd     int

	lock sync.Mutex
}

func newInterleavedWeightedRoundRobinPicker(instances []discovery.Instance) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (ip *InterleavedWeightedRoundRobinPicker) Next(ctx context.Context, request interface{}) discovery.Instance {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}

func newIwrrQueue() *iwrrQueue { _ = "STUB: not implemented"; return nil }

func (q *iwrrQueue) enqueue(node *iwrrNode) { _ = "STUB: not implemented"; return }

func (q *iwrrQueue) dequeue() *iwrrNode { _ = "STUB: not implemented"; return nil }

func (q *iwrrQueue) empty() bool { _ = "STUB: not implemented"; return false }
