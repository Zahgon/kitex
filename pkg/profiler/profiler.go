package profiler

import (
	"bytes"
	"context"
	"sync"
	"time"
)

type profilerContextKey struct{}

type Profiler interface {
	Run(ctx context.Context) (err error)
	Stop()
	Pause()
	Resume()
	Prepare(ctx context.Context) context.Context
	Tag(ctx context.Context, tags ...string) context.Context
	Untag(ctx context.Context)
	Lookup(ctx context.Context, key string) (string, bool)
}

type Processor func(profiles []*TagsProfile) error

func LogProcessor(profiles []*TagsProfile) error { _ = "STUB: not implemented"; return nil }

func NewProfiler(processor Processor, interval, window time.Duration) *profiler {
	_ = "STUB: not implemented"
	return nil
}

var _ Profiler = (*profiler)(nil)

const (
	stateRunning  = 0
	statePausing  = 1
	statePaused   = 2
	stateResuming = 3
	stateStopped  = 4
)

type profiler struct {
	data      bytes.Buffer
	state     int
	stateCond *sync.Cond

	processor Processor
	interval  time.Duration
	window    time.Duration
}

func Tag(ctx context.Context) { _ = "STUB: not implemented"; return }

func IsEnabled(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func Untag(ctx context.Context) { _ = "STUB: not implemented"; return }

type profilerContext struct {
	profiler Profiler
	untagCtx context.Context
	tags     []string
}

func newProfilerContext(profiler Profiler) *profilerContext { _ = "STUB: not implemented"; return nil }

func (p *profiler) Prepare(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (p *profiler) State() (state int) { _ = "STUB: not implemented"; return 0 }

func (p *profiler) Stop() { _ = "STUB: not implemented"; return }

func (p *profiler) Pause() { _ = "STUB: not implemented"; return }

func (p *profiler) Resume() { _ = "STUB: not implemented"; return }

func (p *profiler) Run(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (p *profiler) Tag(ctx context.Context, tags ...string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (p *profiler) Untag(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *profiler) Lookup(ctx context.Context, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (p *profiler) stateChange(from, to int) { _ = "STUB: not implemented"; return }

func (p *profiler) stateWait(to int) { _ = "STUB: not implemented"; return }

func (p *profiler) startProfile() error { _ = "STUB: not implemented"; return nil }

func (p *profiler) stopProfile() { _ = "STUB: not implemented"; return }

func (p *profiler) analyse() ([]*TagsProfile, error) { _ = "STUB: not implemented"; return nil, nil }

type TagsProfile struct {
	Key     string
	Tags    []string
	Value   int64
	Percent float64
}

func labelToTags(label map[string][]string) []string { _ = "STUB: not implemented"; return nil }

func tagsToKey(tags []string) string { _ = "STUB: not implemented"; return "" }
