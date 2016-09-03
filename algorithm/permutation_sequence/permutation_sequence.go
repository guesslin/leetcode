package main

import (
	"fmt"
)

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	nums := make([]int, n)
	res := make([]byte, 0, n)
	k = k - 1
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = i * factorial[i-1]
		nums[i-1] = i
	}
	for i := n - 1; i > 0; i-- {
		t := k / factorial[i]
		k = k % factorial[i]
		res = append(res, '0'+byte(nums[t]))
		nums = append(nums[:t], nums[t+1:]...)
	}
	res = append(res, '0'+byte(nums[0]))
	return string(res)
}

func main() {
	fmt.Println(getPermutation(3, 5))
}
