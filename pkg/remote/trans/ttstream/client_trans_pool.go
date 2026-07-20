package ttstream

import (
	"github.com/cloudwego/netpoll"
)

var dialer = netpoll.NewDialer()

type transPool interface {
	Get(network, addr string) (trans *clientTransport, err error)
	Put(trans *clientTransport)
}
