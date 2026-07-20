package server

import (
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func serverTimeoutMW(next endpoint.Endpoint) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}
