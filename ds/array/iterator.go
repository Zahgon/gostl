package array

import (
	"github.com/liyue201/gostl/utils/iterator"
)

type T any

// ArrayIterator is an implementation of RandomAccessIterator
var _ iterator.RandomAccessIterator[T] = (*ArrayIterator[T])(nil)

// ArrayIterator is an implementation of Array iterator
type ArrayIterator[T any] struct {
	array    *Array[T]
	position int
}

// IsValid returns true if  the iterator is valid, otherwise returns false
func (iter *ArrayIterator[T]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Value returns the value of array at the position of the iterator point to
func (iter *ArrayIterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

// SetValue sets the value of the array at the position of the iterator point to
func (iter *ArrayIterator[T]) SetValue(val T) { _ = "STUB: not implemented"; return }

// Next moves the position of iterator to the next position and returns itself
func (iter *ArrayIterator[T]) Next() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Prev moves the position of iterator to the previous position and returns itself
func (iter *ArrayIterator[T]) Prev() iterator.ConstBidIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Clone clones the iterator to a new iterator
func (iter *ArrayIterator[T]) Clone() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// IteratorAt creates a new iterator with position pos
func (iter *ArrayIterator[T]) IteratorAt(pos int) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Position returns the position of the iterator
func (iter *ArrayIterator[T]) Position() int { _ = "STUB: not implemented"; return 0 }

// Equal returns true if the iterator is equal to the passed iterator, otherwise returns false
func (iter *ArrayIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	_ = "STUB: not implemented"
	return false
}
