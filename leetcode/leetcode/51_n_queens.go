package leetcode

import (
	"fmt"
	"queue/queue"
	"strings"
)

type Cell struct {
	row, col int
}

func NQueens(n int) [][]string {
	var ans [][]string
	var q queue.Queue[Cell]
	var f func(int)

	f = func(col int) {
		if col == n {
			a := make([]string, 4)
			for _, queen := range q.Iter() {
				a[queen.row] = strings.Repeat(".", queen.col) + "Q" + strings.Repeat(".", n-queen.col-1)
			}
			ans = append(ans, a)
			return
		}

		for row := 0; row < n; row++ {
			skip := false
			for _, queen := range q.Iter() {
				if queen.col == col || queen.row == row || queen.col-col == queen.row-row || queen.col-col == row-queen.row {
					skip = true
					break
				}
			}

			if !skip {
				q.Push(Cell{row, col})
				f(col + 1)
				q.Pop()
			}
		}
	}

	f(0)

	fmt.Println("ans", ans)
	return ans
}
