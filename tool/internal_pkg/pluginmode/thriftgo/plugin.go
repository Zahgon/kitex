package thriftgo

import (
	"github.com/cloudwego/thriftgo/plugin"
)

const PluginName = "thrift-gen-kitex"

const TheUseOptionMessage = "kitex_gen is not generated due to the -use option"

func Run() int { _ = "STUB: not implemented"; return 0 }

func (c *converter) failResp(err error) *plugin.Response { _ = "STUB: not implemented"; return nil }

func HandleRequest(req *plugin.Request) *plugin.Response { _ = "STUB: not implemented"; return nil }

func exit(res *plugin.Response) int { _ = "STUB: not implemented"; return 0 }
