package bound

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/limiter"
	"github.com/cloudwego/kitex/pkg/remote"
)

func NewServerLimiterHandler(conLimit limiter.ConcurrencyLimiter, qpsLimit limiter.RateLimiter, reporter limiter.LimitReporter, qpsLimitPostDecode bool) remote.InboundHandler {
	_ = "STUB: not implemented"
	return *new(remote.InboundHandler)
}

type serverLimiterHandler struct {
	connLimit          limiter.ConcurrencyLimiter
	qpsLimit           limiter.RateLimiter
	reporter           limiter.LimitReporter
	qpsLimitPostDecode bool
}

func (l *serverLimiterHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (l *serverLimiterHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (l *serverLimiterHandler) OnInactive(ctx context.Context, conn net.Conn) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (l *serverLimiterHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func DeepEqual(bound1, bound2 remote.InboundHandler) bool { _ = "STUB: not implemented"; return false }
