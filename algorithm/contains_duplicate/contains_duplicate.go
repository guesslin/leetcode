package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if ok := seen[num]; ok {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 2, 3, 4}))
}
