package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) < 2 {
		return 0
	}
	minBuy, maxSell := A[0], 0
	for i := 1; i < len(A); i++ {
		minBuy = min(minBuy, A[i])
		maxSell = max(maxSell, A[i]-minBuy)
	}
	return maxSell
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
