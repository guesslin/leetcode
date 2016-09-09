package main

import (
	"fmt"
)

func rob(nums []int) int {
	houses := len(nums)
	if houses == 0 {
		return 0
	}
	if houses == 1 {
		return nums[0]
	}
	count := make([]int, houses)
	count[0] = nums[0]
	count[1] = max(nums[0], nums[1])
	for i := 2; i < houses; i++ {
		count[i] = max(count[i-1], nums[i]+count[i-2])
	}
	return count[houses-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(rob([]int{0}))
	fmt.Println(rob([]int{2, 1}))
	fmt.Println(rob([]int{2, 6, 1}))
}
