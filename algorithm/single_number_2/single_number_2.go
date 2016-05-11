package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	once, twice := 0, 0
	for _, n := range nums {
		once = (once ^ n) & ^twice
		twice = (twice ^ n) & ^once
	}
	return once
}

func main() {
	testCases := [][]int{
		{1},
		{1, 1, 1, 2},
	}
	for _, c := range testCases {
		fmt.Println(c, singleNumber(c))
	}
}
