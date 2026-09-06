package rpcinfo

import (
	"net"
	"sync"
)

var (
	_            EndpointInfo        = &endpointInfo{}
	_            MutableEndpointInfo = &endpointInfo{}
	endpointPool sync.Pool
)

type endpointInfo struct {
	serviceName string
	method      string
	address     net.Addr
	tags        map[string]string
}

func init() {
	endpointPool.New = newEndpointInfo
}

func newEndpointInfo() interface{} { _ = "STUB: not implemented"; return nil }

func (ei *endpointInfo) ServiceName() string { _ = "STUB: not implemented"; return "" }

func (ei *endpointInfo) Method() string { _ = "STUB: not implemented"; return "" }

func (ei *endpointInfo) Address() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (ei *endpointInfo) Tag(key string) (value string, exist bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ei *endpointInfo) DefaultTag(key, def string) string { _ = "STUB: not implemented"; return "" }

func (ei *endpointInfo) SetServiceName(serviceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ei *endpointInfo) SetMethod(method string) error { _ = "STUB: not implemented"; return nil }

func (ei *endpointInfo) SetAddress(addr net.Addr) error { _ = "STUB: not implemented"; return nil }

func (ei *endpointInfo) SetTag(key, value string) error { _ = "STUB: not implemented"; return nil }

func (ei *endpointInfo) ImmutableView() EndpointInfo {
	_ = "STUB: not implemented"
	return *new(EndpointInfo)
}

func (ei *endpointInfo) Reset() { _ = "STUB: not implemented"; return }

func (ei *endpointInfo) ResetFromBasicInfo(bi *EndpointBasicInfo) {
	_ = "STUB: not implemented"
	return
}

func (ei *endpointInfo) zero() { _ = "STUB: not implemented"; return }

func (ei *endpointInfo) Recycle() { _ = "STUB: not implemented"; return }

func NewMutableEndpointInfo(serviceName, method string, address net.Addr, tags map[string]string) MutableEndpointInfo {
	_ = "STUB: not implemented"
	return *new(MutableEndpointInfo)
}

func NewEndpointInfo(serviceName, method string, address net.Addr, tags map[string]string) EndpointInfo {
	_ = "STUB: not implemented"
	return *new(EndpointInfo)
}

func FromBasicInfo(bi *EndpointBasicInfo) EndpointInfo {
	_ = "STUB: not implemented"
	return *new(EndpointInfo)
}

func EmptyEndpointInfo() EndpointInfo { _ = "STUB: not implemented"; return *new(EndpointInfo) }
