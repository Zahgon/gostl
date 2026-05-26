package deque

// Segment is a fixed capacity ring
type Segment[T any] struct {
	data  []T
	begin int
	end   int
	nSize int
}

func newSegment[T any](capacity int) *Segment[T] { _ = "STUB: not implemented"; return nil }

func (s *Segment[T]) pushBack(value T) { _ = "STUB: not implemented"; return }

func (s *Segment[T]) pushFront(val T) { _ = "STUB: not implemented"; return }

func (s *Segment[T]) insert(position int, value T) { _ = "STUB: not implemented"; return }

//move the front pos items

//move the back pos items

func (s *Segment[T]) popBack() T { _ = "STUB: not implemented"; return *new(T) }

//s.data[s.end] = nil

func (s *Segment[T]) popFront() T {
	_ = "STUB: not implemented"
	return *

	// s.data[s.begin] = nil
	new(T)
}

func (s *Segment[T]) eraseAt(position int) { _ = "STUB: not implemented"; return }

//s.data[s.begin] = nil

//s.data[s.preIndex(s.end)] = nil

func (s *Segment[T]) size() int { _ = "STUB: not implemented"; return 0 }

func (s *Segment[T]) capacity() int { _ = "STUB: not implemented"; return 0 }

func (s *Segment[T]) full() bool { _ = "STUB: not implemented"; return false }

func (s *Segment[T]) empty() bool { _ = "STUB: not implemented"; return false }

func (s *Segment[T]) nextIndex(index int) int { _ = "STUB: not implemented"; return 0 }

func (s *Segment[T]) preIndex(index int) int { _ = "STUB: not implemented"; return 0 }

func (s *Segment[T]) at(position int) T { _ = "STUB: not implemented"; return *new(T) }

//return nil

func (s *Segment[T]) set(position int, val T) { _ = "STUB: not implemented"; return }

func (s *Segment[T]) back() T { _ = "STUB: not implemented"; return *new(T) }

func (s *Segment[T]) front() T { _ = "STUB: not implemented"; return *new(T) }

func (s *Segment[T]) clear() { _ = "STUB: not implemented"; return }

//s.data[i] = nil
