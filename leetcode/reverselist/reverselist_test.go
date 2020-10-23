package reverselist

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	rev := ReverseList(head)
	checkList(t, rev, []interface{}{7, 6, 5, 4, 3, 2, 1})

}

func TestReverseBetween(t *testing.T) {
	head := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	rev := ReverseBetween(head, 2, 4)
	checkList(t, rev, []interface{}{1, 4, 3, 2, 5, 6, 7})
}

func TestReverseKGroup(t *testing.T) {
	head := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	rev := ReverseKGroup(head, 3)
	checkList(t, rev, []interface{}{3, 2, 1, 6, 5, 4, 7})

}
func checkList(t *testing.T, head *ListNode, es []interface{}) {
	i := 0
	n := len(es)
	for head != nil {
		if n-i < 0 {
			t.Errorf("The list len > %d, want %d", i, n)
			return
		}
		if head.Val != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, head.Val, es[i])
			return
		}
		i++
		head = head.Next
	}
	if n != i {
		t.Errorf("The list len is %d, want %d", i, n)
		return
	}
}
