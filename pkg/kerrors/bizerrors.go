package kerrors

import (
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
)

type BizStatusErrorIface interface {
	BizStatusCode() int32
	BizMessage() string
	BizExtra() map[string]string
	Error() string
}

type GRPCStatusIface interface {
	GRPCStatus() *status.Status
	SetGRPCStatus(status *status.Status)
}

type BizStatusError struct {
	code  int32
	msg   string
	extra map[string]string
}

func FromBizStatusError(err error) (bizErr BizStatusErrorIface, ok bool) {
	_ = "STUB: not implemented"
	return *new(BizStatusErrorIface), false
}

func NewBizStatusError(code int32, msg string) BizStatusErrorIface {
	_ = "STUB: not implemented"
	return *new(BizStatusErrorIface)
}

func NewBizStatusErrorWithExtra(code int32, msg string, extra map[string]string) BizStatusErrorIface {
	_ = "STUB: not implemented"
	return *new(BizStatusErrorIface)
}

func (e *BizStatusError) BizStatusCode() int32 { _ = "STUB: not implemented"; return 0 }

func (e *BizStatusError) BizMessage() string { _ = "STUB: not implemented"; return "" }

func (e *BizStatusError) AppendBizMessage(extraMsg string) { _ = "STUB: not implemented"; return }

func (e *BizStatusError) BizExtra() map[string]string { _ = "STUB: not implemented"; return nil }

func (e *BizStatusError) SetBizExtra(key, value string) { _ = "STUB: not implemented"; return }

func (e *BizStatusError) Error() string { _ = "STUB: not implemented"; return "" }

func NewGRPCBizStatusError(code int32, msg string) BizStatusErrorIface {
	_ = "STUB: not implemented"
	return *new(BizStatusErrorIface)
}

func NewGRPCBizStatusErrorWithExtra(code int32, msg string, extra map[string]string) BizStatusErrorIface {
	_ = "STUB: not implemented"
	return *new(BizStatusErrorIface)
}

type GRPCBizStatusError struct {
	BizStatusError

	status *status.Status
}

func (e *GRPCBizStatusError) GRPCStatus() *status.Status { _ = "STUB: not implemented"; return nil }

func (e *GRPCBizStatusError) SetGRPCStatus(status *status.Status) {
	_ = "STUB: not implemented"
	return
}
