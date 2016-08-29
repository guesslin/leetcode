package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))
	record := make(map[string]int)
	// 用 two pointer
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			s2 := target - nums[i] - nums[j]
			for m, n := j+1, len(nums)-1; m < n; {
				if nums[m]+nums[n] > s2 {
					n--
				} else if nums[m]+nums[n] < s2 {
					m++
				} else {
					if nums[m]+nums[n] == s2 {
						key := fmt.Sprintf("%d,%d,%d,%d", nums[i], nums[j], nums[m], nums[n])
						if _, ok := record[key]; !ok {
							tmp := []int{nums[i], nums[j], nums[m], nums[n]}
							res = append(res, tmp)
							record[key]++
						}
					}
					m++
				}
			}
		}
	}
	return res
}

func main() {
	s := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(s, 0))
}
