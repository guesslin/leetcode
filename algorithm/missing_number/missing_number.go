package main

import (
	"fmt"
)

// 要線性時間，所以不可以用 sort 來做排序
func missingNumber(nums []int) int {
	n := len(nums)
	sum := (n + 1) * n / 2
	for _, c := range nums {
		sum -= c
	}
	return sum
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 4}))
	size := 25
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(missingNumber(arr))
}
