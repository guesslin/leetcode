package main

import (
	"fmt"
)

func nthUglyNumber(n int, primes []int) int {
	m := make([]int, len(primes))
	copy(m, primes)
	index := make([]int, len(primes))
	seq := make([]int, n)
	seq[0] = 1
	for i := 1; i < n; i++ {
		minV := min(m)
		seq[i] = minV
		for j := 0; j < len(index); j++ {
			if minV == m[j] {
				index[j]++
				m[j] = primes[j] * seq[index[j]]
			}
		}
	}
	return seq[n-1]
}

func min(nums []int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if m > nums[i] {
			m = nums[i]
		}
	}
	return m
}

func main() {
	n := 7
	fmt.Printf("%d-th Ugly Number is %d\n", n, nthUglyNumber(n, []int{2, 3, 5}))
}
