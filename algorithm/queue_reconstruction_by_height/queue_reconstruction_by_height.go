package main

import (
	"fmt"
	"sort"
)

type lessFunc func(p, q []int) bool

type multiSorter struct {
	people    [][]int
	lessFuncs []lessFunc
}

func orderBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		lessFuncs: less,
	}
}

func (ms *multiSorter) Sort(p [][]int) {
	ms.people = p
	sort.Sort(ms)
}

func (ms *multiSorter) Len() int {
	return len(ms.people)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.people[i], ms.people[j] = ms.people[j], ms.people[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := ms.people[i], ms.people[j]
	var k int
	for k = 0; k < len(ms.lessFuncs)-1; k++ {
		less := ms.lessFuncs[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.lessFuncs[k](p, q)
}

func reconstructQueue(people [][]int) [][]int {
	height := func(p, q []int) bool {
		return p[0] < q[0]
	}
	k := func(p, q []int) bool {
		return p[1] < q[1]
	}
	orderBy(k, height).Sort(people)
	for i := 0; i < len(people); i++ {
		if people[i][1] == 0 {
			continue
		}
		big := 0
		for j := 0; j < i; j++ {
			if people[j][0] >= people[i][0] {
				big++
			}
		}
		diff := big - people[i][1]
		for j := 0; j < diff; j++ {
			people[i-j], people[i-j-1] = people[i-j-1], people[i-j]
		}
	}
	return people
}

func main() {
	people := [][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2},
	}
	fmt.Println(people)
	fmt.Println(reconstructQueue(people))
}
