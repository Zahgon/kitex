package ttstream

type sideType int32

const (
	clientSide sideType = 1
	serverSide sideType = 2

	streamCacheSize = 32
	frameCacheSize  = 256

	connStateOpen   = 0
	connStateClosed = 1
)

func isIgnoreError(err error) bool { _ = "STUB: not implemented"; return false }
