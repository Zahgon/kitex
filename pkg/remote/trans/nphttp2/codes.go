package nphttp2

import (
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
)

var kitexErrConvTab = map[error]codes.Code{
	kerrors.ErrInternalException: codes.Internal,
	kerrors.ErrOverlimit:         codes.ResourceExhausted,
	kerrors.ErrRemoteOrNetwork:   codes.Unavailable,
	kerrors.ErrACL:               codes.PermissionDenied,
}

func convertStatus(err error) *status.Status { _ = "STUB: not implemented"; return nil }

func getStatusForBizErr(dErr *kerrors.DetailedError) *status.Status {
	_ = "STUB: not implemented"
	return nil
}
