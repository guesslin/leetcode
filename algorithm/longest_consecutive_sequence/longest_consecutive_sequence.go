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
	record := make(map[int]int)
	for _, c := range nums {
		if _, ok := record[c]; !ok {
			record[c]++
		}
	}
	cur := 0
	res := 0
	for len(record) != 0 {
		k := 0
		for k = range record {
			break
		}
		for ; record[k+cur] != 0; cur++ {
			delete(record, k+cur)
		}
		for i := 1; record[k-i] != 0; i++ {
			cur++
			delete(record, k-i)
		}
		if res < cur {
			res = cur
		}
		cur = 0
	}
	return res
}

func main() {
	nums := []int{200, 1, 2, 4, 5, 3, 201}
	fmt.Println(longestConsecutive(nums))
	sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)
}
