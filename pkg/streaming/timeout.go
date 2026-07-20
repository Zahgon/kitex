package streaming

import (
	"context"
	"time"
)

type TimeoutConfig struct {
	Timeout time.Duration

	DisableCancelRemote bool
}

func CallWithTimeout(timeout time.Duration, cancel context.CancelFunc, f func() (err error)) error {
	_ = "STUB: not implemented"
	return nil
}
