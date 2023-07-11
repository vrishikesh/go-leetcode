package leetcode

import (
	"fmt"

	"github.com/vrishikesh/go-leetcode/queue/queue"
)

type Point struct {
	x, y, t int
}

func RottenOranges(grid [][]int) int {
	R := len(grid)
	C := len(grid[0])
	v := make([][]int, R)
	for i := range v {
		v[i] = make([]int, C)
	}

	q := queue.Queue[Point]{}
	noOfFreshOranges := 0

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 2 {
				q.Push(Point{i, j, 0})
				v[i][j] = 2
			}

			if grid[i][j] == 1 {
				noOfFreshOranges += 1
			}
		}
	}

	t := 0
	c := 0
	dRow := [4]int{-1, 0, +1, 0}
	dCol := [4]int{0, +1, 0, -1}
	for !q.IsEmpty() {
		p := q.Pop()

		for i := 0; i < 4; i++ {
			nRow := p.x + dRow[i]
			nCol := p.y + dCol[i]

			if nRow >= 0 && nRow < R &&
				nCol >= 0 && nCol < C &&
				grid[nRow][nCol] == 1 &&
				v[nRow][nCol] != 2 {
				if p.t+1 > t {
					t = p.t + 1
				}
				q.Push(Point{nRow, nCol, p.t + 1})
				v[nRow][nCol] = 2
				c += 1
			}
		}
	}

	ans := t
	if c != noOfFreshOranges {
		ans = -1
	}

	fmt.Println(ans)
	return t
}
