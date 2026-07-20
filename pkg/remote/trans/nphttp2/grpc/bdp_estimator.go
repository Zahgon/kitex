package grpc

import (
	"sync"
	"time"
)

const (
	bdpLimit = (1 << 20) * 16

	alpha = 0.9

	beta = 0.66

	gamma = 2
)

var bdpPing = &ping{data: [8]byte{2, 4, 16, 16, 9, 14, 7, 7}}

var bdpPingInterval = time.Second

type bdpEstimator struct {
	sentAt time.Time

	mu sync.Mutex

	bdp uint32

	sample uint32

	bwMax float64

	isSent bool

	updateFlowControl func(n uint32)

	sampleCount uint64

	rtt float64
}

func (b *bdpEstimator) timesnap(d [8]byte) { _ = "STUB: not implemented"; return }

func (b *bdpEstimator) add(n uint32) bool { _ = "STUB: not implemented"; return false }

func (b *bdpEstimator) calculate(d [8]byte) { _ = "STUB: not implemented"; return }
