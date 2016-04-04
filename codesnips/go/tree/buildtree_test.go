package tree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	cases := []struct {
		Input     []string
		Postorder []int
		Inorder   []int
		Preorder  []int
	}{
		{
			Input:     []string{"1", "2", "2", "3", "null", "null", "3", "4", "null", "null", "4"},
			Postorder: []int{4, 3, 2, 4, 3, 2, 1},
			Inorder:   []int{4, 3, 2, 1, 2, 3, 4},
			Preorder:  []int{1, 2, 3, 4, 2, 3, 4}},
		{
			Input:     []string{"1", "2", "3", "4", "5", "null", "7"},
			Postorder: []int{4, 5, 2, 7, 3, 1},
			Inorder:   []int{4, 2, 5, 1, 3, 7},
			Preorder:  []int{1, 2, 4, 5, 3, 7}},
	}
	for _, c := range cases {
		result := BuildTree(c.Input)
		po := PostorderTraversal(result)
		eq := reflect.DeepEqual(po, c.Postorder)
		if !eq {
			t.Errorf("PostOrder Expect %v\nGet %v\n", c.Postorder, po)
		}
		io := InorderTraversal(result)
		eq = reflect.DeepEqual(io, c.Inorder)
		if !eq {
			t.Errorf("InOrder Expect %v\nGet %v\n", c.Inorder, io)
		}
		pre := PreorderTraversal(result)
		eq = reflect.DeepEqual(pre, c.Preorder)
		if !eq {
			t.Errorf("PreOrder Expect %v\nGet %v\n", c.Preorder, pre)
		}
	}
}

func TestToInt(t *testing.T) {
	cases := []struct {
		Input  string
		Result int
	}{
		{"123", 123},
		{"1234567890", 1234567890},
		{"1", 1},
		{"abc", 0},
	}
	for _, c := range cases {
		result := toInt(c.Input)
		if result != c.Result {
			t.Errorf("Input %s, Except %d, Get %d", c.Input, c.Result, result)
		}
	}
}
