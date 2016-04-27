package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	removeCount := 0
	c := nums[0]
	for n := 1; n < length; n++ {
		if c == nums[n] {
			removeCount++
		} else {
			nums[n-removeCount] = nums[n]
			c = nums[n]
		}
	}
	return length - removeCount
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 5}
	nl := removeDuplicates(nums)
	fmt.Println(nums[:nl])
	fmt.Println(nums)
}
