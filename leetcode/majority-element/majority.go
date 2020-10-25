package majority

import "sort"

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

链接：https://leetcode-cn.com/problems/majority-element

1.暴力 两重循环 O(n*n)
2. map count O(N)，空间复杂度 O(N)
3.因为肯定存在，所以sort后中间就是我们要的值 O(NlogN)
4.分治，左边求majority,右边求majority, 若left==right，找到;若不同，return max count of left or right.  O(NLogN)
*/
func majorityElement(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n/2]
}
