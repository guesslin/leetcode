package main

import (
	"fmt"
)

// 1, 2, 3, 4 => 24, 12, 8, 6
// without division
func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	p := 1
	for i := 0; i < n; i++ {
		output[i] = p
		p = p * nums[i]
	}
	p = 1
	for i := n - 1; i >= 0; i-- {
		output[i] = output[i] * p
		p = p * nums[i]
	}
	return output
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
