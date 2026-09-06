package utils

import "sync"

type ring struct {
	l    sync.RWMutex
	arr  []interface{}
	size int
	tail int
	head int
}

func newRing(size int) *ring { _ = "STUB: not implemented"; return nil }

func (r *ring) Push(i interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *ring) Pop() interface{} { _ = "STUB: not implemented"; return nil }

type ringDump struct {
	Array []interface{} `json:"array"`
	Len   int           `json:"len"`
	Cap   int           `json:"cap"`
}

func (r *ring) Dump(m *ringDump) { _ = "STUB: not implemented"; return }

func (r *ring) inc() int { _ = "STUB: not implemented"; return 0 }

func (r *ring) dec() int { _ = "STUB: not implemented"; return 0 }

func (r *ring) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (r *ring) isFull() bool { _ = "STUB: not implemented"; return false }
