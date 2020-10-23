package heap

import "container/heap"

type IntMaxHeap []int
type KthMin struct {
	k int
	H *IntMaxHeap
}

func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *IntMaxHeap) Len() int {
	return len(*h)
}
func (h *IntMaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func (h *IntMaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func ConstructorKmin(k int, nums []int) KthMin {
	max := &IntMaxHeap{}
	heap.Init(max)
	kth := &KthMin{H: max, k: k}
	for _, num := range nums {
		kth.Add(num)
	}

	return *kth
}

func (kth *KthMin) Add(val int) int {
	if kth.H.Len() < kth.k {
		heap.Push(kth.H, val)
	} else if (*kth.H)[0] > val {
		heap.Pop(kth.H)
		heap.Push(kth.H, val)
	}
	return (*kth.H)[0]
}
