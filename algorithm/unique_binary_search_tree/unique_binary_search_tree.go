package main

import (
	"fmt"
)

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	count := make([]int, n+1)
	count[0] = 1
	count[1] = 1
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			sum += count[j-1] * count[i-j]
		}
		count[i] = sum
	}
	return count[n]
}

func main() {
	fmt.Println(numTrees(0))
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
}
