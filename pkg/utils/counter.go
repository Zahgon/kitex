package utils

type AtomicInt int32

func (i *AtomicInt) Inc() { _ = "STUB: not implemented"; return }

func (i *AtomicInt) Dec() { _ = "STUB: not implemented"; return }

func (i *AtomicInt) Value() int { _ = "STUB: not implemented"; return 0 }
