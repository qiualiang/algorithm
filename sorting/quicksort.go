package sorting

import (
	"fmt"
)

/*
快速排序算法通过多次比较和交换来实现排序，其排序流程如下：（分治法）
(1)首先设定一个分界值，通过该分界值将数组分成左右两部分。
(2)将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。
此时，左边部分中各元素都小于或等于分界值，而右边部分中各元素都大于或等于分界值。
(3)然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，
将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理
(4)重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。
当左、右两个部分各数据排序完成后，整个数组的排序也就完成了。
*/
func QuickSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	pivot := input[0]
	less := []int{}
	greater := []int{}
	for i := 1; i < len(input); i++ {
		if input[i] <= pivot {
			less = append(less, input[i])
		}
		if input[i] > pivot {
			greater = append(greater, input[i])
		}
	}
	left := QuickSort(less)
	right := QuickSort(greater)
	left = append(left, pivot)
	return append(left, right...)
}
func main() {
	input := []int{9, 5, 1, 2, 6, 4, 3, 5, 7, 9}
	fmt.Println(QuickSort(input))
}
