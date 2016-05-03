package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	minHeight := heights[0]
	index := 0
	for i := 1; i < len(heights); i++ {
		if minHeight > heights[i] {
			minHeight = heights[i]
			index = i
		}
	}
	left := largestRectangleArea(heights[:index])
	right := largestRectangleArea(heights[index+1:])
	mid := minHeight * len(heights)
	max := mid
	if max < left {
		max = left
	}
	if max < right {
		max = right
	}
	return max
}

func main() {
	rectangle := []int{2, 1, 5, 6, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	fmt.Println(rectangle, largestRectangleArea(rectangle))
}
