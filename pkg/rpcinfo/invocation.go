package rpcinfo

import (
	"sync"
	"sync/atomic"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

const InvocationServiceInfoKey = "service_info_key"

var (
	_              Invocation       = (*invocation)(nil)
	_              InvocationSetter = (*invocation)(nil)
	invocationPool sync.Pool
	globalSeqID    int32 = 0
)

func init() {
	invocationPool.New = newInvocation
}

type InvocationSetter interface {
	SetPackageName(name string)
	SetServiceName(name string)
	SetMethodName(name string)
	SetMethodInfo(methodInfo serviceinfo.MethodInfo)
	SetStreamingMode(mode serviceinfo.StreamingMode)
	SetSeqID(seqID int32)
	SetBizStatusErr(err kerrors.BizStatusErrorIface)
	SetExtra(key string, value interface{})
	Reset()
}
type invocation struct {
	packageName   string
	serviceName   string
	methodInfo    serviceinfo.MethodInfo
	methodName    string
	streamingMode serviceinfo.StreamingMode
	seqID         int32

	bizErr atomic.Pointer[kerrors.BizStatusErrorIface]

	mu    sync.Mutex
	extra map[string]any
}

func NewInvocation(service, method string, pkgOpt ...string) *invocation {
	_ = "STUB: not implemented"
	return nil
}

func NewServerInvocation() Invocation { _ = "STUB: not implemented"; return *new(Invocation) }

func genSeqID() int32 { _ = "STUB: not implemented"; return 0 }

func newInvocation() interface{} { _ = "STUB: not implemented"; return nil }

func (i *invocation) SeqID() int32 { _ = "STUB: not implemented"; return 0 }

func (i *invocation) SetSeqID(seqID int32) { _ = "STUB: not implemented"; return }

func (i *invocation) PackageName() string { _ = "STUB: not implemented"; return "" }

func (i *invocation) SetPackageName(name string) { _ = "STUB: not implemented"; return }

func (i *invocation) ServiceName() string { _ = "STUB: not implemented"; return "" }

func (i *invocation) SetServiceName(name string) { _ = "STUB: not implemented"; return }

func (i *invocation) MethodName() string { _ = "STUB: not implemented"; return "" }

func (i *invocation) SetMethodName(name string) { _ = "STUB: not implemented"; return }

func (i *invocation) MethodInfo() serviceinfo.MethodInfo {
	_ = "STUB: not implemented"
	return *new(serviceinfo.MethodInfo)
}

func (i *invocation) SetMethodInfo(methodInfo serviceinfo.MethodInfo) {
	_ = "STUB: not implemented"
	return
}

func (i *invocation) StreamingMode() serviceinfo.StreamingMode {
	_ = "STUB: not implemented"
	return *new(serviceinfo.StreamingMode)
}

func (i *invocation) SetStreamingMode(mode serviceinfo.StreamingMode) {
	_ = "STUB: not implemented"
	return
}

func (i *invocation) BizStatusErr() kerrors.BizStatusErrorIface {
	_ = "STUB: not implemented"
	return *new(kerrors.BizStatusErrorIface)
}

func (i *invocation) SetBizStatusErr(err kerrors.BizStatusErrorIface) {
	_ = "STUB: not implemented"
	return
}

func (i *invocation) SetExtra(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (i *invocation) Extra(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (i *invocation) Reset() { _ = "STUB: not implemented"; return }

func (i *invocation) Recycle() { _ = "STUB: not implemented"; return }

func (i *invocation) zero() { _ = "STUB: not implemented"; return }
