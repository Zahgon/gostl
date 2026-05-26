package slice

import "github.com/liyue201/gostl/utils/iterator"

type T any

// SliceIterator is an implementation of RandomAccessIterator

var _ iterator.RandomAccessIterator[T] = (*SliceIterator[T])(nil)

// SliceIterator represents a slice iterator
type SliceIterator[T any] struct {
	s        ISlice[T]
	position int
}

// IsValid returns trus if the iterator is valid, othterwise return false
func (iter *SliceIterator[T]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Value returns the value of the iterator point to
func (iter *SliceIterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

// SetValue sets the value of the iterator point to
func (iter *SliceIterator[T]) SetValue(val T) { _ = "STUB: not implemented"; return }

// Next moves the iterator's position to the next position, and returns itself
func (iter *SliceIterator[T]) Next() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Prev move the iterator's position to the previous position, and return itself
func (iter *SliceIterator[T]) Prev() iterator.ConstBidIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Clone clones the iterator into a new one
func (iter *SliceIterator[T]) Clone() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// IteratorAt creates an iterator with the passed position
func (iter *SliceIterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Position returns the position of the iterator
func (iter *SliceIterator[T]) Position() int { _ = "STUB: not implemented"; return 0 }

// Equal returns true if the iterator is equal to the passed iterator
func (iter *SliceIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	_ = "STUB: not implemented"
	return false
}
