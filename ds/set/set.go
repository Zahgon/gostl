package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// constants definition
const (
	Empty = true
)

var (
	defaultLocker sync.FakeLocker
)

// Options holds the Set's options
type Options struct {
	locker sync.Locker
}

// Option is a function  type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to set the set goroutine-safe
// Note that iterators are not goroutine safe, and it is useless to turn on the setting option here.
// so don't use iterators in multi goroutines
func WithGoroutineSafe() Option { _ = "STUB: not implemented"; return *new(Option) }

// Set uses RbTress for internal data structure, and every key can must bee unique.
type Set[T any] struct {
	tree   *rbtree.RbTree[T, bool]
	locker sync.Locker
	keyCmp comparator.Comparator[T]
}

// New creates a new set
func New[T any](cmp comparator.Comparator[T], opts ...Option) *Set[T] {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts an element to the set
func (s *Set[T]) Insert(element T) { _ = "STUB: not implemented"; return }

// Erase erases an element from the set
func (s *Set[T]) Erase(element T) { _ = "STUB: not implemented"; return }

// Find finds the element's node in the set, and return its iterator
func (s *Set[T]) Find(element T) *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// LowerBound finds the first element that equal or greater than the passed element in the set, and returns its iterator
func (s *Set[T]) LowerBound(element T) *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// UpperBound finds the first element that greater than the passed element in the set, and returns its iterator
func (s *Set[T]) UpperBound(element T) *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// Begin returns the iterator with the minimum element in the set
func (s *Set[T]) Begin() *SetIterator[T] {
	_ = "STUB: not implemented"

	// First returns the iterator with the minimum element in the set
	return nil
}

func (s *Set[T]) First() *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// Last returns the iterator with the maximum element in the set
func (s *Set[T]) Last() *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// Clear clears the set
func (s *Set[T]) Clear() { _ = "STUB: not implemented"; return }

// Contains returns true if the passed element is in the Set. otherwise returns false.
func (s *Set[T]) Contains(element T) bool { _ = "STUB: not implemented"; return false }

// Size returns the amount of element in the set
func (s *Set[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Traversal traversals elements in the set, it will not stop until to the end of the set or the visitor returns false
func (s *Set[T]) Traversal(visitor visitor.Visitor[T]) { _ = "STUB: not implemented"; return }

// String returns a string representation of the set
func (s *Set[T]) String() string { _ = "STUB: not implemented"; return "" }

// Intersect returns a new set with the common elements in the set s and the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }

// Union returns a new set with the all elements in the set s and the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set[T]) Union(other *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }

// Diff returns a new set with the elements in the set s but not in the passed set
// Please ensure s set and other set uses the same keyCmp
func (s *Set[T]) Diff(other *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }
