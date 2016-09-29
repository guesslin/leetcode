package main

import (
	"fmt"
)

func grayCode(n int) []int {
	maxGray := 1 << uint(n)
	res := make([]int, maxGray)
	for i := 1; i < maxGray; i++ {
		res[i] = (i >> 1) ^ i
	}
	return res
}

func main() {
	n := 3
	fmt.Println(grayCode(n))
}
