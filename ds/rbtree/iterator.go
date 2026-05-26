package rbtree

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// RbTreeIterator is an iterator implementation of RbTree
type RbTreeIterator[K, V any] struct {
	node *Node[K, V]
}

// NewIterator creates a RbTreeIterator from the passed node
func NewIterator[K, V any](node *Node[K, V]) *RbTreeIterator[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *RbTreeIterator[K, V]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *RbTreeIterator[K, V]) Next() iterator.ConstIterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Prev moves the pointer of the iterator to the previous node, and returns itself
func (iter *RbTreeIterator[K, V]) Prev() iterator.ConstBidIterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Key returns the node's key of the iterator point to
func (iter *RbTreeIterator[K, V]) Key() K {
	_ = "STUB: not implemented"
	return *

	// Value returns the node's value of the iterator point to
	new(K)
}

func (iter *RbTreeIterator[K, V]) Value() V {
	_ = "STUB: not implemented"
	return *

	// SetValue sets the node's value of the iterator point to
	new(V)
}

func (iter *RbTreeIterator[K, V]) SetValue(val V) error { _ = "STUB: not implemented"; return nil }

// Clone clones the iterator into a new RbTreeIterator
func (iter *RbTreeIterator[K, V]) Clone() iterator.ConstIterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *RbTreeIterator[K, V]) Equal(other iterator.ConstIterator[V]) bool {
	_ = "STUB: not implemented"
	return false
}
