package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

// Stable sorts the container by using merge sort
func Stable[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	_ = "STUB: not implemented"
	return
}

func mergeSort[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T], tempSlice []T) {
	_ = "STUB: not implemented"
	return
}

func merge[T any](first, mid, end iterator.RandomAccessIterator[T], cmp comparator.Comparator[T], tempSlice []T) {
	_ = "STUB: not implemented"
	return
}
