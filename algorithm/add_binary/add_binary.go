package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	var result string
	carry := 0
	a = reverse(a)
	b = reverse(b)
	for i := 0; i < min(len(a), len(b)); i++ {
		t := int(a[i]+b[i]) + carry - 96
		if t >= 2 {
			carry = 1
			t = t - 2
		} else {
			carry = 0
		}
		result += string(t + 48)
	}
	if len(a) > len(b) {
		for i := len(b); i < len(a); i++ {
			t := int(a[i]) + carry - 48
			if t >= 2 {
				carry = 1
				t = t - 2
			} else {
				carry = 0
			}
			result += string(t + 48)

		}
	} else {
		for i := len(a); i < len(b); i++ {
			t := int(b[i]) + carry - 48
			if t >= 2 {
				carry = 1
				t = t - 2
			} else {
				carry = 0
			}
			result += string(t + 48)

		}
	}
	if carry == 1 {
		result += string("1")
	}
	return reverse(result)
}

func reverse(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	testCases := []struct {
		a      string
		b      string
		result string
	}{
		{"1", "10", "11"},
		{"1", "0", "1"},
		{"1", "11", "100"},
		{"11", "11", "110"},
		{"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011", "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"},
	}
	for i, c := range testCases {
		fmt.Printf("%d testcase %s + %s\nexcept:\t%s\nget:\t%s\n", i, c.a, c.b, c.result, addBinary(c.a, c.b))
	}
}
