package slice

// SliceWrapper wraps a slice in order to provide functions related to iterators
type SliceWrapper[T any] struct {
	slice []T
}

// NewSliceWrapper creates a SliceWrapper
func NewSliceWrapper[T any](slice []T) *SliceWrapper[T] { _ = "STUB: not implemented"; return nil }

// Attach update the internal slice to newSlice
func (s *SliceWrapper[T]) Attach(newSlice []T) {
	_ = "STUB: not implemented"

	// Len returns the length of s
	return
}

func (s *SliceWrapper[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// At returns the value at position
func (s *SliceWrapper[T]) At(position int) T { _ = "STUB: not implemented"; return *new(T) }

// Set sets value at position
func (s *SliceWrapper[T]) Set(position int, val T) { _ = "STUB: not implemented"; return }

// Begin returns the first iterator of s
func (s *SliceWrapper[T]) Begin() *SliceIterator[T] {
	_ = "STUB: not implemented"

	// End returns the end iterator of s
	return nil
}

func (s *SliceWrapper[T]) End() *SliceIterator[T] { _ = "STUB: not implemented"; return nil }

// First returns the first iterator of s
func (s *SliceWrapper[T]) First() *SliceIterator[T] { _ = "STUB: not implemented"; return nil }

// Last returns the last iterator of s
func (s *SliceWrapper[T]) Last() *SliceIterator[T] { _ = "STUB: not implemented"; return nil }
