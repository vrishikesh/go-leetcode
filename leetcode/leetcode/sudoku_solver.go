package leetcode

import "fmt"

func SolveSudoku(board [][]byte) {
	printBoard(board)
	fmt.Println()

	var f func() bool
	f = func() bool {
		for row := byte(0); row < 9; row++ {
			for col := byte(0); col < 9; col++ {
				if board[row][col] == '.' {
					for c := byte(0); c < 9; c++ {
						if isValid(board, row, col, c+1) {
							board[row][col] = c + 1
							if f() {
								return true
							}
							board[row][col] = '.'
						}
					}
					return false
				}
			}
		}

		return true
	}

	f()
	printBoard(board)
	fmt.Println()
}

func isValid(board [][]byte, row, col, c byte) bool {
	for i := byte(0); i < 9; i++ {
		if board[row][i] == c {
			return false
		}

		if board[i][col] == c {
			return false
		}

		if board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}

	return true
}

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				fmt.Print(" .")
			} else {
				fmt.Print(" ", board[i][j])
			}
		}
		fmt.Println()
	}
}
