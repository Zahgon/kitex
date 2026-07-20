package netpollmux

import (
	"github.com/cloudwego/gopkg/protocol/thrift"
)

type ControlFrame struct{}

func NewControlFrame() *ControlFrame { _ = "STUB: not implemented"; return nil }

func (p *ControlFrame) BLength() int { _ = "STUB: not implemented"; return 0 }

func (p *ControlFrame) FastWrite(b []byte) int { _ = "STUB: not implemented"; return 0 }

func (p *ControlFrame) FastWriteNocopy(b []byte, w thrift.NocopyWriter) int {
	_ = "STUB: not implemented"
	return 0
}

func (p *ControlFrame) FastRead(b []byte) (off int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var _ thrift.FastCodec = &ControlFrame{}
