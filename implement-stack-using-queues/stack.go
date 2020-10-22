package stack

import (
	"container/list"
	"fmt"
)

type MyStack struct {
	l1, l2 list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	var l1, l2 list.List
	return MyStack{l1, l2}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.l1.Len() == 0 {
		this.l1.PushBack(x)
		for e := this.l2.Front(); e != nil; e = e.Next() {
			this.l1.PushBack(e.Value)
		}
		this.l2.Init()
	} else {
		this.l2.PushBack(x)
		for e := this.l1.Front(); e != nil; e = e.Next() {
			this.l2.PushBack(e.Value)
		}
		this.l1.Init()
	}

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.l1.Len() > 0 {
		front := this.l1.Front()
		this.l1.Remove(front)
		return front.Value.(int)
	} else if this.l2.Len() > 0 {
		front := this.l2.Front()
		this.l2.Remove(front)
		return front.Value.(int)
	}
	return -1
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.l1.Len() > 0 {
		return this.l1.Front().Value.(int)
	} else if this.l2.Len() > 0 {
		return this.l2.Front().Value.(int)
	}
	return -1
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.l1.Len() == 0 && this.l2.Len() == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
