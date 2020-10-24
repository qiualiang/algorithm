package twosum

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))

	for i, v := range nums {
		if index, ok := numMap[target-v]; ok {
			return []int{index, i}
		}
		numMap[v] = i
	}
	return nil
}

