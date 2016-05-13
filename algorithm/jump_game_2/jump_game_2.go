package main

import (
	"fmt"
)

func jump(nums []int) int {
	minStep := make([]int, len(nums))
	for i := range minStep {
		minStep[i] = i
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			minStep[j] = min(minStep[j], minStep[i]+1)
		}
	}
	return minStep[len(minStep)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	testCases := [][]int{
		{2, 3, 1, 1, 4},
		{3, 4, 6, 1, 1, 1, 1, 2},
		{3, 2, 1},
		{1, 2, 3},
		{1},
		{1, 2, 1, 1, 1},
	}
	for i, c := range testCases {
		fmt.Printf("%d-th :%v, result %d\n", i+1, c, jump(c))
	}
}
