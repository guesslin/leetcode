package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	var nums []int
	for i := range matrix {
		nums = append(nums, matrix[i]...)
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	fmt.Println(matrix)
	fmt.Println(searchMatrix(matrix, -1), -1)
	// fmt.Println(searchMatrix(matrix, 1), 1)
	// fmt.Println(searchMatrix(matrix, 10), 10)
	// fmt.Println(searchMatrix(matrix, 23), 23)
	fmt.Println(searchMatrix(matrix, 11), 11)
	fmt.Println(searchMatrix(matrix, 20), 20)
	fmt.Println(searchMatrix(matrix, 34), 34)
	fmt.Println(searchMatrix(matrix, 7), 7)
	m2 := [][]int{
		[]int{1},
	}
	fmt.Println(m2)
	fmt.Println(searchMatrix(m2, 2), 2)
}
