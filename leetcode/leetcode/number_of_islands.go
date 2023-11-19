package leetcode

import "fmt"

func NumberOfIslands() {
	{
		grid := [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}
		fmt.Println(numIslands(grid))
	}

	{
		grid := [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '1'},
			{'0', '0', '0', '1', '1'},
		}
		fmt.Println(numIslands(grid))
	}

	{
		grid := [][]byte{
			{'0', '1', '0', '1', '0'},
			{'1', '0', '1', '0', '1'},
			{'0', '1', '1', '1', '0'},
			{'1', '0', '1', '0', '1'},
		}
		fmt.Println(numIslands(grid))
	}
}

func numIslands(grid [][]byte) int {
	M := len(grid)
	N := len(grid[0])
	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	ans := 0
	q := &Queue[[2]int]{}

	var travelBFS = func(m, n int) {
		q.Enqueue([2]int{m, n})
		for q.Len() > 0 {
			x := q.Dequeue()
			m := x[0]
			n := x[1]

			if m < 0 || m >= M || n < 0 || n >= N || grid[m][n] == '0' {
				continue
			}

			grid[m][n] = '0'
			for _, v := range directions {
				q.Enqueue([2]int{m + v[0], n + v[1]})
			}
		}
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '1' {
				ans += 1
				travelBFS(i, j)
			}
		}
	}

	return ans
}
