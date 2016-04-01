package tree

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	cases := [][]string{{"1", "2", "2", "3", "nil", "nil", "3", "4", "nil", "nil", "4"}}
	for _, c := range cases {
		BuildTree(c)
	}
}
