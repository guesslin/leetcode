package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int) // number -> last appear position
	for i, c := range nums {
		if p, ok := record[c]; ok {
			if i-p <= k {
				return true
			}
		}
		record[c] = i
	}
	return false
}

func main() {
	nums := []int{1, 2, 1}
	k := 0
	fmt.Println(nums)
	fmt.Println(containsNearbyDuplicate(nums, k))
}
