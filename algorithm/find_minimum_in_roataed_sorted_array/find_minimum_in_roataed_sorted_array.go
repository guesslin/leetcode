package main

import (
	"fmt"
)

// 我想用 biSearch 找到 index: n, m where nums[n] > nums[m]
func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] <= nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return nums[hi]
}

func main() {
	testCases := []struct {
		input  []int
		expect int
	}{
		{input: []int{4, 5, 6, 7, 1, 2, 3}, expect: 1},
		{input: []int{4, 5, 6, 7, 8, 1, 2, 3}, expect: 1},
		{input: []int{1, 2, 3, 4, 5, 6, 7}, expect: 1},
		{input: []int{2, 3, 4, 5, 6, 7, 1}, expect: 1},
	}
	for i, c := range testCases {
		res := findMin(c.input)
		fmt.Println(i, "input", c.input, "Expect", c.expect, "get", res)
	}

}
