package xds

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

type ClientSuite struct {
	RouterMiddleware endpoint.Middleware
	Resolver         discovery.Resolver
}

func CheckClientSuite(cs ClientSuite) bool { _ = "STUB: not implemented"; return false }
