package leakcheck

import (
	"time"
)

var goroutinesToIgnore = []string{
	"testing.Main(",
	"testing.tRunner(",
	"testing.(*M).",
	"runtime.goexit",
	"created by runtime.gc",
	"created by runtime/trace.Start",
	"interestingGoroutines",
	"runtime.MHeap_Scavenger",
	"signal.signal_recv",
	"sigterm.handler",
	"runtime_mcall",
	"(*loggingT).flushDaemon",
	"goroutine in C code",
	"httputil.DumpRequestOut",
}

func RegisterIgnoreGoroutine(s string) { _ = "STUB: not implemented"; return }

func ignore(g string) bool { _ = "STUB: not implemented"; return false }

func interestingGoroutines() (gs []string) { _ = "STUB: not implemented"; return nil }

type Errorfer interface {
	Errorf(format string, args ...interface{})
}

func check(efer Errorfer, timeout time.Duration) { _ = "STUB: not implemented"; return }

func Check(efer Errorfer) { _ = "STUB: not implemented"; return }
