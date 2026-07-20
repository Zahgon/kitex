package circuitbreak

import (
	"context"

	"github.com/bytedance/gopkg/cloud/circuitbreaker"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/cep"
)

type Parameter struct {
	Enabled bool

	ErrorRate float64

	MinimalSample int64
}

type ErrorType int

const (
	TypeIgnorable ErrorType = iota

	TypeTimeout

	TypeFailure

	TypeSuccess
)

func WrapErrorWithType(err error, errorType ErrorType) CircuitBreakerAwareError {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerAwareError)
}

type GetErrorTypeFunc func(ctx context.Context, request, response interface{}, err error) ErrorType

type Control struct {
	GetKey func(ctx context.Context, request interface{}) (key string, enabled bool)

	GetErrorType GetErrorTypeFunc

	DecorateError func(ctx context.Context, request interface{}, err error) error
}

func NewCircuitBreakerMW(control Control, panel circuitbreaker.Panel) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func NewStreamCircuitBreakerMW(control Control, panel circuitbreaker.Panel) cep.StreamMiddleware {
	_ = "STUB: not implemented"
	return *new(cep.StreamMiddleware)
}

func RecordStat(ctx context.Context, request, response interface{}, err error, cbKey string, ctl *Control, panel circuitbreaker.Panel) {
	_ = "STUB: not implemented"
	return
}

type CircuitBreakerAwareError interface {
	error
	TypeForCircuitBreaker() ErrorType
}

type errorWrapperWithType struct {
	errType ErrorType
	err     error
}

func (e errorWrapperWithType) TypeForCircuitBreaker() ErrorType {
	_ = "STUB: not implemented"
	return *new(ErrorType)
}

func (e errorWrapperWithType) Error() string { _ = "STUB: not implemented"; return "" }

func (e errorWrapperWithType) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e errorWrapperWithType) Is(target error) bool { _ = "STUB: not implemented"; return false }
