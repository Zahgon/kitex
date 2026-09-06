package loadbalance

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
)

type AliasMethodPicker struct {
	instances []discovery.Instance
	weightSum int
	alias     []int
	prob      []float64
}

func newAliasMethodPicker(instances []discovery.Instance, weightSum int) Picker {
	_ = "STUB: not implemented"
	return *new(Picker)
}

func (a *AliasMethodPicker) init() {
	n := len(a.instances)
	a.alias = make([]int, n)
	a.prob = make([]float64, n)

	totalWeight := a.weightSum

	scaledProb := make([]float64, n)
	small := make([]int, 0, n)
	large := make([]int, 0, n)

	for i, instance := range a.instances {
		scaledProb[i] = float64(instance.Weight()) * float64(n) / float64(totalWeight)
		if scaledProb[i] < 1.0 {
			small = append(small, i)
		} else {
			large = append(large, i)
		}
	}

	for len(small) > 0 && len(large) > 0 {
		l := small[len(small)-1]
		small = small[:len(small)-1]
		g := large[len(large)-1]
		large = large[:len(large)-1]

		a.prob[l] = scaledProb[l]
		a.alias[l] = g

		scaledProb[g] -= 1.0 - scaledProb[l]
		if scaledProb[g] < 1.0 {
			small = append(small, g)
		} else {
			large = append(large, g)
		}
	}

	for len(large) > 0 {
		g := large[len(large)-1]
		large = large[:len(large)-1]
		a.prob[g] = 1.0
	}

	for len(small) > 0 {
		l := small[len(small)-1]
		small = small[:len(small)-1]
		a.prob[l] = 1.0
	}
}

func (a *AliasMethodPicker) Next(ctx context.Context, request interface{}) discovery.Instance {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}
