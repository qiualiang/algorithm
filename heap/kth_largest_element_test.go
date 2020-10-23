package heap

import "testing"

func TestKthLarget(t *testing.T) {
	nums := []int{4, 5, 8, 2}
	kth := Constructor(3, nums)
	if v := kth.Add(3); v != 4 {
		t.Errorf("The kth largest should be %d,but it returns %d", 4, v)
	}
	if v := kth.Add(5); v != 5 {
		t.Errorf("The kth largest should be %d,but it returns %d", 5, v)
	}
	if v := kth.Add(10); v != 5 {
		t.Errorf("The kth largest should be %d,but it returns %d", 5, v)
	}
	if v := kth.Add(9); v != 8 {
		t.Errorf("The kth largest should be %d,but it returns %d", 8, v)
	}
	if v := kth.Add(4); v != 8 {
		t.Errorf("The kth largest should be %d,but it returns %d", 8, v)
	}
}
