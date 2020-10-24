package foursum

import (
	"algorithm/util"
	"testing"
)

func TestFourSum(t *testing.T) {
	nums := []int{-2, -1, -1, 1, 1, 2, 2}

	// nums := []int{0, 0, 0, 0}
	target := 0
	slice1, slice2 := []int{-2, -1, 1, 2}, []int{-1, -1, 1, 1}
	answer := [][]int{slice1, slice2}

	res := FourSum(nums, target)
	if len(res) != len(answer) {
		t.Fatalf("FourSum() return %v,but it should return %v", res, answer)
	}

	if !(util.CheckSliceEq(res[0], slice1) && util.CheckSliceEq(res[1], slice2)) {
		t.Errorf("FourSum() return %v,but it should return %v", res, answer)
	}
}
