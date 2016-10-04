package main

import (
	"fmt"
)

type Node struct {
	sqrtArea int
	upper    int
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	record := make([][]Node, len(matrix))
	maxArea := 0
	// first row = 0
	for col := 0; col < len(matrix[0]); col++ {
		tmp := int(matrix[0][col] - '0')
		record[0] = append(record[0], Node{sqrtArea: tmp, upper: tmp})
		// most situation should be useless, unless there all be zero except the first row
		if tmp > maxArea {
			maxArea = tmp
		}
	}
	for row := 1; row < len(matrix); row++ {
		curRight := 0
		for col := 0; col < len(matrix[row]); col++ {
			curVal := int(matrix[row][col] - '0')
			curUp := curVal
			curArea := 0
			if curVal == 0 || col-1 < 0 {
				curRight = curVal
				curArea = curVal
			} else {
				curArea = min(curRight, record[row-1][col].upper, record[row-1][col-1].sqrtArea) + 1
				curUp += record[row-1][col].upper
				curRight++
			}
			record[row] = append(record[row], Node{sqrtArea: curArea, upper: curUp})
			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}
	return maxArea * maxArea
}

func min(nums ...int) int {
	mins := nums[0]
	for _, c := range nums {
		if mins > c {
			mins = c
		}
	}
	return mins
}

func main() {
	matrix := [][]byte{
		[]byte("0001"),
		[]byte("1101"),
		[]byte("1111"),
		[]byte("0111"),
		[]byte("0111"),
	}
	for _, r := range matrix {
		fmt.Println(string(r))
	}
	fmt.Println(maximalSquare(matrix))
	matrix2 := [][]byte{
		[]byte("101101"),
		[]byte("111111"),
		[]byte("011011"),
		[]byte("111010"),
		[]byte("011111"),
		[]byte("110111"),
	}
	for _, r := range matrix2 {
		fmt.Println(string(r))
	}
	fmt.Println(maximalSquare(matrix2))
	matrix3 := [][]byte{
		[]byte("00010111"),
		[]byte("01100101"),
		[]byte("10111101"),
		[]byte("00010000"),
		[]byte("00100010"),
		[]byte("11100111"),
		[]byte("10011001"),
		[]byte("01001100"),
		[]byte("10010000"),
	}
	for _, r := range matrix3 {
		fmt.Println(string(r))
	}
	fmt.Println(maximalSquare(matrix3))
	matrix4 := [][]byte{
		[]byte("11"),
		[]byte("11"),
	}
	for _, r := range matrix4 {
		fmt.Println(string(r))
	}
	fmt.Println(maximalSquare(matrix4))
}
