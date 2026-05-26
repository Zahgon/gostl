package queue

import (
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/utils/sync"
)

var (
	defaultLocker sync.FakeLocker
)

// Options holds Queue's options
type Options[T any] struct {
	locker    sync.Locker
	container container.Container[T]
}

// Option is a function type used to set Options
type Option[T any] func(option *Options[T])

// WithGoroutineSafe is used to set a Queue goroutine-safe
func WithGoroutineSafe[T any]() Option[T] { _ = "STUB: not implemented"; return nil }

// WithContainer is used to set a Queue's underlying container
func WithContainer[T any](c container.Container[T]) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// WithListContainer is used to set List as a Queue's underlying container
func WithListContainer[T any]() Option[T] { _ = "STUB: not implemented"; return nil }

// Queue is a first-in-first-out data structure
type Queue[T any] struct {
	container container.Container[T]
	locker    sync.Locker
}

// New creates a new queue
func New[T any](opts ...Option[T]) *Queue[T] { _ = "STUB: not implemented"; return nil }

// Size returns the amount of elements in the queue
func (q *Queue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Empty returns true if the queue is empty, otherwise returns false
func (q *Queue[T]) Empty() bool { _ = "STUB: not implemented"; return false }

// Push pushes a value to the end of the queue
func (q *Queue[T]) Push(value T) { _ = "STUB: not implemented"; return }

// Front returns the front value in the queue
func (q *Queue[T]) Front() T { _ = "STUB: not implemented"; return *new(T) }

// Back returns the back value in the queue
func (q *Queue[T]) Back() T { _ = "STUB: not implemented"; return *new(T) }

// Pop removes the the front element in the queue, and returns its value
func (q *Queue[T]) Pop() T { _ = "STUB: not implemented"; return *new(T) }

// Clear clears all elements in the queue
func (q *Queue[T]) Clear() { _ = "STUB: not implemented"; return }

// String returns a string representation of the queue
func (q *Queue[T]) String() string { _ = "STUB: not implemented"; return "" }
