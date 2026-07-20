package limiter

import (
	"github.com/cloudwego/configmanager/iface"
	"github.com/cloudwego/configmanager/util"
)

var _ iface.ConfigValueItem = (*LimiterConfig)(nil)

const TypeLimiter iface.ItemType = "limiter_config"

var defaultLimiterConfig = &LimiterConfig{}

type LimiterConfig struct {
	ConnectionLimit int64 `json:"connection_limit"`
	QPSLimit        int64 `json:"qps_limit"`
}

var NewLimiterConfig = util.JsonInitializer(func() iface.ConfigValueItem {
	return &LimiterConfig{}
})

func CopyDefaultLimitConfig() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

func (l *LimiterConfig) DeepCopy() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

func (l *LimiterConfig) EqualsTo(item iface.ConfigValueItem) bool {
	_ = "STUB: not implemented"
	return false
}
