package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		mid := (s + e) / 2
		if nums[mid] >= target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return s
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		posi   int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, posi: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, posi: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, posi: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, posi: 0},
	}
	for i, c := range testCases {
		res := searchInsert(c.nums, c.target)
		if res != c.posi {
			fmt.Println(i, "case is wrong while Expect", c.posi, "get", res)
		}
	}
}
