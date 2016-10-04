package main

func longestPalindrome(s string) int {
	chars := make(map[byte]int)
	maxPalin := 0
	for _, c := range []byte(s) {
		chars[c]++
		if chars[c] == 2 {
			maxPalin += 2
			delete(chars, c)
		}
	}
	if len(chars) > 0 {
		maxPalin++
	}
	return maxPalin
}
