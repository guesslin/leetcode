package main

import (
	"fmt"
)

type Sudoku [][]byte

func (s Sudoku) Solve() {
}

func (s Sudoku) Valid() bool {
	return isValidSudoku(s)
}

func (s Sudoku) Show() {
	fmt.Printf("||=========||=========||=========||\n")
	for i := 0; i < 9; i++ {
		fmt.Printf("||")
		for j := 0; j < 9; j++ {
			fmt.Printf(" %s ", string(s[i][j]))
			if j%3 == 2 {
				fmt.Printf("||")
			}
		}
		fmt.Printf("\n")
		if i%3 == 2 {
			fmt.Printf("||=========||=========||=========||\n")
		}
	}
}

func solveSudoku(board [][]byte) {
	solver(board, 0, 0)
}

func solver(board [][]byte, row, col int) bool {
	if row >= 9 {
		return true
	}
	if col >= 9 {
		return solver(board, row+1, 0)
	}
	if board[row][col] != '.' {
		return solver(board, row, col+1)
	}
	for i := 1; i <= 9; i++ {
		if isValid(board, row, col, byte(i)+'0') {
			board[row][col] = byte(i) + '0'
			if solver(board, row, col+1) {
				return true
			}
			board[row][col] = '.'
		}
	}
	return false
}

func isValid(board [][]byte, row, col int, target byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == target {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == target {
			return false
		}
	}
	for i := (row / 3) * 3; i < (row/3)*3+3; i++ {
		for j := (col / 3) * 3; j < (col/3)*3+3; j++ {
			if board[i][j] == target {
				return false
			}
		}
	}
	return true
}

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
	board := Sudoku{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	board.Show()
	solveSudoku(board)
	board.Show()
}
