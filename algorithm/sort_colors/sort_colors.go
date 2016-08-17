package main

import (
	"fmt"
)

func sortColors(nums []int) {
	count := make([]int, 3)
	for _, c := range nums {
		count[c]++
	}
	p := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < count[i]; j++ {
			nums[p+j] = i
		}
		p += count[i]
	}
}

func main() {
	colors := []int{1, 2, 0, 1, 0, 2, 0, 1, 0, 2, 1, 0, 2}
	fmt.Println(colors)
	sortColors(colors)
	fmt.Println(colors)
}
