package solution

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution1 = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		buf4 := make([]byte, 4)
		i := 0
		isContinue := true
		remains := n
		for isContinue {
			step := read4(buf4)
			if step < 4 {
				isContinue = false
			}
			if remains < step {
				step = remains
				isContinue = false
			}
			copy(buf[i:i+step], buf4[:step])
			remains -= step
			i += step
		}
		return i
	}
}

/*
Solution2:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Read N Characters Given Read4.
Memory Usage: 1.9 MB, less than 85.71% of Go online submissions for Read N Characters Given Read4.
*/
var solution2 = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		copied := 0
		step := 4
		remains := min(n, cap(buf))

		for remains >= 4 && step == 4 {
			step = read4(buf[copied : copied+step])
			remains -= step
			copied += step
		}
		if remains > 0 {
			buf4 := make([]byte, 4)
			step = read4(buf4)
			r := min(step, remains)
			copy(buf[copied:copied+r], buf4[:r])
			copied += r
		}
		return copied
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
