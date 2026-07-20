package streaming

import (
	"sync"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

const KitexUnusedProtection = 0

var userStreamNotImplementingWithDoFinish sync.Once

func UnaryCompatibleMiddleware(mode serviceinfo.StreamingMode, allow bool) bool {
	_ = "STUB: not implemented"
	return false
}

func FinishStream(s Stream, err error) { _ = "STUB: not implemented"; return }

func FinishClientStream(s ClientStream, err error) { _ = "STUB: not implemented"; return }

func GetServerStreamFromArg(arg interface{}) (ServerStream, error) {
	_ = "STUB: not implemented"
	return *new(ServerStream), nil
}
