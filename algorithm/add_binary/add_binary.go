package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	byteSliceA := reverse([]byte(a))
	byteSliceB := reverse([]byte(b))
	lenA, lenB := len(a), len(b)
	result := make([]byte, 0, max(lenA, lenB)+1)
	carry := 0
	for i := 0; i < min(lenA, lenB); i++ {
		t := int(byteSliceA[i]+byteSliceB[i]) + carry - 96
		if t >= 2 {
			carry = 1
			t = t - 2
		} else {
			carry = 0
		}
		result = append(result, byte(t+48))
	}
	if lenA > lenB {
		for i := lenB; i < lenA; i++ {
			t := int(byteSliceA[i]) + carry - 48
			if t >= 2 {
				carry = 1
				t = t - 2
			} else {
				carry = 0
			}
			result = append(result, byte(t+48))
		}
	} else {
		for i := lenA; i < lenB; i++ {
			t := int(byteSliceB[i]) + carry - 48
			if t >= 2 {
				carry = 1
				t = t - 2
			} else {
				carry = 0
			}
			result = append(result, byte(t+48))
		}
	}
	if carry == 1 {
		result = append(result, '1')
	}
	return string(reverse(result))
}

func reverse(s []byte) []byte {
	result := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
