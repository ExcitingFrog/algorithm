package utils

import (
	"container/heap"
	"fmt"
)

type H []int

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *H) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func curHeap() {
	h := &H{}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 2)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

}
