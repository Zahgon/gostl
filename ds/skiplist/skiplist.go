package skiplist

import (
	"errors"
	"math/rand"

	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

var (
	defaultMaxLevel = 10
	defaultLocker   sync.FakeLocker
)
var ErrorNotFound = errors.New("not found")

// Options holds Skiplist's options
type Options struct {
	maxLevel int
	locker   sync.Locker
}

// Option is a function used to set Options
type Option func(option *Options)

// WithGoroutineSafe sets Skiplist goroutine-safety,
func WithGoroutineSafe() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxLevel sets max level of Skiplist
func WithMaxLevel(maxLevel int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Node is a list node
type Node[K, V any] struct {
	next []*Element[K, V]
}

// Element is a kind of node with key-value data
type Element[K, V any] struct {
	Node[K, V]
	key   K
	value V
}

// Skiplist is a kind of data structure which can search quickly by exchanging space for time
type Skiplist[K, V any] struct {
	locker         sync.Locker
	head           Node[K, V]
	maxLevel       int
	keyCmp         comparator.Comparator[K]
	len            int
	prevNodesCache []*Node[K, V]
	rander         *rand.Rand
}

// New news a Skiplist
func New[K, V any](cmp comparator.Comparator[K], opts ...Option) *Skiplist[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts a key-value pair into the skiplist
func (sl *Skiplist[K, V]) Insert(key K, value V) { _ = "STUB: not implemented"; return }

//same key, update value

// Get returns the value associated with the passed key if the key is in the skiplist, otherwise returns error
func (sl *Skiplist[K, V]) Get(key K) (V, error) { _ = "STUB: not implemented"; return *new(V), nil }

// Remove removes the key-value pair associated with the passed key and returns true if the key is in the skiplist, otherwise returns false
func (sl *Skiplist[K, V]) Remove(key K) bool { _ = "STUB: not implemented"; return false }

// Len returns the amount of key-value pair in the skiplist
func (sl *Skiplist[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (sl *Skiplist[K, V]) randomLevel() int { _ = "STUB: not implemented"; return 0 }

// 2^n-1

func (sl *Skiplist[K, V]) findPrevNodes(key K) []*Node[K, V] { _ = "STUB: not implemented"; return nil }

// Traversal traversals elements in the skiplist, it will stop until to the end or the visitor returns false
func (sl *Skiplist[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	_ = "STUB: not implemented"
	return
}

// Keys returns all keys in the skiplist
func (sl *Skiplist[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }
