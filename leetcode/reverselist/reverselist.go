package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

//ReverseList the key is preserve the head of the new list
func ReverseList(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		head.Next, head, cur = cur, head.Next, head
	}
	return cur
}

/**
 * Definition for singly-linked list.
 * 1. 头插法反转链表, 每次将遍历的当前节点插入开始反转的位置。
 * 2. 通过哨兵节点处理 m == 1 的情况
 * 3. 通过m的值确定头插法的头节点位置
 * 4. 通过n-m的值确定执行几次头插操作
 */
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	dummyhead := &ListNode{-1, head}
	// dummyhead.Next = head
	dummy := dummyhead
	i, j := m, n-m
	for i > 1 {
		dummy = dummy.Next
		i--
	}

	cur := dummy.Next.Next
	pre := dummy.Next
	for j > 0 {
		cur.Next, pre, cur = pre, cur, cur.Next
		j--
	}
	mm := dummy.Next
	dummy.Next = pre
	mm.Next = cur
	return dummyhead.Next

}
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}
	dummyhead := &ListNode{-1, head}
	pre := dummyhead
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyhead.Next
			}

		}
		next := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}

	return dummyhead.Next
}
func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	p := head
	nex := tail.Next
	tail.Next = nil
	var cur *ListNode
	for head != nil {
		head.Next, head, cur = cur, head.Next, head
	}
	p.Next = nex
	return cur, p
}
