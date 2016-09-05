package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	n := len(nums)
	k = k % n
	k = n - k
	tmp := append(nums[k:], nums[:k]...)
	for i := 0; i < n; i++ {
		nums[i] = tmp[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	rotate(nums, 1)
	fmt.Println(nums)
}
