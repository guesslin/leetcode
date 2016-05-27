package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	words := strings.Split(s, " ")
	pos := len(words) - 1
	for pos >= 0 && len(words[pos]) == 0 {
		pos -= 1
	}
	if pos < 0 {
		return 0
	}
	return len(words[pos])
}

func main() {
	testCases := []string{
		"Hello world",
		"",
		"a ",
		"b   a    ",
		" ",
	}
	for _, c := range testCases {
		fmt.Printf("%s last word length is %d\n", c, lengthOfLastWord(c))
	}

}
