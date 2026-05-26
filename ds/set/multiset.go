package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// MultiSet uses RbTress for internal data structure, and keys can bee repeated.
type MultiSet[T any] struct {
	tree   *rbtree.RbTree[T, bool]
	locker sync.Locker
}

// NewMultiSet creates a new MultiSet
func NewMultiSet[T any](cmp comparator.Comparator[T], opts ...Option) *MultiSet[T] {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts an element to the MultiSet
func (ms *MultiSet[T]) Insert(element T) { _ = "STUB: not implemented"; return }

// Erase erases the first node with passed element in the MultiSet
func (ms *MultiSet[T]) Erase(element T) { _ = "STUB: not implemented"; return }

// Erase erases all node with passed element in the MultiSet
func (ms *MultiSet[T]) EraseAll(element T) { _ = "STUB: not implemented"; return }

// Find finds the first element that is equal to the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) Find(element T) *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// LowerBound finds the first element that is equal to or greater than the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) LowerBound(element T) *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// UpperBound finds the first element that is greater than the passed element in the MultiSet, and returns its iterator
func (ms *MultiSet[T]) UpperBound(element T) *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// Begin returns the iterator with the minimum element in the MultiSet
func (ms *MultiSet[T]) Begin() *SetIterator[T] {
	_ = "STUB: not implemented"

	// First returns the iterator with the minimum element in the MultiSet
	return nil
}

func (ms *MultiSet[T]) First() *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// Last returns the iterator with the maximum element in the MultiSet
func (ms *MultiSet[T]) Last() *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// Count returns the amount of elements that are equal to the passed element in the MultiSet
func (ms *MultiSet[T]) Count(element T) int { _ = "STUB: not implemented"; return 0 }

// Clear clears all elements in the MultiSet
func (ms *MultiSet[T]) Clear() { _ = "STUB: not implemented"; return }

// Contains returns true if the passed element is in the MultiSet. otherwise returns false.
func (ms *MultiSet[T]) Contains(element T) bool { _ = "STUB: not implemented"; return false }

// Size returns the amount of elements in the MultiSet
func (ms *MultiSet[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Traversal traversals elements in the MultiSet, it will not stop until to the end of the MultiSet or the visitor returns false
func (ms *MultiSet[T]) Traversal(visitor visitor.Visitor[T]) { _ = "STUB: not implemented"; return }

// String returns s string representation of the MultiSet
func (ms *MultiSet[T]) String() string { _ = "STUB: not implemented"; return "" }
