package invoke

import (
	"github.com/cloudwego/kitex/pkg/remote"
)

type ivkTransHandlerFactory struct{}

func NewIvkTransHandlerFactory() remote.ServerTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandlerFactory)
}

func (f *ivkTransHandlerFactory) NewTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}

func newIvkTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}
