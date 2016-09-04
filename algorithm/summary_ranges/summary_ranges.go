package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	res := make([]string, 0)
	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			if start != end {
				res = append(res, fmt.Sprintf("%d->%d", start, end))
			} else {
				res = append(res, fmt.Sprintf("%d", start))
			}
			start, end = nums[i], nums[i]
		}
	}
	if start != end {
		res = append(res, fmt.Sprintf("%d->%d", start, end))
	} else {
		res = append(res, fmt.Sprintf("%d", start))
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7, 8, 9}
	// nums := []int{1, 3}
	fmt.Println(summaryRanges(nums))
}
