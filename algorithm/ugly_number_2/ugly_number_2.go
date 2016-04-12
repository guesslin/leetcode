package main

import (
	"fmt"
)

func nthUglyNumber(n int) int {
	m2, m3, m5 := 2, 3, 5
	var id2, id3, id5 int
	seq := make([]int, n, n)
	seq[0] = 1
	for i := 1; i < n; i++ {
		minV := min(min(m2, m3), m5)
		seq[i] = minV
		if minV == m2 {
			id2++
			m2 = 2 * seq[id2]
		}
		if minV == m3 {
			id3++
			m3 = 3 * seq[id3]
		}
		if minV == m5 {
			id5++
			m5 = 5 * seq[id5]
		}
	}
	return seq[n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	n := 1500
	fmt.Printf("%d-th Ugly Number is %d\n", n, nthUglyNumber(n))
}
