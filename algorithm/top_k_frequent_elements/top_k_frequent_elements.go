package main

import (
	"fmt"
	"reflect"
)

func topKFrequent(nums []int, k int) []int {
	record := make(map[int]int)
	for _, c := range nums {
		record[c]++
	}
	n := len(nums)
	freq := make([][]int, len(nums)+1)
	res := make([]int, 0, k)
	for key, value := range record {
		freq[value] = append(freq[value], key)
	}
	for i := n; i >= 0; i-- {
		res = append(res, freq[i]...)
		if len(res) >= k {
			break
		}
	}
	fmt.Println(res)
	return res[:k]
}

func main() {
	testCases := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{nums: []int{1}, k: 1, expect: []int{1}},
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expect: []int{1, 2}},
		{nums: []int{2, 3, 4, 1, 4, 0, 4, -1, -2, -1}, k: 2, expect: []int{4, -1}},
	}
	for i, c := range testCases {
		res := topKFrequent(c.nums, c.k)
		eq := reflect.DeepEqual(res, c.expect)
		if !eq {
			fmt.Println(i, "res", res, "expect", c.expect, c.k, c.nums)
		}
	}

}
