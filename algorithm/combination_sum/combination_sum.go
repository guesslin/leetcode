package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sol := make([]int, 0)
	record := make(map[string]bool)
	var foo func(int)
	foo = func(current int) {
		for _, c := range candidates {
			if current+c == target {
				sol = append(sol, c)
				tmp := make([]int, len(sol))
				for i := 0; i < len(sol); i++ {
					tmp[i] = sol[i]
				}
				sort.Sort(sort.IntSlice(tmp))
				cur := fmt.Sprintf("%v", tmp)
				if _, ok := record[cur]; !ok {
					res = append(res, tmp)
					record[cur] = true
				}
				sol = sol[:len(sol)-1]
			} else if current+c < target {
				sol = append(sol, c)
				foo(current + c)
				sol = sol[:len(sol)-1]
			}
		}
	}
	for _, c := range candidates {
		sol = append(sol, c)
		if c == target {
			tmp := make([]int, len(sol))
			for i := 0; i < len(sol); i++ {
				tmp[i] = sol[i]
			}
			sort.Sort(sort.IntSlice(tmp))
			cur := fmt.Sprintf("%v", tmp)
			if _, ok := record[cur]; !ok {
				res = append(res, tmp)
				record[cur] = true
			}
		} else if c < target {
			foo(c)
		}
		sol = sol[:len(sol)-1]
	}
	return res
}

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(nums, target))
}
