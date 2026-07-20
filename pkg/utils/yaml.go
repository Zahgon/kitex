package utils

import (
	"time"

	"github.com/cloudwego/kitex/internal/configutil"
)

var (
	_ configutil.Config         = &YamlConfig{}
	_ configutil.RichTypeConfig = &YamlConfig{}
)

type YamlConfig struct {
	configutil.RichTypeConfig
	data map[interface{}]interface{}
}

func ReadYamlConfigFile(yamlFile string) (*YamlConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (yc *YamlConfig) Get(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (yc *YamlConfig) GetDuration(key string) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}
