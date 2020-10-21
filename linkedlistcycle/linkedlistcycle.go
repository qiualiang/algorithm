package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GoThrough(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	hash := make(map[*ListNode]int)
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = head.Val
		head = head.Next
	}
	return false
}

func FastSlowPoint(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil && slow != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
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
	n5.Next = n3
	fmt.Println(GoThrough(head))
	fmt.Println(FastSlowPoint(head))

}
