package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	sol := make([]int, 0)
	seen := make([]bool, len(nums))
	var f func(int)
	f = func(n int) {
		if n == len(nums) {
			tmp := make([]int, len(nums))
			fmt.Println(sol)
			for i, c := range sol {
				tmp[i] = c
			}
			res = append(res, tmp)
			return
		}
		for i := range nums {
			if !seen[i] {
				seen[i] = true
				sol = append(sol, nums[i])
				f(n + 1)
				sol = sol[:len(sol)-1]
				seen[i] = false
			}
		}
	}
	f(0)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
