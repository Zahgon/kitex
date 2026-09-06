package streamcall

import (
	"strings"

	"github.com/cloudwego/kitex/client/callopt"
)

type Option struct {
	f func(o *callopt.CallOptions, di *strings.Builder)
}

func (o Option) GetCallOption() callopt.Option {
	_ = "STUB: not implemented"
	return *new(callopt.Option)
}

func ConvertOptionFrom(option callopt.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func GetCallOptions(ops []Option) []callopt.Option { _ = "STUB: not implemented"; return nil }
