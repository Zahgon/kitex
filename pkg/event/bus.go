package event

import (
	"sync"
)

type Callback func(*Event)

type Bus interface {
	Watch(event string, callback Callback)
	Unwatch(event string, callback Callback)
	Dispatch(event *Event)
	DispatchAndWait(event *Event)
}

func NewEventBus() Bus { _ = "STUB: not implemented"; return *new(Bus) }

type bus struct {
	callbacks sync.Map
}

func (b *bus) Watch(event string, callback Callback) { _ = "STUB: not implemented"; return }

func (b *bus) Unwatch(event string, callback Callback) { _ = "STUB: not implemented"; return }

func (b *bus) Dispatch(event *Event) { _ = "STUB: not implemented"; return }

func (b *bus) DispatchAndWait(event *Event) { _ = "STUB: not implemented"; return }
