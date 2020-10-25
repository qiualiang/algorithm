package majority

import "testing"

func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	if n := majorityElement(nums); n != 3 {
		t.Error("get error element:", n)
	}
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	if n := majorityElement(nums); n != 2 {
		t.Error("get error element:", n)
	}
}
