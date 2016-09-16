package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	sums := make([][]int, 0, n)
	row := make([]int, m)
	row[0] = grid[0][0]
	for j := 1; j < m; j++ {
		row[j] = row[j-1] + grid[0][j]
	}
	sums = append(sums, row)
	for i := 1; i < n; i++ {
		row := make([]int, m)
		row[0] = sums[i-1][0] + grid[i][0]
		for j := 1; j < m; j++ {
			row[j] = min(row[j-1], sums[i-1][j]) + grid[i][j]
		}
		sums = append(sums, row)
	}
	return sums[n-1][m-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	grid := [][]int{
		[]int{0},
	}
	fmt.Println(grid)
	fmt.Println(minPathSum(grid))
}
