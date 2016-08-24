package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	var rows [9][9]int
	var cols [9][9]int
	var box [9][9]int
	var pos int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			pos = int(board[i][j] - '1')
			rows[i][pos]++
			cols[j][pos]++
			box[((i/3)*3)+(j/3)][pos]++
			if rows[i][pos] > 1 || cols[j][pos] > 1 || box[((i/3)*3)+(j/3)][pos] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	/*
		t := [][]byte{
			[]byte(".87654321"),
			[]byte("2........"),
			[]byte("3........"),
			[]byte("4........"),
			[]byte("5........"),
			[]byte("6........"),
			[]byte("7........"),
			[]byte("8........"),
			[]byte("9........"),
		}
	*/
	t2 := [][]byte{
		[]byte("....5..1."),
		[]byte(".4.3....."),
		[]byte(".....3..1"),
		[]byte("8......2."),
		[]byte("..2.7...."),
		[]byte(".15......"),
		[]byte(".....2..."),
		[]byte(".2.9....."),
		[]byte("..4......"),
	}
	// fmt.Println(isValidSudoku(t))
	fmt.Println(isValidSudoku(t2))
}
