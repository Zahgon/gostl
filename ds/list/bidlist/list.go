package bidlist

import (
	"errors"

	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/utils/visitor"
)

// List is an implementation of Container
type T any

var ErrorOutOfRange = errors.New("out of range")

var _ container.Container[T] = (*List[T])(nil)

// Node is a list node
type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	Value T
	list  *List[T]
}

// Next returns the next list node or nil.
func (n *Node[T]) Next() *Node[T] { _ = "STUB: not implemented"; return nil }

// Prev returns the previous list node or nil.
func (n *Node[T]) Prev() *Node[T] { _ = "STUB: not implemented"; return nil }

// List represents a bidirectional list:
//
//	head -> node1 -- node2 --  node3
//	         |                   |
//	        node6 -- node5 --  node4
type List[T any] struct {
	head *Node[T] // point to the front Node
	len  int      // current list length
}

// New creates a list
func New[T any]() *List[T] { _ = "STUB: not implemented"; return nil }

// Len returns the amount of list nodes.
func (l *List[T]) Len() int {
	_ = "STUB: not implemented"

	// Size returns the amount of list nodes.
	return 0
}

func (l *List[T]) Size() int {
	_ = "STUB: not implemented"

	// Empty returns true if the list is empty
	return 0
}

func (l *List[T]) Empty() bool {
	_ = "STUB: not implemented"

	// FrontNode returns the front node of the list or nil if the list is empty
	return false
}

func (l *List[T]) FrontNode() *Node[T] {
	_ = "STUB: not implemented"

	// BackNode returns the last node of the list or nil if the list is empty
	return nil
}

func (l *List[T]) BackNode() *Node[T] { _ = "STUB: not implemented"; return nil }

// Front returns the value of the front node
func (l *List[T]) Front() T { _ = "STUB: not implemented"; return *new(T) }

// Back returns the value of the last node
func (l *List[T]) Back() T { _ = "STUB: not implemented"; return *new(T) }

// PushBack inserts a new node n with value v at the back of the list
func (l *List[T]) PushBack(v T) {
	_ = "STUB: not implemented"

	// PushBack inserts a new node n with value v at the back of the list and returns n.
	return
}

func (l *List[T]) pushBack(v T) *Node[T] { _ = "STUB: not implemented"; return nil }

// PushFront inserts a new node n with value v at the front of the list.
func (l *List[T]) PushFront(v T) { _ = "STUB: not implemented"; return }

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Node[T]) *Node[T] { _ = "STUB: not implemented"; return nil }

// InsertBefore inserts a new node n with value v immediately before mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(v T, mark *Node[T]) *Node[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] { _ = "STUB: not implemented"; return nil }

// Remove removes n from l list if n is a node of l list.
// It returns the n value n.Value.
// The node must not be nil.
func (l *List[T]) Remove(n *Node[T]) T { _ = "STUB: not implemented"; return *new(T) }

func (l *List[T]) remove(n *Node[T]) *Node[T] { _ = "STUB: not implemented"; return nil }

// avoid memory leaks
// avoid memory leaks

// Clear removes all nodes
func (l *List[T]) Clear() { _ = "STUB: not implemented"; return }

// PopBack removes the last node in the list and returns its value
func (l *List[T]) PopBack() T { _ = "STUB: not implemented"; return *new(T) }

// PopFront removes the first node in the list and returns its value
func (l *List[T]) PopFront() T { _ = "STUB: not implemented"; return *new(T) }

// MoveToFront moves node n to the front of the list.
// If n is not a node of the list, the list is not modified.
// The n must not be nil.
func (l *List[T]) MoveToFront(n *Node[T]) { _ = "STUB: not implemented"; return }

// MoveToBack moves node  n to the back of the list.
// If e is not a node of the list, the list is not modified.
// The node must not be nil.
func (l *List[T]) MoveToBack(n *Node[T]) { _ = "STUB: not implemented"; return }

// MoveAfter moves node n to its new position after mark.
// If n or mark is not a node of the list, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List[T]) MoveAfter(n, mark *Node[T]) { _ = "STUB: not implemented"; return }

func (l *List[T]) moveToAfter(n, at *Node[T]) { _ = "STUB: not implemented"; return }

// PushBackList inserts a copy of an other list at the back of the list.
// The list and other may be the same. They must not be nil.
func (l *List[T]) PushBackList(other *List[T]) { _ = "STUB: not implemented"; return }

// PushFrontList inserts a copy of another list at the front of the list.
// The list and other may be the same. They must not be nil.
func (l *List[T]) PushFrontList(other *List[T]) { _ = "STUB: not implemented"; return }

// String returns a string representation of the list
func (l *List[T]) String() string { _ = "STUB: not implemented"; return "" }

// Traversal traversals elements in the list, it will not stop until to the end of the list or the visitor returns false
func (l *List[T]) Traversal(visitor visitor.Visitor[T]) { _ = "STUB: not implemented"; return }
