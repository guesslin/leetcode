package main

import (
	"fmt"
)

func numSquares(n int) int {
	count := make([]int, 1, n+1)
	for len(count) <= n {
		m := len(count)
		currentMin := m
		for i := 1; i*i <= m; i++ {
			currentMin = min(currentMin, count[m-i*i]+1)
		}
		count = append(count, currentMin)
	}
	return count[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	num := 100
	for i := 1; i < num; i++ {
		fmt.Println(i, "has been composed by", numSquares(i))
	}
}
