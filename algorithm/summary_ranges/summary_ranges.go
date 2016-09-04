package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	res := make([]string, 0)
	rang := make([]int, 2)
	rang[0] = nums[0]
	rang[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == rang[len(rang)-1]+1 {
			rang[len(rang)-1] = nums[i]
		} else {
			rang = append(rang, nums[i])
			rang = append(rang, nums[i])
		}
	}
	for i := 0; i < len(rang); i += 2 {
		if rang[i] != rang[i+1] {
			res = append(res, fmt.Sprintf("%d->%d", rang[i], rang[i+1]))
		} else {
			res = append(res, fmt.Sprintf("%d", rang[i]))
		}
	}

	return res
}

func main() {
	// nums := []int{0, 1, 2, 4, 5, 7, 8, 9, 11}
	nums := []int{1, 3}
	fmt.Println(summaryRanges(nums))
}
