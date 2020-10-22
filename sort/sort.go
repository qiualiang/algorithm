package sort

/*
BubbleSort 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/
func BubbleSort(raw []int) []int {
	for i := 0; i < len(raw); i++ {
		for j := i + 1; j < len(raw); j++ {
			if raw[i] > raw[j] {
				raw[i], raw[j] = raw[j], raw[i]
			}
		}
	}
	return raw
}

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
func QuickSort(raw []int) []int {
	if len(raw) < 2 {
		return raw
	}
	pivot := raw[0]
	less := []int{}
	greater := []int{}
	for i := 1; i < len(raw); i++ {
		if raw[i] <= pivot {
			less = append(less, raw[i])
		}
		if raw[i] > pivot {
			greater = append(greater, raw[i])
		}
	}
	left := QuickSort(less)
	right := QuickSort(greater)
	left = append(left, pivot)
	return append(left, right...)
}

/*
插入排序，一般也被称为直接插入排序。对于少量元素的排序，它是一个有效的算法 。
插入排序是一种最简单的排序方法，它的基本思想是将一个记录插入到已经排好序的有序表中，从而一个新的、记录数增1的有序表。
在其实现过程使用双层循环，外层循环对除了第一个元素之外的所有元素，内层循环对当前元素前面有序表进行待插入位置查找，并进行移动 。
*/
func InsertSort(raw []int) []int {
	for i := 1; i < len(raw); i++ {
		for j := 0; j < i; j++ {
			if raw[i] < raw[j] {
				raw[j], raw[i] = raw[i], raw[j]
			}
		}
	}
	return raw
}

/*
选择排序（Selection sort）是一种简单直观的排序算法。
它的工作原理是：第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，
然后再从剩余的未排序元素中寻找到最小（大）元素，然后放到已排序的序列的末尾。
以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法。
*/
func SelectionSort(raw []int) []int {
	for i := 0; i < len(raw)-1; i++ {
		pos := i
		min := raw[i]
		for j := i + 1; j < len(raw); j++ {
			if raw[j] < min {
				min = raw[j]
				pos = j
			}
		}
		raw[i], raw[pos] = raw[pos], raw[i]
	}
	return raw
}

/*
归并排序（Merge Sort）是建立在归并操作上的一种有效，稳定的排序算法。
该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。
若将两个有序表合并成一个有序表，称为二路归并。
*/
func MergeSort(raw []int) []int {
	length := len(raw)
	if length <= 1 {
		return raw
	}
	num := length / 2
	left := MergeSort(raw[:num])
	right := MergeSort(raw[num:])

	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
