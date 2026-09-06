package ttstream

import (
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/streaming"
)

var (
	errApplicationException = newException("application exception", nil, 12001)
	errUnexpectedHeader     = newException("unexpected header frame", kerrors.ErrStreamingProtocol, 12002)
	errIllegalBizErr        = newException("illegal bizErr", kerrors.ErrStreamingProtocol, 12003)
	errIllegalFrame         = newException("illegal frame", kerrors.ErrStreamingProtocol, 12004)
	errIllegalOperation     = newException("illegal operation", kerrors.ErrStreamingProtocol, 12005)
	errTransport            = newException("transport is closing", kerrors.ErrStreamingProtocol, 12006)

	errBizCancel = newException("user code invoking stream RPC with context processed by context.WithCancel or context.WithTimeout, then invoking cancel() actively",
		kerrors.ErrStreamingCanceled, 12007)
	errBizCancelWithCause     = newException("user code canceled with cancelCause(error)", kerrors.ErrStreamingCanceled, 12008)
	errDownstreamCancel       = newException("canceled by downstream", kerrors.ErrStreamingCanceled, 12009)
	errUpstreamCancel         = newException("canceled by upstream", kerrors.ErrStreamingCanceled, 12010)
	errInternalCancel         = newException("internal canceled", kerrors.ErrStreamingCanceled, 12011)
	errBizHandlerReturnCancel = newException("canceled by business handler returning", kerrors.ErrStreamingCanceled, 12012)
	errConnectionClosedCancel = newException("canceled by connection closed", kerrors.ErrStreamingCanceled, 12013)
)

var errServerSideBizHandlerReturnCancel = errBizHandlerReturnCancel.newBuilder().withSide(serverSide)

func newStreamRecvTimeoutException(cfg streaming.TimeoutConfig) *Exception {
	_ = "STUB: not implemented"
	return nil
}

var notSetStreamTimeout = time.Duration(-1)

func newStreamTimeoutException(tm time.Duration) *Exception { _ = "STUB: not implemented"; return nil }

const (
	setSide = 1 << iota
	setCancelPath
	setCause
	setCanRetry
)

type Exception struct {
	message string
	typeId  int32

	side       sideType
	cancelPath string

	parent error

	cause error

	canRetry bool

	bitSet uint8
}

func newException(message string, parent error, typeId int32) *Exception {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exception) newBuilder() *Exception { _ = "STUB: not implemented"; return nil }

func (e *Exception) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Exception) withCause(cause error) *Exception { _ = "STUB: not implemented"; return nil }

func (e *Exception) withCauseAndTypeId(cause error, typeId int32) *Exception {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exception) isCauseSet() bool { _ = "STUB: not implemented"; return false }

func (e *Exception) withSide(side sideType) *Exception { _ = "STUB: not implemented"; return nil }

func (e *Exception) isSideSet() bool { _ = "STUB: not implemented"; return false }

func (e *Exception) setOrAppendCancelPath(cancelPath string) *Exception {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exception) isCancelPathSet() bool { _ = "STUB: not implemented"; return false }

func (e *Exception) withCanRetry() *Exception { _ = "STUB: not implemented"; return nil }

func (e *Exception) isCanRetrySet() bool { _ = "STUB: not implemented"; return false }

func (e *Exception) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *Exception) getMessage() string { _ = "STUB: not implemented"; return "" }

func (e *Exception) TypeId() int32 { _ = "STUB: not implemented"; return 0 }

func appendCancelPath(oriCp, node string) string { _ = "STUB: not implemented"; return "" }

func formatCancelPath(cancelPath string) string { _ = "STUB: not implemented"; return "" }

func checkCanRetry(err error) bool { _ = "STUB: not implemented"; return false }
