package main

import (
	"fmt"
	"sort"
)

// 轉換成 Coin Change 問題，用 DP 找當下數字的最多組合
func combinationSum4(nums []int, target int) int {
	sortedNums := make([]int, len(nums))
	for i := range nums {
		sortedNums[i] = nums[i]
	}
	sort.Sort(sort.IntSlice(sortedNums))
	res := make([]int, target+1)
	res[0] = 1
	for i := 1; i <= target; i++ {
		for _, c := range sortedNums {
			if c > i {
				break
			}
			res[i] += res[i-c]
		}
	}
	return res[target]
}

func main() {
	// nums := []int{2, 3, 6, 7}
	// target := 7
	nums := []int{4, 2, 1}
	target := 32
	fmt.Println(combinationSum4(nums, target))
}
