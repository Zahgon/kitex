package ttstream

func newShortConnTransPool() transPool { _ = "STUB: not implemented"; return *new(transPool) }

type shortConnTransPool struct{}

func (p *shortConnTransPool) Get(network, addr string) (*clientTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *shortConnTransPool) Put(trans *clientTransport) { _ = "STUB: not implemented"; return }
