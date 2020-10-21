package main

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
func main() {
	head := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	fmt.Println("Befor swap:")
	tmp := head
	for tmp != nil {

		fmt.Printf("%d ", tmp.Val)

		tmp = tmp.Next
	}
	swap := SwapPairs(head)
	fmt.Println("After swap:")

	for swap != nil {
		fmt.Printf("%d ", swap.Val)
		swap = swap.Next
	}
}
