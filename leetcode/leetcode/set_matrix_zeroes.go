package leetcode

import "fmt"

func SetZeroes() {
	{
		matrix := [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		setZeroes(matrix)
		fmt.Println(matrix)
	}

	// {
	// 	matrix := [][]int{
	// 		{0, 1, 2, 0},
	// 		{3, 4, 5, 2},
	// 		{1, 3, 1, 5},
	// 	}
	// 	setZeroes(matrix)
	// 	fmt.Println(matrix)
	// }
}

func setZeroes(matrix [][]int) {
	M := len(matrix)
	N := len(matrix[0])
	rows := make([]int, M)
	cols := make([]int, N)

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				cols[j] = 1
			}
		}
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if rows[i] == 1 || cols[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
