package main

import "fmt"

func singleNumber(nums []int) []int {
	key := 0
	// first phase, find the key bit
	for i := range nums {
		key = key ^ nums[i]
	}
	key = key & (-key)
	one, two := 0, 0
	for i := range nums {
		if key&nums[i] == 0 {
			one = one ^ nums[i]
		} else {
			two = two ^ nums[i]
		}
	}
	return []int{one, two}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}
