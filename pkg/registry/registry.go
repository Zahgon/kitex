package registry

import (
	"net"
	"time"
)

type Registry interface {
	Register(info *Info) error
	Deregister(info *Info) error
}

type Info struct {
	ServiceName string

	Addr net.Addr

	PayloadCodec string

	Weight    int
	StartTime time.Time
	WarmUp    time.Duration

	Tags map[string]string

	SkipListenAddr bool
}

var NoopRegistry Registry = &noopRegistry{}

type noopRegistry struct{}

func (e noopRegistry) Register(*Info) error { _ = "STUB: not implemented"; return nil }

func (e noopRegistry) Deregister(*Info) error { _ = "STUB: not implemented"; return nil }
