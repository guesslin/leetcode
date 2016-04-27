package main

import (
	"fmt"
	"os"
	"strings"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minBuy := prices[0]
	maxDiff := 0
	for i := 1; i < len(prices); i++ {
		minBuy = min(minBuy, prices[i])
		maxDiff = max(maxDiff, prices[i]-minBuy)
		fmt.Println(minBuy, maxDiff)
	}
	return maxDiff
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(strings.Split(os.Args[0], "/")[1])
	fmt.Println(maxProfit([]int{1, 0, 3}))
	fmt.Println(maxProfit([]int{17, 24, 1, 5}))
	fmt.Println(maxProfit([]int{1, 5, 17, 25}))
	fmt.Println(maxProfit([]int{17, 24, 1, 10}))
}
