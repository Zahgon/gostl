package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

// Sort sorts the container by using quick sort
func Sort[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	_ = "STUB: not implemented"
	return
}

func quickSort[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	_ = "STUB: not implemented"
	return
}

func doPivot[T any](first, mid, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	_ = "STUB: not implemented"
	return
}

func swapValue[T any](a, b iterator.RandomAccessIterator[T]) { _ = "STUB: not implemented"; return }
