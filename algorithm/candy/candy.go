package main

import (
	"fmt"
)

func candy(ratings []int) int {
	count := make([]int, len(ratings))
	count[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			count[i] = count[i-1] + 1
		} else {
			count[i] = 1
		}
	}
	c := count[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && count[i] <= count[i+1] {
			count[i] = count[i+1] + 1
		}
		c += count[i]
	}
	return c
}

func main() {
	fmt.Println(candy([]int{1, 2, 3, 9, 4, 5}))
}
