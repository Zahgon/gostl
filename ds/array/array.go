package array

// Array is a fixed size slice
type Array[T any] struct {
	values []T
}

// New creates a new array with passed size
func New[T any](size int) *Array[T] { _ = "STUB: not implemented"; return nil }

// NewFromArray creates a new array from another array, and copy its values
func NewFromArray[T any](other *Array[T]) *Array[T] { _ = "STUB: not implemented"; return nil }

// Fill fills Array a with value val
func (a *Array[T]) Fill(val T) { _ = "STUB: not implemented"; return }

// Set sets value val to the position pos of the array
func (a *Array[T]) Set(pos int, val T) { _ = "STUB: not implemented"; return }

// At returns the value at position pos in the array
func (a *Array[T]) At(pos int) T { _ = "STUB: not implemented"; return *new(T) }

// Front returns the first value in the array
func (a *Array[T]) Front() T {
	_ = "STUB: not implemented"

	// Back returns the last value in the array
	return *new(T)
}

func (a *Array[T]) Back() T { _ = "STUB: not implemented"; return *new(T) }

// Size returns number of elements within the array
func (a *Array[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Empty returns whether the array is empty or not
func (a *Array[T]) Empty() bool { _ = "STUB: not implemented"; return false }

// SwapArray swaps the values of two arrays
func (a *Array[T]) SwapArray(other *Array[T]) { _ = "STUB: not implemented"; return }

// Data returns the internal values of the array
func (a *Array[T]) Data() []T {
	_ = "STUB: not implemented"

	// Begin returns an iterator of the array with the first position
	return nil
}

func (a *Array[T]) Begin() *ArrayIterator[T] {
	_ = "STUB: not implemented"

	// End returns an iterator of the array with the position a.Size()
	return nil
}

func (a *Array[T]) End() *ArrayIterator[T] { _ = "STUB: not implemented"; return nil }

// First returns an iterator of the array with the first position
func (a *Array[T]) First() *ArrayIterator[T] {
	_ = "STUB: not implemented"

	// Last returns an iterator of the array with the last position
	return nil
}

func (a *Array[T]) Last() *ArrayIterator[T] { _ = "STUB: not implemented"; return nil }

// IterAt returns an iterator of the array with position pos
func (a *Array[T]) IterAt(pos int) *ArrayIterator[T] { _ = "STUB: not implemented"; return nil }

// String returns a string representation of the array
func (a *Array[T]) String() string { _ = "STUB: not implemented"; return "" }
