package heap

import "testing"

func TestKthMinimum(t *testing.T) {
	nums := []int{4, 5, 8, 2}
	kth := ConstructorMax(3, nums)
	if v := kth.Add(3); v != 4 {
		t.Errorf("The kth minimum should be %d,but it returns %d", 4, v)
	}
	if v := kth.Add(5); v != 4 {
		t.Errorf("The kth minimum should be %d,but it returns %d", 4, v)
	}
	if v := kth.Add(10); v != 4 {
		t.Errorf("The kth minimum should be %d,but it returns %d", 4, v)
	}
	if v := kth.Add(1); v != 3 {
		t.Errorf("The kth largest should be %d,but it returns %d", 3, v)
	}
	if v := kth.Add(4); v != 3 {
		t.Errorf("The kth largest should be %d,but it returns %d", 3, v)
	}
}
