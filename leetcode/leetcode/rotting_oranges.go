package leetcode

import "fmt"

func OrangesRotting() {
	{
		grid := [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}
		fmt.Println(orangesRotting(grid)) // 4
	}

	{
		grid := [][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		}
		fmt.Println(orangesRotting(grid)) // -1
	}

	{
		grid := [][]int{
			{0, 2},
		}
		fmt.Println(orangesRotting(grid)) // 0
	}

	{
		grid := [][]int{
			{0},
		}
		fmt.Println(orangesRotting(grid)) // 0
	}

	{
		grid := [][]int{
			{1},
		}
		fmt.Println(orangesRotting(grid)) // -1
	}
}

func orangesRotting(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
	ans := -1
	freshOranges := 0

	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	q := &Queue[[2]int]{}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 2 {
				q.Enqueue([2]int{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	if freshOranges == 0 {
		return 0
	}

	for q.Len() > 0 {
		ans += 1
		L := q.Len()
		for i := 0; i < L; i++ {
			cell := q.Dequeue()
			m, n := cell[0], cell[1]

			for _, v := range directions {
				nextM := m + v[0]
				nextN := n + v[1]
				if nextM < 0 || nextM >= M || nextN < 0 || nextN >= N || grid[nextM][nextN] != 1 {
					continue
				}

				freshOranges--
				grid[nextM][nextN] = 2
				q.Enqueue([2]int{nextM, nextN})
			}
		}
	}

	if freshOranges > 0 {
		return -1
	}
	return ans
}
