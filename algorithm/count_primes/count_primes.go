package main

import (
	"fmt"
)

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	marked := make([]bool, n)
	marked[0] = true
	marked[1] = true
	for i := 2; i*i < n; i++ {
		if marked[i] {
			continue
		}
		for j := 2; i*j < n; j++ {
			marked[i*j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !marked[i] {
			count++
		}
	}
	return count
}

func main() {
	testCases := []struct {
		input  int
		result int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{100, 25},
	}
	for i, c := range testCases {
		fmt.Println(i, c.result, countPrimes(c.input))
	}
}
