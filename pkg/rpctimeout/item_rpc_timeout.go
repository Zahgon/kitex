package rpctimeout

import (
	"sync/atomic"
	"time"

	"github.com/cloudwego/configmanager/iface"
	"github.com/cloudwego/configmanager/util"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var (
	_ iface.ConfigValueItem   = (*RPCTimeout)(nil)
	_ rpcinfo.Timeouts        = (*RPCTimeout)(nil)
	_ rpcinfo.TimeoutProvider = (*Container)(nil)
)

const TypeRPCTimeout iface.ItemType = "rpc_timeout"

const wildcardMethod = "*"

var defaultRPCTimeout = &RPCTimeout{
	RPCTimeoutMS:  1000,
	ConnTimeoutMS: 50,
}

type RPCTimeout struct {
	RPCTimeoutMS  int `json:"rpc_timeout_ms"`
	ConnTimeoutMS int `json:"conn_timeout_ms"`
}

var NewRPCTimeout = util.JsonInitializer(func() iface.ConfigValueItem {
	return &RPCTimeout{}
})

func CopyDefaultRPCTimeout() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

func (r *RPCTimeout) DeepCopy() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

func (r *RPCTimeout) EqualsTo(other iface.ConfigValueItem) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RPCTimeout) RPCTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *RPCTimeout) ConnectTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *RPCTimeout) ReadWriteTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func NewContainer() *Container { _ = "STUB: not implemented"; return nil }

type rpcTimeoutConfig struct {
	configs      map[string]*RPCTimeout
	globalConfig *RPCTimeout
}

type Container struct {
	config atomic.Value
}

func (c *Container) NotifyPolicyChange(configs map[string]*RPCTimeout) {
	_ = "STUB: not implemented"
	return
}

func (c *Container) Timeouts(ri rpcinfo.RPCInfo) rpcinfo.Timeouts {
	_ = "STUB: not implemented"
	return *new(rpcinfo.Timeouts)
}
