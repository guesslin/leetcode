package main

import (
	"fmt"
)

func buildVersionList(version string) []int {
	res := make([]int, 0)
	num := 0
	for i := 0; i < len(version); i++ {
		if version[i] == '.' {
			res = append(res, num)
			num = 0
		} else {
			num = num*10 + int(version[i]-'0')
		}
	}
	res = append(res, num)
	return res
}

func compareVersion(version1 string, version2 string) int {
	v1n := buildVersionList(version1)
	v2n := buildVersionList(version2)
	l1 := len(v1n)
	l2 := len(v2n)
	length := l1
	if l1 > l2 {
		length = l2
	}
	for i := 0; i < length; i++ {
		if v1n[i] == v2n[i] {
			continue
		} else if v1n[i] > v2n[i] {
			return 1
		}
		return -1
	}
	if l1 > l2 {
		if v1n[l1-1] == 0 {
			return 0
		}
		return 1
	} else if l1 < l2 {
		if v2n[l2-1] == 0 {
			return 0
		}
		return -1
	}
	return 0
}

func main() {
	fmt.Println(-1, compareVersion("0.1", "0.2"))
	fmt.Println(1, compareVersion("1.1", "0.2"))
	fmt.Println(0, compareVersion("1.1", "1.1"))
	fmt.Println(-1, compareVersion("1.1", "1.1.1"))
	fmt.Println(1, compareVersion("1.2", "1.1.1"))
	fmt.Println(0, compareVersion("1.2", "1.2.0"))
}
