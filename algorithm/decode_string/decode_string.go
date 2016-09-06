package main

import (
	"fmt"
)

// 3[a]2[bc] 	    => aaabcbc
// 3[a2[c]]  	    => accaccacc
// 2[abc]3[cd]ef    => abcabccdcdcdef
func decodeString(s string) string {
	res := make([]byte, 0)
	nStack := make([]int, 0)
	cStack := make([][]byte, 0)
	number := 0
	var cs []byte
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			number = 10*number + int(s[i]-'0')
		} else if s[i] == '[' {
			// push number to stack
			nStack = append(nStack, number)
			number = 0
			cStack = append(cStack, cs)
			cs = nil
		} else if s[i] == ']' {
			// pop from stack
			number, nStack = nStack[len(nStack)-1], nStack[:len(nStack)-1]
			var tmp []byte
			for i := 0; i < number; i++ {
				tmp = append(tmp, cs...)
			}
			if len(nStack) == 0 {
				res = append(res, tmp...)
				cs = make([]byte, 0)
			} else if len(cStack) != 0 {
				cs, cStack = cStack[len(cStack)-1], cStack[:len(cStack)-1]
				cs = append(cs, tmp...)
			} else {
				cs = tmp
			}
			number = 0
		} else {
			if len(nStack) == 0 {
				res = append(res, s[i])
			} else {
				cs = append(cs, s[i])
			}
		}
	}
	res = append(res, cs...)
	return string(res)
}

// 3[a]2[bc] 	    => aaabcbc
// 3[a2[c]]  	    => accaccacc
// 2[abc]3[cd]ef    => abcabccdcdcdef
func main() {
	testCases := []struct {
		Input  string
		Expect string
	}{
		{Input: "3[a]2[bc]", Expect: "aaabcbc"},
		{Input: "3[a2[c]]", Expect: "accaccacc"},
		{Input: "2[abc]3[cd]ef", Expect: "abcabccdcdcdef"},
		{Input: "2[2[b]]", Expect: "bbbb"},
		{Input: "abc3[cd]xyz", Expect: "abccdcdcdxyz"},
	}
	for i, c := range testCases {
		res := decodeString(c.Input)
		fmt.Println("Input", i, ":", c.Input)
		fmt.Println("Expect:", c.Expect)
		fmt.Println("Res   :", res)
	}
}
