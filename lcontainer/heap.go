package lcontainer

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[0 : n-1]
	return x
}

func Test() {
	h := IntHeap{1, 4, 6, 8, 9, 10, 33, 41, 256, 78}
	heap.Init(&h)
	heap.Push(&h, 99)
	heap.Pop(&h)
	fmt.Println(h)
}
