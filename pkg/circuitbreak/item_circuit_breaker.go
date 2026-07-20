package circuitbreak

import (
	"github.com/cloudwego/configmanager/iface"
	"github.com/cloudwego/configmanager/util"
)

var _ iface.ConfigValueItem = (*CBConfigItem)(nil)

const TypeCircuitBreaker iface.ItemType = "cb_config"

type CBConfigItem CBConfig

func CopyDefaultCBConfig() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

var NewCBConfig = util.JsonInitializer(func() iface.ConfigValueItem {
	return &CBConfigItem{}
})

func (c *CBConfigItem) DeepCopy() iface.ConfigValueItem {
	_ = "STUB: not implemented"
	return *new(iface.ConfigValueItem)
}

func (c *CBConfigItem) EqualsTo(other iface.ConfigValueItem) bool {
	_ = "STUB: not implemented"
	return false
}
