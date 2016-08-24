package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		for i < len(nums) && nums[i] != val {
			i++
		}
		for j >= 0 && nums[j] == val {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return j + 1
}

func main() {
	testCases := []struct {
		nums   []int
		val    int
		expcet int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, expcet: 2},
		{nums: []int{1}, val: 1, expcet: 0},
		{nums: []int{}, val: 1, expcet: 0},
	}
	for _, c := range testCases {
		fmt.Printf("Nums: %v\n", c.nums)
		res := removeElement(c.nums, c.val)
		fmt.Printf("Nums: %v, Len: %d, Same: %v\n", c.nums, res, res == c.expcet)
	}
}
