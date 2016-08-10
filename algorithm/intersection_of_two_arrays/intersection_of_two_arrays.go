package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	sort.IntSlice(nums1).Sort()
	sort.IntSlice(nums2).Sort()
	fmt.Println(nums1)
	fmt.Println(nums2)
	numMap := make(map[int]int)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			numMap[nums1[i]] += 1
			j++
		} else if nums1[i] > nums2[j] {
			j++
			continue
		}
		i++
	}
	res := make([]int, 0, len(numMap))
	for k := range numMap {
		res = append(res, k)
	}
	return res
}

func main() {
	fmt.Println(intersection([]int{9, 1, 2, 2, 3, 1}, []int{2, 2, 8, 7, 13, 3, 9}))
}
