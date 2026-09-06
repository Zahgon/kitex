package syscall

import (
	"net"
	"syscall"
	"time"
)

func GetCPUTime() int64 { _ = "STUB: not implemented"; return 0 }

type Rusage = syscall.Rusage

func GetRusage() *Rusage { _ = "STUB: not implemented"; return nil }

func CPUTimeDiff(first, latest *Rusage) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

func SetTCPUserTimeout(conn net.Conn, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func GetTCPUserTimeout(conn net.Conn) (opt int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
