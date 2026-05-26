package hamt

import (
	"errors"

	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

// Some constants
const (
	BITMAP_NODE = 0
	KV_NODE     = 1
	Fanout      = 6 //each bitmap node has 6 bits, so the max depth of tree is 64/6 = 10.666 = 11
	Mask        = (1 << Fanout) - 1
)

var ErrorNotFound = errors.New("not found")

// Key is a redefinition of []byte
type Key []byte

var (
	defaultLocker sync.FakeLocker
)

// Options holds Hamt's options
type Options struct {
	locker sync.Locker
}

// Option is a function type used to set Options
type Option func(option *Options)

// WithGoroutineSafe is used to config a Hamt with goroutine-safe
func WithGoroutineSafe() Option { _ = "STUB: not implemented"; return *new(Option) }

// Entry is a tree node interface
type Entry interface {
	// Type returns the node type
	Type() int

	// BitPosNum returns number from a bit position
	BitPosNum(depth int) uint64
}

// BitmapNode defines Hamt's bitmap node
type BitmapNode[T any] struct {
	bitmap   uint64
	children []Entry
	pos      uint8 //position in parent array, in range [0, 64)
}

// KvPair is a list node with actually value
type KvPair[T any] struct {
	key   Key
	value T
	next  *KvPair[T]
}

// KvNode is Hamt's key-value node
type KvNode[T any] struct {
	hash   uint64
	kvList *KvPair[T]
}

// Hamt is an implementation of hash-array-mapped-trie
type Hamt[T any] struct {
	root   BitmapNode[T]
	locker sync.Locker
}

// Type returns the node type
func (h *BitmapNode[T]) Type() int {
	_ = "STUB: not implemented"

	// BitPosNum returns the number from a bit position
	return 0
}

func (h *BitmapNode[T]) BitPosNum(int) uint64 { _ = "STUB: not implemented"; return 0 }

// Index returns the index of a bitPos int bitmap
func (h *BitmapNode[T]) Index(bitPos uint64) int { _ = "STUB: not implemented"; return 0 }

func (h *BitmapNode[T]) insert(depth int, hash uint64, kv *KvPair[T]) {
	_ = "STUB: not implemented"
	return
	//hash in current node's position
}

//hash in current bitmap's position in bit

func (h *BitmapNode[T]) find(depth int, hash uint64, key Key) (T, error) {
	_ = "STUB: not implemented"
	return *
	//hash in current node's position
	new(T), nil
}

//hash in current bitmap's position in bit

func (h *BitmapNode[T]) traversal(visitor visitor.KvVisitor[Key, T]) {
	_ = "STUB: not implemented"
	return
}

func (h *BitmapNode[T]) erase(depth int, hash uint64, key Key) bool {
	_ = "STUB: not implemented"
	return false
	//hash in current node's position
}

//hash in current bitmap's position in bit

// remove iter

// change bitmapNode to kvNode, if a bitmapNode has only one kvNode

// Type returns the node type
func (h *KvNode[T]) Type() int {
	_ = "STUB: not implemented"

	// BitPosNum returns the bit position
	return 0
}

func (h *KvNode[T]) BitPosNum(depth int) uint64 { _ = "STUB: not implemented"; return 0 }

// New creates a Hamt(hash array mapped trie) instance
func New[T any](opts ...Option) *Hamt[T] { _ = "STUB: not implemented"; return nil }

// Insert inserts a key-value pair into the hamt
func (h *Hamt[T]) Insert(key Key, value T) { _ = "STUB: not implemented"; return }

// Get returns the value by the passed key if the key is in the hamt, otherwise returns nil
func (h *Hamt[T]) Get(key Key) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// Erase erases the key-value pair in hamt, and returns true if succeed.
func (h *Hamt[T]) Erase(key Key) bool { _ = "STUB: not implemented"; return false }

// Keys returns keys in Hamt
func (h *Hamt[T]) Keys() []Key { _ = "STUB: not implemented"; return nil }

// StringKeys returns keys in Hamt
func (h *Hamt[T]) StringKeys() []string { _ = "STUB: not implemented"; return nil }

// Traversal traversals elements in Hamt, it will not stop until to the end or the visitor returns false
func (h *Hamt[T]) Traversal(visitor visitor.KvVisitor[Key, T]) { _ = "STUB: not implemented"; return }

func hash(a []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func pos(hash uint64, depth int) uint8 { _ = "STUB: not implemented"; return 0 }

func bitPos(pos uint8) uint64 { _ = "STUB: not implemented"; return 0 }
