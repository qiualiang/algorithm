package swappairs

import "testing"

func TestSwapPairs(t *testing.T) {
	head := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	swap := SwapPairs(head)
	checkList(t, swap, []interface{}{2, 1, 4, 3, 5})

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
