package stack

import (
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/utils/sync"
)

var (
	defaultLocker sync.FakeLocker
)

// Options holds the Stack's options
type Options[T any] struct {
	locker    sync.Locker
	container container.Container[T]
}

// Option is a function type used to set Options
type Option[T any] func(option *Options[T])

// WithGoroutineSafe is used to set a stack goroutine-safe
func WithGoroutineSafe[T any]() Option[T] { _ = "STUB: not implemented"; return nil }

// WithContainer is used to set a stack's underlying container
func WithContainer[T any](c container.Container[T]) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// WithListContainer is used to set List for a stack's underlying container
func WithListContainer[T any]() Option[T] { _ = "STUB: not implemented"; return nil }

// Stack is a last-in-first-out data structure
type Stack[T any] struct {
	container container.Container[T]
	locker    sync.Locker
}

// New creates a new stack
func New[T any](opts ...Option[T]) *Stack[T] { _ = "STUB: not implemented"; return nil }

// Size returns the amount of elements in the stack
func (s *Stack[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Empty returns true if the stack is empty, otherwise returns false
func (s *Stack[T]) Empty() bool { _ = "STUB: not implemented"; return false }

// Push pushes a value to the stack
func (s *Stack[T]) Push(value T) { _ = "STUB: not implemented"; return }

// Top returns the top value in the stack
func (s *Stack[T]) Top() T { _ = "STUB: not implemented"; return *new(T) }

// Pop removes the top value in the stack and returns it
func (s *Stack[T]) Pop() T { _ = "STUB: not implemented"; return *new(T) }

// Clear clears all elements in the stack
func (s *Stack[T]) Clear() { _ = "STUB: not implemented"; return }

// String returns a string representation of the stack
func (s *Stack[T]) String() string { _ = "STUB: not implemented"; return "" }
