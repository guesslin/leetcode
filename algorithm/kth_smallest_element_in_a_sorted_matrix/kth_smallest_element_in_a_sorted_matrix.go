package main

import (
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	if k == 1 {
		return matrix[0][0]
	}
	n := len(matrix)
	lo, hi := matrix[0][0], matrix[n-1][n-1]
	for lo <= hi {
		count := 0
		mid := (lo + hi) / 2
		for i := 0; i < n; i++ {
			count += upperBound(matrix[i], mid)
		}
		if count > k-1 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func upperBound(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	matrix := [][]int{
		[]int{1, 5, 9},
		[]int{10, 11, 13},
		[]int{12, 13, 15},
	}
	nums := []int{1, 1, 1, 3, 3, 3, 7, 7, 7}
	fmt.Println(nums)
	fmt.Println(upperBound(nums, 0))
	fmt.Println(upperBound(nums, 1))
	fmt.Println(upperBound(nums, 4))
	fmt.Println(upperBound(nums, 5))
	fmt.Println(matrix)
	fmt.Println(kthSmallest(matrix, 8))
}
