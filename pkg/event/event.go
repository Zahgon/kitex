package event

import "time"

type Event struct {
	Name   string
	Time   time.Time
	Detail string
	Extra  interface{}
}

type lazyExtra interface {
	KitexDumpLazyExtra() interface{}
}
