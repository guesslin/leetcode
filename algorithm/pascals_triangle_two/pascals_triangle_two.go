package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	var t int
	for i := 0; i <= rowIndex; i++ {
		result[0] = 1
		result[i] = 1
		mid := result[0]
		for j, k := 1, i-1; j <= k; j, k = j+1, k-1 {
			t, mid = mid+result[j], result[j]
			result[j] = t
			result[k] = t
		}
	}
	return result
}

func main() {
	fmt.Println(getRow(10))
}
