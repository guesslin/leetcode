package search

import (
	"testing"
)

func TestBiSearch(t *testing.T) {
	testCases := []struct {
		Nums   []int
		Target int
		Expect int
	}{
		{Nums: []int{1, 2, 3, 4, 5, 6, 7}, Target: 8, Expect: -1},
		{Nums: []int{1, 2, 3, 4, 5, 6, 7}, Target: 7, Expect: 6},
		{Nums: []int{1, 2, 3, 4, 5, 6, 7}, Target: 1, Expect: 0},
		{Nums: []int{1, 2, 3, 4, 5, 6, 7}, Target: 6, Expect: 5},
	}
	for i, c := range testCases {
		t.Logf("Case %d, Expect %d", i, c.Expect)
		res := BiSearch(c.Nums, c.Target)
		if res != c.Expect {
			t.Errorf("TestCases %d expect %d but get %d\n", i, c.Expect, res)
		}
	}
}

func TestUpperBound(t *testing.T) {
	testCases := []struct {
		Nums   []int
		Target int
		Expect int
	}{
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 7, Expect: 7},
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 6, Expect: 7},
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 1, Expect: 2},
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 3, Expect: 4},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 7, Expect: 6},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 0, Expect: 0},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 1, Expect: 2},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 5, Expect: 4},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 20, Expect: 6},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 19, Expect: 3},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 21, Expect: 6},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 0, Expect: 0},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 31, Expect: 8},
	}
	for i, c := range testCases {
		t.Logf("Case %d, Expect %d", i, c.Expect)
		res := UpperBound(c.Nums, c.Target)
		if res != c.Expect {
			t.Errorf("TestCases %d expect %d but get %d\n", i, c.Expect, res)
		}
	}
}

func TestLowerBound(t *testing.T) {
	testCases := []struct {
		Nums   []int
		Target int
		Expect int
	}{
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 7, Expect: 7},
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 6, Expect: 6},
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 1, Expect: 1},
		{Nums: []int{0, 1, 2, 3, 4, 5, 6}, Target: 3, Expect: 3},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 7, Expect: 6},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 0, Expect: 0},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 1, Expect: 0},
		{Nums: []int{1, 1, 2, 2, 6, 6, 8}, Target: 5, Expect: 4},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 20, Expect: 3},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 19, Expect: 3},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 21, Expect: 6},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 0, Expect: 0},
		{Nums: []int{10, 10, 10, 20, 20, 20, 30, 30}, Target: 31, Expect: 8},
	}
	for i, c := range testCases {
		t.Logf("Case %d, Expect %d", i, c.Expect)
		res := LowerBound(c.Nums, c.Target)
		if res != c.Expect {
			t.Errorf("TestCases %d expect %d but get %d\n", i, c.Expect, res)
		}
	}
}
