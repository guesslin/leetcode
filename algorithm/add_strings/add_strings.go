package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	n1, n2 := []byte(num1), []byte(num2)
	l1, l2 := len(n1), len(n2)
	minLen := l1 + l2 - max(l1, l2)
	res := make([]byte, 0, max(l1, l2)+1) // 注意進位狀態
	i, carry := 0, 0
	for ; i < minLen; i++ {
		t := int(n1[l1-i-1]-'0') + int(n2[l2-i-1]-'0') + carry
		carry = t / 10
		t = t % 10
		res = append(res, byte(t+'0'))
	}
	ptr := n1
	if minLen == l1 {
		ptr = n2
	}
	maxLen := max(l1, l2)
	for i = minLen; i < maxLen; i++ {
		t := carry + int(ptr[maxLen-i-1]-'0')
		carry = t / 10
		t = t % 10
		res = append(res, byte(t+'0'))
	}
	if carry > 0 {
		res = append(res, byte(carry+'0'))
	}
	reverse(res)
	return string(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := addStrings("700", "321")
	fmt.Println(s)
}
