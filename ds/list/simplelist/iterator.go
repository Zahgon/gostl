package simplelist

import "github.com/liyue201/gostl/utils/iterator"

// ListIterator is an implementation of Iterator
var _ iterator.Iterator[any] = (*ListIterator[any])(nil)

// ListIterator is an iterator for list
type ListIterator[T any] struct {
	node *Node[T]
}

// NewIterator news a ListIterator
func NewIterator[T any](node *Node[T]) *ListIterator[T] { _ = "STUB: not implemented"; return nil }

// IsValid returns whether iter is valid
func (iter *ListIterator[T]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Next returns the next iterator
func (iter *ListIterator[T]) Next() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Value returns the internal value of iter
func (iter *ListIterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

// SetValue sets the value of iter
func (iter *ListIterator[T]) SetValue(value T) { _ = "STUB: not implemented"; return }

// Clone clones iter to a new ListIterator
func (iter *ListIterator[T]) Clone() iterator.ConstIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Equal returns whether iter is equal to other
func (iter *ListIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	_ = "STUB: not implemented"
	return false
}
