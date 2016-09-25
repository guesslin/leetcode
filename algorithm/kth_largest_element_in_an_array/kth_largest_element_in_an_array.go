package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.IntSlice(nums))
	return nums[len(nums)-k]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(nums, k, findKthLargest(nums, k))
}
