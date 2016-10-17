package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	res := 0
	for i := 1; i <= N; i++ {
		res = res ^ i
		res = res ^ A[i-1]
	}
	res = res ^ (N + 1)
	return res
}
