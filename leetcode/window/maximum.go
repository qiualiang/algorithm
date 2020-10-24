package window

import (
	"fmt"
)

func MaxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	var max, index int
	window := []int{}
	j := 0
	for i := 0; i < len(nums); i++ {
		if i < k {
			window = nums[:k]
			i = k - 1
		}
		// fmt.Println("i:", i)
		if i == (k - 1) { //First window
			max, index = getMax(window, -1, -1)
			// fmt.Printf("max:%d,index:%d\n", max, index)
			result = append(result, max)
		}
		if i >= k {
			j++
			window = nums[j : i+1]
			max, index = getMax(window, index, nums[i])
			result = append(result, max)
		}

	}

	return result
}
func getMax(nums []int, lastMaxIndex int, elem int) (max int, index int) {
	// fmt.Println(nums, lastMaxIndex, elem)
	if lastMaxIndex == -1 {
		index = 0
		max = nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > max {
				max, index = nums[i], i
			}
		}
		return max, index
	} else {
		if lastMaxIndex-1 > 0 { //the last max elem is in current window
			lastMax := nums[lastMaxIndex-1]
			if lastMax > elem {
				max, index = lastMax, lastMaxIndex-1
				return

			} else {
				max = elem
				index = len(nums) - 1
				return
			}
		} else { //the last max elem is not in current window
			index = 0
			max = nums[0]
			for i := 1; i < len(nums); i++ {
				if nums[i] > max {
					max, index = nums[i], i
				}
			}
			return max, index
		}
	}
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(MaxSlidingWindow(nums, 3))
}
