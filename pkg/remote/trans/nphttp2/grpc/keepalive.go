package grpc

import (
	"time"
)

type ClientKeepalive struct {
	Time time.Duration

	Timeout time.Duration

	PermitWithoutStream bool
}

type ServerKeepalive struct {
	MaxConnectionIdle time.Duration

	MaxConnectionAge time.Duration

	MaxConnectionAgeGrace time.Duration

	Time time.Duration

	Timeout time.Duration
}

type EnforcementPolicy struct {
	MinTime time.Duration

	PermitWithoutStream bool
}
