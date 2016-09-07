package main

import (
	"fmt"
)

// UpperBound return the number of element is smaller or equal than target in list
func UpperBound(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		mid := (s + e) / 2
		if nums[mid] > target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	if 0 <= e && nums[e] == target {
		return e
	}
	return -1
}

// LowerBound return the number of element is smaller than target in list
func LowerBound(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		mid := (s + e) / 2
		if nums[mid] >= target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	if s < len(nums) && nums[s] == target {
		return s
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	l, r := LowerBound(nums, target), UpperBound(nums, target)
	return []int{l, r}
}

func main() {
	fmt.Println(searchRange([]int{1}, 0), false)
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
