package linkedlistcycle

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
