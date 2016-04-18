package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	count := make(map[int]int)
	length := len(nums)
	var result int
	for _, n := range nums {
		count[n]++
	}
	for key, n := range count {
		if n*2 >= length {
			result = key
			break
		}
	}
	return result
}

func main() {
	fmt.Println(majorityElement([]int{5, 1, 5, 5, 2, 2}))
	fmt.Println(majorityElement([]int{6, 5, 5}))
}
