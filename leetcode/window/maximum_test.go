package window

import "testing"

func TestMaximum(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	target := []int{3, 3, 5, 5, 6, 7}
	max := MaxSlidingWindow(nums, 3)

	if !testEq(max, target) {
		t.Errorf("The result should be %v,but it returns %v", target, max)
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func BenchmarkMaximum(b *testing.B) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	for i := 0; i < b.N; i++ {
		MaxSlidingWindow(nums, 3)
	}
}
