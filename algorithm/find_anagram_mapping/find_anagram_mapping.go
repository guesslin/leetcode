package main

// https://leetcode.com/problems/find-anagram-mappings/description/
// 760. Find Anagram Mappings

func anagramMappings(A []int, B []int) []int {
	locate := make(map[int]int)
	for i, n := range B {
		locate[n] = i
	}
	res := make([]int, 0, len(A))
	for _, n := range A {
		res = append(res, locate[n])
	}
	return res
}
