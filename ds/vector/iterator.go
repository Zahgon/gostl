package vector

import (
	"github.com/liyue201/gostl/utils/iterator"
)

type T any

// ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator[T] = (*VectorIterator[T])(nil)

// VectorIterator represents a vector iterator
type VectorIterator[T any] struct {
	vec      *Vector[T]
	position int // the position of iterator point to
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *VectorIterator[T]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Value returns the value of the iterator point to
func (iter *VectorIterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

// SetValue sets the value of the iterator point to
func (iter *VectorIterator[T]) SetValue(val T) { _ = "STUB: not implemented"; return }

// Next moves the position of iterator to the next position and returns itself
func (iter *VectorIterator[T]) Next() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Prev moves the position of the iterator to the previous position and returns itself
func (iter *VectorIterator[T]) Prev() iterator.ConstBidIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Clone clones the iterator into a new iterator
func (iter *VectorIterator[T]) Clone() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// IteratorAt creates an iterator with the passed position
func (iter *VectorIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Position return the position of the iterator point to
func (iter *VectorIterator[T]) Position() int { _ = "STUB: not implemented"; return 0 }

// Equal returns true if the iterator is equal to the passed iterator
func (iter *VectorIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	_ = "STUB: not implemented"
	return false
}
