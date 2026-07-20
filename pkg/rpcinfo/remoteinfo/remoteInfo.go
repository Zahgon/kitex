package remoteinfo

import (
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type RemoteInfo interface {
	rpcinfo.EndpointInfo
	SetServiceName(name string)
	SetTag(key, value string) error
	ForceSetTag(key, value string)

	SetTagLock(key string)
	GetInstance() discovery.Instance
	SetInstance(ins discovery.Instance)

	SetRemoteAddr(addr net.Addr) (ok bool)
	ImmutableView() rpcinfo.EndpointInfo
}

type RefreshableInstance interface {
	RefreshInstanceWithAddr(addr net.Addr) (newInstance discovery.Instance)
}

var (
	_              RemoteInfo = &remoteInfo{}
	remoteInfoPool sync.Pool
)

type remoteInfo struct {
	sync.RWMutex
	serviceName string
	method      string
	tags        map[string]string

	instance discovery.Instance
	tagLocks map[string]struct{}
}

func init() {
	remoteInfoPool.New = newRemoteInfo
}

func newRemoteInfo() interface{} { _ = "STUB: not implemented"; return nil }

func (ri *remoteInfo) GetInstance() (ins discovery.Instance) {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}

func (ri *remoteInfo) SetInstance(ins discovery.Instance) { _ = "STUB: not implemented"; return }

func (ri *remoteInfo) ServiceName() string { _ = "STUB: not implemented"; return "" }

func (ri *remoteInfo) SetServiceName(name string) { _ = "STUB: not implemented"; return }

func (ri *remoteInfo) Method() string { _ = "STUB: not implemented"; return "" }

func (ri *remoteInfo) Tag(key string) (value string, exist bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ri *remoteInfo) DefaultTag(key, def string) string { _ = "STUB: not implemented"; return "" }

func (ri *remoteInfo) SetRemoteAddr(addr net.Addr) bool { _ = "STUB: not implemented"; return false }

func (ri *remoteInfo) Address() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (ri *remoteInfo) SetTagLock(key string) { _ = "STUB: not implemented"; return }

func (ri *remoteInfo) SetTag(key, value string) error { _ = "STUB: not implemented"; return nil }

func (ri *remoteInfo) ForceSetTag(key, value string) { _ = "STUB: not implemented"; return }

func (ri *remoteInfo) ImmutableView() rpcinfo.EndpointInfo {
	_ = "STUB: not implemented"
	return *new(rpcinfo.EndpointInfo)
}

func (ri *remoteInfo) zero() { _ = "STUB: not implemented"; return }

func (ri *remoteInfo) Recycle() { _ = "STUB: not implemented"; return }

func NewRemoteInfo(basicInfo *rpcinfo.EndpointBasicInfo, method string) RemoteInfo {
	_ = "STUB: not implemented"
	return *new(RemoteInfo)
}

func AsRemoteInfo(r rpcinfo.EndpointInfo) RemoteInfo {
	_ = "STUB: not implemented"
	return *new(RemoteInfo)
}
