package loadbalance

type round struct {
	state uint64
	_     [7]uint64
}

func (r *round) Next() uint64 { _ = "STUB: not implemented"; return 0 }

func newRound() *round { _ = "STUB: not implemented"; return nil }

func newRandomRound() *round { _ = "STUB: not implemented"; return nil }
