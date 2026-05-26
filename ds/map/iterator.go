package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// MapIterator is a map iterator
type MapIterator[K, V any] struct {
	node *rbtree.Node[K, V]
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *MapIterator[K, V]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *MapIterator[K, V]) Next() iterator.ConstIterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Prev moves the pointer of the iterator to the previous node, and returns itseft
func (iter *MapIterator[K, V]) Prev() iterator.ConstBidIterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Key returns the node's key of the iterator point to
func (iter *MapIterator[K, V]) Key() K {
	_ = "STUB: not implemented"
	return *

	// Value returns the node's value of the iterator point to
	new(K)
}

func (iter *MapIterator[K, V]) Value() V {
	_ = "STUB: not implemented"
	return *

	// SetValue sets the node's value of the iterator point to
	new(V)
}

func (iter *MapIterator[K, V]) SetValue(val V) { _ = "STUB: not implemented"; return }

// Clone clones the iterator to a new MapIterator
func (iter *MapIterator[K, V]) Clone() iterator.ConstIterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *MapIterator[K, V]) Equal(other iterator.ConstIterator[V]) bool {
	_ = "STUB: not implemented"
	return false
}
