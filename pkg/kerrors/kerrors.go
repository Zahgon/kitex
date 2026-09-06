package kerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInternalException  = &basicError{"internal exception"}
	ErrServiceDiscovery   = &basicError{"service discovery error"}
	ErrGetConnection      = &basicError{"get connection error"}
	ErrLoadbalance        = &basicError{"loadbalance error"}
	ErrNoMoreInstance     = &basicError{"no more instances to retry"}
	ErrRPCTimeout         = &basicError{"rpc timeout"}
	ErrCanceledByBusiness = &basicError{"canceled by business"}
	ErrTimeoutByBusiness  = &basicError{"timeout by business"}
	ErrACL                = &basicError{"request forbidden"}
	ErrCircuitBreak       = &basicError{"forbidden by circuitbreaker"}
	ErrRemoteOrNetwork    = &basicError{"remote or network error"}
	ErrOverlimit          = &basicError{"request over limit"}
	ErrPanic              = &basicError{"panic"}
	ErrBiz                = &basicError{"biz error"}

	ErrRetry = &basicError{"retry error"}

	ErrRPCFinish = &basicError{"rpc call finished"}

	ErrRoute = &basicError{"rpc route failed"}

	ErrPayloadValidation = &basicError{"payload validation error"}
)

var (
	ErrNotSupported         = ErrInternalException.WithCause(errors.New("operation not supported"))
	ErrNoResolver           = ErrInternalException.WithCause(errors.New("no resolver available"))
	ErrNoDestService        = ErrInternalException.WithCause(errors.New("no dest service"))
	ErrNoDestAddress        = ErrInternalException.WithCause(errors.New("no dest address"))
	ErrNoConnection         = ErrInternalException.WithCause(errors.New("no connection available"))
	ErrConnOverLimit        = ErrOverlimit.WithCause(errors.New("to many connections"))
	ErrQPSOverLimit         = ErrOverlimit.WithCause(errors.New("request too frequent"))
	ErrNoIvkRequest         = ErrInternalException.WithCause(errors.New("invoker request not set"))
	ErrServiceCircuitBreak  = ErrCircuitBreak.WithCause(errors.New("service circuitbreak"))
	ErrInstanceCircuitBreak = ErrCircuitBreak.WithCause(errors.New("instance circuitbreak"))
	ErrNoInstance           = ErrServiceDiscovery.WithCause(errors.New("no instance available"))
)

func ErrNonExistentMethod(svcName, method string) error { _ = "STUB: not implemented"; return nil }

func ErrNotStreamingMethod(svcName, method string) error { _ = "STUB: not implemented"; return nil }

func ErrNotUnaryMethod(svcName, method string) error { _ = "STUB: not implemented"; return nil }

type basicError struct {
	message string
}

func (be *basicError) Error() string { _ = "STUB: not implemented"; return "" }

func (be *basicError) WithCause(cause error) error { _ = "STUB: not implemented"; return nil }

func (be *basicError) WithCauseAndStack(cause error, stack string) error {
	_ = "STUB: not implemented"
	return nil
}

func (be *basicError) WithCauseAndExtraMsg(cause error, extraMsg string) error {
	_ = "STUB: not implemented"
	return nil
}

func (be *basicError) Timeout() bool { _ = "STUB: not implemented"; return false }

type DetailedError struct {
	basic    *basicError
	cause    error
	stack    string
	extraMsg string
}

func (de *DetailedError) Error() string { _ = "STUB: not implemented"; return "" }

func (de *DetailedError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (de *DetailedError) ErrorType() error { _ = "STUB: not implemented"; return nil }

func (de *DetailedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (de *DetailedError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (de *DetailedError) As(target interface{}) bool { _ = "STUB: not implemented"; return false }

func (de *DetailedError) Timeout() bool { _ = "STUB: not implemented"; return false }

func (de *DetailedError) Stack() string { _ = "STUB: not implemented"; return "" }

func (de *DetailedError) WithExtraMsg(extraMsg string) { _ = "STUB: not implemented"; return }

func appendErrMsg(errMsg, extra string) string { _ = "STUB: not implemented"; return "" }

func IsKitexError(err error) bool { _ = "STUB: not implemented"; return false }

var TimeoutCheckFunc func(err error) bool

func IsTimeoutError(err error) bool { _ = "STUB: not implemented"; return false }
