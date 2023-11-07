package leetcode

import (
	"container/list"
	"fmt"
)

type RO_Point struct {
	x, y, t int
}

func RottenOranges() {
	fmt.Println(rottenOranges([][]int{
		{0, 1, 2},
		{0, 1, 2},
		{2, 1, 1},
	}))
}

func rottenOranges(grid [][]int) int {
	R := len(grid)
	C := len(grid[0])
	v := make([][]int, R)
	for i := range v {
		v[i] = make([]int, C)
	}

	q := list.New()
	noOfFreshOranges := 0

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 2 {
				q.PushBack(RO_Point{i, j, 0})
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
	for q.Len() > 0 {
		p := q.Front().Value.(RO_Point)
		q.Remove(q.Front())

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
				q.PushBack(RO_Point{nRow, nCol, p.t + 1})
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
