package invoke

import (
	"net"

	"github.com/cloudwego/kitex/pkg/remote/trans/invoke"
)

type Message = invoke.Message

func NewMessage(local, remote net.Addr) Message { _ = "STUB: not implemented"; return *new(Message) }
