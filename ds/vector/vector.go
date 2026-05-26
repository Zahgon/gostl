package vector

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// Options holds the Vector's options
type Options struct {
	capacity int
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithCapacity is used to set the capacity of a Vector
func WithCapacity(capacity int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Vector is a linear data structure, the internal is a slice
type Vector[T any] struct {
	data []T
}

// New creates a new Vector
func New[T any](opts ...Option) *Vector[T] { _ = "STUB: not implemented"; return nil }

// NewFromVector news a Vector from other Vector
func NewFromVector[T any](other *Vector[T]) *Vector[T] { _ = "STUB: not implemented"; return nil }

// Size returns the size of the vector
func (v *Vector[T]) Size() int {
	_ = "STUB: not implemented"

	// Capacity returns the capacity of the vector
	return 0
}

func (v *Vector[T]) Capacity() int {
	_ = "STUB: not implemented"

	// Empty returns true if the vector is empty, otherwise returns false
	return 0
}

func (v *Vector[T]) Empty() bool { _ = "STUB: not implemented"; return false }

// PushBack pushes val to the back of the vector
func (v *Vector[T]) PushBack(val T) { _ = "STUB: not implemented"; return }

// SetAt sets the value val to the vector at position pos
func (v *Vector[T]) SetAt(pos int, val T) { _ = "STUB: not implemented"; return }

// InsertAt inserts the value val to the vector at position pos
func (v *Vector[T]) InsertAt(pos int, val T) { _ = "STUB: not implemented"; return }

// EraseAt erases the value at position pos
func (v *Vector[T]) EraseAt(pos int) { _ = "STUB: not implemented"; return }

// EraseIndexRange erases values at range[first, last)
func (v *Vector[T]) EraseIndexRange(first, last int) { _ = "STUB: not implemented"; return }

// At returns the value at position pos, returns nil if pos is out off range .
func (v *Vector[T]) At(pos int) T { _ = "STUB: not implemented"; return *new(T) }

// Front returns the first value in the vector, returns nil if the vector is empty.
func (v *Vector[T]) Front() T {
	_ = "STUB: not implemented"

	// Back returns the last value in the vector, returns nil if the vector is empty.
	return *new(T)
}

func (v *Vector[T]) Back() T {
	_ = "STUB: not implemented"
	return *

	// PopBack returns the last value of the vector and erase it, returns nil if the vector is empty.
	new(T)
}

func (v *Vector[T]) PopBack() T { _ = "STUB: not implemented"; return *new(T) }

// Reserve makes a new space for the vector with passed capacity
func (v *Vector[T]) Reserve(capacity int) { _ = "STUB: not implemented"; return }

// ShrinkToFit shrinks the capacity of the vector to the fit size
func (v *Vector[T]) ShrinkToFit() { _ = "STUB: not implemented"; return }

// Clear clears all data in the vector
func (v *Vector[T]) Clear() { _ = "STUB: not implemented"; return }

// Data returns internal data of the vector
func (v *Vector[T]) Data() []T {
	_ = "STUB: not implemented"

	// Begin returns the first iterator of the vector
	return nil
}

func (v *Vector[T]) Begin() *VectorIterator[T] {
	_ = "STUB: not implemented"

	// End returns the end iterator of the vector
	return nil
}

func (v *Vector[T]) End() *VectorIterator[T] { _ = "STUB: not implemented"; return nil }

// First returns the first iterator of the vector
func (v *Vector[T]) First() *VectorIterator[T] {
	_ = "STUB: not implemented"

	// Last returns the last iterator of the vector
	return nil
}

func (v *Vector[T]) Last() *VectorIterator[T] { _ = "STUB: not implemented"; return nil }

// IterAt  returns the iterator at position of the vector
func (v *Vector[T]) IterAt(pos int) *VectorIterator[T] { _ = "STUB: not implemented"; return nil }

// Insert inserts a value val to the vector at the position of the iterator iter point to
func (v *Vector[T]) Insert(iter iterator.ConstIterator[T], val T) *VectorIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Erase erases the element of the iterator iter point to
func (v *Vector[T]) Erase(iter iterator.ConstIterator[T]) *VectorIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// EraseRange erases all elements in the range[first, last)
func (v *Vector[T]) EraseRange(first, last iterator.ConstIterator[T]) *VectorIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Resize resizes the size of the vector to the passed size
func (v *Vector[T]) Resize(size int) { _ = "STUB: not implemented"; return }

// String returns a string representation of the vector
func (v *Vector[T]) String() string { _ = "STUB: not implemented"; return "" }
