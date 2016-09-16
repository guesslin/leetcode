package main

import (
	"fmt"
)

func rob(nums []int) int {
	houses := len(nums)
	// there is no house to rob
	if houses == 0 {
		return 0
	}
	// there is only one house to rob
	if houses == 1 {
		return nums[0]
	}
	first := make([]int, houses)
	notFirst := make([]int, houses)
	first[0] = nums[0]
	first[1] = max(nums[0], nums[1])
	notFirst[0] = 0
	notFirst[1] = nums[1]
	for i := 2; i < houses; i++ {
		first[i] = max(first[i-1], nums[i]+first[i-2])
		notFirst[i] = max(notFirst[i-1], nums[i]+notFirst[i-2])
	}
	return max(first[houses-2], notFirst[houses-1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(rob([]int{2, 2, 4, 3, 2, 5}))
	fmt.Println(rob([]int{2, 1}))
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(rob([]int{0}))
	fmt.Println(rob([]int{2, 1}))
	fmt.Println(rob([]int{2, 6, 1}))
}
