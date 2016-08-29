package main

import (
	"fmt"
	"sort"
)

// threeSum find elemenets' sum equal to zero
func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	uniq := make(map[string]struct{})
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		target := 0 - nums[i]
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				key := fmt.Sprintf("%d%d%d", nums[i], nums[j], nums[k])
				if _, ok := uniq[key]; !ok {
					tmp := []int{nums[i], nums[j], nums[k]}
					res = append(res, tmp)
					uniq[key] = struct{}{}
				}
				j++
			}
		}
	}
	return res
}

func main() {
	s := []int{-1, 0, 1, 2, -1, -4, -2}
	fmt.Println(s)
	fmt.Println(threeSum(s))
}
