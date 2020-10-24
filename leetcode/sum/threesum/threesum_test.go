package threesum

import (
	"algorithm/util"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	slice1, slice2 := []int{-1, -1, 2}, []int{-1, 0, 1}
	target := [][]int{slice1, slice2}

	res := ThreeSum(nums)
	if len(res) != len(target) {
		t.Fatalf("ThreeSum() return %v,but it should return %v", res, target)
	}
	if !(util.CheckSliceEq(res[0], slice1) && util.CheckSliceEq(res[1], slice2)) {
		t.Errorf("ThreeSum() return %v,but it should return %v", res, target)
	}

}
