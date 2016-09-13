package main

import (
	"fmt"
)

func maxRotateFunction(A []int) int {
	if len(A) == 0 {
		return 0
	}
	maxSum, sum := 0, 0
	n := len(A)
	// compute f(0)
	for i := range A {
		sum += A[i]
		maxSum += A[i] * i
	}
	fSum := maxSum
	for i := 0; i < n; i++ {
		fSum = fSum - sum + n*A[i]
		if fSum > maxSum {
			maxSum = fSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{4, 3, 2, 6}
	fmt.Println(maxRotateFunction(nums))
}
