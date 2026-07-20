package retry

import (
	"context"
	"sync"

	"github.com/bytedance/gopkg/cloud/circuitbreaker"

	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type RPCCallFunc func(ctx context.Context, retryer Retryer, request, response interface{}) (rpcinfo rpcinfo.RPCInfo, err error)

type GenRetryKeyFunc func(ctx context.Context, ri rpcinfo.RPCInfo) string

type Retryer interface {
	AllowRetry(ctx context.Context) (msg string, ok bool)
	UpdatePolicy(policy Policy) error

	Do(ctx context.Context, rpcCall RPCCallFunc, firstRI rpcinfo.RPCInfo, request, response interface{}) (lastRI rpcinfo.RPCInfo, recycleRI bool, err error)
	AppendErrMsgIfNeeded(ctx context.Context, err error, ri rpcinfo.RPCInfo, msg string)

	Prepare(ctx context.Context, prevRI, retryRI rpcinfo.RPCInfo)
	Dump() map[string]interface{}
	Type() Type
}

func NewRetryContainerWithCB(cc *circuitbreak.Control, cp circuitbreaker.Panel) *Container {
	_ = "STUB: not implemented"
	return nil
}

func newCBSuite(opts []circuitbreak.CBSuiteOption) *circuitbreak.CBSuite {
	_ = "STUB: not implemented"
	return nil
}

func NewRetryContainerWithCBStat(cc *circuitbreak.Control, cp circuitbreaker.Panel) *Container {
	_ = "STUB: not implemented"
	return nil
}

func NewRetryContainerWithPercentageLimit() *Container { _ = "STUB: not implemented"; return nil }

type ContainerOption func(rc *Container)

func WithContainerCBSuite(cbs *circuitbreak.CBSuite) ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

func WithCustomizeKeyFunc(fn GenRetryKeyFunc) ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

func WithContainerCBSuiteOptions(opts ...circuitbreak.CBSuiteOption) ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

func WithContainerCBControl(ctrl *circuitbreak.Control) ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

func WithContainerCBPanel(panel circuitbreaker.Panel) ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

func WithContainerCBStat() ContainerOption { _ = "STUB: not implemented"; return *new(ContainerOption) }

func WithContainerEnablePercentageLimit() ContainerOption {
	_ = "STUB: not implemented"
	return *new(ContainerOption)
}

func NewRetryContainer(opts ...ContainerOption) *Container { _ = "STUB: not implemented"; return nil }

func defaultGenRetryKey(_ context.Context, rpcInfo rpcinfo.RPCInfo) string {
	_ = "STUB: not implemented"
	return ""
}

type Container struct {
	hasCodeCfg  bool
	retryerMap  sync.Map
	cbContainer *cbContainer
	msg         string
	sync.RWMutex

	genRetryKey GenRetryKeyFunc

	shouldResultRetry *ShouldResultRetry
}

type cbContainer struct {
	cbSuite *circuitbreak.CBSuite

	cbCtl   *circuitbreak.Control
	cbPanel circuitbreaker.Panel

	cbStat bool

	enablePercentageLimit bool

	cbSuiteOptions []circuitbreak.CBSuiteOption
}

func (c *cbContainer) IsValid() bool { _ = "STUB: not implemented"; return false }

func (rc *Container) InitWithPolicies(methodPolicies map[string]Policy) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *Container) DeletePolicy(key string) { _ = "STUB: not implemented"; return }

func (rc *Container) NotifyPolicyChange(key string, p Policy) { _ = "STUB: not implemented"; return }

func (rc *Container) Init(mp map[string]Policy, rr *ShouldResultRetry) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type retryContext struct {
	context.Context

	reqOp  int32
	respOp int32
}

func (p *retryContext) Value(k any) any { _ = "STUB: not implemented"; return *new(any) }

func PrepareRetryContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (rc *Container) WithRetryIfNeeded(ctx context.Context, callOptRetry *Policy, rpcCall RPCCallFunc, ri rpcinfo.RPCInfo, request, response interface{}) (lastRI rpcinfo.RPCInfo, recycleRI bool, err error) {
	_ = "STUB: not implemented"
	return *new(rpcinfo.RPCInfo), false, nil
}

func NewRetryer(p Policy, r *ShouldResultRetry, cbC *cbContainer) (retryer Retryer, err error) {
	_ = "STUB: not implemented"
	return *new(Retryer), nil
}

func (rc *Container) getRetryer(ctx context.Context, ri rpcinfo.RPCInfo) Retryer {
	_ = "STUB: not implemented"
	return *new(Retryer)
}

func (rc *Container) Dump() interface{} { _ = "STUB: not implemented"; return nil }

func (rc *Container) initRetryer(method string, p Policy) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *Container) updateRetryer(rr *ShouldResultRetry) { _ = "STUB: not implemented"; return }

func (rc *Container) Close() (err error) { _ = "STUB: not implemented"; return nil }
