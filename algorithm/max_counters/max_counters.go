package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	res := make([]int, N)
	curMax := 0
	// maxCounter = N+1
	for _, op := range A {
		if op <= N {
			res[op-1]++
			if res[op-1] > curMax {
				curMax = res[op-1]
			}
		} else {
			for i := range res {
				res[i] = curMax
			}
		}
	}
	return res
}
