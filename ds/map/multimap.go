package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// MultiMap uses RbTress for internal data structure, and keys can bee repeated.
type MultiMap[K, V any] struct {
	tree   *rbtree.RbTree[K, V]
	locker sync.Locker
}

// NewMultiMap creates a new MultiMap
func NewMultiMap[K, V any](cmp comparator.Comparator[K], opts ...Option) *MultiMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts a key-value to the MultiMap
func (mm *MultiMap[K, V]) Insert(key K, value V) { _ = "STUB: not implemented"; return }

// Get returns the first node's value by the passed key if the key is in the MultiMap, otherwise returns nil
func (mm *MultiMap[K, V]) Get(key K) (V, error) { _ = "STUB: not implemented"; return *new(V), nil }

// Erase erases the key in the MultiMap
func (mm *MultiMap[K, V]) Erase(key K) { _ = "STUB: not implemented"; return }

// Find finds the node by the passed key in the MultiMap and returns its iterator
func (mm *MultiMap[K, V]) Find(key K) *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// LowerBound find the first node that its key is equal or greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap[K, V]) LowerBound(key K) *MapIterator[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// UpperBound find the first node that its key is greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap[K, V]) UpperBound(key K) *MapIterator[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Begin returns the first node's iterator
func (mm *MultiMap[K, V]) Begin() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// First returns the first node's iterator
func (mm *MultiMap[K, V]) First() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// Last returns the last node's iterator
func (mm *MultiMap[K, V]) Last() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// Clear clears the MultiMap
func (mm *MultiMap[K, V]) Clear() { _ = "STUB: not implemented"; return }

// Contains returns true if the passed value is in the MultiMap. otherwise returns false.
func (mm *MultiMap[K, V]) Contains(key K) bool { _ = "STUB: not implemented"; return false }

// Size returns the amount of elements in the MultiMap
func (mm *MultiMap[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

// Traversal traversals elements in the MultiMap, it will not stop until to the end of the MultiMap or the visitor returns false
func (mm *MultiMap[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	_ = "STUB: not implemented"
	return
}
