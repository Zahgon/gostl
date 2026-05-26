package priorityqueue

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
)

var (
	defaultLocker sync.FakeLocker
)

// ElementHolder holds elements of the PriorityQueue
type ElementHolder[T any] struct {
	elements []T
	cmpFun   comparator.Comparator[T]
}

// Push pushes an element to the ElementHolder
func (h *ElementHolder[T]) Push(element T) { _ = "STUB: not implemented"; return }

// Pop pops an element from the ElementHolder
func (h *ElementHolder[T]) Pop() T { _ = "STUB: not implemented"; return *new(T) }

func (h *ElementHolder[T]) top() T { _ = "STUB: not implemented"; return *new(T) }

// Len returns the amount of elements in ElementHolder
func (h *ElementHolder[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// Len compare two elements at position i and j , and returns true if elements[i] < elements[j]
func (h *ElementHolder[T]) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap swaps two elements at position i and j
func (h *ElementHolder[T]) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Options holds PriorityQueue's options
type Options struct {
	locker sync.Locker
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set the PriorityQueue goroutine-safe
func WithGoroutineSafe() Option { _ = "STUB: not implemented"; return *new(Option) }

// PriorityQueue is an implementation of priority queue
type PriorityQueue[T any] struct {
	holder *ElementHolder[T]
	locker sync.Locker
}

// New creates a PriorityQueue
func New[T any](cmp comparator.Comparator[T], opts ...Option) *PriorityQueue[T] {
	_ = "STUB: not implemented"
	return nil
}

// Push pushes an element to the PriorityQueue
func (q *PriorityQueue[T]) Push(e T) { _ = "STUB: not implemented"; return }

// Pop pops an element from the PriorityQueue
func (q *PriorityQueue[T]) Pop() T { _ = "STUB: not implemented"; return *new(T) }

// Top returns the top element in the PriorityQueue
func (q *PriorityQueue[T]) Top() T { _ = "STUB: not implemented"; return *new(T) }

// Empty returns true if the PriorityQueue is empty, otherwise returns false
func (q *PriorityQueue[T]) Empty() bool { _ = "STUB: not implemented"; return false }

// Clear clears all elements in the priority queue
func (q *PriorityQueue[T]) Clear() { _ = "STUB: not implemented"; return }

// reset cap to zero

// Size returns the amount of elements in the queue
func (q *PriorityQueue[T]) Size() int { _ = "STUB: not implemented"; return 0 }
