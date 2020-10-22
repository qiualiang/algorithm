package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	if v := stack.Top(); v != 2 {
		t.Errorf("The top value should be 2, but it returns %d", v)
	}
	if v := stack.Pop(); v != 2 {
		t.Errorf("The first pop value should be 2, but it returns %d", v)
	}
	if v := stack.Pop(); v != 1 {
		t.Errorf("The second pop value should be 1, but it returns %d", v)
	}
	if !stack.Empty() {
		t.Error("The stack should be empty!")
	}
}
