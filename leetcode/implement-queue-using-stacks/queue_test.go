package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	if v := queue.Peek(); v != 1 {
		t.Errorf("The peek value should be 1, but it returns %d", v)
	}
	if v := queue.Pop(); v != 1 {
		t.Errorf("The first pop value should be 1, but it returns %d", v)
	}
	if v := queue.Pop(); v != 2 {
		t.Errorf("The second pop value should be 2, but it returns %d", v)
	}
	if !queue.Empty() {
		t.Error("The queue should be empty!")
	}
}
