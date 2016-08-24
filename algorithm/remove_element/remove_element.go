package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			if i == len(nums)-1 {
				nums = nums[:i]
			} else {
				nums = append(nums[:i], nums[i+1:]...)
			}
		} else {
			i++
		}
	}
	return len(nums)
}

func main() {
	testCases := []struct {
		nums   []int
		val    int
		expcet int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, expcet: 2},
	}
	for _, c := range testCases {
		fmt.Printf("Nums: %v\n", c.nums)
		res := removeElement(c.nums, c.val)
		fmt.Printf("Nums: %v, Len: %d\n", c.nums, res)
	}
}
