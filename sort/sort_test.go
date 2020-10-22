package sort

import "testing"

var raw = []int{9, 11, 5, 6, 3, 1, 7, 2, 4, 20, 13}
var correct = []int{1, 2, 3, 4, 5, 6, 7, 9, 11, 13, 20}

func checkLen(t *testing.T, raw []int, sorted []int) bool {
	if len(raw) != len(sorted) {
		t.Error("The length of raw array and sorted array are not equal.")
		return false
	}
	return true
}
func checkSort(t *testing.T, sorted []int, correct []int) {
	for i := 0; i < len(sorted); i++ {
		if sorted[i] != correct[i] {
			t.Errorf("sorted[%d] = %v, want %v", i, sorted[i], correct[i])
			return
		}
	}

}
func TestBubbleSort(t *testing.T) {
	sorted := BubbleSort(raw)
	if checkLen(t, raw, sorted) {
		checkSort(t, sorted, correct)
	}

}
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(raw)
	}

}
func TestQuickSort(t *testing.T) {
	sorted := QuickSort(raw)
	if checkLen(t, raw, sorted) {
		checkSort(t, sorted, correct)
	}

}
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(raw)
	}
}
func TestInsertSort(t *testing.T) {
	sorted := InsertSort(raw)
	if checkLen(t, raw, sorted) {
		checkSort(t, sorted, correct)
	}
}
func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(raw)
	}
}

func TestSelectionSort(t *testing.T) {
	sorted := SelectionSort(raw)
	if checkLen(t, raw, sorted) {
		checkSort(t, sorted, correct)
	}
}
func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(raw)
	}
}

func TestMergeSort(t *testing.T) {
	sorted := MergeSort(raw)
	if checkLen(t, raw, sorted) {
		checkSort(t, sorted, correct)
	}
}
func BenchmarkSMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(raw)
	}
}
