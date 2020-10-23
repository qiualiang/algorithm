package heap

import "container/heap"

type IntHeap []int
type KthLargest struct {
	k int
	H *IntHeap
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *IntHeap) Len() int {
	return len(*h)
}
func (h *IntHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func Constructor(k int, nums []int) KthLargest {
	min := &IntHeap{}
	heap.Init(min)
	kth := &KthLargest{H: min, k: k}
	for _, num := range nums {
		kth.Add(num)
	}

	return *kth
}

func (kth *KthLargest) Add(val int) int {
	if kth.H.Len() < kth.k {
		heap.Push(kth.H, val)
	} else if (*kth.H)[0] < val {
		heap.Pop(kth.H)
		heap.Push(kth.H, val)
	}
	return (*kth.H)[0]
}
