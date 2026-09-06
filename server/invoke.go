package server

import (
	"github.com/cloudwego/kitex/pkg/remote/trans/invoke"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

type InvokeCaller interface {
	Call(invoke.Message) error
}

type Invoker interface {
	RegisterService(svcInfo *serviceinfo.ServiceInfo, handler interface{}, opts ...RegisterOption) error
	Init() (err error)
	InvokeCaller
}

type tInvoker struct {
	invoke.Handler
	*server
}

func NewInvoker(opts ...Option) Invoker { _ = "STUB: not implemented"; return *new(Invoker) }

func (s *tInvoker) Init() (err error) { _ = "STUB: not implemented"; return nil }

func (s *tInvoker) Call(msg invoke.Message) error { _ = "STUB: not implemented"; return nil }

func (s *tInvoker) newInvokeHandler() (handler invoke.Handler, err error) {
	_ = "STUB: not implemented"
	return *new(invoke.Handler), nil
}
