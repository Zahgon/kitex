package test

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
)

func newGenericClient(destService string, g generic.Generic, targetIPPort string) genericclient.Client {
	_ = "STUB: not implemented"
	return *new(genericclient.Client)
}

func newGenericServer(g generic.Generic, addr net.Addr, handler generic.Service) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

type GenericServiceEchoImpl struct{}

func (g *GenericServiceEchoImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
