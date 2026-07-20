package gonet

import (
	"github.com/cloudwego/kitex/pkg/remote"
)

type cliTransHandlerFactory struct{}

func NewCliTransHandlerFactory() remote.ClientTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandlerFactory)
}

func (f *cliTransHandlerFactory) NewTransHandler(opt *remote.ClientOption) (remote.ClientTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandler), nil
}

func newCliTransHandler(opt *remote.ClientOption) (remote.ClientTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandler), nil
}
