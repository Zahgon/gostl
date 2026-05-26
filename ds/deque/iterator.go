package deque

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// DequeIterator is an implementation of RandomAccessIterator
type T any

var _ iterator.RandomAccessIterator[T] = (*DequeIterator[T])(nil)

// DequeIterator is an implementation of Deque iterator
type DequeIterator[T any] struct {
	dq       *Deque[T]
	position int
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *DequeIterator[T]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Value returns the value of the deque at the position of the iterator point to
func (iter *DequeIterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

// SetValue sets the value of the deque at the position of the iterator point to
func (iter *DequeIterator[T]) SetValue(val T) { _ = "STUB: not implemented"; return }

// Next moves the position of the iterator to the next position and returns itself
func (iter *DequeIterator[T]) Next() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *DequeIterator[T]) Prev() iterator.ConstBidIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Clone clones the iterator to a new iterator
func (iter *DequeIterator[T]) Clone() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// IteratorAt creates a new iterator with the passed position
func (iter *DequeIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Position returns the position of iterator
func (iter *DequeIterator[T]) Position() int { _ = "STUB: not implemented"; return 0 }

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *DequeIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	_ = "STUB: not implemented"
	return false
}
