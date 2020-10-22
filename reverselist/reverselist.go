package reverselist

import (
	"fmt"
)

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
func main() {
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
	tmp := head
	fmt.Println("Befor reverse:")
	for tmp != nil {

		fmt.Printf("%d ", tmp.Val)

		tmp = tmp.Next
	}
	// fmt.Println("\nAfter reverse:")
	// rev := ReverseList(head)

	// for rev != nil {
	// 	fmt.Printf("%d ", rev.Val)
	// 	rev = rev.Next
	// }
	// fmt.Println("\nReverse betwwen [2,4]")
	// rev2 := reverseBetween(head, 2, 4)
	// for rev2 != nil {
	// 	fmt.Printf("%d ", rev2.Val)
	// 	rev2 = rev2.Next
	// }
	fmt.Println("\nAfter reverse:")
	rev := ReverseKGroup(head, 2)

	for rev != nil {
		fmt.Printf("%d ", rev.Val)
		rev = rev.Next
	}

}
