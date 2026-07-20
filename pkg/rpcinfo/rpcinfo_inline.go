package rpcinfo

import (
	"sync"
)

var inlineRPCInfoPool sync.Pool

func init() {
	inlineRPCInfoPool.New = newInlineRPCInfo
}

type inlineRPCInfo struct {
	from       endpointInfo
	to         endpointInfo
	invocation invocation
	config     rpcConfig
	stats      rpcStats
}

func (r *inlineRPCInfo) From() EndpointInfo { _ = "STUB: not implemented"; return *new(EndpointInfo) }

func (r *inlineRPCInfo) To() EndpointInfo { _ = "STUB: not implemented"; return *new(EndpointInfo) }

func (r *inlineRPCInfo) Invocation() Invocation { _ = "STUB: not implemented"; return *new(Invocation) }

func (r *inlineRPCInfo) Config() RPCConfig { _ = "STUB: not implemented"; return *new(RPCConfig) }

func (r *inlineRPCInfo) Stats() RPCStats { _ = "STUB: not implemented"; return *new(RPCStats) }

func (r *inlineRPCInfo) Recycle() { _ = "STUB: not implemented"; return }

func NewRPCInfoWithInlineFields() RPCInfo { _ = "STUB: not implemented"; return *new(RPCInfo) }

func newInlineRPCInfo() interface{} { _ = "STUB: not implemented"; return nil }
