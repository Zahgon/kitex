package testutils

import (
	"context"
)

const DefaultChanBufferSize = 1

type Channel struct {
	ch chan interface{}
}

func (c *Channel) Send(value interface{}) { _ = "STUB: not implemented"; return }

func (c *Channel) SendContext(ctx context.Context, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) SendOrFail(value interface{}) bool { _ = "STUB: not implemented"; return false }

func (c *Channel) ReceiveOrFail() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (c *Channel) Receive(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) Replace(value interface{}) { _ = "STUB: not implemented"; return }

func NewChannel() *Channel { _ = "STUB: not implemented"; return nil }

func NewChannelWithSize(bufSize int) *Channel { _ = "STUB: not implemented"; return nil }
