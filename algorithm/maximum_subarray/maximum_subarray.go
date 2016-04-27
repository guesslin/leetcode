package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	countSum := make([]int, len(nums)+1)
	maxNum := nums[0]
	for i := range nums {
		maxNum = max(maxNum, nums[i])
		countSum[i+1] = countSum[i] + nums[i]
	}
	sum := maxProfit(countSum)
	if sum == 0 {
		return maxNum
	}
	return sum
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minBuy := prices[0]
	maxDiff := 0
	for i := 1; i < len(prices); i++ {
		minBuy = min(minBuy, prices[i])
		maxDiff = max(maxDiff, prices[i]-minBuy)
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
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{-2, -1}))
	fmt.Println(maxSubArray([]int{-2, 2}))
	fmt.Println(maxSubArray([]int{-2, -1, 1}))
	fmt.Println(maxSubArray([]int{1, 2}))
}
