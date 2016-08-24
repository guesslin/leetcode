package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				tmp[j] = 1
			} else {
				tmp[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result = append(result, tmp)
	}
	return result
}

func main() {
	n := 5
	r := generate(n)
	for i := 0; i < n; i++ {
		fmt.Println(r[i])
	}
}
