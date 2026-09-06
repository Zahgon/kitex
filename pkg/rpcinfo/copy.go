package rpcinfo

type plainRPCInfo struct {
	from struct{ EndpointInfo }
	to   struct{ EndpointInfo }
	inv  struct{ Invocation }
	cfg  struct{ RPCConfig }
}

func (p *plainRPCInfo) From() EndpointInfo { _ = "STUB: not implemented"; return *new(EndpointInfo) }

func (p *plainRPCInfo) To() EndpointInfo { _ = "STUB: not implemented"; return *new(EndpointInfo) }

func (p *plainRPCInfo) Invocation() Invocation { _ = "STUB: not implemented"; return *new(Invocation) }

func (p *plainRPCInfo) Config() RPCConfig { _ = "STUB: not implemented"; return *new(RPCConfig) }

func (p *plainRPCInfo) Stats() RPCStats { _ = "STUB: not implemented"; return *new(RPCStats) }

func freeze(ri RPCInfo) RPCInfo { _ = "STUB: not implemented"; return *new(RPCInfo) }

func copyMap(src map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func copyEndpointInfo(ei EndpointInfo) EndpointInfo {
	_ = "STUB: not implemented"
	return *new(EndpointInfo)
}

func copyInvocation(i Invocation) Invocation { _ = "STUB: not implemented"; return *new(Invocation) }

func copyRPCConfig(cfg RPCConfig) RPCConfig { _ = "STUB: not implemented"; return *new(RPCConfig) }
