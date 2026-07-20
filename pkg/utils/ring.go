package utils

import (
	"errors"
)

var ErrRingFull = errors.New("ring is full")

func NewRing(size int) *Ring { _ = "STUB: not implemented"; return nil }

type Ring struct {
	length int
	rings  []*ring
}

func (r *Ring) Push(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *Ring) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (r *Ring) Dump() interface{} { _ = "STUB: not implemented"; return nil }
