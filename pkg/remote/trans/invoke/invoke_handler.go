package invoke

import (
	"github.com/cloudwego/kitex/pkg/remote"
)

type Handler interface {
	Call(Message) error
}

type ivkHandler struct {
	opt       *remote.ServerOption
	transHdlr remote.ServerTransHandler
}

func NewIvkHandler(opt *remote.ServerOption, transHdlr remote.ServerTransHandler) (Handler, error) {
	_ = "STUB: not implemented"
	return *new(Handler), nil
}

func (s *ivkHandler) Call(msg Message) (err error) { _ = "STUB: not implemented"; return nil }
