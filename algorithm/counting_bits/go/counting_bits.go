package main

import (
	"fmt"
)

func countBits(num int) []int {
	result := make([]int, num+1)
	if num == 0 {
		return result
	}
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

func main() {
	r := countBits(2)
	fmt.Println(r)
}
