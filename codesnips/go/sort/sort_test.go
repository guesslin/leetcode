package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		Nums []int
	}{
		{Nums: []int{5, 4, 3, 2, 1}},
		{Nums: []int{5, 4, 7, 63, 21, 1}},
	}
	for i, c := range testCases {
		QuickSort(c.Nums)
		if !IsSorted(c.Nums) {
			t.Errorf("%d -> %v is fail to sort\n", i, c.Nums)
		}
	}

}

func TestIsSorted(t *testing.T) {
	testCases := []struct {
		Nums   []int
		Expect bool
	}{
		{Nums: []int{1, 2, 3, 4, 5, 6, 7}, Expect: true},
		{Nums: []int{1, 2, 2, 4, 5, 6, 7}, Expect: true},
		{Nums: []int{1, 2, 2, 4, 8, 6, 7}, Expect: false},
	}
	for i, c := range testCases {
		res := IsSorted(c.Nums)
		if res != c.Expect {
			t.Errorf("%d -> %v is fail to IsSorted\n", i, c.Nums)
		}
	}
}
