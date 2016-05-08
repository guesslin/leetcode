package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	if n == 0 {
		return nil
	}
	var result [][]int
	tmp := make([]int, k)
	var rc func(int, int)
	last := k - 1
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			tmp[i] = j + 1
			if i == last {
				r := make([]int, k)
				for n := range tmp {
					r[n] = tmp[n]
				}
				result = append(result, r)
			} else {
				rc(i+1, j+1)
			}
		}
	}
	rc(0, 0)
	return result
}

func main() {
	combs := combine(5, 3)
	for _, c := range combs {
		fmt.Println(c)
	}
}
