package heap

import (
	"container/heap"
)

type Item struct {
	value    int
	priority int
	index    int
}
type PriorityQueue []*Item
type KthMinimum struct {
	k int
	H *PriorityQueue
}

func (h *PriorityQueue) Less(i, j int) bool {
	return (*h)[i].priority > (*h)[j].priority
}

func (h *PriorityQueue) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *PriorityQueue) Len() int {
	return len(*h)
}
func (h *PriorityQueue) Pop() (v interface{}) {
	var item *Item
	*h, item = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return item.value
}
func (h *PriorityQueue) Push(v interface{}) {
	n := len(*h)
	item := Item{value: v.(int), priority: v.(int), index: n}
	*h = append(*h, &item)
}

func ConstructorMax(k int, nums []int) KthMinimum {
	max := &PriorityQueue{}
	heap.Init(max)
	kth := &KthMinimum{H: max, k: k}
	for _, num := range nums {
		kth.Add(num)
	}
	return *kth
}

func (kth *KthMinimum) Add(val int) int {
	if kth.H.Len() < kth.k {
		heap.Push(kth.H, val)
	} else if (*kth.H)[0].value > val {
		heap.Pop(kth.H)
		heap.Push(kth.H, val)
	}
	return (*kth.H)[0].value
}
