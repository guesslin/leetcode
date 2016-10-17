package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	rangeSum := make([]int, len(A)+1)
	for i := 1; i <= len(A); i++ {
		rangeSum[i] = rangeSum[i-1] + A[i-1]
	}
	totalSum := rangeSum[len(A)]
	minAbsSum := abs(totalSum - 2*rangeSum[1])
	for i := 2; i < len(A); i++ {
		minAbsSum = min(minAbsSum, abs(totalSum-2*rangeSum[i]))
	}
	return minAbsSum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
