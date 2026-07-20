package peer

import (
	"context"
	"net"
)

type peerKey struct{}

type Peer struct {
	Addr net.Addr
}

func GRPCPeer(ctx context.Context, p *Peer) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetPeerFromContext(ctx context.Context) (peer *Peer, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
