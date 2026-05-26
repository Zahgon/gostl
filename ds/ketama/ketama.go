package ketama

import (
	"github.com/liyue201/gostl/utils/sync"
)

var (
	defaultReplicas = 10
	defaultLocker   sync.FakeLocker
)

const salt = "ni9fkh72hgh1g"

// Options hold Ketama's options
type Options struct {
	replicas int
	locker   sync.Locker
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to config a Ketama with goroutine-safe
func WithGoroutineSafe() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReplicas is used to config the hash replicas of a Ketama
func WithReplicas(replicas int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Ketama is an implementation of consistent-hash
type Ketama struct {
	locker   sync.Locker
	replicas int
	m        *treemap.Map[uint64, string]
}

// New creates a new ketama
func New(opts ...Option) *Ketama { _ = "STUB: not implemented"; return nil }

// Empty returns true if the ketama is empty, otherwise returns false
func (k *Ketama) Empty() bool { _ = "STUB: not implemented"; return false }

// Add adds nodes to the ketama ring
func (k *Ketama) Add(nodes ...string) { _ = "STUB: not implemented"; return }

// Remove removes nodes from the ketama ring
func (k *Ketama) Remove(nodes ...string) { _ = "STUB: not implemented"; return }

// Get returns the node closest to key in the clockwise direction
func (k *Ketama) Get(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }
