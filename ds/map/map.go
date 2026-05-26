package treemap

import (
	"errors"

	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

var (
	defaultLocker sync.FakeLocker
)

var ErrorNotFound = errors.New("not found")

// Options holds Map's options
type Options struct {
	locker sync.Locker
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set a map goroutine-safe
// Note that iterators are not goroutine safe, and it is useless to turn on the setting option here.
// so don't use iterator in multi goroutines
func WithGoroutineSafe() Option { _ = "STUB: not implemented"; return *new(Option) }

// Map uses RbTress for internal data structure, and every key can must bee unique.
type Map[K, V any] struct {
	tree   *rbtree.RbTree[K, V]
	locker sync.Locker
}

// New creates a new map
func New[K, V any](cmp comparator.Comparator[K], opts ...Option) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts a key-value to the map
func (m *Map[K, V]) Insert(key K, value V) { _ = "STUB: not implemented"; return }

// Get returns the value of the passed key if the key is in the map, otherwise returns nil
func (m *Map[K, V]) Get(key K) (V, error) { _ = "STUB: not implemented"; return *new(V), nil }

// Erase erases the node by the passed key from the map if the key in the Map
func (m *Map[K, V]) Erase(key K) { _ = "STUB: not implemented"; return }

// EraseIter erases the node that iterator iter point to from the map
func (m *Map[K, V]) EraseIter(iter iterator.ConstKvIterator[K, V]) {
	_ = "STUB: not implemented"
	return
}

// Find finds a node by the passed key and returns its iterator
func (m *Map[K, V]) Find(key K) *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// LowerBound finds a node that its key is equal or greater than the passed key and returns its iterator
func (m *Map[K, V]) LowerBound(key K) *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// UpperBound finds a node that its key is greater than the passed key and returns its iterator
func (m *Map[K, V]) UpperBound(key K) *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// Begin returns the first node's iterator
func (m *Map[K, V]) Begin() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// First returns the first node's iterator
func (m *Map[K, V]) First() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// Last returns the last node's iterator
func (m *Map[K, V]) Last() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// Clear clears the map
func (m *Map[K, V]) Clear() { _ = "STUB: not implemented"; return }

// Contains returns true if the key is in the map. otherwise returns false.
func (m *Map[K, V]) Contains(key K) bool { _ = "STUB: not implemented"; return false }

// Size returns the amount of elements in the map
func (m *Map[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

// Traversal traversals elements in the map, it will not stop until to the end or the visitor returns false
func (m *Map[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) { _ = "STUB: not implemented"; return }
