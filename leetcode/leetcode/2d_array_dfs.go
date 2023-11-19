package leetcode

import "fmt"

func TwoDimensionalArrayDFSTraversal() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	fmt.Println(twoDimensionalArrayDFSTraversal(matrix))
}

func twoDimensionalArrayDFSTraversal(matrix [][]int) []int {
	M := len(matrix)
	N := len(matrix[0])

	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	seen := make([][]bool, M)
	for i := 0; i < M; i++ {
		seen[i] = make([]bool, N)
	}
	values := make([]int, 0, M*N)

	var travelDFS func(m, n int)
	travelDFS = func(m, n int) {
		if m < 0 || m >= M || n < 0 || n >= N || seen[m][n] {
			return
		}

		values = append(values, matrix[m][n])
		seen[m][n] = true

		for _, v := range directions {
			travelDFS(m+v[0], n+v[1])
		}
	}
	travelDFS(0, 0)

	fmt.Println(matrix)
	return values
}
