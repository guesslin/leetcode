package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	steps := make([][]int, m)
	for col := 0; col < n; col++ {
		steps[0] = append(steps[0], 1) // first row can only reach by move only right from left point
	}
	for row := 1; row < m; row++ {
		steps[row] = append(steps[row], 1) // first col can only reach by move only down from up point
		for col := 1; col < n; col++ {
			steps[row] = append(steps[row], steps[row][col-1]+steps[row-1][col])
		}
	}
	return steps[m-1][n-1]
}

func main() {
	testCases := []struct {
		m, n int
	}{
		{m: 1, n: 2},
		{m: 3, n: 13},
	}
	for i, c := range testCases {
		res := uniquePaths(c.m, c.n)
		fmt.Println(i, c.m, "*", c.n, "=>", res)
	}

}
