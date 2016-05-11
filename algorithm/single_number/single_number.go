package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	r := 0
	for _, c := range nums {
		r = r ^ c
	}
	return r
}

func main() {
	testCases := [][]int{
		{2, 2, 1},
		{1, 1, 2},
	}
	for _, c := range testCases {
		fmt.Println(c, singleNumber(c))
	}
}
