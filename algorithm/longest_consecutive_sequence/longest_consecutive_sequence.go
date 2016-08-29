package main

import (
	"fmt"
	"sort"
)

// using counting sort?
func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	sort.Sort(sort.IntSlice(nums))
	max := 1
	cur := 1
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			continue
		} else if nums[i] == pre+1 {
			cur++
			pre = nums[i]
		} else {
			if max < cur {
				max = cur
			}
			cur = 1
			pre = nums[i]
		}
	}
	if max < cur {
		max = cur
	}

	return max
}

func main() {
	nums := []int{200, 1, 2, 4, 5, 3, 201}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(nums)
}
