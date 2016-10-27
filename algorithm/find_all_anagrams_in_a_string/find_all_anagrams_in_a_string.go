package main

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	pMemo := make(map[byte]int)
	sMemo := make(map[byte]int)
	for _, c := range []byte(p) {
		pMemo[c]++
	}
	for i, c := range []byte(s) {
		if i < len(p) {
			sMemo[c]++
		} else {
			if isEqual(sMemo, pMemo) {
				res = append(res, i-len(p))
			}
			sMemo[c]++
			sMemo[s[i-len(p)]]--
			if sMemo[s[i-len(p)]] == 0 {
				delete(sMemo, s[i-len(p)])
			}
		}
	}
	if isEqual(sMemo, pMemo) {
		res = append(res, len(s)-len(p))
	}
	return res
}

func isEqual(s, p map[byte]int) bool {
	if len(s) != len(p) {
		return false
	}
	for k, v := range p {
		if s[k] != v {
			return false
		}
	}
	return true
}
