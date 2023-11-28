package leetcode

import "fmt"

func SudokuSolver() {
	sudokuSolver([][]byte{
		{53, 51, '.', '.', 55, '.', '.', '.', '.'},
		{54, '.', '.', 49, 57, 53, '.', '.', '.'},
		{'.', 57, 56, '.', '.', '.', '.', 54, '.'},
		{56, '.', '.', '.', 54, '.', '.', '.', 51},
		{52, '.', '.', 56, '.', 51, '.', '.', 49},
		{55, '.', '.', '.', 50, '.', '.', '.', 54},
		{'.', 54, '.', '.', '.', '.', 50, 56, '.'},
		{'.', '.', '.', 52, 49, 57, '.', '.', 53},
		{'.', '.', '.', '.', 56, '.', '.', 55, 57},
	})
}

func sudokuSolver(board [][]byte) {
	printBoard(board)
	fmt.Println()
	var recurse func() bool
	recurse = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					for k := byte('1'); k <= '9'; k++ {
						if isValidSudoku(board, i, j, k) {
							board[i][j] = k
							if recurse() {
								return true
							}
							board[i][j] = '.'
						}
					}
					return false
				}
			}
		}
		return true
	}
	recurse()
	printBoard(board)
}

func isValidSudoku(board [][]byte, r, c int, k byte) bool {
	for i := 0; i < 9; i++ {
		if board[r][i] == k {
			return false
		}
		if board[i][c] == k {
			return false
		}
		if board[3*(r/3)+i/3][3*(c/3)+i%3] == k {
			return false
		}
	}
	return true
}
