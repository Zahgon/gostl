package rbtree

// Color defines node color type
type Color bool

// Define node 's colors
const (
	RED   = false
	BLACK = true
)

// Node is a tree node
type Node[K, V any] struct {
	parent *Node[K, V]
	left   *Node[K, V]
	right  *Node[K, V]
	color  Color
	key    K
	value  V
}

// Key returns node's key
func (n *Node[K, V]) Key() K {
	_ = "STUB: not implemented"

	// Value returns node's value
	return *new(K)
}

func (n *Node[K, V]) Value() V {
	_ = "STUB: not implemented"

	// SetValue sets node's value
	return *new(V)
}

func (n *Node[K, V]) SetValue(val V) {
	_ = "STUB: not implemented"

	// Next returns the Node's successor as an iterator.
	return
}

func (n *Node[K, V]) Next() *Node[K, V] { _ = "STUB: not implemented"; return nil }

// Prev returns the Node's predecessor as an iterator.
func (n *Node[K, V]) Prev() *Node[K, V] { _ = "STUB: not implemented"; return nil }

// successor returns the successor of the Node
func successor[K, V any](x *Node[K, V]) *Node[K, V] { _ = "STUB: not implemented"; return nil }

// presuccessor returns the presuccessor of the Node
func presuccessor[K, V any](x *Node[K, V]) *Node[K, V] { _ = "STUB: not implemented"; return nil }

// minimum finds the minimum Node of subtree n.
func minimum[K any, V any](n *Node[K, V]) *Node[K, V] { _ = "STUB: not implemented"; return nil }

// maximum finds the maximum Node of subtree n.
func maximum[K any, V any](n *Node[K, V]) *Node[K, V] { _ = "STUB: not implemented"; return nil }
