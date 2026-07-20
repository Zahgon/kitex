package streamclient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/utils"
)

type Option struct {
	F func(*client.Options, *utils.Slice)
}

func (o *Option) GetClientOption() client.Option {
	_ = "STUB: not implemented"
	return *new(client.Option)
}

func ConvertOptionFrom(opt client.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func GetClientOptions(ops []Option) []client.Option { _ = "STUB: not implemented"; return nil }
