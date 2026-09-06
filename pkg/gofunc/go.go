package gofunc

import (
	"context"
	"sync"

	"github.com/bytedance/gopkg/util/gopool"

	"github.com/cloudwego/kitex/pkg/profiler"
)

type GoTask func(context.Context, func())

var GoFunc GoTask

func init() {
	GoFunc = func(ctx context.Context, f func()) {
		gopool.CtxGo(ctx, func() {
			profiler.Tag(ctx)
			f()
			profiler.Untag(ctx)
		})
	}
}

func RecoverGoFuncWithInfo(ctx context.Context, task func(), info *Info) {
	_ = "STUB: not implemented"
	return
}

func SetPanicHandler(hdlr func(info *Info, panicErr interface{}, panicStack string)) {
	_ = "STUB: not implemented"
	return
}

func NewBasicInfo(remoteService, remoteAddr string) *Info { _ = "STUB: not implemented"; return nil }

type Info struct {
	RemoteService string
	RemoteAddr    string
}

var (
	EmptyInfo = &Info{}

	panicHandler func(info *Info, panicErr interface{}, panicStack string)
	phLock       sync.RWMutex

	infoPool = &sync.Pool{
		New: func() interface{} {
			return new(Info)
		},
	}
)
