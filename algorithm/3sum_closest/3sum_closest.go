package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// threeSumClosest use two pointers
func threeSumClosest(nums []int, target int) int {
	// Element less than three can't give three elements' sum
	if len(nums) < 3 {
		return 0
	}
	sort.IntSlice(nums).Sort()
	best := nums[0] + nums[1] + nums[2]
	if best == target {
		return best
	}
	min := abs(target - best)
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		n, m := i+1, len(nums)-1
		for n < m {
			sum := tmp + nums[n] + nums[m]
			if sum == target {
				return sum
			}
			diff := abs(target - sum)
			if min > diff {
				best = sum
				min = diff
			}
			if sum > target {
				m--
			} else {
				n++
			}
		}
	}
	return best
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		expect int
	}{
		{nums: []int{1}, target: 1, expect: 0},
		{nums: []int{-1, 2, 1, -4}, target: 1, expect: 2},
		{nums: []int{1, 1, 2, 3}, target: 4, expect: 4},
		{nums: []int{0, 2, 1, -3}, target: 1, expect: 0},
	}
	for i, c := range testCases {
		r := threeSumClosest(c.nums, c.target)
		if r != c.expect {
			fmt.Printf("Case %02d: target %d\tGet %d\tExpect %d\tInput %v\n", i, c.target, r, c.expect, c.nums)
		}

	}
}
