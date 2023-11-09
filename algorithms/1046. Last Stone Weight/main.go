package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func main() {
	stones := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if y > x {
			heap.Push(&h, y-x)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(&h).(int)
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

// Swap swaps the elements with indexes i and j.
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
