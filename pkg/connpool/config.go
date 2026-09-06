package connpool

import "time"

type IdleConfig struct {
	MinIdlePerAddress int
	MaxIdlePerAddress int
	MaxIdleGlobal     int
	MaxIdleTimeout    time.Duration
}

const (
	defaultMaxIdleTimeout = 30 * time.Second
	minMaxIdleTimeout     = 2 * time.Second
	maxMinIdlePerAddress  = 5
	defaultMaxIdleGlobal  = 1 << 20
)

func CheckPoolConfig(config IdleConfig) *IdleConfig { _ = "STUB: not implemented"; return nil }
