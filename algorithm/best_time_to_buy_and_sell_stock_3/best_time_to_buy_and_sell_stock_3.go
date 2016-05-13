package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buy1, buy2 := -prices[0], -prices[0]
	sell1, sell2 := 0, 0
	for _, p := range prices {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, p+buy1)
		buy2 = max(buy2, sell1-p)
		sell2 = max(sell2, p+buy2)
	}
	return sell2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	testCases := [][]int{
		{1, 2, 3, 4, 1, 5, 6, 2, 5, 30, 24, 29, 23},
		{2, 4, 1},
		{6, 1, 3, 2, 4, 7},
		{},
		{3, 3},
		{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
	}
	for i, c := range testCases {
		fmt.Printf("%d-th case: %v, result %d\n", i+1, c, maxProfit(c))
	}
}
