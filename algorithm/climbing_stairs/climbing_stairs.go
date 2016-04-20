package main

import (
	"fmt"
)

func climbStairs(n int) int {
	fib := []int{1, 1, 2, 3, 5, 8, 13}
	for i := 7; i <= n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib[n]
}

func main() {
	fmt.Println(13, climbStairs(13))
}
