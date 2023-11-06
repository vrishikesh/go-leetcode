package leetcode

import (
	"fmt"
)

func RatInAMaze() {
	ratInAMaze([][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	})
}

func ratInAMaze(maze [][]int) {
	N := len(maze)
	dir := "DLRU"
	dr := []int{1, 0, 0, -1}
	dc := []int{0, -1, 1, 0}
	v := make([][]int, N)
	for i := 0; i < N; i++ {
		v[i] = make([]int, N)
	}

	var f func(int, int, string)
	var isValidMove func(int, int) bool
	printMatrix(maze)
	fmt.Println()

	f = func(row, col int, d string) {
		if row == N-1 && col == N-1 {
			fmt.Println(d)
			return
		}

		for i := 0; i < 4; i++ {
			nextr := row + dr[i]
			nextc := col + dc[i]
			if isValidMove(nextr, nextc) {
				v[row][col] = 1
				f(nextr, nextc, d+string(dir[i]))
				v[row][col] = 0
			}
		}
	}

	isValidMove = func(r, c int) bool {
		return r >= 0 && r < N &&
			c >= 0 && c < N &&
			v[r][c] == 0 &&
			maze[r][c] == 1
	}

	if maze[0][0] == 1 {
		f(0, 0, "")
	}
}

func printMatrix(m [][]int) {
	for _, r := range m {
		fmt.Println(r)
	}
}
