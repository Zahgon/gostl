package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// SetIterator is an iterator implementation of set
type SetIterator[T any] struct {
	node *rbtree.Node[T, bool]
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *SetIterator[K]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Next moves the pointer of the iterator to the next node and returns itself
func (iter *SetIterator[T]) Next() iterator.ConstIterator[T] { _ = "STUB: not implemented"; return nil }

// Prev moves the pointer of the iterator to the previous node and returns itself
func (iter *SetIterator[T]) Prev() iterator.ConstBidIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Value returns the element of the iterator point to
func (iter *SetIterator[T]) Value() T {
	_ = "STUB: not implemented"
	return *

	// Clone clones the iterator into a new SetIterator
	new(T)
}

func (iter *SetIterator[T]) Clone() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SetIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	_ = "STUB: not implemented"
	return false
}
