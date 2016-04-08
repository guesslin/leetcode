package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	notFound := []int{-1, -1}
	l, r := 0, len(nums)-1
	for l < r {
		m := (l+r)/2 + 1
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	if nums[r] != target {
		return notFound
	}
	right := r
	l = 0
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return []int{l, right}
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
