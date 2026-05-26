package simplelist

import (
	"github.com/liyue201/gostl/utils/visitor"
)

// Node is a list node
type Node[T any] struct {
	next  *Node[T]
	Value T
}

// Next returns the next list node or nil.
func (n *Node[T]) Next() *Node[T] {
	_ = "STUB: not implemented"

	// List represents a single direction list:
	//
	//	head -> node1 --> node2 --> node3 <- tail
	return nil
}

type List[T any] struct {
	head *Node[T] // point to the front Node
	tail *Node[T] // point to the back Node
	len  int      // current list length
}

// New creates a list
func New[T any]() *List[T] { _ = "STUB: not implemented"; return nil }

// Len returns the amount of list nodes.
func (l *List[T]) Len() int {
	_ = "STUB: not implemented"

	// FrontNode returns the front node of the list or nil if the list is empty
	return 0
}

func (l *List[T]) FrontNode() *Node[T] {
	_ = "STUB: not implemented"

	// BackNode returns the last node of the list or nil if the list is empty
	return nil
}

func (l *List[T]) BackNode() *Node[T] {
	_ = "STUB: not implemented"

	// PushFront inserts a new node n with value v at the front of the list.
	return nil
}

func (l *List[T]) PushFront(v T) { _ = "STUB: not implemented"; return }

// PushBack inserts a new node n with value v at the back of the list.
func (l *List[T]) PushBack(v T) { _ = "STUB: not implemented"; return }

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of the list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Node[T]) *Node[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] { _ = "STUB: not implemented"; return nil }

// Remove removes node n from the list.
// The node must not be nil.
func (l *List[T]) Remove(pre, n *Node[T]) T { _ = "STUB: not implemented"; return *new(T) }

// MoveToFront moves node n to the front of the list.
// The n must not be nil.
func (l *List[T]) MoveToFront(pre, n *Node[T]) { _ = "STUB: not implemented"; return }

// MoveToBack moves node n to the back of the list.
// The n must not be nil.
func (l *List[T]) MoveToBack(pre, n *Node[T]) { _ = "STUB: not implemented"; return }

// String returns a string representation of the list
func (l *List[T]) String() string { _ = "STUB: not implemented"; return "" }

// Traversal traversals elements in the list, it will not stop until to the end of the list or the visitor returns false
func (l *List[T]) Traversal(visitor visitor.Visitor[T]) { _ = "STUB: not implemented"; return }
