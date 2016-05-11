package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	step := 0
	for step < len(nums) {
		step += nums[step]
		if step >= len(nums)-1 {
			return true
		}
		if nums[step] == 0 {
			need := 2
			for {
				step--
				if step < 0 {
					return false
				}
				if nums[step] >= need {
					break
				}
				need++
			}
		}
	}
	return true
}

func main() {
	testCases := []struct {
		input  []int
		expect bool
	}{
		{input: []int{2, 3, 1, 1, 4}, expect: true},
		{input: []int{3, 3, 0, 0, 1}, expect: true},
		{input: []int{3, 2, 1, 0, 4}, expect: false},
		{input: []int{0}, expect: true},
	}
	for i, c := range testCases {
		fmt.Printf("%d-th case: %v, is match to expect result %v\n", i+1, c.input, canJump(c.input) == c.expect)
	}
}
