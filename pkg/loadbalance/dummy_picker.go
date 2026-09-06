package loadbalance

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
)

var _ Picker = &DummyPicker{}

type DummyPicker struct{}

func (np *DummyPicker) Next(ctx context.Context, request interface{}) (ins discovery.Instance) {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}
