package main

import (
	"fmt"
)

func binaryGap(N int) int {
	l := 0
	c := 0
	// shift to first left 1's
	for N > 0 {
		if N&1 == 1 {
			break
		}
		N = N >> 1
	}
	for N > 0 {
		if N&1 == 1 {
			l = max(l, c)
			c = 1
		} else {
			c = c + 1
		}
		N = N >> 1
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	cases := []struct {
		input    int
		expected int
	}{
		{input: 22, expected: 2}, // 0b10110
		{input: 5, expected: 2},  // 0b00101
		{input: 6, expected: 1},  // 0b00110
		{input: 8, expected: 0},  // 0b01000
	}
	for _, c := range cases {
		result := binaryGap(c.input)
		fmt.Printf("input: %5d, expected: %5d, result: %5d, Passed: %+v\n", c.input, c.expected, result, c.expected == result)
	}
}
