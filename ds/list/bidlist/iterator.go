package bidlist

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// ListIterator is an implementation of BidIterator
var _ iterator.BidIterator[T] = (*ListIterator[T])(nil)

// ListIterator is an implementation of list iterator
type ListIterator[T any] struct {
	node *Node[T]
}

// NewIterator creates a ListIterator
func NewIterator[T any](node *Node[T]) *ListIterator[T] { _ = "STUB: not implemented"; return nil }

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *ListIterator[T]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Next moves the pointer of iterator to the next node and returns itself
func (iter *ListIterator[T]) Next() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Prev moves the pointer of iterator to the previous node and returns itself
func (iter *ListIterator[T]) Prev() iterator.ConstBidIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Value returns the node's value of the iterator point to
func (iter *ListIterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

// SetValue sets the node's value of the iterator point to
func (iter *ListIterator[T]) SetValue(value T) { _ = "STUB: not implemented"; return }

// Clone clones the iterator to a new iterator
func (iter *ListIterator[T]) Clone() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ListIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	_ = "STUB: not implemented"
	return false
}
