package leetcode

import "fmt"

func WallsAndGates() {
	{
		rooms := [][]int{
			{2147483647, -1, 0, 2147483647},
			{-1, 2147483647, 2147483647, -1},
			{2147483647, -1, 2147483647, -1},
			{0, -1, 2147483647, 2147483647},
		}
		wallsAndGates(rooms)
		fmt.Println()
	}

	{
		rooms := [][]int{
			{2147483647, -1, 0, 2147483647},
			{2147483647, 2147483647, 2147483647, -1},
			{2147483647, -1, 2147483647, -1},
			{0, -1, 2147483647, 2147483647},
		}
		wallsAndGates(rooms)
		fmt.Println()
	}

	{
		rooms := [][]int{
			{-1},
		}
		wallsAndGates(rooms)
		fmt.Println()
	}
}

func wallsAndGates(rooms [][]int) {
	M := len(rooms)
	N := len(rooms[0])

	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	var travelDFS func(m, n, d int)
	travelDFS = func(m, n, d int) {
		if m < 0 || m >= M ||
			n < 0 || n >= N ||
			rooms[m][n] == -1 ||
			(d > rooms[m][n]) {
			return
		}

		rooms[m][n] = d
		for _, v := range directions {
			travelDFS(m+v[0], n+v[1], d+1)
		}
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if rooms[i][j] == 0 {
				travelDFS(i, j, 0)
			}
		}
	}

	for _, v := range rooms {
		fmt.Println(v)
	}
}
