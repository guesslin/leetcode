package main

func maxProfit(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	if k < len(prices)/2 {
		maxSell := 0
		sells := make([]int, k+1)
		buys := make([]int, k+1)
		for i := 1; i <= k; i++ {
			buys[i] = -prices[0]
		}
		for _, p := range prices {
			for i := 1; i <= k; i++ {
				sells[i] = max(sells[i], p+buys[i])
				buys[i] = max(buys[i], sells[i-1]-p)
			}
		}
		for _, s := range sells {
			if s > maxSell {
				maxSell = s
			}
		}
		return maxSell
	}
	prof := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			prof += prices[i] - prices[i-1]
		}
	}
	return prof
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
