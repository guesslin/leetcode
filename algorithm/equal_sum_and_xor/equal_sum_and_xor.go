package main

import (
	"fmt"
)

// countValues return the count of i, which i satisfy 1 <= i <= n and n+i = n^i
// for example if n = 7, r can be 0 => return 1
func countValues(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		if (n + i) == (n ^ i) {
			count++
		}
	}
	return count
}

// countValues2 use bit
// a + b = a^b + a&b => n+i = n^i + n&i ==> let n&i = 0 >> n+i = n^i
func countValues2(n int) int {
	l := 1
	var bit uint
	for l < n {
		if l&n == 0 {
			bit++
		}
		l = l << 1
	}
	return 1 << bit
}

func main() {
	testCases := []struct {
		num int
	}{
		{num: 7},
		{num: 5},
		{num: 123},
		{num: 129},
		{num: 13},
		{num: 2},
	}
	for i, c := range testCases {
		r1 := countValues(c.num)
		r2 := countValues2(c.num)
		fmt.Println(i, r1, r2, r1 == r2)
	}
}
