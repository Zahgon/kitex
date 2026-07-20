package netpollmux

import (
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"
)

type EventHandler interface {
	Recv(bufReader remote.ByteBuffer, err error) error
}

type shardMap struct {
	size   int32
	shards []*shard
}

type shard struct {
	msgs map[int32]EventHandler
	sync.RWMutex
}

func newShardMap(size int) *shardMap { _ = "STUB: not implemented"; return nil }

func (m *shardMap) getShard(seqID int32) *shard { _ = "STUB: not implemented"; return nil }

func (m *shardMap) store(seqID int32, msg EventHandler) { _ = "STUB: not implemented"; return }

func (m *shardMap) load(seqID int32) (msg EventHandler, ok bool) {
	_ = "STUB: not implemented"
	return *new(EventHandler), false
}

func (m *shardMap) delete(seqID int32) { _ = "STUB: not implemented"; return }

func (m *shardMap) rangeMap(fn func(seqID int32, msg EventHandler)) {
	_ = "STUB: not implemented"
	return
}

func abs(n int32) int32 { _ = "STUB: not implemented"; return 0 }
