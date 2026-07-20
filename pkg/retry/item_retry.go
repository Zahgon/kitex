package retry

import (
	"github.com/cloudwego/configmanager/iface"
	"github.com/cloudwego/configmanager/util"
)

var _ iface.ConfigValueItem = (*RetryConfig)(nil)

const TypeRetry iface.ItemType = "retry_config"

var defaultRetry = &RetryConfig{
	Config: &Policy{},
}

type RetryConfig struct {
	Config *Policy `json:"config"`
}

var NewRetryConfig = util.JsonInitializer(func() iface.ConfigValueItem {
	return &RetryConfig{
		Config: &Policy{},
	}
})

func CopyDefaultRetryConfig() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

func (c *RetryConfig) DeepCopy() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

func (c *RetryConfig) EqualsTo(other iface.ConfigValueItem) bool {
	_ = "STUB: not implemented"
	return false
}
