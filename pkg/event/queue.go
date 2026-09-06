package event

import (
	"sync"
)

var defaultEventNum = int64(MaxEventNum)

func GetDefaultEventNum() int { _ = "STUB: not implemented"; return 0 }

func SetDefaultEventNum(num int) { _ = "STUB: not implemented"; return }

const (
	MaxEventNum = 200
)

type Queue interface {
	Push(e *Event)
	Dump() interface{}
}

type queue struct {
	ring []*Event
	tail uint32
	mu   sync.RWMutex
}

func NewQueue(cap int) Queue { _ = "STUB: not implemented"; return *new(Queue) }

func (q *queue) Push(e *Event) { _ = "STUB: not implemented"; return }

func (q *queue) Dump() interface{} { _ = "STUB: not implemented"; return nil }
