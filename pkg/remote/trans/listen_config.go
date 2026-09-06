//go:build !windows
// +build !windows

package trans

import (
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
)

func NewListenConfig(opt *remote.ServerOption) net.ListenConfig {
	_ = "STUB: not implemented"
	return *new(net.ListenConfig)
}
