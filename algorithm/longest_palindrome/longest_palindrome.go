package main

func longestPalindrome(s string) int {
	chars := make(map[byte]int)
	maxPalin := 0
	for _, c := range []byte(s) {
		chars[c]++
	}
	for k, v := range chars {
		maxPalin += v / 2 * 2
		if v%2 == 0 {
			delete(chars, k)
		}
	}
	if len(chars) > 0 {
		maxPalin++
	}
	return maxPalin
}
