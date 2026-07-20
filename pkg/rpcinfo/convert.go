package rpcinfo

type Taggable interface {
	SetTag(key, value string) error
}

func AsTaggable(i interface{}) Taggable { _ = "STUB: not implemented"; return *new(Taggable) }

func AsMutableEndpointInfo(ei EndpointInfo) MutableEndpointInfo {
	_ = "STUB: not implemented"
	return *new(MutableEndpointInfo)
}

func AsMutableRPCStats(r RPCStats) MutableRPCStats {
	_ = "STUB: not implemented"
	return *new(MutableRPCStats)
}

func AsMutableRPCConfig(r RPCConfig) MutableRPCConfig {
	_ = "STUB: not implemented"
	return *new(MutableRPCConfig)
}
