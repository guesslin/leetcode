package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums2 is empty, just return nums1 as sorted array
	if n == 0 {
		return
	}
	// nums1 is empty, copy nums2 element to nums1 and return
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	// shift nums1 by n
	for i := m + n - 1; i >= n; i-- {
		nums1[i] = nums1[i-n]
	}
	i, j, p := n, 0, 0
	for ; p < n+m; p++ {
		fmt.Println(nums1)
		if nums1[i] < nums2[j] {
			nums1[p] = nums1[i]
			i++
			if i >= n+m {
				break
			}
		} else {
			nums1[p] = nums2[j]
			j++
			if j >= n {
				break
			}
		}
	}
	p++
	if i >= n+m {
		fmt.Println(nums1)
		for ; p < m+n; p++ {
			nums1[p] = nums2[j]
			j++
		}
	}
}

func main() {
	m, n := 5, 1
	nums1 := []int{1, 2, 4, 5, 6, 0}
	nums2 := []int{3}
	fmt.Println(nums1)
	fmt.Println(nums2)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
