package main

import (
	"fmt"
)

type stack []int

func (s stack) Top() int {
	return s[len(s)-1]
}

func largestRectangleArea(heights []int) int {
	var s stack
	maxArea := 0
	i := 0
	length := len(heights)
	for i < length {
		if len(s) == 0 || heights[s.Top()] <= heights[i] {
			s = append(s, i)
			i++
		} else {
			top := s.Top()
			s = s[:len(s)-1]
			wide := i
			if len(s) != 0 {
				wide = i - s.Top() - 1
			}
			tmpArea := heights[top] * wide
			if maxArea < tmpArea {
				maxArea = tmpArea
			}
		}
	}
	for len(s) > 0 {
		top := s.Top()
		s = s[:len(s)-1]
		wide := i
		if len(s) != 0 {
			wide = i - s.Top() - 1
		}
		tmpArea := heights[top] * wide
		if maxArea < tmpArea {
			maxArea = tmpArea
		}
	}
	return maxArea
}

func main() {
	rectangle := []int{2, 1, 5, 6, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	fmt.Println(rectangle, largestRectangleArea(rectangle))
}
