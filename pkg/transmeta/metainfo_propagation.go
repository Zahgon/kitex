package transmeta

import (
	"context"
	"sync/atomic"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

type MetainfoPropagationMode int32

const (
	MetainfoPropagationLegacy MetainfoPropagationMode = iota

	MetainfoPropagationSingleHop
)

var (
	metainfoPropagationMode atomic.Int32

	metainfoPropagationModeProvider atomic.Value
)

func init() {
	metainfoPropagationMode.Store(int32(MetainfoPropagationLegacy))
	SetMetainfoPropagationModeProvider(nil)
}

func SetMetainfoPropagationMode(mode MetainfoPropagationMode) { _ = "STUB: not implemented"; return }

func GetMetainfoPropagationMode() MetainfoPropagationMode {
	_ = "STUB: not implemented"
	return *new(MetainfoPropagationMode)
}

func SetMetainfoPropagationModeProvider(provider func(context.Context) MetainfoPropagationMode) {
	_ = "STUB: not implemented"
	return
}

func getMetainfoPropagationMode(ctx context.Context) MetainfoPropagationMode {
	_ = "STUB: not implemented"
	return *new(MetainfoPropagationMode)
}

func TransferForwardMetaInfo(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func saveOutboundMetaInfoToHTTPHeader(ctx context.Context, h metainfo.HTTPHeaderSetter) {
	_ = "STUB: not implemented"
	return
}
