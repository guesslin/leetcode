package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	loRow, hiRow := 0, len(matrix)-1
	for loRow <= hiRow {
		midRow := loRow + (hiRow-loRow)/2
		if matrix[midRow][0] == target {
			return true
		}
		if matrix[midRow][0] > target {
			hiRow = midRow - 1
		} else {
			loRow = midRow + 1
		}
	}
	if hiRow < 0 {
		return false
	}
	if loRow == len(matrix) {
		loRow -= 1
	}
	lo, hi := 0, len(matrix[loRow])-1
	row := matrix[hiRow]
	fmt.Println(row)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if row[mid] == target {
			return true
		}
		if row[mid] > target {
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
