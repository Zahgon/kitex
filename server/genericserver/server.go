package genericserver

import (
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/server"
)

func NewServer(handler generic.Service, g generic.Generic, opts ...server.Option) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

func NewServerWithServiceInfo(handler generic.Service, g generic.Generic, svcInfo *serviceinfo.ServiceInfo, opts ...server.Option) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

func NewServerV2(handler *generic.ServiceV2, g generic.Generic, opts ...server.Option) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

func RegisterService(svr server.Server, handler *generic.ServiceV2, g generic.Generic, opts ...server.RegisterOption) error {
	_ = "STUB: not implemented"
	return nil
}
