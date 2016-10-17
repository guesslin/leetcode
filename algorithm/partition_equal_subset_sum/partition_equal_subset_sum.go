package main

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum&1 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, c := range nums {
		for j := sum; j >= c; j-- {
			dp[j] = dp[j] || dp[j-c]
		}
	}
	return dp[sum]
}
