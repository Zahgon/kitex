package loadbalance

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
)

type Picker interface {
	Next(ctx context.Context, request interface{}) discovery.Instance
}

type Loadbalancer interface {
	GetPicker(discovery.Result) Picker
	Name() string
}

type Rebalancer interface {
	Rebalance(discovery.Change)
	Delete(discovery.Change)
}
