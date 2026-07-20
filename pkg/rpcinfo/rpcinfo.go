package rpcinfo

import (
	"os"
	"sync"
)

var (
	rpcInfoPool sync.Pool
	enablePool  int32 = 1
)

func init() {

	if os.Getenv("KITEX_DISABLE_RPCINFO_POOL") != "" {
		EnablePool(false)
	}
	rpcInfoPool.New = newRPCInfo
}

func EnablePool(enable bool) { _ = "STUB: not implemented"; return }

func PoolEnabled() bool { _ = "STUB: not implemented"; return false }

type rpcInfo struct {
	from       EndpointInfo
	to         EndpointInfo
	invocation Invocation
	config     RPCConfig
	stats      RPCStats
}

func (r *rpcInfo) From() EndpointInfo { _ = "STUB: not implemented"; return *new(EndpointInfo) }

func (r *rpcInfo) To() EndpointInfo { _ = "STUB: not implemented"; return *new(EndpointInfo) }

func (r *rpcInfo) Config() RPCConfig { _ = "STUB: not implemented"; return *new(RPCConfig) }

func (r *rpcInfo) Invocation() Invocation { _ = "STUB: not implemented"; return *new(Invocation) }

func (r *rpcInfo) Stats() RPCStats { _ = "STUB: not implemented"; return *new(RPCStats) }

func (r *rpcInfo) zero() { _ = "STUB: not implemented"; return }

func (r *rpcInfo) Recycle() { _ = "STUB: not implemented"; return }

func NewRPCInfo(from, to EndpointInfo, ink Invocation, config RPCConfig, stats RPCStats) RPCInfo {
	_ = "STUB: not implemented"
	return *new(RPCInfo)
}

func newRPCInfo() interface{} { _ = "STUB: not implemented"; return nil }
