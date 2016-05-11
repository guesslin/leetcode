package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	sep := "."
	var v1n, v2n []int
	for _, v := range strings.Split(version1, sep) {
		tmp, _ := strconv.Atoi(v)
		v1n = append(v1n, tmp)
	}
	for _, v := range strings.Split(version2, sep) {
		tmp, _ := strconv.Atoi(v)
		v2n = append(v2n, tmp)
	}
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
