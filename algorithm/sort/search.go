package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

// BinarySearch returns true if exist an element witch value is val in the range [first, last), or false if not exist
func BinarySearch[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) bool {
	_ = "STUB: not implemented"
	return false
}

func binarySearch[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) bool {
	_ = "STUB: not implemented"
	return false
}

// LowerBound returns the iterator pointing to the first element greater than or equal to value passed in the range [first, last), or iterator last if not exist.
func LowerBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

func lowerBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

// UpperBound returns the iterator pointing to the first element greater than val in the range [first, last), or iterator last if not exist.
func UpperBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

func upperBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	_ = "STUB: not implemented"
	return nil
}
