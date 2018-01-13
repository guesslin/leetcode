package main

// https://leetcode.com/problems/assign-cookies/description/
// 455. Assign Cookies

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Sort(sort.IntSlice(g))
	sort.Sort(sort.IntSlice(s))
	content := 0
	cur := 0
	for i := 0; i < len(s) && cur < len(g); i++ {
		if s[i] >= g[cur] {
			content++
			cur++
		}
	}
	return content
}
