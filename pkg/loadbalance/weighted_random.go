package loadbalance

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
)

type weightedRandomPicker struct {
	instances []discovery.Instance
	weightSum int
}

func newWeightedRandomPickerWithSum(instances []discovery.Instance, weightSum int) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (wp *weightedRandomPicker) Next(ctx context.Context, request interface{}) (ins discovery.Instance) {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}

type randomPicker struct {
	instances []discovery.Instance
}

func newRandomPicker(instances []discovery.Instance) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (rp *randomPicker) Next(ctx context.Context, request interface{}) (ins discovery.Instance) {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}
