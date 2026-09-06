package utils

type MaxCounter struct {
	now int64
	max int
}

func NewMaxCounter(max int) *MaxCounter { _ = "STUB: not implemented"; return nil }

func (cl *MaxCounter) Inc() bool { _ = "STUB: not implemented"; return false }

func (cl *MaxCounter) Dec() { _ = "STUB: not implemented"; return }

func (cl *MaxCounter) DecN(n int64) { _ = "STUB: not implemented"; return }

func (cl *MaxCounter) Now() int64 { _ = "STUB: not implemented"; return 0 }
