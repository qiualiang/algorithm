package linkedlistcycle

import "testing"

func TestGoThrough(t *testing.T) {
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
	if !GoThrough(head) {
		t.Error("The GoThrough functiong should return true.")
	}
}

func TestFastSlowPoint(t *testing.T) {
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
	if !FastSlowPoint(head) {
		t.Error("The FastSlowPoint functiong should return true.")
	}
}
