package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}

func main() {
	testCases := [][]int{
		{1, 2, 3, 4, 1, 5, 6, 2, 4, 30, 25, 29, 23},
	}
	for i, c := range testCases {
		fmt.Printf("%d-th test case: %v, result: %d\n", i+1, c, maxProfit(c))
	}
}
