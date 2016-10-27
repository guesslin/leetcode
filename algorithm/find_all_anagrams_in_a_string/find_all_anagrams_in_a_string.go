package main

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	var pMemo [26]int
	var sMemo [26]int
	for _, c := range []byte(p) {
		pMemo[int(c-'a')]++
	}
	for i, c := range []byte(s) {
		if i < len(p) {
			sMemo[int(c-'a')]++
		} else {
			if isEqual(sMemo, pMemo) {
				res = append(res, i-len(p))
			}
			sMemo[int(c-'a')]++
			sMemo[int(s[i-len(p)]-'a')]--
		}
	}
	if isEqual(sMemo, pMemo) {
		res = append(res, len(s)-len(p))
	}
	return res
}

func isEqual(s, p [26]int) bool {
	for i := 0; i < 26; i++ {
		if s[i] != p[i] {
			return false
		}
	}
	return true
}
