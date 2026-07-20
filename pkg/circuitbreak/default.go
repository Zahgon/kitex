package circuitbreak

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

var ignoreErrTypes = map[error]ErrorType{
	kerrors.ErrInternalException: TypeIgnorable,
	kerrors.ErrServiceDiscovery:  TypeIgnorable,
	kerrors.ErrACL:               TypeIgnorable,
	kerrors.ErrLoadbalance:       TypeIgnorable,
}

func ErrorTypeOnServiceLevel(ctx context.Context, request, response interface{}, err error) ErrorType {
	_ = "STUB: not implemented"
	return *new(ErrorType)
}

func ErrorTypeOnInstanceLevel(ctx context.Context, request, response interface{}, err error) ErrorType {
	_ = "STUB: not implemented"
	return *new(ErrorType)
}

func FailIfError(ctx context.Context, request, response interface{}, err error) ErrorType {
	_ = "STUB: not implemented"
	return *new(ErrorType)
}

func NoDecoration(ctx context.Context, request interface{}, err error) error {
	_ = "STUB: not implemented"
	return nil
}
