package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}

	t1 := TwoSum(nums, 9)
	if t1[0] != 0 && t1[1] != 1 {
		t.Errorf("t1[0]:%d,t1[1]:%d,want 0 and 1", t1[0], t1[1])
	}
	t2 := TwoSum(nums, 10)
	if len(t2) != 0 {
		t.Errorf("t1:%v,want []", t1)
	}
}
