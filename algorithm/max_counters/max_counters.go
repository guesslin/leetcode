package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	res := make([]int, N)
	curMax, curBase := 0, 0
	// maxCounter = N+1
	for _, op := range A {
		if op <= N {
			if res[op-1] >= curBase {
				res[op-1]++
			} else {
				res[op-1] = curBase + 1
			}
			if res[op-1] > curMax {
				curMax = res[op-1]
			}
		} else {
			curBase = curMax
		}
	}
	for i := 0; i < N; i++ {
		if res[i] < curBase {
			res[i] = curBase
		}
	}
	return res
}
