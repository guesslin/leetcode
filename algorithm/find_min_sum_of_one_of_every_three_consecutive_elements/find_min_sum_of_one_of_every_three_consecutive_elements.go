package main

import (
	"fmt"
)

func findMinSum(nums []int) int {
	length := len(nums)
	sums := make([]int, length)
	sums[0] = nums[0]
	sums[1] = nums[1]
	sums[2] = nums[2]
	for i := 3; i < length; i++ {
		sums[i] = nums[i] + minNum(sums[i-3], sums[i-2], sums[i-1])
	}
	return minNum(sums[length-3], sums[length-2], sums[length-1])
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minNum(x, y, z int) int {
	return min(min(x, y), z)
}

func main() {
	nums := []int{1, 2, 3, 6, 7, 1}
	fmt.Println(nums)
	fmt.Println(findMinSum(nums))
}
