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
