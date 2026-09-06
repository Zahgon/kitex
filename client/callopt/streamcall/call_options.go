package streamcall

import (
	"time"

	"github.com/cloudwego/kitex/pkg/streaming"
)

func WithHostPort(hostPort string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnectTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTag(key, val string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRecvTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRecvTimeoutConfig(cfg streaming.TimeoutConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBinaryGenericIDLService(svcName string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
