package grpc

import (
	"math"
	"time"

	"github.com/cloudwego/kitex"
)

const (
	defaultWindowSize = uint32(65535)

	initialWindowSize = defaultWindowSize

	Infinity                     = time.Duration(math.MaxInt64)
	defaultMaxStreamsClient      = 100
	defaultMaxConnectionIdle     = Infinity
	defaultMaxConnectionAge      = Infinity
	defaultMaxConnectionAgeGrace = Infinity

	defaultClientKeepaliveTime    = Infinity
	defaultClientKeepaliveTimeout = 20 * time.Second
	defaultServerKeepaliveTime    = 2 * time.Hour
	defaultServerKeepaliveTimeout = 20 * time.Second
	defaultKeepalivePolicyMinTime = 5 * time.Minute

	maxWindowSize = math.MaxInt32

	defaultWriteQuota              = 64 * 1024
	defaultClientMaxHeaderListSize = uint32(16 << 20)
	defaultServerMaxHeaderListSize = uint32(16 << 20)

	defaultWriteBufferSize = uint32(32 * 1024)

	defaultReadBufferSize = uint32(32 * 1024)

	defaultUserAgent = "kitex/" + kitex.Version
)

const (
	KeepaliveMinPingTime = 10 * time.Second
)
