package swappairs

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	self := new(ListNode)
	pre := self
	pre.Next = head
	self.Next = pre.Next
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
	}
	return self.Next
}

