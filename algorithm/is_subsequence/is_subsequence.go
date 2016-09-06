package main

// isSubsequence check if s is subsequence of t
// s := ace, t := abcde -> true
// s := aec, t := abcde -> false
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i, j := 0, 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
