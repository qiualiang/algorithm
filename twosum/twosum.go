package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))

	for i, v := range nums {
		if index, ok := numMap[target-v]; ok {
			return []int{index, i}
		}
		numMap[v] = i
	}
	return nil
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	if result := twoSum(nums, target); result != nil {
		fmt.Println(result)
	} else {
		fmt.Println("Could not find the solution.")
	}
}
