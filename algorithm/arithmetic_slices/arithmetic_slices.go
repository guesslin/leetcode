package main

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	count := make([]int, n)
	result := 0
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			count[i] = count[i-1] + 1
			result += count[i]
		}
	}
	return result
}
