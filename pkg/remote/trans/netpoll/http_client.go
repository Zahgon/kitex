package netpoll

import (
	"github.com/cloudwego/kitex/pkg/remote"
)

type httpCliTransHandlerFactory struct{}

func NewHTTPCliTransHandlerFactory() remote.ClientTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandlerFactory)
}

func (f *httpCliTransHandlerFactory) NewTransHandler(opt *remote.ClientOption) (remote.ClientTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandler), nil
}
