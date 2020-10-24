package foursum

import "sort"

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，
判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
找出所有满足条件且不重复的四元组。
注意：答案中不可以包含重复的四元组。
示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
链接：https://leetcode-cn.com/problems/4sum
特别要注意去重以及 i和 j 的边界判断，否则会出错！
*/
func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	// if len(nums) == 4 {
	// 	if nums[0]+nums[1]+nums[2]+nums[3] != target {
	// 		return nil
	// 	} else {
	// 		return [][]int{{nums[0], nums[1], nums[2], nums[3]}}
	// 	}
	// }
	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if nums[i]+nums[i+1]+nums[i+2]+nums[3] > target {
			continue
		}
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				continue
			}
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				s := nums[i] + nums[j] + nums[left] + nums[right]
				if s < target {
					left++
				} else if s > target {
					right--
				} else { //s==0
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] { //去重
						right--
					}
					left++
					right--
				}

			}
			for j < n-2 && nums[j] == nums[j+1] {
				j++
			}
		}
		for i < n-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
