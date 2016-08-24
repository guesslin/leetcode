package main

import (
	"fmt"
)

func isPowerOfFour(num int) bool {
	compare := num - 1
	return num > 0 && (num&compare == 0) && (compare%3 == 0)
}

func main() {
	testCases := []struct {
		input  int
		expect bool
	}{
		{input: 16, expect: true},
		{input: 19, expect: false},
	}
	for _, c := range testCases {
		res := isPowerOfFour(c.input)
		if res != c.expect {
			fmt.Printf("Input: %d, Expect: %v, Get: %v\n", c.input, c.expect, res)
		}

	}
}
