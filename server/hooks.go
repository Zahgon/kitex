package server

import "sync"

func RegisterStartHook(h func()) { _ = "STUB: not implemented"; return }

func RegisterShutdownHook(h func()) { _ = "STUB: not implemented"; return }

type Hooks []func()

func (h *Hooks) add(g func()) { _ = "STUB: not implemented"; return }

var (
	onServerStart   Hooks
	muStartHooks    sync.Mutex
	onShutdown      Hooks
	muShutdownHooks sync.Mutex
)
