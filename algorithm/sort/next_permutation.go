package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

// NextPermutation transform range [first last) to next permutation,return true if success, or false if failure
func NextPermutation[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) bool {
	_ = "STUB: not implemented"
	return false
}

func reverse[T any](s, e iterator.RandomAccessIterator[T]) { _ = "STUB: not implemented"; return }
