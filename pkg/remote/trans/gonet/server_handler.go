package gonet

import (
	"github.com/cloudwego/kitex/pkg/remote"
)

type svrTransHandlerFactory struct{}

func NewSvrTransHandlerFactory() remote.ServerTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandlerFactory)
}

func (f *svrTransHandlerFactory) NewTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}

func newSvrTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}
