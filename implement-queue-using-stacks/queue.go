package main

import (
	"fmt"
)

type MyQueue struct {
	input  []int
	output []int
}

/** Initialize your data structure here. */
func Constructor() *MyQueue {
	input := []int{}
	output := []int{}
	return &MyQueue{input, output}

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.input) == 0 && len(this.output) == 0 {
		return -1
	}
	if len(this.output) == 0 {
		for i := len(this.input) - 1; i >= 0; i-- {
			this.output = append(this.output, this.input[i])
		}
		this.input = []int{}
	}
	result := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.input) == 0 && len(this.output) == 0 {
		return -1
	}
	if len(this.output) == 0 {
		for i := len(this.input) - 1; i >= 0; i-- {
			this.output = append(this.output, this.input[i])
		}
		this.input = []int{}
	}
	return this.output[len(this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.input) == 0 && len(this.output) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
	fmt.Println(queue.Pop())
	fmt.Println("empty?", queue.Empty())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
}
