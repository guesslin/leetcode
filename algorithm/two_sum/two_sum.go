package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	records := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		search := target - nums[i]
		if j, ok := records[search]; ok {
			return []int{j, i}
		}
		records[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{0, 4, 3, 0}, 0))
}
