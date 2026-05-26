package rbtree

import (
	"errors"

	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/visitor"
)

var ErrorNotFound = errors.New("not found")

// RbTree is a kind of self-balancing binary search tree in computer science.
// Each node of the binary tree has an extra bit, and that bit is often interpreted
// as the color (red or black) of the node. These color bits are used to ensure the tree
// remains approximately balanced during insertions and deletions.
type RbTree[K, V any] struct {
	root   *Node[K, V]
	size   int
	keyCmp comparator.Comparator[K]
}

// New creates a new RbTree
func New[K, V any](cmp comparator.Comparator[K]) *RbTree[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Clear clears the RbTree
func (t *RbTree[K, V]) Clear() { _ = "STUB: not implemented"; return }

// Find finds the first node that the key is equal to the passed key, and returns its value
func (t *RbTree[K, V]) Find(key K) (V, error) { _ = "STUB: not implemented"; return *new(V), nil }

// FindNode the first node that the key is equal to the passed key and return it
func (t *RbTree[K, V]) FindNode(key K) *Node[K, V] { _ = "STUB: not implemented"; return nil }

// Compare compares two keys wrt the RbTree's key comparator
func (t *RbTree[K, V]) Compare(key1, key2 K) int { _ = "STUB: not implemented"; return 0 }

// Begin returns the node with minimum key in the RbTree
func (t *RbTree[K, V]) Begin() *Node[K, V] {
	_ = "STUB: not implemented"

	// First returns the node with minimum key in the RbTree
	return nil
}

func (t *RbTree[K, V]) First() *Node[K, V] { _ = "STUB: not implemented"; return nil }

// RBegin returns the Node with maximum key in the RbTree
func (t *RbTree[K, V]) RBegin() *Node[K, V] {
	_ = "STUB: not implemented"

	// Last returns the Node with maximum key in the RbTree
	return nil
}

func (t *RbTree[K, V]) Last() *Node[K, V] { _ = "STUB: not implemented"; return nil }

// IterFirst returns the iterator of first node
func (t *RbTree[K, V]) IterFirst() *RbTreeIterator[K, V] { _ = "STUB: not implemented"; return nil }

// IterLast returns the iterator of first node
func (t *RbTree[K, V]) IterLast() *RbTreeIterator[K, V] { _ = "STUB: not implemented"; return nil }

// Empty returns true if Tree is empty,otherwise returns false.
func (t *RbTree[K, V]) Empty() bool {
	_ = "STUB: not implemented"

	// Size returns the size of the rbtree.
	return false
}

func (t *RbTree[K, V]) Size() int {
	_ = "STUB: not implemented"

	// Insert inserts a key-value pair into the RbTree.
	return 0
}

func (t *RbTree[K, V]) Insert(key K, value V) { _ = "STUB: not implemented"; return }

func (t *RbTree[K, V]) rbInsertFixup(z *Node[K, V]) { _ = "STUB: not implemented"; return }

// Delete deletes node from the RbTree
func (t *RbTree[K, V]) Delete(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (t *RbTree[K, V]) rbDeleteFixup(x, parent *Node[K, V]) { _ = "STUB: not implemented"; return }

func (t *RbTree[K, V]) rbFixupLeft(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *RbTree[K, V]) rbFixupRight(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *RbTree[K, V]) leftRotate(x *Node[K, V]) { _ = "STUB: not implemented"; return }

func (t *RbTree[K, V]) rightRotate(x *Node[K, V]) { _ = "STUB: not implemented"; return }

// findNode finds the node that its key is equal to the passed key, and returns it.
func (t *RbTree[K, V]) findNode(key K) *Node[K, V] { _ = "STUB: not implemented"; return nil }

// findNode finds the first node that its key is equal to the passed key, and returns it
func (t *RbTree[K, V]) findFirstNode(key K) *Node[K, V] { _ = "STUB: not implemented"; return nil }

// FindLowerBoundNode finds the first node that its key is equal or greater than the passed key, and returns it
func (t *RbTree[K, V]) FindLowerBoundNode(key K) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (t *RbTree[K, V]) findLowerBoundNode(x *Node[K, V], key K) *Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// FindUpperBoundNode finds the first node that its key is greater than the passed key, and returns it
func (t *RbTree[K, V]) FindUpperBoundNode(key K) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (t *RbTree[K, V]) findUpperBoundNode(x *Node[K, V], key K) *Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Traversal traversals elements in the RbTree, it will not stop until to the end of RbTree or the visitor returns false
func (t *RbTree[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	_ = "STUB: not implemented"
	return
}

// IsRbTree is a function use to test whether t is a RbTree or not
func (t *RbTree[K, V]) IsRbTree() (bool, error) {
	_ = "STUB: not implemented"
	// Properties:
	// 1. Each node is either red or black.
	// 2. The root is black.
	// 3. All leaves (NIL) are black.
	// 4. If a node is red, then both its children are black.
	// 5. Every path from a given node to any of its descendant NIL nodes contains the same number of black nodes.
	return false, nil
}

func (t *RbTree[K, V]) test(n *Node[K, V]) (int, int, bool) {
	_ = "STUB: not implemented"

	// property 3:
	return 0, 0, false
}

// property 2:

// property 5:

// property 4:

// if n == t.root {
// 	fmt.Printf("blackCount:%v \n", blackCount)
// }

// getColor returns the node's color
func getColor[K, V any](n *Node[K, V]) Color { _ = "STUB: not implemented"; return *new(Color) }
