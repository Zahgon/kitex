package acl

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

type RejectFunc func(ctx context.Context, request interface{}) (reason error)

func ApplyRules(ctx context.Context, request interface{}, rules []RejectFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func NewACLMiddleware(rules []RejectFunc) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}
