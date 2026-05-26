package deque

import (
	"errors"
)

// Constants definition
const (
	SegmentCapacity = 128
)

// Define internal errors
var (
	ErrOutOfRange = errors.New("out off range")
)

// Deque is double-ended queue supports efficient data insertion from the head and tail, random access and iterator access.
type Deque[T any] struct {
	pool  *Pool[T]
	segs  []*Segment[T]
	begin int
	end   int
	size  int
}

// New creates a new deque
func New[T any]() *Deque[T] { _ = "STUB: not implemented"; return nil }

// Size returns the amount of values in the deque
func (d *Deque[T]) Size() int {
	_ = "STUB: not implemented"

	// Empty returns true if the deque is empty,otherwise returns false.
	return 0
}

func (d *Deque[T]) Empty() bool { _ = "STUB: not implemented"; return false }

func (d *Deque[T]) segUsed() int { _ = "STUB: not implemented"; return 0 }

// PushFront pushed a value to the front of the deque
func (d *Deque[T]) PushFront(value T) { _ = "STUB: not implemented"; return }

// PushBack pushed a value to the back of deque
func (d *Deque[T]) PushBack(value T) { _ = "STUB: not implemented"; return }

// Insert inserts a value to the position pos of the deque
func (d *Deque[T]) Insert(pos int, value T) { _ = "STUB: not implemented"; return }

// seg is closer to the front

// seg is closer to the back

func (d *Deque[T]) moveFrontInsert(seg, pos int, value T) { _ = "STUB: not implemented"; return }

func (d *Deque[T]) moveBackInsert(seg, pos int, value T) {
	_ = "STUB: not implemented"
	// move back
	return
}

// Front returns the value at the first position of the deque
func (d *Deque[T]) Front() T { _ = "STUB: not implemented"; return *new(T) }

// Back returns the value at the last position of the deque
func (d *Deque[T]) Back() T { _ = "STUB: not implemented"; return *new(T) }

// At returns the value at position pos of the deque
func (d *Deque[T]) At(pos int) T { _ = "STUB: not implemented"; return *new(T) }

// Set sets the value of the deque's position pos with value val
func (d *Deque[T]) Set(pos int, val T) error { _ = "STUB: not implemented"; return nil }

// PopFront returns the value at the first position of the deque and removes it
func (d *Deque[T]) PopFront() T { _ = "STUB: not implemented"; return *new(T) }

// PopBack returns the value at the lase position of the deque and removes it
func (d *Deque[T]) PopBack() T { _ = "STUB: not implemented"; return *new(T) }

// EraseAt erases the element at the position pos
func (d *Deque[T]) EraseAt(pos int) { _ = "STUB: not implemented"; return }

// EraseRange erases elements in range [firstPos, lastPos)
func (d *Deque[T]) EraseRange(firstPos, lastPos int) { _ = "STUB: not implemented"; return }

// move back

// move front

// Clear erases all elements in the deque
func (d *Deque[T]) Clear() { _ = "STUB: not implemented"; return }

func (d *Deque[T]) putToPool(s *Segment[T]) { _ = "STUB: not implemented"; return }

func (d *Deque[T]) firstAvailableSegment() *Segment[T] { _ = "STUB: not implemented"; return nil }

func (d *Deque[T]) lastAvailableSegment() *Segment[T] { _ = "STUB: not implemented"; return nil }

func (d *Deque[T]) firstSegment() *Segment[T] { _ = "STUB: not implemented"; return nil }

func (d *Deque[T]) lastSegment() *Segment[T] { _ = "STUB: not implemented"; return nil }

func (d *Deque[T]) segmentAt(seg int) *Segment[T] { _ = "STUB: not implemented"; return nil }

func (d *Deque[T]) pos(position int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (d *Deque[T]) expand() { _ = "STUB: not implemented"; return }

// shrinkIfNeeded shrinks the deque if it has too many unused space.
func (d *Deque[T]) shrinkIfNeeded() { _ = "STUB: not implemented"; return }

func (d *Deque[T]) nextIndex(index int) int { _ = "STUB: not implemented"; return 0 }

func (d *Deque[T]) preIndex(index int) int { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of the deque
func (d *Deque[T]) String() string { _ = "STUB: not implemented"; return "" }

// Begin returns an iterator of the deque with the first position
func (d *Deque[T]) Begin() *DequeIterator[T] {
	_ = "STUB: not implemented"

	// End returns an iterator of the deque with the position d.Size()
	return nil
}

func (d *Deque[T]) End() *DequeIterator[T] { _ = "STUB: not implemented"; return nil }

// First returns an iterator of the deque with the first position
func (d *Deque[T]) First() *DequeIterator[T] {
	_ = "STUB: not implemented"

	// Last returns an iterator of the deque with the last position
	return nil
}

func (d *Deque[T]) Last() *DequeIterator[T] { _ = "STUB: not implemented"; return nil }

// IterAt returns an iterator of the deque with the position pos
func (d *Deque[T]) IterAt(pos int) *DequeIterator[T] { _ = "STUB: not implemented"; return nil }
