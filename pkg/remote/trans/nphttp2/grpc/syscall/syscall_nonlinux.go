//go:build !linux
// +build !linux

package syscall

import (
	"net"
	"sync"
	"time"
)

var once sync.Once

func log() { _ = "STUB: not implemented"; return }

func GetCPUTime() int64 { _ = "STUB: not implemented"; return 0 }

type Rusage struct{}

func GetRusage() *Rusage { _ = "STUB: not implemented"; return nil }

func CPUTimeDiff(first, latest *Rusage) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

func SetTCPUserTimeout(conn net.Conn, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func GetTCPUserTimeout(conn net.Conn) (int, error) { _ = "STUB: not implemented"; return 0, nil }
