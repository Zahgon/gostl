package main

import (
	"fmt"
	"math/rand"

	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	q := deque.New[int]()
	for i := 0; i < 100; i++ {
		r := rand.Int() % 100
		q.PushBack(r)
		q.PushFront(r)
	}
	fmt.Printf("%v\n", q)

	sort.Sort[int](q.Begin(), q.End(), comparator.IntComparator)
	fmt.Printf("%v\n", q)

	for !q.Empty() {
		r := rand.Int() % q.Size()
		q.EraseAt(r)
	}
	fmt.Printf("%v\n", q)
}
