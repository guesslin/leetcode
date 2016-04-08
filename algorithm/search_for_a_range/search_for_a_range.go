package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	fmt.Println(nums, target)
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	ptr := (left + right) / 2
	for nums[ptr] != target {
		if left == right {
			return []int{-1, -1}
		}
		if nums[ptr] > target {
			right = ptr
			ptr = (left + right) / 2
		} else {
			left = ptr + 1
			ptr = (left + right) / 2
		}
	}
	mid := ptr
	for mid = ptr; mid <= right && nums[mid] == target; mid++ {
	}
	right = mid - 1
	for mid = ptr; mid >= left && nums[mid] == target; mid-- {
	}
	left = mid + 1
	return []int{left, right}
}

func main() {
	fmt.Println(searchRange([]int{1, 4}, 4), true)
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 10}, 8), true)
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6), false)
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 9), false)
	fmt.Println(searchRange([]int{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 10}, 6), true)
	fmt.Println(searchRange([]int{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 10}, 9), false)
	fmt.Println(searchRange([]int{6, 6, 6, 6, 6, 6, 6, 6, 6, 6}, 6), true)
	fmt.Println(searchRange([]int{1, 4, 5}, 4), true)
	fmt.Println(searchRange([]int{1, 2, 3}, 1), true)
	fmt.Println(searchRange([]int{2, 2}, 2), true)
	fmt.Println(searchRange([]int{2, 2}, 3), false)
	fmt.Println(searchRange([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 3), true)
}
