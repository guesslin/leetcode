package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	distance := Y - X
	steps := distance / D
	if distance%D > 0 {
		steps++
	}
	return steps
}
