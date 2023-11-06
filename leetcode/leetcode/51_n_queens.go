package leetcode

import (
	"container/list"
	"fmt"
	"strings"
)

type Cell struct {
	row, col int
}

func NQueens() {
	fmt.Println(nQueens(4))
}

func nQueens(n int) [][]string {
	var ans [][]string
	var f func(int)
	q := list.New()

	f = func(col int) {
		if col == n {
			a := make([]string, 4)
			for e := q.Front(); e != nil; e = e.Next() {
				queen := e.Value.(Cell)
				a[queen.row] = strings.Repeat(".", queen.col) + "Q" + strings.Repeat(".", n-queen.col-1)
			}
			ans = append(ans, a)
			return
		}

		for row := 0; row < n; row++ {
			skip := false
			for e := q.Front(); e != nil; e = e.Next() {
				queen := e.Value.(Cell)
				if queen.col == col || queen.row == row || queen.col-col == queen.row-row || queen.col-col == row-queen.row {
					skip = true
					break
				}
			}

			if !skip {
				q.PushBack(Cell{row, col})
				f(col + 1)
				q.Remove(q.Front())
			}
		}
	}

	f(0)

	fmt.Println("ans", ans)
	return ans
}
